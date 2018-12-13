package function

import "testing"

func TestHoly(t *testing.T) {
	if Holy() != "Holy" {
		t.Error("Holy crap")
	}
}
