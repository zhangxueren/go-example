package main

import (
	"fmt"
	"go-example/zyb/wxgroup"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go resetRecoverTask")
		return
	}

	action := args[1]
	if action == "resetRecoverTask" {
		wxgroup.ResetRecoverTask()
	}

	if action == "updateGroupClassify" {
		wxgroup.UpdateGroupClassifyByClassify2Id(23523, 23508, 23523, 0)
	}

}
