package deleter

import (
	"os"
	"log"
	
)
func DeleteFiles(files []string) {
	for _, file := range files {
		e := os.Remove(file)
		if e != nil { 
			log.Fatal(e) 
		}  
	}
}