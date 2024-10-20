package main

import (
	"github.com/alisher-baizhumanov/word-of-wisdom/internal/application"
)

func main() {
	app := application.NewApp()
	app.Run()
}
