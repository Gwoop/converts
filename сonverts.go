package сonverts

import (
	"github.com/google/uuid"
	"log"
	"strconv"
)

func StringToInt64(str string) (number int64, err error) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}
	return i, err
}

func StringToInt(str string) (number int, err error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Println(err)
		return
	}
	return i, err
}

func IntToString(number int) (str string) {
	t := strconv.Itoa(number)
	return t
}

func IntToInt64(number int) (number64 int64) {
	number64 = int64(number)
	return number64
}

func StringToUuid(str string) (UUID uuid.UUID) {
	UUID = uuid.Must(uuid.Parse(str))
	return UUID
}
func StringToUint(str string) (Uint uint, err error) {
	u64, err := strconv.ParseUint(str, 10, 32)
	Uint = uint(u64)

	return Uint, err
}

func UintToInt(uNumber uint) (intResult int) {
	intResult = int(uNumber)
	return intResult
}
func IntToUint(Number int) (uintResult uint) {
	uintResult = uint(Number)
	return uintResult
}
