package main

import (
	"fmt"

	"github.com/smoreg/simpleCash/simpleCash"
)

func main() {
	str := SomeName()
	fmt.Printf("%v", str)
	str = SomeDefer()
	fmt.Printf("%v", str)

}

func SomeName() (ex CoolEX) {
	cash := simpleCash.Cash{}
	exist := cash.GetCashStruct(0, &ex)
	if exist {
		return ex
	}
	ex = CoolEX{
		X: 1,
		Y: "3",
	}
	cash.SetCash(0, ex, false)
	return ex
}
func SomeDefer() (ex CoolEX) {
	cash := simpleCash.Cash{}
	exist := cash.GetCashStruct(0, &ex)
	if exist {
		return ex
	}
	defer func() { cash.SetCash(0, ex, true) }()
	ex = CoolEX{
		X: 1,
		Y: "3",
	}

	return ex
}

type CoolEX struct {
	X int
	Y string
}
