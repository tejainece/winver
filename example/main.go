package main

import (
	"fmt"

	"github.com/tejainece/winver"
)

func main() {
	ver := winver.GetVersion()
	fmt.Println(ver)
	fmt.Println(ver.IsWindows7())
	fmt.Println(ver.IsVista())
}
