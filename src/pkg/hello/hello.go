package hello

import (
	"strings"

	"github.com/EnricoPicci/go-class-hello-with-with-internal/src/internal/print"
)

func Upper(greeting string) {
	gUpper := strings.ToUpper(greeting)
	print.Print(gUpper)
}
