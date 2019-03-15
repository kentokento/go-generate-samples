package sample

import "time"

type Sample interface {
	Update(str string, num int, t time.Time) error
	Get(str string) int
}
