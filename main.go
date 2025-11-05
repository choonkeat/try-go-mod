package main

import "github.com/choonkeat/try-go-util/pkg/logutil"

func main() {
	l := &logutil.FmtLogger{}
	l.Log("basic")
}
