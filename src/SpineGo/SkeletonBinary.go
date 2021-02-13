package SpineGo

import (
	"SpineGo/utlis/charArray"
	"SpineGo/utlis/strArray"
	"bufio"
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
		panic("EOFException: ReadByte Error")
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
	charCount := 0
	for index := 0; index < byteCount; {
		b, err := i.ReadByte()
		var tmp byte
		switch b >> 4 {
		case 12, 13:
			tmp, err = i.ReadByte()
			i.chars.Replace(charCount+1, rune((b&0x1F)<<6|tmp&0x3F))
			index += 2
		case 14:
			var tmp2 byte
			tmp, err = i.ReadByte()
			tmp2, err = i.ReadByte()
			i.chars.Replace(charCount+1, rune((int(b)&0x0F)<<12|(int(tmp)&0x3F)<<6|int(tmp2)&0x3F))
			index += 3
		default:
			i.chars.Replace(charCount+1, rune(b))
			index++
		}
		if err != nil {
			panic("EOFException: ReadByte Error")
		}
	}
	return i.chars.ToString(), true
}

func ReadSkeletonData(file *os.File) {

}
