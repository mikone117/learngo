package printer

import (
	"fmt"
	"runtime"
)

//Hello is an exported function
func Hello() {
	fmt.Println("Exported hello")
}

//Version returns the installed Go Version
func Version() string {
	return runtime.Version()
}
