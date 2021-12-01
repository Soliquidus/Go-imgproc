package main

import (
	"imgproc/filter"
	"imgproc/task"
)

func main() {
	var f filter.Filter = filter.Grayscale{}
	t := task.NewWaitGrpTask("./img", "output", f)
	err := t.Process()
	if err != nil {
		return
	}

}
