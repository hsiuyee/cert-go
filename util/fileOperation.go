package util

import (
	"io/fs"
	"os"
	"path/filepath"

	logger "github.com/Alonza0314/logger-go"
)

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func FileWrite(filePath string, data []byte, code fs.FileMode) error {
	err := os.WriteFile(filePath, data, code)
	if err != nil {
		logger.Error("FileWrite", err.Error())
	}
	return err
}

func FileDelete(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		logger.Error("FileDelete", err.Error())
	}
	return err
}

func FileDir(filePath string) string {
	return filepath.Dir(filePath)
}

func FileDirCreate(filePath string) error {
	err := os.MkdirAll(filepath.Dir(filePath), 0775)
	if err != nil {
		logger.Error("FileDirCreate", err.Error())
	}
	return err
}

func FileDirExists(filePath string) bool {
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}
