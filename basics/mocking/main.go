package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// CountDown will print 3 2 1 Go!
func CountDown(out io.Writer) {
	for i := 3; i > 0; i-- {
		time.Sleep(time.Second * 1)
		fmt.Fprintln(out, i)
	}
	time.Sleep(time.Second * 1)
	fmt.Fprint(out, "Go!")
}

func main() {
	CountDown(os.Stdout)
}
