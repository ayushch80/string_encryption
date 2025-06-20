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

	for _, file := range files {
		_strings := utils.FindStrings(file.Code)
		fmt.Println(file.Name, ":", _strings)
	}
}
