package сonverts

import (
	"strconv"
)

func StringToInt64(str string) (number int64) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
		return
	}
	return i
}

func StringToInt(str string) (number int) {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
		return
	}
	return i
}

func IntToString(number int) (str string) {
	t := strconv.Itoa(number)
	return t
}

func IntToInt64(number int) (number64 int64) {
	number64 = int64(number)
	return number64
}

func UintToInt(uNumber uint) (intResult int) {
	intResult = int(uNumber)
	return intResult
}
func IntToUint(Number int) (uintResult uint) {
	uintResult = uint(Number)
	return uintResult
}
