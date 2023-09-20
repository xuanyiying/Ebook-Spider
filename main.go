package main

import (
	"ebook-spider/parser"
	"syscall"
)

func main() {
	parser.Parse()
	syscall.Exit(0)
}
