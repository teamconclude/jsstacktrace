# jsstacktrace

jsstacktrace converts stack traces generated inside the browser using locally available map files

This makes it easy to convert log data containing stack traces of minimized/packaged JS to the contain
the correct line number and source file information on your server.

## Usage

```Go
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
```


