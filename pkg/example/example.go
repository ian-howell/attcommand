package example

import (
	"fmt"
	"io"
)

func Fprint(out io.Writer, msg string) {
	fmt.Println(msg)
}
