package enum

import "strconv"

type AntString string

func (v AntString) Int() int {
	val, _ := strconv.ParseInt(string(v), 10, 32)
	return int(val)
}

func (v AntString) String() string {
	return string(v)
}

type AntInt int

func (v AntInt) Int() int {
	return int(v)
}

func (v AntInt) String() string {
	return strconv.Itoa(int(v))
}
