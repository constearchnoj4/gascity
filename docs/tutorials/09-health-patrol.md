---
title: Tutorial 09 - Health Patrol
description: Understand the controller's tick loop — how it monitors agent health, reconciles sessions, dispatches orders, and keeps the city running.
---

# Tutorial 09: Health Patrol

When you run `gc start`, you're launching a controller. The controller is a single process that wakes up every 30 seconds (by default), looks at the state of the city, and takes action: starting agents that should be running, stopping ones that shouldn't, firing due orders, cleaning up finished work. This periodic check is the *health patrol*.

You've already seen pieces of this. Orders fire because the controller evaluates gates on each tick. Agents restart after crashes because the controller notices they're gone. Wisps get garbage-collected because the controller sweeps for expired ones. This chapter shows the full picture.

## The tick loop

Every 30 seconds — or when poked by `gc sling` — the controller runs through these phases in order:

```
┌─ Pool death detection
│    Dead pool instances → run on_death handlers
│
├─ Config reload (if changed)
│    city.toml modified → rebuild all trackers
│
├─ Session bead sync
│    Running sessions ↔ bead store reconciliation
│
├─ Bead-driven reconciliation
│    Wake/sleep decisions based on demand
│
├─ Wisp GC
│    Closed wisps older than TTL → deleted
│
├─ Order dispatch
│    Evaluate gates → fire due orders
│
├─ Service tick
│    Runtime service management
│
├─ Chat session auto-suspend
│    Idle detached sessions → suspended
│
└─ Convergence tick
     Active workflow state machines → advance
```

Each phase is independent. A failure in one doesn't block the others (though they run sequentially within a tick).

## Agent reconciliation

The core of the patrol is deciding which agents should be running. This is *demand-driven* — agents wake when there's a reason to be awake, and sleep when there isn't.

On each tick, the controller evaluates *wake reasons* for every session bead:

| Wake reason | What it means |
|---|---|
| **WakeWork** | Work query found pending work for this agent |
| **WakeConfig** | Pool desires more instances of this agent template |
| **WakeSession** | An active session exists and matches desired config |
| **WakeKeepWarm** | Agent has a `keep_warm` policy (stay running even without work) |
| **WakeAttached** | A user has a tmux pane attached to this session |
| **WakePending** | A pending interaction is ready for the agent |
| **WakeCreate** | A pending session creation claim exists |
| **WakeWait** | Agent is in a ready-wait set |

If any wake reason applies, the agent stays running (or gets started). If none apply, the agent drains and sleeps.

This is how Gas City avoids burning API credits on idle agents. An agent with no work and no config demand simply goes to sleep. When work arrives — via `gc sling`, a formula dispatch, or mail — the next tick detects it through the work query and wakes the agent.

### Dependency ordering

Agents can declare dependencies:

```toml
[[agent]]
name = "reviewer"
depends_on = ["worker"]
```

The reconciler topologically sorts agents by their dependencies before making wake/sleep decisions. The `worker` is evaluated first, ensuring it's running before `reviewer` tries to start.

## Crash loop detection

Agents crash. That's expected — the prompt might be wrong, the tool might fail, the provider might hiccup. The controller restarts crashed agents on the next tick. But if an agent keeps crashing, something is fundamentally broken, and restarting it indefinitely is wasteful.

The crash tracker uses a sliding window:

```toml
[daemon]
max_restarts = 5        # quarantine after this many restarts
restart_window = "1h"   # within this time window
```

If an agent restarts 5 times within an hour, it's *quarantined* — the controller stops trying to start it. The quarantine clears when enough time passes for old restarts to fall outside the window.

This follows the Erlang/OTP supervision model: let it crash, but don't let it crash forever. The crash tracker is in-memory — a controller restart clears all quarantine state, which is intentional (the Erlang parallel: restarting the supervisor clears child restart counts).

You can see quarantined agents in the event log:

```
~/my-city
$ gc events --type session.quarantined
SEQ  TYPE                   ACTOR  SUBJECT    MESSAGE                    TIME
8    session.quarantined    gc     worker-1   crash loop detected        10:05:30
```

## Idle timeout

Agents that are running but not doing anything can be stopped to free resources:

```toml
[[agent]]
name = "worker"
idle_timeout = "30m"
```

On each tick, the controller checks how long since the agent's last activity. If it exceeds `idle_timeout`, the session is killed. The event log records `session.idle_killed`.

This only works if the provider supports activity tracking. If it doesn't, the idle check is skipped gracefully.

## Wisp garbage collection

Wisps are ephemeral — they're meant to be created, executed, and forgotten. But "forgotten" doesn't mean instant deletion. The wisp GC runs periodically and deletes closed wisps older than a TTL:

```toml
[daemon]
wisp_gc_interval = "5m"    # check every 5 minutes
wisp_ttl = "24h"           # keep closed wisps for 24 hours
```

Both settings must be present to enable wisp GC. If either is missing, closed wisps accumulate until you clean them up manually.

The GC is purely mechanical: list closed molecules, delete any older than the TTL. No intelligence, no judgment.

## Config reload

