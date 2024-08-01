package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/teamconclude/jsstacktrace"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: simple <baseDir>")
		os.Exit(1)
	}
	jsMap := jsstacktrace.NewJSMap(os.Args[1])

	reader := bufio.NewReader(os.Stdin)
	data, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	mappedStack := jsMap.ConvertStackTraceString(string(data))

	fmt.Print(mappedStack)
}
