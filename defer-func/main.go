package main

import "fmt"

func main() {
	basicDefer()
	deferStack()
	deferInCondition()
	deferInSwitch()
	deferInLoop()
		fmt.Println("callback in blacklist https://sandboxqc-acnotify.zpapps.vn/v2/callback\ngit.zapa.cloud/acquiring-platform/callback-proxy/business.validateCallbackURL\n\t/home/lap11757/Workspace/vng/MEP/acquiring-platform/callback-proxy/business/merchant.go:241\ngit.zapa.cloud/acquiring-platform/callback-proxy/business.callbackToMerchant\n\t/home/lap11757/Workspace/vng/MEP/acquiring-platform/callback-proxy/business/merchant.go:141\ngit.zapa.cloud/acquiring-platform/callback-proxy/business.forwardOrderCallback\n\t/home/lap11757/Workspace/vng/MEP/acquiring-platform/callback-proxy/business/merchant.go:64\ngit.zapa.cloud/acquiring-platform/callback-proxy/business.handleForwardMerchant\n\t/home/lap11757/Workspace/vng/MEP/acquiring-platform/callback-proxy/business/merchant.go:37\ngit.zapa.cloud/acquiring-platform/callback-proxy/business.(*Kafka).handleRetry\n\t/home/lap11757/Workspace/vng/MEP/acquiring-platform/callback-proxy/business/kafka.go:119\ngit.zapa.cloud/acquiring-platform/callback-proxy/database.(*Broker).ListenMessage\n\t/home/lap11757/Workspace/vng/MEP/acquiring-platform/callback-proxy/database/kafka.go:85\nruntime.goexit\n\t/home/lap11757/.gvm/gos/go1.13/src/runtime/asm_amd64.s:1357\nfail: callback URL invalid (https://sandboxqc-acnotify.zpapps.vn/v2/callback)\ngit.zapa.cloud/acquiring-platform/callback-proxy/business.callbackToMerchant\n\t/home/lap11757/Workspace/vng/MEP/acquiring-platform/callback-proxy/business/merchant.go:145\ngit.zapa.cloud/acquiring-platform/callback-proxy/business.forwardOrderCallback\n\t/home/lap11757/Workspace/vng/MEP/acquiring-platform/callback-proxy/business/merchant.go:64\ngit.zapa.cloud/acquiring-platform/callback-proxy/business.handleForwardMerchant\n\t/home/lap11757/Workspace/vng/MEP/acquiring-platform/callback-proxy/business/merchant.go:37\ngit.zapa.cloud/acquiring-platform/callback-proxy/business.(*Kafka).handleRetry\n\t/home/lap11757/Workspace/vng/MEP/acquiring-platform/callback-proxy/business/kafka.go:119\ngit.zapa.cloud/acquiring-platform/callback-proxy/database.(*Broker).ListenMessage\n\t/home/lap11757/Workspace/vng/MEP/acquiring-platform/callback-proxy/database/kafka.go:85\nruntime.goexit\n\t/home/lap11757/.gvm/gos/go1.13/src/runtime/asm_amd64.s:1357\nfail: CallbackToMerchant\ngit.zapa.cloud/acquiring-platform/callback-proxy/business.forwardOrderCallback\n\t/home/lap11757/Workspace/vng/MEP/acquiring-platform/callback-proxy/business/merchant.go:66\ngit.zapa.cloud/acquiring-platform/callback-proxy/business.handleForwardMerchant\n\t/home/lap11757/Workspace/vng/MEP/acquiring-platform/callback-proxy/business/merchant.go:37\ngit.zapa.cloud/acquiring-platform/callback-proxy/business.(*Kafka).handleRetry\n\t/home/lap11757/Workspace/vng/MEP/acquiring-platform/callback-proxy/business/kafka.go:119\ngit.zapa.cloud/acquiring-platform/callback-proxy/database.(*Broker).ListenMessage\n\t/home/lap11757/Workspace/vng/MEP/acquiring-platform/callback-proxy/database/kafka.go:85\nruntime.goexit\n\t/home/lap11757/.gvm/gos/go1.13/src/runtime/asm_amd64.s:1357")
	m:= map[int]bool{1:true,2:true}
	if m[3]{
		fmt.Println("error")
	}else{
		fmt.Println("ok")
	}

	if m[2]{
		fmt.Println("ok")
	}else{
		fmt.Println("error")
	}
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
