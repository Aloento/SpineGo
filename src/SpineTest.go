package main

import (
	. "SpineGo"
	"os"
)

func main() {
	file, err := os.Open("spineboy.skel")
	if err != nil {
		println("读取错误", err)
	}
	ReadSkeletonData(file)
}
