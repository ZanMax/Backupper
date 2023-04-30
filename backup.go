package main

import "fmt"

func main() {
	config := getConfigs()
	fmt.Println(createBackupName())
	for i := 0; i < len(config.Files); i++ {
		exist := checkFileExists(config.Files[i])
		if exist {
			fmt.Println("File exists: ", config.Files[i])
		} else {
			fmt.Println("File does not exist: ", config.Files[i])
		}
	}
}
