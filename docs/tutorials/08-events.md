---
title: Tutorial 08 - Events
description: Observe system activity through Gas City's append-only event bus — query, filter, watch, and emit events from the CLI and API.
---

# Tutorial 08: Events

Every significant thing that happens in Gas City — an agent starting, a bead closing, an order firing, mail being sent — gets recorded as an event. The event bus is an append-only log of system activity, and it's what makes the system observable.

You've already seen events at work without knowing it. When a bead closes and an order's event gate fires — that's the event bus. When the controller records `order.completed` after dispatching a formula — event bus. The hooks chapter showed bead hooks emitting `bead.created` and `bead.closed` events. This chapter shows you how to see and use those events directly.

## Viewing events

```
~/my-city
$ gc events
SEQ  TYPE              ACTOR        SUBJECT    MESSAGE                          TIME
1    controller.started gc                                                      10:00:00
2    session.woke      gc           worker-1   agent started successfully       10:00:01
3    bead.created      gc           gc-15      Fix the login bug                10:00:05
4    order.fired       controller   health-ch  order dispatched                 10:00:30
5    bead.closed       gc           gc-15      Fix the login bug                10:02:14
6    order.completed   controller   health-ch  order completed                  10:00:32
```

Every event has:

- **Seq** — a monotonically increasing sequence number. Events are never reordered.
- **Type** — what happened (`bead.closed`, `order.fired`, `session.woke`, etc.)
- **Actor** — who did it (an agent name, `gc`, `controller`, `convergence`)
- **Subject** — what was affected (a bead ID, agent name, order name)
- **Message** — human-readable description
- **Time** — when it happened

## Filtering events

The event list can be filtered by type, time, or sequence:

```
~/my-city
$ gc events --type bead.closed
SEQ  TYPE          ACTOR  SUBJECT  MESSAGE              TIME
5    bead.closed   gc     gc-15    Fix the login bug    10:02:14
12   bead.closed   gc     gc-20    Update API docs      10:15:33

~/my-city
$ gc events --since 5m
SEQ  TYPE              ACTOR        SUBJECT    MESSAGE              TIME
10   order.fired       controller   cleanup    order dispatched     10:12:00
11   bead.created      gc           gc-20      Update API docs      10:14:01
12   bead.closed       gc           gc-20      Update API docs      10:15:33

~/my-city
$ gc events --after 10
SEQ  TYPE          ACTOR  SUBJECT  MESSAGE              TIME
11   bead.created  gc     gc-20    Update API docs      10:14:01
12   bead.closed   gc     gc-20    Update API docs      10:15:33
```

You can also filter by payload fields:

```
~/my-city
$ gc events --type order.fired --payload-match name=health-check
```

For machine consumption, add `--json` to get a JSON array.

## Watching for events

To block until a specific event arrives:

```
~/my-city
$ gc events --watch --type bead.closed
```

This waits (30 seconds by default) for the next `bead.closed` event, prints it, and exits. Useful in scripts that need to react to a specific state change.

To continuously stream events as they happen:

```
~/my-city
$ gc events --follow
SEQ  TYPE              ACTOR        SUBJECT    MESSAGE              TIME
13   order.fired       controller   health-ch  order dispatched     10:15:00
14   order.completed   controller   health-ch  order completed      10:15:02
15   session.woke      gc           worker-2   agent started        10:15:30
...
```

`--follow` never exits on its own — it streams until you interrupt it. This is the real-time view of what the city is doing.

## Emitting custom events

You can record your own events from scripts, hooks, or the command line:

```
~/my-city
$ gc event emit deploy.started --subject my-app --message "Deploying v2.1"
```

The event appears in the log alongside system events:

```
~/my-city
$ gc events --type deploy.started
SEQ  TYPE             ACTOR   SUBJECT  MESSAGE           TIME
16   deploy.started   human   my-app   Deploying v2.1    10:20:00
```

The `--actor` flag controls who the event is attributed to. Inside an agent session, it defaults to the agent's name. From the command line, it defaults to `human`. You can also attach arbitrary JSON:

```
~/my-city
$ gc event emit build.completed --subject my-app --payload '{"version":"2.1","duration":"45s"}'
```

`gc event emit` is best-effort — it always exits 0, even if recording fails. This makes it safe to use in bead hooks and scripts where you don't want a logging failure to break the workflow.

## Event types

Here are the event types you'll encounter:

| Category | Types |
|---|---|
| Agent lifecycle | `session.woke`, `session.stopped`, `session.crashed`, `session.draining`, `session.quarantined`, `session.idle_killed`, `session.suspended` |
| Work | `bead.created`, `bead.closed`, `bead.updated` |
| Mail | `mail.sent`, `mail.read`, `mail.replied`, `mail.deleted` |
| Orders | `order.fired`, `order.completed`, `order.failed` |
| Convoys | `convoy.created`, `convoy.closed` |
| Controller | `controller.started`, `controller.stopped` |
| City | `city.suspended`, `city.resumed` |

Event types are free-form strings. The ones above are the system-defined types, but `gc event emit` can record any type you want. There's no schema enforcement.

## Events and orders

The event bus is what makes event-gated orders work. When you write:

