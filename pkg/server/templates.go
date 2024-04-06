package server

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var templates = make([]string, 0)

func LoadTemplates(patterns ...string) {
	for _, pattern := range patterns {
		err := filepath.Walk(pattern,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() && strings.HasSuffix(info.Name(), ".html") {
					templates = append(templates, path)
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
	}
}

func GetTemplates() []string {
	return templates
}
