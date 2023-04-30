package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Configs struct {
	Files []string `json:"files"`
	DBs   []DB     `json:"dbs"`
}

type DB struct {
	Type       string   `json:"type"`
	ConnString string   `json:"connString"`
	DbNames    []string `json:"dbNames"`
}

func getConfigs() Configs {
	// Read configs
	config := Configs{}
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func(configFile *os.File) {
		errConfigClose := configFile.Close()
		if errConfigClose != nil {
			fmt.Println("Error while closing config file: ", err)
		}
	}(configFile)
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		log.Fatal(err)
	}
	return config
}

func createArchive() {
	config := getConfigs()
	for i := 0; i < len(config.Files); i++ {
		fmt.Println(config.Files[i])
	}
}

func CreateArchiveFromFiles(archiveName string, files []string) error {
	archiveFile, err := os.Create(archiveName)
	if err != nil {
		return err
	}
	defer archiveFile.Close()

	archive := zip.NewWriter(archiveFile)
	defer archive.Close()

	for _, fileName := range files {
		err = addFileToArchive(archive, fileName)
		if err != nil {
			return err
		}
	}

	return nil
}

func addFileToArchive(archive *zip.Writer, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}
	header.Name = filepath.Base(fileName)
	header.Method = zip.Deflate

	writer, err := archive.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	return err
}

func checkFileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err == nil {
		return true
	}
	return false
}

func getCurrentDateTime() string {
	currentTime := time.Now()
	return currentTime.Format("20060102150405")
}

func createBackupName() string {
	return "backup_" + getCurrentDateTime() + ".zip"
}
