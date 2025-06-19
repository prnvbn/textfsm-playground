//go:build js && wasm

package main

func main() {}

//export add
func add(x int, y int) int {
	return x + y
}
