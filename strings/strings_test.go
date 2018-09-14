package strings

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	Repeat()
}

func TestNilToString(t *testing.T) {
	NilToString()
}

func TestRandom(t *testing.T) {
	length := 26
	random := Random(length)
	if len(random) < length {
		t.Fatal("random string failed len is %d should be %d", len(random), length)
	}
}
