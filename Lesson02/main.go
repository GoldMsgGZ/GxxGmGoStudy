package main

import (
	_ "Lesson02/boot"
	_ "Lesson02/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
