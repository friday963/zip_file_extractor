package unzipper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/artdarek/go-unzip"
)

func Unzipfiles(files []string) {
	path, _ := os.Getwd()
	for _, file := range files {
		folderArray := strings.Split(file, ".")
		folder := folderArray[0]
		err := os.Mkdir(folder, 0755)
		if err != nil {
			fmt.Println(err)
		}
		fullfolderpath := filepath.Join(path, folder)
		uz := unzip.New(file, fullfolderpath)
		err = uz.Extract()
		if err != nil {
			fmt.Println(err)
		}
	}

}
