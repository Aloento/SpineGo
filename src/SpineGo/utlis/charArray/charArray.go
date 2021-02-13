package charArray

import (
	"strconv"
	"sync"
	"unicode/utf8"
)

type CharArray struct {
	array []rune
	len   int
	cap   int
	lock  sync.Mutex
}

func Make(len, cap int) *CharArray {
	s := new(CharArray)
	if len > cap {
		panic("len large than cap")
	}
	array := make([]rune, cap, cap)

	s.array = array
	s.cap = cap
	s.len = 0
	return s
}

func (a *CharArray) Append(element rune) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.len == a.cap {
		newCap := 2 * a.len

		if a.cap == 0 {
			newCap = 1
		}
		newArray := make([]rune, newCap, newCap)

		for k, v := range a.array {
			newArray[k] = v
		}
		a.array = newArray
		a.cap = newCap
	}
	a.array[a.len] = element
	a.len = a.len + 1
}

func (a *CharArray) AppendMany(element ...rune) {
	for _, v := range element {
		a.Append(v)
	}
}

func (a *CharArray) Replace(index int, element rune) {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.Get(index)
	a.array[index] = element
}

func (a *CharArray) Get(index int) rune {
	if a.len == 0 || index >= a.len {
		panic("Index over len: Request " + strconv.Itoa(index) + " But only " + strconv.Itoa(a.len))
	}
	return a.array[index]
}

func (a *CharArray) Len() int {
	return a.len
}

func (a *CharArray) Cap() int {
	return a.cap
}

func (a *CharArray) ToString() string {
	var p []byte
	buffer := make([]byte, 3)
	for _, r := range a.array {
		n := utf8.EncodeRune(buffer, r)
		p = append(p, buffer[:n]...)
	}
	return string(p)
}
