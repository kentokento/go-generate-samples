package sample

import "time"

//go:generate go get github.com/cweill/gotests
//go:generate gotests -all -i -w ./
//go:generate go test ./

func Sample(str string, num int, t time.Time) bool {
	return true
}
