package main

import (
	"fmt"
	ewallet "unit-test/appEwallet"
)

func main() {
	fmt.Println(ewallet.JalankanPerintah([]string{"deposit", "deposit", "withdraw"}))
}
