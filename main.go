package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Printf("Operating System: %s\nArchitecture: %s\nNumber of CPUs: %d\n", runtime.GOOS, runtime.GOARCH, runtime.NumCPU())
}
