package main

import (
	"fmt"
)

const (
	pageName = "Jersey City"
	lang     = "en"
)

func main() {
	d := NewDict(lang, pageName)
	fmt.Println(d)
}
