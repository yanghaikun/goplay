package main

import (
	"fmt"
	"container/heap"
	"github.com/yanghaikun/goplay/container"
	"time"
	"github.com/yanghaikun/utils"
	"log"
	"path/filepath"
)



func main() {
	log.Printf("1 = ", 1)
	log.Println("2")
	log.Println("3")
	h := &container.IntHeap{2, 1, 5, 4}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

	//test filepath
	var test = "hello/world/1"
	hello := filepath.Base(test)
	fmt.Printf("the value of hello is %v", hello)
}

func printjson() {
	var print = func() {
		for {
			name := utils.RandomStr(5)
			age := utils.RandomInt(100)
			fmt.Printf("{\"name\":\"%v\", \"age\":\"%v\"}\n", name, age)
			time.Sleep(time.Second)
		}
	}

	//go print()
	print()
}
