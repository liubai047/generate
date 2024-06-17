package enum

import "strconv"

type ValString string

func (v ValString) Int() int {
	val, _ := strconv.ParseInt(string(v), 10, 32)
	return int(val)
}

func (v ValString) String() string {
	return string(v)
}

type ValInt int

func (v ValInt) Int() int {
	return int(v)
}

func (v ValInt) String() string {
	return strconv.Itoa(int(v))
}
