package main

import (
	"fmt"

	"github.com/traesk/kafft/util"
)

func main() {
	//cmd.Execute()
	err := util.ZipFiles("util")
	if err != nil {
		fmt.Println(err)
	}
}
