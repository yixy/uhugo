package util

import (
	"strings"
)

// Get filename without numeric prefix and .md suffix
func GetMDRealName(name string) (realName string, isMd bool) {
	end := strings.LastIndex(name, ".md")
	if end == -1 {
		//end = len(name)
		return "", false
	}
	start := strings.Index(name, "|")
	if start == -1 {
		start = 0
	} else {
		start += 1
	}
	return name[start:end], true
}

var sizeTable []uint

const MaxUint uint = ^uint(0)

func init() {
	max := GetMaxUintSize()
	base := 1
	for i := 1; i <= max; i++ {
		base *= 10
		sizeTable = append(sizeTable, uint(base-1))
	}
	sizeTable = append(sizeTable, MaxUint)
}

func GetMaxUintSize() int {
	maxUint := MaxUint
	count := 0
	for {
		count++
		maxUint = maxUint / 10
		if maxUint/10 == 0 {
			break
		}
	}
	return count
}

func StringSize(u uint) (size int) {
	for i, v := range sizeTable {
		if u <= v {
			return i + 1
		}
	}
	return len(sizeTable)
}
