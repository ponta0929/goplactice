package parallel

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

// SampleChannel 簡単なチャネルのサンプル
func SampleChannel() {
	stringCh := make(chan string) //バッファ宣言をしないので0と一緒
	go func() {
		stringCh <- "sample"
	}()
	fmt.Println(<-stringCh)
}

// DeclarationChannel チャネルの初期化サンプル
func DeclarationChannel() {
	//空インタフェース型のチャネル宣言
	var myCh chan interface{}
	myCh = make(chan interface{}, 10)
	myCh <- "first"

	var recvCh <-chan interface{}
	//recvCh = make(<-chan interface{})
	recvCh = myCh

	go func() {
		var sendCh chan<- interface{}
		//sendCh = make(chan<- interface{})
		sendCh = myCh
		sendCh <- "second"
		close(sendCh)
	}()

	/*for {
		str, ok := <-recvCh
		if ok == false {
			break
		}
		fmt.Println(str)
	}*/
	//↑のコメントアウトしたforと同じ
	for str := range recvCh {
		fmt.Println(str)
	}
}

// ChannelCloseSample closeにより複数のゴールーチンを同時に走らせる
func ChannelCloseSample() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}
	fmt.Println("Unblocking goroutines...")
	time.Sleep(2 * time.Second)
	close(begin)
	wg.Wait()
}

// SampleBafferChannel バッファつきチャネルの挙動確認
func SampleBafferChannel() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()
	for i := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", i)
	}
}
