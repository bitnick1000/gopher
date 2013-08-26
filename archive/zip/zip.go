package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Unzip(file string) (err error) {
	baseDir := filepath.Dir(file)

	reader, err := zip.OpenReader(file)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			os.Mkdir(filepath.Join(baseDir, file.Name), 0666)
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		dstFile, err := os.OpenFile(filepath.Join(baseDir, file.Name), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, fileReader)
		if err != nil {
			return err
		}
	}
	return nil
}
