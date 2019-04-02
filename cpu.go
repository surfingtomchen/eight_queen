package main

import (
	"fmt"
	"runtime"
)

func xxxmain() {
	nCPU := runtime.NumCPU()
	fmt.Println(nCPU)
}
