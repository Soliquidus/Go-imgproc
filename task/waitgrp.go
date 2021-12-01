package task

import (
	"fmt"
	"imgproc/filter"
	"path"
	"path/filepath"
	"sync"
)

type WaitGrpTask struct {
	dirCtx
	Filter filter.Filter
}

func NewWaitGrpTask(srcDir, dstDir string, filter filter.Filter) Tasker {
	return &WaitGrpTask{
		Filter: filter,
		dirCtx: dirCtx{
			SrcDir: srcDir,
			DstDir: dstDir,
			files:  buildFileList(srcDir),
		},
	}
}

func (w *WaitGrpTask) Process() error {
	var wg sync.WaitGroup
	size := len(w.files)
	for i, f := range w.files {
		filename := filepath.Base(f)
		dst := path.Join(w.DstDir, filename)
		wg.Add(1)

		go w.applyFilter(f, dst, &wg, i+1, size)
	}

	wg.Wait()
	fmt.Println("Done processing files!")
	return nil
}

func (w *WaitGrpTask) applyFilter(src, dst string, wg *sync.WaitGroup, i, size int) {
	err := w.Filter.Process(src, dst)
	if err != nil {
		fmt.Printf("Error while Process(), %v", err)
	}
	fmt.Printf("Processed [%d/%d] %v => %v\n", i, size, src, dst)
	wg.Done()
}
