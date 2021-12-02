package main

import (
	"fmt"
	"imgproc/filter"
	"imgproc/task"
	"time"
)

func main() {
	var f filter.Filter = filter.Grayscale{}
	//t := task.NewWaitGrpTask("./img", "output", f)
	t := task.NewChanTask("./img", "output", f, 16)

	start := time.Now()
	err := t.Process()
	if err != nil {
		fmt.Printf("Error while processing, %v", err)
	}
	elapsed := time.Since(start)
	fmt.Printf("Image processing took %s\n", elapsed)
}
