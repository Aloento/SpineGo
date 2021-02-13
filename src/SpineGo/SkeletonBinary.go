package SpineGo

import (
	"bufio"
	"log"
	"os"
)

var byteSlice = make([]byte, 512)

type skeletonInput struct {
	bufio.Reader
	Strings []string
	chars   [32]rune
}

func (i *skeletonInput) skeletonInput(file *os.File) {
	i.Reader = *bufio.NewReader(file)
}

func (i skeletonInput) readInt(optimizePositive bool) (result int) {
	b, err := i.ReadByte()
	result = int(b) & 0x7F
	if (b & 0x80) != 0 {
		b, err = i.ReadByte()
		result |= int(b&0x7F) << 14
		if (b & 0x80) != 0 {
			b, err = i.ReadByte()
			result |= int(b&0x7F) << 21
			if (b & 0x80) != 0 {
				b, err = i.ReadByte()
				result |= int(b&0x7F) << 28
			}
		}
	}
	if err != nil {
		log.Println("ReadByte Error:", err)
		return -1
	} else {
		if optimizePositive {
			return
		} else {
			return int(uint(result)>>1) ^ -(result & 1)
		}
	}
}

func (i skeletonInput) readStringRef() string {
	return ""

}

func ReadSkeletonData(file *os.File) {

}
