package main

import (
	"fmt"

	"string_encryption/utils"
)

func main() {
	folderPath := "/Users/brluser/Desktop/sampleApp"
	prodBundle := utils.GetProjectBundle(folderPath)

	projName := strings.Split(prodBundle, ".")[len(strings.Split(prodBundle, "."))-1]

	fmt.Println(projName)
}
