package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Printf("Operating System: %s\nArchitecture: %s\nNumber of CPUs: %d\nVersion: %s\n", runtime.GOOS, runtime.GOARCH, runtime.NumCPU(), runtime.Version())
}
