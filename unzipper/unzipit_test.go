package unzipper_test

import (
	"zipextractor/unzipper"
	"testing"
	"zipextractor/globber"
)

func TestUnzipfile(* testing.T) {
	files := globber.Globzipfiles()
	unzipper.Unzipfiles(files)
}