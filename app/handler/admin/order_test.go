package admin

import (
	"testing"
)

func TestNormalizePaidRefHashRejectsBlankValue(t *testing.T) {
	_, err := normalizePaidRefHash(" \t\n")
	if err == nil {
		t.Fatal("expected blank transaction hash to be rejected")
	}
}

func TestNormalizePaidRefHashTrimsValue(t *testing.T) {
	const hash = "0xd949d4dd02cceb58b8e352e901025a8987e9c0c3e091a1ff09c5339efacd2646"

	got, err := normalizePaidRefHash("  " + hash + "\n")
	if err != nil {
		t.Fatalf("normalize transaction hash: %v", err)
	}
	if got != hash {
		t.Fatalf("expected %q, got %q", hash, got)
	}
}
