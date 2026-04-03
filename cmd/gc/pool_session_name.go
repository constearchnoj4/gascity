package main

import (
	"path"
	"strings"

	"github.com/gastownhall/gascity/internal/beads"
)

// PoolSessionName derives the tmux session name for a pool worker session.
// Format: {basename(template)}-{beadID} (e.g., "claude-mc-xyz").
// Named sessions with an alias use the alias instead.
func PoolSessionName(template, beadID string) string {
	base := path.Base(template)
	// Sanitize: replace "/" with "--" for tmux compatibility.
	base = strings.ReplaceAll(base, "/", "--")
	return base + "-" + beadID
}

// GCSweepSessionBeads closes open session beads that have no remaining
// assigned work beads (all assigned beads are closed). Returns the IDs
// of session beads that were closed.
func GCSweepSessionBeads(store beads.Store, sessionBeads []beads.Bead, allWorkBeads []beads.Bead) []string {
	// Index work beads by assignee.
	assigneeHasWork := make(map[string]bool)
	for _, wb := range allWorkBeads {
		if wb.Status == "closed" {
			continue
		}
		assignee := strings.TrimSpace(wb.Assignee)
		if assignee != "" {
			assigneeHasWork[assignee] = true
		}
	}

	var closed []string
	for _, sb := range sessionBeads {
		if sb.Status == "closed" {
			continue
		}
		// Check if any non-closed work bead is assigned to this session
		// via any identifier: bead ID, session name, or named identity (alias).
		if sessionHasAssignedWork(sb, assigneeHasWork) {
			continue
		}
		if err := store.SetMetadata(sb.ID, "state", "gc_swept"); err != nil {
			continue
		}
		if err := store.Close(sb.ID); err != nil {
			continue
		}
		closed = append(closed, sb.ID)
	}
	return closed
}

// ClearOrphanedWorkAssignees finds open/in_progress work beads assigned to
// sessions that no longer exist and clears their assignee. This prevents
// work from getting stuck when a session drains before continuation-group
// beads become ready. Returns the IDs of beads whose assignees were cleared.
func ClearOrphanedWorkAssignees(store beads.Store, allBeads []beads.Bead, workBeads []beads.Bead) []string {
	// Build set of all known session identifiers (bead ID, session name, alias)
	// from open session beads.
	knownSession := make(map[string]bool)
	for _, b := range allBeads {
		if b.Type != "session" || b.Status == "closed" {
			continue
		}
		knownSession[b.ID] = true
		if sn := strings.TrimSpace(b.Metadata["session_name"]); sn != "" {
			knownSession[sn] = true
		}
		if ni := strings.TrimSpace(b.Metadata["configured_named_identity"]); ni != "" {
			knownSession[ni] = true
		}
	}

	var cleared []string
	for _, wb := range workBeads {
		if wb.Status == "closed" || wb.Type == "session" {
			continue
		}
		assignee := strings.TrimSpace(wb.Assignee)
		if assignee == "" {
			continue
		}
		// Skip if the assignee matches any known live session.
		if knownSession[assignee] {
			continue
		}
		// Assignee points to a dead session — clear it so the bead
		// becomes visible to EffectiveWorkQuery tier 3 (unassigned).
		empty := ""
		if err := store.Update(wb.ID, beads.UpdateOpts{Assignee: &empty}); err != nil {
			continue
		}
		cleared = append(cleared, wb.ID)
	}
	return cleared
}

// sessionHasAssignedWork checks whether any work bead is assigned to this
// session bead via any of its identifiers: bead ID, session name, or
// named identity (alias).
func sessionHasAssignedWork(sb beads.Bead, assigneeHasWork map[string]bool) bool {
	if assigneeHasWork[sb.ID] {
		return true
	}
	if sn := strings.TrimSpace(sb.Metadata["session_name"]); sn != "" && assigneeHasWork[sn] {
		return true
	}
	if ni := strings.TrimSpace(sb.Metadata["configured_named_identity"]); ni != "" && assigneeHasWork[ni] {
		return true
	}
	return false
}
