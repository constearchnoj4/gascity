package main

import (
	"context"
	"testing"
	"time"

	"github.com/gastownhall/gascity/internal/runtime"
)

func TestMemoryIdleTrackerCheckIdle(t *testing.T) {
	sp := runtime.NewFake()
	sessionName := "gc-test-mayor"
	if err := sp.Start(context.Background(), sessionName, runtime.Config{}); err != nil {
		t.Fatalf("Start: %v", err)
	}

	tracker := newIdleTracker()
	now := time.Date(2026, 4, 20, 12, 0, 0, 0, time.UTC)
	tracker.setTimeout(sessionName, 10*time.Minute)

	sp.SetActivity(sessionName, now.Add(-11*time.Minute))
	if !tracker.checkIdle(sessionName, sp, now) {
		t.Fatal("checkIdle = false, want true for activity older than timeout")
	}

	sp.SetActivity(sessionName, now.Add(-5*time.Minute))
	if tracker.checkIdle(sessionName, sp, now) {
		t.Fatal("checkIdle = true, want false for recent activity")
	}

	tracker.setTimeout(sessionName, 0)
	if tracker.checkIdle(sessionName, sp, now) {
		t.Fatal("checkIdle = true, want false after clearing timeout")
	}
}
