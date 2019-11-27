package main

import "fmt"

// 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 管道过滤器: 删除能被素数整除的数
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	ch := GenerateNatural() // 自然数序列: 2, 3, 4, ...
	prime := 0
	for i := 0; i < 1000000; i++ {
		prime = <-ch // 新出现的素数
		//fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime)
	}
	fmt.Printf("%v\n", prime)
}

// 解释:
// 首先第一个prime是素数(2)
// 经过PrimeFilter过滤后返回的第一个不能被上一个prime整除的数必然是素数
// 不断过滤取第一个数就能得到所有素数.
