package main

import (
	"github.com/ghost-circuit/word-of-wisdom/internal/application"
)

func main() {
	app := application.NewApp()
	app.Run()
}
