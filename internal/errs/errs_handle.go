package errs

import (
	"fmt"
	"log"
)

func HandleFatal(err error, text string) {
	if err != nil {
		log.Fatal(fmt.Errorf("%s: %v\n", text, err))
	}
}
