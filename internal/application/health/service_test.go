package health

import "testing"

func TestServiceStatus(t *testing.T) {
	service := NewService()

	status := service.Status()

	if status.Status != "ok" {
		t.Fatalf("expected status ok, got %q", status.Status)
	}

	found := false
	for _, m := range memeMessages {
		if status.Message == m {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected message to be one of the meme messages, got %q", status.Message)
	}
}

func TestServiceStatusIsRandomized(t *testing.T) {
	service := NewService()

	seen := make(map[string]bool)
	for i := 0; i < 200; i++ {
		seen[service.Status().Message] = true
	}

	if len(seen) < 2 {
		t.Fatalf("expected multiple distinct messages across requests, got %d", len(seen))
	}
}
