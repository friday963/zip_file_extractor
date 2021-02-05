package main

import (
	"zipextractor/deleter"
	"zipextractor/globber"
	"zipextractor/unzipper"
)

func main() {
	files := globber.Globzipfiles()
	unzipper.Unzipfiles(files)
	deleter.DeleteFiles(files)
}