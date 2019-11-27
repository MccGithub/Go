package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)
import "fmt"

// 生产者: 生成 factor 整数倍的序列
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		time.Sleep(time.Millisecond * 300)
		out <- i * factor
	}
}

// 消费者
func Consumer(in <-chan int) {
	for v := range in {
		time.Sleep(time.Millisecond * 300)
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 64) // 成果队列

	go Producer(3, ch) // 生成 3 的倍数的序列
	go Producer(5, ch) // 生成 5 的倍数的序列
	go Consumer(ch)    // 消费生成的序列

	// 运行一定时间后退出.
	//time.Sleep(5 * time.Second)

	// Ctrl+C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}

// 用Ctrl+C可以终止程序运行.
// 但是通过kill可能只能终止主进程的运行, 子进程会被init接管导致程序一直输出无法终止.
// 需要通过ps -eLf | grep go找到在运行的子进程并kill掉.
// 本次ps查到的进程为:
// m         3638     1  3638  0   10 22:40 pts/0    00:00:00 /tmp/go-build459633756/b001/exe/main
