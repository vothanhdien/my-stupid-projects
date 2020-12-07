package main

import "fmt"

func main() {
	basicDefer()
	deferStack()
	deferInCondition()
	deferInSwitch()
	deferInLoop()
}

func deferStack() {
	fmt.Println("deferStack start")
	defer fmt.Println("deferStack", "first")
	defer fmt.Println("deferStack", "second")
	defer fmt.Println("deferStack", "third")
	defer fmt.Println("deferStack", "fourth")
	defer fmt.Println("deferStack", "firth")
	fmt.Println("deferStack end")
}

func basicDefer() {
	fmt.Println("basicDefer", "start")
	defer fmt.Println("basicDefer", "defer", "end")
	fmt.Println("basicDefer", "end")
}

func deferInCondition() {
	fmt.Println("deferInCondition", "start")
	if true {
		defer fmt.Println("deferInCondition", "condition clause")
	}
	fmt.Println("deferInCondition", "end")
}

func deferInLoop() {
	fmt.Println("deferInLoop", "start")
	for i := 0; i < 1; i++ {
		defer fmt.Println("deferInLoop", "loop clause")
	}
	fmt.Println("deferInLoop", "end")
}

func deferInSwitch() {
	fmt.Println("deferInSwitch", "start")
	switch 1 {
	default:
		defer fmt.Println("deferInSwitch", "switch clause")
	}
	fmt.Println("deferInSwitch", "end")
}
