package repository

import "testing"

func TestActiveVideoTaskStatusesIncludeProviderNotStartedStates(t *testing.T) {
	active := make(map[string]bool, len(activeVideoTaskStatuses))
	for _, status := range activeVideoTaskStatuses {
		active[status] = true
	}
	for _, status := range []string{"not_start", "not_started"} {
		if !active[status] {
			t.Fatalf("activeVideoTaskStatuses does not include %q", status)
		}
	}
}
