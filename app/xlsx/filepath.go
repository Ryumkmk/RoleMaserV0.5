package xlsx

import (
	"log"
	"os"
	"path/filepath"
)

func GetXlsxPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	absPath := filepath.Join(dir, "app/xlsx")
	// files, err := os.ReadDir(absPath)
	// fmt.Printf(files[0].Name())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return absPath
}
