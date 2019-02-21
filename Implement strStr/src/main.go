package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println(strStr("abc","c"))


}

func strStr(haystack string, needle string) int {
	return strings.Index(haystack,needle)
}
