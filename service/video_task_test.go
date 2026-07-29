package service

import "testing"

func TestNormalizeVideoTaskStatusKeepsNotStartedTasksQueued(t *testing.T) {
	for _, status := range []string{"NOT_START", "not_started"} {
		if got := NormalizeVideoTaskStatus(status); got != "queued" {
			t.Fatalf("NormalizeVideoTaskStatus(%q) = %q, want queued", status, got)
		}
	}
}