The controller watches `city.toml` via filesystem notifications. When the file changes, the next tick reloads the config and rebuilds all trackers — crash tracker, idle tracker, wisp GC, and order dispatcher.

This means you can change patrol settings, agent definitions, or order overrides without restarting the city. The controller picks up changes automatically.

However, some changes require a full restart — adding a new formula layer directory, for instance, since that requires rescanning for orders. If you're unsure, `gc stop && gc start` is always safe.

## Configuration reference

```toml
[daemon]
patrol_interval = "30s"      # tick frequency (default: 30s)
max_restarts = 5             # crash loop threshold (default: 5, 0 = unlimited)
restart_window = "1h"        # sliding window for crash counting (default: 1h)
shutdown_timeout = "5s"      # grace period before force-kill (default: 5s)
wisp_gc_interval = "5m"      # wisp GC frequency (disabled if unset)
wisp_ttl = "24h"             # closed wisp retention (disabled if unset)
drift_drain_timeout = "2m"   # config-drift restart grace period (default: 2m)
```

Most cities don't need to touch these. The defaults are reasonable for development and small production setups. You might lower `patrol_interval` if you need faster reaction times, or increase `wisp_ttl` if you want more history in your bead store.

## Nondeterministic Idempotence

The health patrol embodies Gas City's NDI principle: the system converges to correct outcomes because work is persistent and observers are idempotent.

If the controller crashes mid-tick, nothing is lost. Beads persist in the store. Agent processes persist in the process table. On restart, the controller evaluates the same state and makes the same decisions. There are no status files to go stale, no in-memory queues to lose, no distributed locks to expire.

Multiple ticks with the same inputs produce the same outputs. An agent that should be running will be started on every tick until it's running. An order that's due will fire on every tick until the tracking bead prevents re-fire. The patrol is a convergence loop, not a sequence of irreversible steps.

This is why Gas City is crash-safe by design rather than by accident.

## Observing the patrol

The patrol doesn't have its own CLI command — it's internal to the controller. But you can observe its effects through events:

```
~/my-city
$ gc events --since 10m
SEQ  TYPE                  ACTOR        SUBJECT      MESSAGE                  TIME
1    controller.started    gc                                                 10:00:00
2    session.woke          gc           worker-1     agent started            10:00:01
3    order.fired           controller   health-ch    order dispatched         10:00:30
4    session.crashed       gc           worker-1     session died             10:02:15
5    session.woke          gc           worker-1     agent restarted          10:02:30
6    order.fired           controller   health-ch    order dispatched         10:03:00
```

Every lifecycle event — agent starts, stops, crashes, quarantines, idle kills — is recorded. The event log is the patrol's audit trail.

---

<!--
BONEYARD — draft material for future sections. Not part of the published tutorial.

### Pool death detection

Pool instances are tracked via poolDeathHandlers. When an instance dies
between ticks, the controller runs on_death scripts configured on the pool.
This happens before reconciliation so that death-triggered cleanup
(worktree salvage, bead reassignment) completes before new instances start.

### The poke mechanism

When gc sling assigns work, it sends a poke signal to the controller via
the Unix socket. This wakes the controller immediately instead of waiting
for the next 30-second tick. The poke triggers a full tick cycle, ensuring
the target agent is started quickly.

### Config drift detection

When an agent's config changes (prompt, environment, provider), the
controller detects the drift via SHA-256 fingerprinting of the agent's
configuration fields. Drifted agents receive a drain signal with a
configurable timeout (drift_drain_timeout, default 2m), allowing in-flight
work to complete before the agent restarts with the new config.

### Convergence tick

The convergence subsystem (formula v2) has its own state machine that
advances on each controller tick. It processes active convergence loops,
evaluates gate conditions, manages manual approval queues, and handles
iteration retries. This is separate from orders but shares the same
tick-driven execution model.

### Chat session auto-suspend

Sessions that are detached (no user attached) and idle beyond a threshold
are auto-suspended. This is separate from idle_timeout (which kills) —
auto-suspend preserves the session state for later resume.

### Service tick

Runtime services (API server, etc.) get a tick callback for periodic
maintenance tasks. This is a thin hook for service-layer concerns that
don't belong in the reconciliation loop.

### Shutdown sequence

On gc stop or SIGINT:
1. Controller records controller.stopped event
2. All agents receive graceful drain
3. Wait for shutdown_timeout (default 5s)
4. Force-kill remaining processes
5. Release flock on .gc/controller.lock

### Controller singleton

The controller uses flock on .gc/controller.lock to ensure only one
controller runs per city. If gc start finds the lock held, it exits
with an error. This prevents split-brain reconciliation.

### Bead-driven vs config-driven

Reconciliation is bead-driven, not config-driven. The controller
doesn't compare "what agents should exist" against "what agents do
exist." Instead, it evaluates wake reasons for each session bead. A
bead with a wake reason gets woken. A bead without wake reasons gets
drained. The config influences wake reasons (WakeConfig for pool
sizing, agent definitions for templates) but doesn't directly drive
lifecycle decisions.

This is the key insight: beads are the single source of truth for
desired state, not city.toml. The config creates beads, and beads
drive reconciliation.
-->
