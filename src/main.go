package main

import (
	"fmt"
	"php"
)

func main() {
	fmt.Println(php.Date("Y-m-d H:i:s", php.Time()))

	php.ShowError()
}

