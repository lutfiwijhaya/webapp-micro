package utils

import (
	"mime/multipart"
	"os"
	"path/filepath"
)

func SaveUploadedFile(file *multipart.FileHeader, destDir string) (string, error) {
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return "", err
	}

	filePath := filepath.Join(destDir, file.Filename)
	if err := saveFile(file, filePath); err != nil {
		return "", err
	}
	return filePath, nil
}

func saveFile(file *multipart.FileHeader, filePath string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = dst.ReadFrom(src)
	return err
}
