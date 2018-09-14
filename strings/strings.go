package strings

import (
	"strings"
	"fmt"
	"math/rand"
)

func Repeat() {
	var list = []string{"one", "two", "three"}
	var s = strings.Join(list, ",")
	fmt.Println(s)
}

func NilToString() {
	var bytes []byte = nil
	fmt.Printf("result is %s", string(bytes))
}

var baseStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
func Random(length int) string{
	var result = ""
	for len(result) < length {
		var r = rand.Intn(len(baseStr))
		result += baseStr[r : r+1]
	}
	return result
}
