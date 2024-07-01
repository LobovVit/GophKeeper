package utils

import (
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func CreateStorageUser(dirPath string, id int64) error {
	userID := strconv.Itoa(int(id))
	path := filepath.Join(dirPath, userID)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func CreateStorageNotExistsUser(dirPath string, id int64) error {
	userID := strconv.Itoa(int(id))
	path := filepath.Join(dirPath, userID)
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func UploadFile(dirPath string, id int64, name string, data []byte) error {
	userID := strconv.Itoa(int(id))
	path := filepath.Join(dirPath, userID, "/", name)
	// Write data to file
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func DownloadFile(dirPath string, id int64, name string) ([]byte, error) {
	userID := strconv.Itoa(int(id))
	path := filepath.Join(dirPath, userID, "/", name)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func RemoveFile(dirPath string, id int64, name string) error {
	userID := strconv.Itoa(int(id))
	path := filepath.Join(dirPath, userID, "/", name)
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
