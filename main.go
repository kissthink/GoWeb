package main

import (
	"fmt"
	"./web"
)

func main() {

	fmt.Println(web.Date("Y-m-d H:i:s", web.Time()))

	web.ShowError()
}

