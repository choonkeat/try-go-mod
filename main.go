package main

import "github.com/choonkeat/try-go-util/pkg/logutil/logrusutil"

func main() {
	l := &logrusutil.LogrusLogger{}
	l.Log("logrus")
}
