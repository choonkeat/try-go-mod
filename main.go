package main

import (
	"fmt"

	"github.com/choonkeat/try-go-util/pkg/slicesutil"
)

func main() {
	fmt.Printf("%#v\n", slicesutil.Map([]float64{
		1, 2.3, 4.56,
	}, func(f float64) string {
		return fmt.Sprintf("%0.5f", f)
	}))
}
