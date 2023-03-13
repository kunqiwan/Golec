package main

import (
	"fmt"
	"strings"
)

func main() {

	s := strings.Split("dog cat cat dog", " ")
	fmt.Println(s)
}
func wordPattern(pattern string, s string) bool {
	sslice := strings.Split(s, " ")
	plen :=len(pattern)
	slen  := len(sslice)
	pair :=make(map[byte]string)
	if plen !=slen{
		return false
	}
	for i:=0;i<plen;i++{
		key := pattern[i]
		value := sslice[i]
		_,ok:= pair[key]
		if !ok {
			pair[key] = value
			size := getKeyByValue(pair, value)
		if len(size)>1{
			return false
		}
		} else {
			if value ==pair[key]{

			} else {
				return false
			}
		}
	}
	return true
}

func getKeyByValue(m map[byte]string, value string) []byte {
	res :=[]byte{}
	for k, v := range m {
		if v == value {
			res  =append(res,k)
		}
	}
	return res
}