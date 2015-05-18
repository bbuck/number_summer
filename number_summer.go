package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/seer-server/script-engine"
)

var (
	defaultScript = `
		function process(nums)
			sum = 0
			for i = 1, #nums do
				sum = sum + nums[i]
			end

			print("The sum of the given values was " .. sum)
		end
	`
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("USAGE: %s --process PROCESS num[, num[, num[, ...]]]\n", os.Args[0])

		return
	}

	file := flag.String("file", "", "Script file to load with process defined.")
	code := flag.String("code", "", "Lua code that contains a process function.")
	flag.Parse()

	e := lua.NewEngine()
	nums := e.NewTable()

	numStrings := flag.Args()
	for _, ns := range numStrings {
		intVal, err := strconv.Atoi(ns)
		if err != nil {
			fmt.Printf("Invalid number %q was given!\n", ns)

			return
		}
		nums.Append(intVal)
	}

	script := ""
	if len(*file) > 0 {
		if err := e.LoadFile(*file); err != nil {
			fmt.Printf("Script Error: %s\n", err)
		}
	} else if len(*code) > 0 {
		script = *code
	} else {
		script = defaultScript
	}

	if len(script) > 0 {
		if err := e.LoadString(script); err != nil {
			fmt.Printf("Script Error: %d\n", err)

			return
		}
	}

	_, err := e.Call("process", 0, nums)
	if err != nil {
		fmt.Printf("Script Error: %s\n", err)

		return
	}
}
