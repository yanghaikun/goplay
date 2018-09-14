package _map

import "fmt"

func TestDefaultValue() {
	test := make(map[string]string)
	test["one"] = "1"
	test["two"] = ""

	two, ok := test["two"]
	if !ok || len(two) <= 0 {
		fmt.Printf("value of two is %s", two)
	}
}
