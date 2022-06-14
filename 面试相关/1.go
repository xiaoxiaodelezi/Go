// 问题描述使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母
// 最终效果如下：
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

package main

import (
	"fmt"
	"sync"
)

func main() {

	wait := sync.WaitGroup{}

	chn, chs := make(chan bool), make(chan bool)

	//print s
	go func() {

		s := 'A'

		for {
			if s >= 'Z' {
				wait.Done()
				return
			}
			select {
			case <-chs:

				{
					fmt.Print(string(s))
					s++
					fmt.Print(string(s))
					s++
					chn <- true
				}

			}
		}

	}()

	//print n
	wait.Add(1)
	go func() {
		n := 1

		for {
			select {
			case <-chn:
				{
					fmt.Print(n)
					n++
					fmt.Print(n)
					n++
					chs <- true
				}
			}
		}

	}()

	chn <- true

	wait.Wait()

}
