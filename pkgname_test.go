package template

import (
	"testing"
)

func TestDemoFunc0(t *testing.T) {
	e := DemoFunc0(2, 3)
	r := 5
	if e != r {
		t.Errorf("expected: %v, real: %v", e, r)
	}
}
