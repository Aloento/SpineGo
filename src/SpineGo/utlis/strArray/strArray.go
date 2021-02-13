package strArray

import (
	"fmt"
	"strconv"
	"sync"
)

type StrArray struct {
	array []string
	len   int
	cap   int
	lock  sync.Mutex
}

func Make(len, cap int) *StrArray {
	s := new(StrArray)
	if len > cap {
		panic("len large than cap")
	}
	array := make([]string, cap, cap)

	s.array = array
	s.cap = cap
	s.len = 0
	return s
}

func (a *StrArray) Append(element string) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.len == a.cap {
		newCap := 2 * a.len

		if a.cap == 0 {
			newCap = 1
		}
		newArray := make([]string, newCap, newCap)

		for k, v := range a.array {
			newArray[k] = v
		}
		a.array = newArray
		a.cap = newCap
	}
	a.array[a.len] = element
	a.len = a.len + 1
}

func (a *StrArray) AppendMany(element ...string) {
	for _, v := range element {
		a.Append(v)
	}
}

func (a *StrArray) Get(index int) string {
	if a.len == 0 || index >= a.len {
		panic("Index over len: Request " + strconv.Itoa(index) + " But only " + strconv.Itoa(a.len))
	}
	return a.array[index]
}

func (a *StrArray) Len() int {
	return a.len
}

func (a *StrArray) Cap() int {
	return a.cap
}

func Print(array *StrArray) (result string) {
	result = "["
	for i := 0; i < array.Len(); i++ {
		if i == 0 {
			result = fmt.Sprintf("%s %s", result, array.Get(i))
			continue
		}

		result = fmt.Sprintf("%s %s", result, array.Get(i))
	}
	result = result + "]"
	return
}
