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
		s := utils.FindStrings(file.Code)

		for i := range len(s) {
			s[i].Encrypted = s[i].Encrypt()
			fmt.Println(s[i].String)
		}

		fmt.Println(file.Name, ":", s)
	}
}
