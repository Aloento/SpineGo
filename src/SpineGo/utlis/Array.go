package utlis

import (
	"strconv"
	"sync"
	"unicode/utf8"
)

type Array struct {
	array []interface{}
	len   int
	cap   int
	lock  sync.Mutex
}

func NewArray(len, cap int) *Array {
	s := new(Array)
	if len > cap {
		panic("len large than cap")
	}
	array := make([]interface{}, cap, cap)

	s.array = array
	s.cap = cap
	s.len = 0
	return s
}

func (a *Array) Append(element interface{}) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.len == a.cap {
		newCap := 2 * a.len

		if a.cap == 0 {
			newCap = 1
		}
		newArray := make([]interface{}, newCap, newCap)

		for k, v := range a.array {
			newArray[k] = v
		}
		a.array = newArray
		a.cap = newCap
	}
	a.array[a.len] = element
	a.len = a.len + 1
}

func (a *Array) AppendMany(element ...interface{}) {
	for _, v := range element {
		a.Append(v)
	}
}

func (a *Array) Replace(index int, element interface{}) {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.Get(index)
	a.array[index] = element
}

func (a *Array) Get(index int) interface{} {
	if a.len == 0 || index >= a.len {
		panic("Index over len: Request " + strconv.Itoa(index) + " But only " + strconv.Itoa(a.len))
	}
	return a.array[index]
}

func (a *Array) Len() int {
	return a.len
}

func (a *Array) Cap() int {
	return a.cap
}

func (a *Array) byteToStr() string {
	var p []byte
	buffer := make([]byte, 3)
	for _, r := range a.array {
		n := utf8.EncodeRune(buffer, r.(rune))
		p = append(p, buffer[:n]...)
	}
	return string(p)
}
