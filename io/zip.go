package io

import (
	"archive/zip"
	"github.com/macococo/go-webbase/utils"
	"io"
	"os"
	"path/filepath"
)

func Unzip(src string, dest string) []string {
	files := []string{}

	r, err := zip.OpenReader(src)
	utils.HandleError(err)

	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		utils.HandleError(err)

		defer rc.Close()

		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			utils.HandleError(err)

			defer f.Close()

			_, err = io.Copy(f, rc)
			utils.HandleError(err)

			files = append(files, path)
		}
	}

	return files
}
