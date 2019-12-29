package demoCHANNEL

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

type commResource struct {
	wg        sync.WaitGroup
	biStrChan chan string
}

func PingPongUsage() {
	var res commResource
	res.biStrChan = make(chan string)
	res.wg.Add(2)
	go ping(&res)
	go pong(&res)
	res.wg.Wait()
}

func ping(res *commResource) {

	var count = 10

	defer res.wg.Done()

	for {
		count--
		if count < 0 {
			break
		}
		res.biStrChan <- "hello"
		time.Sleep(1000 * time.Millisecond)
	}

	close(res.biStrChan)
}

func ChannelDirectionUsage() {
	var res commResource
	res.biStrChan = make(chan string)
	res.wg.Add(2)
	go sendOnlyChannel(res.biStrChan, &res.wg)
	go recvOnlyChannel(res.biStrChan, &res.wg)
	res.wg.Wait()
}

func SelectCaseUsage() {
	var res commResource
	var msg, message string
	var count int
	var toBreak bool

	fmt.Println("demoSelectCaseUsage at beginning")

	count = 0
	toBreak = false
	msg = "hello, world!"
	res.biStrChan = make(chan string)

	go sendChannelMSG(&msg, res.biStrChan)

	for {
		select {
		case message, toBreak = <-res.biStrChan:
			if toBreak {
				fmt.Println("message = ", message)
			}
		case <-time.After(time.Second):
			count++
			fmt.Println("count = ", count)
		}
		if !toBreak {
			break
		}
	}
	fmt.Println("demoSelectCaseUsage at ending")
}

func BufferedChannelUsage() {
	var res commResource
	var message string

	message = "hello, world!"
	res.biStrChan = make(chan string, 10)

	go sendChannelMSG(&message, res.biStrChan)
	recvChannelMSG(res.biStrChan)
}

func pong(res *commResource) {

	var str string
	var ok bool

	defer res.wg.Done()

	for {
		str, ok = <-res.biStrChan
		if !ok {
			break
		}

		fmt.Println("str = ", str)
	}
}

func recvOnlyChannel(ch <-chan string, wg *sync.WaitGroup) {

	var str string
	var ok bool

	defer wg.Done()

	fmt.Println("recvOnlyChannel at beginning")
	for {
		str, ok = <-ch
		if !ok {
			fmt.Println("ok is false")
			break
		}
		fmt.Println("str = ", str)
	}
	fmt.Println("recvOnlyChannel at ending")
}

func sendOnlyChannel(ch chan<- string, wg *sync.WaitGroup) {

	defer wg.Done()
	defer close(ch)

	fmt.Println("sendOnlyChannel at beginning")
	count := 10
	for {
		count--
		if count < 0 {
			break
		}
		ch <- "hello"
		time.Sleep(time.Second)
	}
	fmt.Println("sendOnlyChannel at ending")
}

func sendChannelMSG(msg *string, ch chan<- string) {

	defer close(ch)

	fmt.Println("sendChannelMSG at beginning")

	count := 10

	for {
		count--
		if count < 0 {
			break
		}

		ch <- *msg
		result, _ := rand.Int(rand.Reader, big.NewInt(1000000000))
		//fmt.Println("result = ", time.Duration(result.Uint64()))
		time.Sleep(time.Duration(result.Uint64()))
	}

	fmt.Println("sendChannelMSG at ending")
}

func recvChannelMSG(ch <-chan string) {

	var ok bool
	var message string

	fmt.Println("recvChannelMSG at beginning")
	for {
		message, ok = <-ch
		if !ok {
			fmt.Println("ok is false")
			break
		}
		fmt.Println("message = ", message)
		time.Sleep(time.Second)
	}

	fmt.Println("recvChannelMSG at ending")
}
