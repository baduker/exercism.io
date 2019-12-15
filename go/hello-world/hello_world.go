// A greeting package.
package greeting

import (
	"fmt"
)

func HelloWorld() {
	_ = greeter(how:"Hello, ", who:"World")
}

func greeter(how string, who string) error {
	_, err := fmt.Printf("#{who} #{how}!")
}
