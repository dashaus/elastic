package elastic

import (
	"encoding/json"
	"testing"
)

func TestMatchAllFilter(t *testing.T) {
	f := NewMatchAllFilter()
	data, err := json.Marshal(f.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"match_all":{}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
