package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func main() {
	result := testing.Benchmark(func(b *testing.B) { run() })
	fmt.Println(result.T)
}

// WaitGroupを利用した並列処理
func run() {
	// 並列処理数
	const processCnt = 5

	// WaitGroupを作成する
	wg := new(sync.WaitGroup)

	fmt.Println("Start!")
	for i := 0; i < processCnt; i++ {
		wg.Add(1)
		go process(i, wg)
	}

	// 同期待ち
	wg.Wait()

	fmt.Println("Finish!")
}

func process(name int, wg *sync.WaitGroup) {
	// 処理終了時、wgの数を1つ減らす
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println(name)
}
