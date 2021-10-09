package main

import "fmt"

type Simpler interface {
	Get() int
	Set() int
}

type Simple struct {
	a int
	b int
}

func (s Simple) Get() int {
	return s.a
}

func (s Simple) Set() int {
	return s.b
}

func showValue(asset Simpler) {
	fmt.Printf("Value of the asset is %d/%d\n", asset.Get(), asset.Set())
}

func main() {
	var show Simpler = Simple{1, 2}
	showValue(show)
}
