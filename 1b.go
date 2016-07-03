package main

import "fmt"
import "os"
import "bufio"
import "strconv"
import "regexp"

// convert from RxCy to <col><row> letter number system
func convertToCoord(row string, col string) string {
	cInt, _ := strconv.Atoi(col)
	return toBase26(cInt) + row
}

// convert from <col><row> letter number system to RxCy system
func convertToRC(column string, row string) string {
	return "R" + row + "C" + fromBase26(column)
}

// convert a number like 75 to base 26 using A-Z as characters
func toBase26(num int) string {
	ans := ""
	for num != 0 {
		rem := (num - 1) % 26
		num = (num - 1) / 26
		ans = string('A'+rem) + ans
	}
	return ans
}

func fromBase26(str string) string {
	v := 0
	multiplicand := 1
	for idx := len(str) - 1; idx >= 0; idx-- {
		temp := int(str[idx] - 'A' + 1)
		v += (temp * multiplicand)
		multiplicand *= 26
	}
	return strconv.Itoa(v)
}

func main() {
	rcRegexp := regexp.MustCompile(`^R(?P<Row>[[:digit:]]+)C(?P<Column>[[:digit:]]+)$`)
	coordRegexp := regexp.MustCompile(`^(?P<Column>[[:upper:]]+)(?P<Row>[[:digit:]]+)$`)

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	val, _ := strconv.Atoi(s.Text())
	//fmt.Println("Reading", val, "values")

	for i := 0; i < val; i++ {
		s.Scan()
		coord := s.Text()

		isRc := rcRegexp.FindStringSubmatch(coord)
		isCoord := coordRegexp.FindStringSubmatch(coord)

		if len(isRc) > 0 {
			fmt.Println(convertToCoord(isRc[1], isRc[2]))
		} else if len(isCoord) > 0 {
			fmt.Println(convertToRC(isCoord[1], isCoord[2]))
		}
	}
}
