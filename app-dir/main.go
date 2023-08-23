package main

import "github.com/jakegee/toolkit"

func main() {
	var tools toolkit.Tools

	tools.CreateDirIfNotExists("./test-dir/")
}
