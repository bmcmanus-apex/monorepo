package fortune

import "testing"

func TestGet(t *testing.T) {
	fortune := Get()
	if fortune == "" {
		t.Errorf("Get() returned an empty string")
	}
}
