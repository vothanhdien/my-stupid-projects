package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	//wg := sync.WaitGroup{}
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		s, err := runTest()
	//		if err != nil {
	//			fmt.Printf("%v : err %+v", i, err)
	//		} else {
	//			fmt.Printf("%v : val %v", i, s)
	//		}
	//		wg.Done()
	//	}(i)
	//}
	//
	//wg.Wait()
	var (
		code = "1"
	)

	defer func(s time.Time) {
		fmt.Printf("finish: %v code %v", s, code)
	}(time.Now())

	code = "999"
}

var (
	once sync.Once
	s    string
)

func runTest() (string, error) {

	once.Do(func() {
		s = doInit()
	})
	if s == "" {
		return "", errors.New("s has not init yet")
	}

	return s, nil
}

func doInit() string {
	time.Sleep(1 * time.Minute)
	return "initial"
}
