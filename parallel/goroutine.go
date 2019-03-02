package parallel

import (
	"fmt"
	"runtime"
	"sync"
)

// CallGoroutine ゴールーチンの分岐と合流のサンプル
func CallGoroutine() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Printf("Hello!\n")
	}
	wg.Add(1)
	go sayHello()
	wg.Wait()
}

// GoClosure クロージャ―の挙動確認
// welcomeが出力される（つまり元の変数に対して操作する）
func GoClosure() {
	var wg sync.WaitGroup
	solution := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		solution = "welcome"
	}()
	wg.Wait()
	fmt.Println(solution)
}

// GoClosureRangeBad クロージャ―の挙動確認（配列版の悪い例）
// メインの実行がゴールーチンが開始する前に終わり、ガベージコレクションの挙動によって、
// 最後の値（この場合good day）を参照した挙動を行う
func GoClosureRangeBad() {
	var wg sync.WaitGroup
	for _, solution := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(solution)
		}()
	}
	wg.Wait()
}

// GoClosureRangeGood クロージャ―の挙動確認（配列版の良い例）
// ゴールーチンに対して値を渡すことで後に残った変数を参照しない
func GoClosureRangeGood() {
	var wg sync.WaitGroup
	for _, solution := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(solution string) {
			defer wg.Done()
			fmt.Println(solution)
		}(solution)
	}
	wg.Wait()
}

// GoroutineMemory ゴールーチンのメモリ消費を計測する
func GoroutineMemory() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}
	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c }

	const numGoroutine = 1e4
	wg.Add(numGoroutine)
	before := memConsumed()
	for i := numGoroutine; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb\n", float64(after-before)/numGoroutine/1000)
}
