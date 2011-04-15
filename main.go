package main

import (
	"fmt"
	"flag"
	"os"
)

const usage = `usage: wptrans LANG NAME
where LANG is a 2- or 3-letter Wikipedia language code
and NAME is the name of a page on the LANG Wikipedia`

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	fmt.Println(NewDict(flag.Arg(0), flag.Arg(1)))
}
