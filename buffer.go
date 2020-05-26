package main

import (
	"bufio"
	"os"
	"fmt"
	"time"
)


func buffer(size int, loop int) {
	target, _ := os.Create("./target")
	writeBuff := bufio.NewWriterSize(target, size)
	
	str := "1"
	for i := 0; i < 10; i ++ {
		str += str
	}

	t1 := time.Now()
	for i := 0; i < loop; i++ {
		writeBuff.WriteString(str)
	}
	elapsed := time.Since(t1)
    fmt.Println("App elapsed: ", elapsed)
}

func main()  {
	buffer(1024000, 1000);	// 386.949µs
	buffer(10240, 1000);	// 783.338µs
	buffer(1024, 1000);		// 2.945506ms

	buffer(1024000, 10000);	// 4.425895ms
	buffer(10240, 10000);	// 6.974164ms
	buffer(1024, 10000);		// 37.774934ms

	buffer(1024000, 100000);	// 34.711955ms
	buffer(10240, 100000);	// 104.748874ms
	buffer(1024, 100000);		// 589.203223ms
}