package strings

import (
	"strings"
	"fmt"
	"math/rand"
	"time"
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
	rand.Seed(time.Now().UnixNano())
	var result = ""
	var n = len(baseStr)
	for len(result) < length {
		var r = rand.Intn(n)
		result += baseStr[r : r+1]
	}
	return result
}
