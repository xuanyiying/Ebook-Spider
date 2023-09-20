package main

import (
	"ebook-spider/parser"
	"syscall"
)

func main() {
	parser.Run()
	syscall.Exit(0)
}
