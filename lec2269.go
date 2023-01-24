package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	i:="100"
	b, _ := strconv.Atoi(i)
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
}
