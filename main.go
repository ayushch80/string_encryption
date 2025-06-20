package main

import (
	"fmt"

	"string_encryption/utils"
)

func main() {
	folderPath := "/Users/brluser/Desktop/sampleApp"
	prodBundle := utils.GetProjectBundle(folderPath)

	fmt.Println(prodBundle)
}
