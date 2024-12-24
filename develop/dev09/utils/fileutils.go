package utils

import (
	"io/ioutil"
	"path/filepath"
)

func SaveHTML(dirName, html string) error {
	filePath := filepath.Join(dirName, "index.html")
	return ioutil.WriteFile(filePath, []byte(html), 0644)
}
