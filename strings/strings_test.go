package strings

import (
	"testing"
	"fmt"
)

func TestRepeat(t *testing.T) {
	Repeat()
}

func TestNilToString(t *testing.T) {
	NilToString()
}

func TestRandom(t *testing.T) {
	length := 10
	random := Random(length)
	if len(random) < length {
		t.Fatal("random string failed len is %d should be %d", len(random), length)
	}
	fmt.Println(random)
}
