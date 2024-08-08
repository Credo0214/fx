package main

import (
	"fmt"
)

// test
// test WSL add

var (
	//dp float32 = 0.01
	days int = 20
)

func main() {
	var capitall float64
	var count int
	var dp float64

	fmt.Print("初期資金\n")
	fmt.Scan(&capitall)
	fmt.Print("日利\n")
	fmt.Scan(&dp)
	fmt.Print("何カ月\n")
	fmt.Scan(&count)

	for i := 0; i < count; i++ {
		months := i + 1
		dpm := capitall * dp
		mp := dpm * (float64(days))
		capitall = capitall + mp
		fmt.Println(months, "ヶ月", int(capitall), "円", "利益", int(mp), "円")
	}
}
