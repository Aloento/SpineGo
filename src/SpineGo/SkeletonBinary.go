package SpineGo

import (
	"SpineGo/utlis/charArray"
	"SpineGo/utlis/strArray"
	"bufio"
	"log"
	"os"
)

type skeletonInput struct {
	bufio.Reader
	Strings strArray.StrArray
	chars   charArray.CharArray
}

func (i *skeletonInput) skeletonInput(file *os.File) {
	i.Reader = *bufio.NewReader(file)
}

// Reads a 1-5 byte int.
func (i *skeletonInput) readInt(optimizePositive bool) (result int) {
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

func (i *skeletonInput) readStringRef() (string, bool) {
	index := i.readInt(true)
	if index == 0 {
		return "", false
	} else {
		return i.Strings.Get(index - 1), true
	}
}

func (i *skeletonInput) readString() (string, bool) {
	byteCount := i.readInt(true)
	switch byteCount {
	case 0:
		return "", false
	case 1:
		return "", true
	}
	byteCount--
	if i.chars.Len() < byteCount {
		i.chars = *charArray.Make(0, byteCount)
	}
}

func ReadSkeletonData(file *os.File) {

}
