//go:build acceptance_a

// Handoff command acceptance tests.
//
// These exercise gc handoff error paths. Handoff is the context
// continuation mechanism — it sends mail and restarts a session.
// Full lifecycle tests belong in Tier B (needs real sessions).
// Tier A covers argument validation and missing-context errors.
package acceptance_test

import (
	"testing"

	helpers "github.com/gastownhall/gascity/test/acceptance/helpers"
)

func TestHandoff_NoArgs_ReturnsError(t *testing.T) {
	c := helpers.NewCity(t, testEnv)
	c.Init("claude")

	// cobra.RangeArgs(1,2) rejects zero args.
	_, err := c.GC("handoff")
	if err == nil {
		t.Fatal("expected error for handoff without args, got success")
	}
}

func TestHandoff_NoSessionContext_ReturnsError(t *testing.T) {
	c := helpers.NewCity(t, testEnv)
	c.Init("claude")

	// Self-handoff requires GC_ALIAS/GC_SESSION_ID which aren't set in tests.
	_, err := c.GC("handoff", "test subject")
	if err == nil {
		t.Fatal("expected error for handoff without session context, got success")
	}
}

func TestHandoff_RemoteNonexistent_ReturnsError(t *testing.T) {
	c := helpers.NewCity(t, testEnv)
	c.Init("claude")

	_, err := c.GC("handoff", "--target", "nonexistent-session", "test subject")
	if err == nil {
		t.Fatal("expected error for handoff to nonexistent target, got success")
	}
}

func TestHandoff_TooManyArgs_ReturnsError(t *testing.T) {
	c := helpers.NewCity(t, testEnv)
	c.Init("claude")

	// cobra.RangeArgs(1,2) rejects three args.
	_, err := c.GC("handoff", "subject", "message", "extra")
	if err == nil {
		t.Fatal("expected error for too many args, got success")
	}
}

func TestHandoff_RemoteWithMessage_ReturnsError(t *testing.T) {
	c := helpers.NewCity(t, testEnv)
	c.Init("claude")

	// Remote handoff with nonexistent target + message body.
	_, err := c.GC("handoff", "--target", "nonexistent", "subject", "body message")
	if err == nil {
		t.Fatal("expected error for remote handoff to nonexistent, got success")
	}
}

// --- gc graph (bonus: related command) ---

func TestGraph_EmptyCity_Succeeds(t *testing.T) {
	c := helpers.NewCity(t, testEnv)
	c.Init("claude")

	// graph should not crash on an empty city.
	_, _ = c.GC("graph")
}

func TestGraph_NotInitialized_ReturnsError(t *testing.T) {
	emptyDir := t.TempDir()
	_, err := helpers.RunGC(testEnv, emptyDir, "graph")
	if err == nil {
		t.Fatal("expected error for graph on non-city directory, got success")
	}
}
