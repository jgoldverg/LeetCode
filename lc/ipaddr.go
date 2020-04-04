package lc

import (
	"fmt"
	"strconv"
	"strings"
)

func checkIPV4(addr string) bool {
	parts := strings.Split(addr, ".")
	if len(parts) > 3 || len(parts) < 1 {
		return false
	}
	for i := 0; i < len(parts); i++ {
		num, err := strconv.Atoi(parts[i])
		if err != nil {
			fmt.Printf("%q check IPV4 loop.", err)
		}
		if num < 0 || num > 255 {
			return false
		}
		if val := checkPrecedingZeroes(parts[i]); val {
			return false
		}
	}
	return true
}

func checkPrecedingZeroes(part string) bool {
	if res := strings.Compare(part, "0"); res == 1 {
		return false
	}
	ret := strings.HasPrefix(part, "0")
	return ret
}

func ipv6map() map[string]int {
	return map[string]int{
		"1": 0,
		"2": 0,
		"3": 0,
		"4": 0,
		"5": 0,
		"6": 0,
		"7": 0,
		"8": 0,
		"9": 0,
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
	}
}

func checkIPV6(addr string) bool {
	parts := strings.Split(addr, "\\:")
	lu := ipv6map()
	for i := 0; i < len(parts); i++ {
		if val := checkPrecedingZeroes(parts[i]); val {
			return false
		}
		if _, ok := lu[parts[i]]; !ok {
			return false
		}

	}
	return false
}

func validIPAddresss(IP string) string {
	if count := strings.Count(IP, "\\."); count == 3 {
		if val := checkIPV4(IP); val {
			return "IPV4"
		} else {
			return "Neither"
		}
	} else if num := strings.Count(IP, "\\:"); num == 8 {
		if val := checkIPV6(IP); val {
			return "IPV6"
		} else {
			return "Neither"
		}
	}
	return "Neither"
}
