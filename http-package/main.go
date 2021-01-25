package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

const (
	url = "http://localhost:8080/ping"
	//url = "https://api.ipify.org?format=json"
)

func main() {
	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	cli := &http.Client{
		Transport: tr,
		//Timeout:   10 * time.Second,
	}
	reader := bytes.NewReader([]byte("{\"resp_time\":20}"))
	resp, err := cli.Post("http://localhost:8080/delay", "application/json", reader)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if _, err := ioutil.ReadAll(resp.Body); err != nil {
		panic(err)
	}

	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt)
	//
	//if err := errPkg.BasicSendPostWithBody(url); err != nil{
	//	//fmt.Printf("%+v",err)
	//	s:=fmt.Sprintf("%+v",err)
	//	fmt.Printf("%v",s)
	//}
	//<-c

	//testCloseResponseImmediately()
	//// connection release
	//<-c
	//testCloseAfterRead()
	//// connection keep
	//<-c
	//testReadResponseAfterClose()
	//// connection keep
	//<-c
	//testCloseWhenGoOutOfFunc()
	//// connection keep
	//<-c
	//testNotClose()
	//// connection keep
	//<-c
}

func testReadResponseAfterClose() {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		fmt.Println("testReadResponseAfterClose", err)
		return
	}

	fmt.Println("testReadResponseAfterClose", resp.Status)
	fmt.Println("testReadResponseAfterClose", resp.Header)
	fmt.Println("testReadResponseAfterClose", string(body))
}

func testCloseWhenGoOutOfFunc() {
	resp, err := http.Get(url)
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
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	_ = resp.Body.Close()
	if resp.Body != nil {
		fmt.Println("CloseAfterRead", "body not nil after close")
	} else {
		fmt.Println("CloseAfterRead", "body nil after close")
	}
	if resp != nil {
		fmt.Println("CloseAfterRead", "response not nil after close")
	} else {
		fmt.Println("CloseAfterRead", "response nil after close")
	}

	if err != nil {
		fmt.Println("CloseAfterRead", err)
		return
	}

	fmt.Println("CloseAfterRead", string(body))
}

func testNotClose() {
	resp, err := http.Get(url)
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

	resp, err := http.Get(url)
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
