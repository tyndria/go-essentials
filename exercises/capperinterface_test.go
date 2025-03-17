package exercises

import (
	"fmt"
	"os"
)

func ExampleCapper() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "hello there")
    // Output: HELLO THERE
}