---
title: Tutorial 07 - Hooks
description: Understand how agents discover work, receive context, and stay connected to the city through provider hooks and bead hooks.
---

# Tutorial 07: Hooks

You've been using hooks since Tutorial 03 without thinking about them. Every time an agent checked for mail, picked up a formula step, or got its prompt — hooks made that happen. This tutorial explains the mechanism.

Gas City has two kinds of hooks: *provider hooks* that run between agent turns and inject context, and *bead hooks* that fire when work items change state. Provider hooks are how agents stay aware. Bead hooks are how the system stays consistent.

## How agents stay aware

An agent session is a conversation with an AI provider — Claude, Gemini, Codex, Copilot, or others. Between turns, the provider runs shell commands at specific lifecycle points. Gas City installs these commands so that agents automatically receive work, mail, and prompt context without anyone typing commands.

The lifecycle looks like this:

```
Session starts
  └── SessionStart hook → gc prime --hook
        (agent receives its full behavioral prompt)

User/system sends a message
  └── UserPromptSubmit hook → gc nudge drain --inject
                             → gc mail check --inject
        (agent receives pending nudges and unread mail)

Session compresses context
  └── PreCompact hook → gc prime --hook
        (agent re-receives its prompt before old context is pruned)

Agent finishes a turn
  └── Stop hook → gc hook --inject
        (agent receives any pending work items)
```

The exact hook points vary by provider — Claude has `Stop`, Gemini has `SessionEnd`, Codex has `Stop` — but the pattern is the same: the provider calls a Gas City command, and the output is injected into the agent's context as a `<system-reminder>`.

## Installing hooks

Hook installation happens automatically when the city starts or when a rig is added. You control which providers get hooks via `install_agent_hooks` in `city.toml`:

```toml
[workspace]
name = "my-city"
provider = "claude"
install_agent_hooks = ["claude", "codex", "gemini"]
```

This installs hook configuration files in the provider-specific locations:

| Provider | Hook file location |
|---|---|
| claude | `.gc/settings.json` and `hooks/claude.json` |
| codex | `.codex/hooks.json` |
| gemini | `.gemini/settings.json` |
| copilot | `.github/hooks/gascity.json` |
| cursor | `.cursor/hooks.json` |
| opencode | `.opencode/plugins/gascity.js` |

Installation is idempotent — existing files are never overwritten. If you've customized a provider's hook config, Gas City won't clobber it.

