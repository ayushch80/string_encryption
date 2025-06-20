package main

import (
	"fmt"
	"strings"
	"path/filepath"

	"string_encryption/utils"
)

func main() {
	folderPath := "/Users/brluser/Desktop/sampleApp"
	prodBundle := utils.GetProjectBundle(folderPath)

	projName := strings.Split(prodBundle, ".")[len(strings.Split(prodBundle, "."))-1]

	files := utils.ReadSwiftFiles(filepath.Join(folderPath, projName))
	fmt.Println(files)
}
