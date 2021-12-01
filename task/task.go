package task

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Tasker interface {
	Process() error
}

type dirCtx struct {
	SrcDir string
	DstDir string
	files  []string
}

func buildFileList(srcDir string) []string {
	var files []string
	fmt.Println("Generating file list...")
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix(path, ".jpg") {
			return err
		}

		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil
	}
	return files
}