```toml
[order]
formula = "convoy-check"
gate = "event"
on = "bead.closed"
pool = "worker"
```

The controller evaluates this by querying the event bus for `bead.closed` events after a cursor position. The cursor tracks the highest sequence number already processed, so each `bead.closed` event triggers the order at most once.

This is the reactive pattern: instead of polling on a timer, the order fires in direct response to system activity. A bead closes → event is recorded → event gate opens → order dispatches formula → agent does work → closes another bead → cycle continues.

## Storage

Events are stored in `.gc/events.jsonl` — a newline-delimited JSON file. Each line is one event. The format is append-only: events are never modified or deleted after recording.

The file recorder uses `O_APPEND` for cross-process safety — multiple processes can record events concurrently without corruption. Sequence numbers are assigned monotonically within a process, and the file is the single source of truth.

For most cities, the default file backend works fine. You can also configure an external script as the event provider:

```toml
[events]
provider = "exec:scripts/my-events.sh"
```

The script receives subcommands (`record`, `list`, `latest-seq`, `watch`) and communicates via JSON on stdin/stdout. This is the escape hatch for integrating with external logging systems.

## Events as debugging tool

When something isn't working — an agent isn't picking up work, an order isn't firing, a convoy isn't closing — the event log is where you look first:

```
~/my-city
$ gc events --since 1h --type order.fired,order.failed,order.completed
```

Did the order fire? Did it fail? When? What was the subject? The event log gives you the timeline without digging through agent sessions or bead state.

For agent issues:

```
~/my-city
$ gc events --type session.crashed,session.quarantined --since 2h
```

For work flow:

```
~/my-city
$ gc events --type bead.created,bead.closed --since 30m
```

The event bus is the audit trail for everything the city does.

---

| Command | What it does |
|---|---|
| `gc events` | List all events |
| `gc events --type <type>` | Filter by event type |
| `gc events --since <duration>` | Show events from duration ago |
| `gc events --after <seq>` | Resume from sequence number |
| `gc events --watch --type <type>` | Block until matching event arrives |
| `gc events --follow` | Continuously stream events |
| `gc events --json` | Output as JSON array |
| `gc events --seq` | Print current head sequence number |
| `gc event emit <type>` | Record a custom event |

---

<!--
BONEYARD — draft material for future sections. Not part of the published tutorial.

### Best-effort semantics

Event recording is best-effort by design. The Recorder.Record() method never
returns an error — failures are logged to stderr. This means:

- Recording an event never fails the caller's operation
- A bead can close successfully even if the event recording fails
- Scripts using gc event emit always exit 0

This is intentional. The event bus is an observation layer, not a
transactional log. It should never block or break the operations it observes.

### Two-tier delivery (convergence)

The convergence subsystem uses a tiered event model:
- TierCritical — at-least-once for state changes (iteration, terminated)
- TierRecoverable — best-effort with reconciliation
- TierBestEffort — fire-and-forget for post-durable-state events

This is internal to convergence. Regular events are all best-effort.

### Multiplexer

The event multiplexer (internal/events/multiplexer.go) aggregates events
from multiple cities into a single stream. This powers the dashboard's
cross-city event feed.

Each event is tagged with its source city. Cursors are formatted as
comma-separated "city:seq" pairs for persistence across connections.

### API endpoints

The API server exposes event endpoints:
- POST /v0/events — emit a custom event
- GET /v0/events — list/filter events (supports wait= for long-polling)
- GET /v0/events/stream — Server-Sent Events stream

The SSE endpoint uses Watcher internally, polling at 250ms for the
FileRecorder. The API is how the web dashboard gets real-time updates.

### FileRecorder internals

The FileRecorder uses a mutex for in-process serialization and O_APPEND
for cross-process safety. On startup, it reads the existing file to
determine the next sequence number (resumeSeq).

The fileWatcher polls every 250ms using ReadFrom() for incremental reads.
ReadFrom reads from a byte offset and handles partial writes gracefully
(incomplete last lines are ignored until the next poll).

### exec provider wire protocol

The exec provider delegates to a user-supplied script:
- "script ensure-running" — health check
- "script record" — record event (JSON on stdin)
- "script list" — query events (JSON filter on stdin, JSON array on stdout)
- "script latest-seq" — get highest seq (number on stdout)
- "script watch" — long-running subprocess, events on stdout

Exit code 2 = unknown operation (forward compatible). Record is
fire-and-forget. List/LatestSeq fork per call. Watch spawns a
long-running subprocess.

### Convergence events

The convergence subsystem records its own event types:
- convergence.created, convergence.iteration, convergence.terminated
- convergence.waiting_manual, convergence.manual_approve
- convergence.manual_iterate, convergence.manual_stop

These carry structured payloads with formula name, iteration count,
gate results, token usage, and duration. They use stable event IDs
(converge:<beadID>:<event>) for deduplication.

### Provider events

- provider.swapped — session provider changed (e.g., claude → gemini)

### External messaging events

- extmsg.bound, extmsg.unbound — external message service lifecycle
- extmsg.group_created — group created
- extmsg.adapter_added, extmsg.adapter_removed — adapter lifecycle
-->
