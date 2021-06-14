package logger

import "fmt"

func Log(a ...interface{}) {
	fmt.Fprintln(logger.out, a)
}

func Err(a ...interface{}) {
	fmt.Fprintln(logger.out, "[Error]", a)
}

func Info(a ...interface{}) {
	fmt.Fprintln(logger.out, "[Info]", a)
}
