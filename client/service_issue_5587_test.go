package client

import (
	"os"
	"strings"
	"testing"
)

func TestLocalRyanIssue5587NetpollWorkaround(t *testing.T) {
	val := os.Getenv("GODEBUG")
	if val == "" {
		t.Fatalf("GODEBUG not set, got empty, want netpoll2=0")
	}
	if !strings.Contains(val, "netpoll2=0") {
		t.Fatalf("GODEBUG=%q, want to contain netpoll2=0", val)
	}
}
