package main

import (
	"github.com/pr0head/micro-old-server/internal"
)

func main() {
	app := internal.NewApplication()
	app.Run()
}
