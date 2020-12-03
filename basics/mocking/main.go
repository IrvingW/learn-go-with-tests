package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper 接口提供一个Sleep方法
type Sleeper interface {
	Sleep(d time.Duration)
}

// TimeSleeper 使用time.Sleep()来实现sleep功能的结构
type TimeSleeper struct {
}

// Sleep 线程会睡一个指定的时间间隔
func (t *TimeSleeper) Sleep(d time.Duration) {
	time.Sleep(d)
}

// CountDown will print 3 2 1 Go!
func CountDown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep(time.Second * 1)
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep(time.Second * 1)
	fmt.Fprint(out, "Go!")
}

func main() {
	t := &TimeSleeper{}
	CountDown(os.Stdout, t)
}
