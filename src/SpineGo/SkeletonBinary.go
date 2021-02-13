package SpineGo

import (
	"bufio"
	"os"
)

type skeletonInput struct {
	bufio.Reader
	Strings []string
	chars   []rune
}

func (input *skeletonInput) skeletonInput(file *os.File) {
	input.Reader = *bufio.NewReader(file)

}

func ReadSkeletonData(file *os.File) {

}
