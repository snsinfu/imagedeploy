package imagedeploy

import (
	"fmt"
	"os"
)

func trace(msg interface{}) {
	if s, ok := msg.(string); ok {
		if s == "" {
			fmt.Fprintln(os.Stderr)
		} else {
			fmt.Fprintln(os.Stderr, "*", s)
		}
		return
	}

	if err, ok := msg.(error); ok {
		fmt.Fprintln(os.Stderr, "! error:", err)
		return
	}

	fmt.Fprintln(os.Stderr, ">", msg)
}
