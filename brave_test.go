package go_llama_agentic_system

import (
	"os"
	"testing"
)

func TestBraveSearch_Query(t *testing.T) {
	apiKey := os.Getenv("BRAVE_API_KEY")
	if os.Getenv("BRAVE_API_KEY") == "" {
		t.Fatal("environment variable BRAVE_API_KEY not set")
	}
	bs := NewBraveSearch(apiKey)
	sr, err := bs.Query("brave search", 3)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if len(sr) == 0 {
		t.Errorf("expected search results to be greater than zero")
	}
}
