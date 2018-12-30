package learning

import (
	"testing"
)

func TestHello(t *testing.T) {
	res := Hello()
	if res != "say Hello." {
		t.Error("some err")

	}
}

