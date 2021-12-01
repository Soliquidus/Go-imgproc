package main

import (
	"fmt"
	"imgproc/task"
)

func main() {
	f := task.BuildFileList("./img")

	fmt.Println(f)
}
