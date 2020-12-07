package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	testCloseResponseImmediately()

	testCloseAfterRead()

	testCloseWhenGoOutOfFunc()

	testNotClose()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func testCloseWhenGoOutOfFunc() {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close() //not ok

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("CloseWhenGoOutOfFunc", err)
		return
	}

	fmt.Println("CloseWhenGoOutOfFunc", string(body))
}

func testCloseAfterRead() {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("CloseAfterRead", err)
		return
	}

	fmt.Println("CloseAfterRead", string(body))
}

func testNotClose() {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("NotClose", err)
		return
	}

	fmt.Println("NotClose", string(body))
}

func testCloseResponseImmediately() {

	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		fmt.Println(err)
		return
	}
	resp.Body.Close() //not ok

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("CloseResponseImmediately", err)
		return
	}

	fmt.Println("CloseResponseImmediately", string(body))
}
