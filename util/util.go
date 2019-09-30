package util

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const EMPTY = "empty"

// Get filename without numeric prefix and .md suffix
func GetMDRealName(name string) (realName string, isMd bool) {
	end := strings.LastIndex(name, ".md")
	if end == -1 {
		//end = len(name)
		return "", false
	}
	return name[:end], true
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

func GetFileMd5(filepath string) (sum string, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	if len(bytes) == 0 {
		return EMPTY, nil
	}
	return fmt.Sprintf("%x", md5.Sum(bytes)), nil
}
