package deleter_test

import (
	"zipextractor/deleter"	
	"zipextractor/globber"
	"testing"
	
)


func TestDeleter(*testing.T) {
	files := globber.Globzipfiles()
	deleter.DeleteFiles(files)
}