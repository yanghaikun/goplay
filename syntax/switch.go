package syntax

import (
	"log"
	"fmt"
)

func Switch() {
	var s = 0
	switch s {
	case 0:
		log.Println(0)
	case 1:
		log.Println(1)
		fallthrough
	case 2:
		log.Println(2)
		goto A
	goto B
	}

A:
	fmt.Println("A")
	fmt.Println("test")
B:
	fmt.Println("B")
}
