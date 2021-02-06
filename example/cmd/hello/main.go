package main

import (
	"fmt"

	"golang.org/x/xerrors"
)

func foobar() error {
	return xerrors.New("error")
}

func main() {
	fmt.Println(foobar())
}
