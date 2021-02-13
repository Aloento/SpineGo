package utlis

import (
	"SpineGo/utlis/charArray"
	"SpineGo/utlis/strArray"
	"bufio"
	"bytes"
	"encoding/binary"
	"os"
)

type SkeletonInput struct {
	bufio.Reader
	Strings strArray.StrArray
	chars   charArray.CharArray
}

func (i *SkeletonInput) skeletonInput(file *os.File) {
	i.Reader = *bufio.NewReader(file)
}

// Reads a 1-5 byte int.
func (i *SkeletonInput) readInt(optimizePositive bool) (result int) {
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

func (i *SkeletonInput) readBool() bool {
	ch, err := i.ReadByte()
	if err != nil {
		panic("EOFException: ReadByte Error")
	}
	return ch != 0
}

func (i *SkeletonInput) readFloat() float32 {
	buffer := make([]byte, 4)
	ch1, err := i.ReadByte()
	ch2, err := i.ReadByte()
	ch3, err := i.ReadByte()
	ch4, err := i.ReadByte()
	if err != nil {
		panic("EOFException: ReadByte Error")
	}
	buffer[0] = ch1
	buffer[1] = ch2
	buffer[2] = ch3
	buffer[3] = ch4

	var float float32
	if binary.Read(bytes.NewReader(buffer), binary.LittleEndian, &float) != nil {
		println("binary.Read failed:", err)
	}
	return float
}

func (i *SkeletonInput) readStringRef() (string, bool) {
	index := i.readInt(true)
	if index == 0 {
		return "", false
	} else {
		return i.Strings.Get(index - 1), true
	}
}

func (i *SkeletonInput) readString() (string, bool) {
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
