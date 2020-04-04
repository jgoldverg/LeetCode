package main

import (
	"lc"
	"os"
)
import ()

func main() {
	if bo := lc.IsValid("{}"); bo {
		os.exit(0)
	} else {
		os.exit(1)
	}
}