Claude hooks are installed by default (you don't need to list `"claude"` explicitly if your provider is `claude`).

You can override hook installation per agent:

```toml
[[agent]]
name = "worker"
install_agent_hooks = ["codex"]    # only codex hooks, not the workspace default
```

Or mark an agent as already having hooks installed (useful for hand-managed configs):

```toml
[[agent]]
name = "special"
hooks_installed = true
```

## The work pickup loop

The most important hook is the one that picks up work. Here's what happens when an agent's Stop hook fires:

1. The provider calls `gc hook --inject`
2. `gc hook` runs the agent's *work query* — a command that checks for available beads
3. If work exists, it wraps the result in `<system-reminder>` tags
4. The provider injects that into the agent's context
5. The agent sees the work and acts on it

This is GUPP in action: "If you find work on your hook, YOU RUN IT." The hook finding work *is* the assignment. No confirmation, no waiting.

You can run the hook command yourself to see what an agent would receive:

```
~/my-city
$ gc hook worker
[{"id":"gc-15","title":"Fix the login bug","status":"open"}]

~/my-city
$ gc hook --inject worker
<system-reminder>
You have pending work. Pick up the next item:

<work-items>
[{"id":"gc-15","title":"Fix the login bug","status":"open"}]
</work-items>

Claim it and start working. Run 'gc hook' to see the full queue.
</system-reminder>
```

Without `--inject`, you get raw output and an exit code (0 if work exists, 1 if empty). With `--inject`, output is always wrapped for agent consumption and always exits 0.

If no agent name is given, `gc hook` uses the `$GC_AGENT` environment variable — which is set automatically inside agent sessions.

## Agents without hooks

Not every provider supports hooks. If an agent doesn't have hooks installed, the *beacon* — the startup instruction in its prompt — includes a fallback:

> Run `gc prime` to refresh your context. Run `gc hook` to check for work.

The agent has to poll manually instead of being notified automatically. It works, but it's slower and less reliable. This is why `install_agent_hooks` matters — it's the difference between push and pull.

## The prime command

`gc prime` outputs the full behavioral prompt for an agent — its identity, capabilities, work instructions, and current context. It's called by the SessionStart hook to give the agent its initial briefing, and by the PreCompact hook to restore context when the conversation window is compressed.

```
~/my-city
$ gc prime worker
You are worker, a coding agent in the city "my-city"...
[full prompt template output]
```

With `--hook`, the output is formatted for hook injection rather than terminal display.

## Bead hooks

Provider hooks run between agent turns. Bead hooks run when beads change state in the store. They're shell scripts in `.beads/hooks/`:

| Hook | Fires when | What it does |
|---|---|---|
| `on_create` | A bead is created | Emits a `bead.created` event |
| `on_close` | A bead is closed | Emits `bead.closed`, auto-closes parent convoy if all children done, auto-closes orphaned wisps |
| `on_update` | A bead is updated | Emits a `bead.updated` event |

These hooks are installed automatically by `gc start`. They're the bridge between the bead store and the event bus — when a bead closes, the `on_close` hook fires an event that the event gate in an order might be watching. That's how `gate = "event"` with `on = "bead.closed"` works.

You don't need to manage bead hooks directly. They're infrastructure plumbing that keeps the system consistent.

## How work queries work

When `gc hook` checks for work, it runs the agent's configured work query. The default work query checks three things:

1. **Assigned work** — beads explicitly assigned to this agent (crash recovery picks these up)
2. **Pre-assigned work** — beads routed to this agent but not yet claimed
3. **Pool work** — beads labeled with the agent's pool (for pool agents)

The query returns JSON — an array of bead objects. `gc hook` normalizes the output (single objects become arrays, empty results become no output) and wraps it for injection.

For pool agents, the query uses label-based matching:

```
bd ready --label=pool:my-project/worker --unassigned --limit=1
```

This is the "pull" model: agents check for work rather than having work pushed to them. It's crash-safe (queued work survives restarts), simple to reason about, and scales naturally.

## Putting it together

Here's the full loop for an order dispatching work that an agent picks up via hooks:

```
Controller tick
  └── Order gate opens (cooldown elapsed)
        └── Controller creates wisp from formula
              └── Wisp bead labeled pool:my-project/worker

Agent finishes current turn
  └── Stop hook fires
        └── gc hook --inject
              └── Work query finds wisp bead
                    └── Output injected as <system-reminder>

Agent reads system reminder
  └── GUPP: claims bead, starts working
        └── Agent works through formula steps
              └── Closes bead when done
                    └── on_close bead hook fires
                          └── bead.closed event emitted
                                └── Another order's event gate opens
```

Orders create work. Hooks deliver it to agents. Bead hooks feed events back to the system. The controller evaluates gates on the next tick. The loop continues without human intervention.

---

| Command | What it does |
|---|---|
| `gc hook [agent]` | Check for available work (raw output, exit 0/1) |
| `gc hook --inject [agent]` | Check for work, wrap in system-reminder (always exit 0) |
| `gc prime [agent]` | Output the full behavioral prompt for an agent |
| `gc prime --hook` | Output prompt formatted for hook injection |

---

<!--
BONEYARD — draft material for future sections. Not part of the published tutorial.

### Supported providers

Gas City supports hooks for 8 providers: claude, codex, gemini, copilot,
cursor, opencode, pi, omp. Two providers (amp, auggie) have no hook mechanism
and fall back to beacon-based polling.

### Provider-specific hook points

Claude: SessionStart, PreCompact, UserPromptSubmit, Stop
Codex: SessionStart, Stop
Gemini: SessionStart, PreCompress, BeforeAgent, SessionEnd
Copilot: sessionStart (30s timeout), userPromptSubmitted (10s), sessionEnd (10s)
Cursor: sessionStart, preCompact, beforeSubmitPrompt, stop
OpenCode: JavaScript plugin in .opencode/plugins/gascity.js
Pi: JavaScript extension in .pi/extensions/gc-hooks.js
OMP: TypeScript hook in .omp/hooks/gc-hook.ts

### Hook installation internals

hooks.Install() is called from startBeadsLifecycle() during gc start or
gc rig add. It reads embedded config files from internal/hooks/config/ and
writes them to provider-specific locations. The write is idempotent — if the
file exists and isn't stale (based on a version marker), it's skipped.

For Gemini, there's upgrade logic: if the installed file is from an older
version, it's replaced. Other providers don't have this yet.

### AgentHasHooks() decision tree

The system determines if an agent has hooks via this logic:
1. If agent.HooksInstalled is explicitly set → use that value
2. If provider is "claude" → true (Claude always has hooks)
3. If provider is in the resolved install_agent_hooks list → true
4. Otherwise → false

When an agent has hooks, the beacon omits the manual polling instruction.
When it doesn't, the beacon includes "Run gc prime to refresh your context."

### Work query details

The work query is a shell command configured on the agent. The default query
runs bd ready with filters for assignee, pool labels, and unassigned work.
The output is JSON — an array of bead objects with id, title, status, type,
and labels.

gc hook normalizes the output:
- Empty string or "[]" → no work (exit 1 without --inject)
- Single JSON object → wrapped in array
- String containing "No ready work found" → treated as empty

### UserPromptSubmit hooks

The UserPromptSubmit hook (Claude) / BeforeAgent (Gemini) runs before the
agent processes each user message. It calls:
- gc nudge drain --inject — delivers any pending nudge messages
- gc mail check --inject — delivers any unread mail

This ensures the agent sees nudges and mail before responding to the user,
not just between turns.

### Bead hook scripts

The bead hooks are shell scripts that run in the background (& at the end).
They call gc event emit to record events, then gc convoy autoclose and
gc wisp autoclose for cascading lifecycle management.

The on_close hook receives the bead ID as $1 and the bead data as stdin JSON.
It extracts the title for the event message, then fires all three commands
asynchronously.

These scripts are installed to .beads/hooks/ by installBeadHooks() in
cmd/gc/hooks.go. They're regenerated on each gc start (not idempotent
like provider hooks — always overwritten).
-->
