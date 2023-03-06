package main

import "fmt"


func similarPairs(words []string) int {
	save := map[[26]int]int{}
for n:=0;n<len(words);n++{
	curr := words[n]
	occ :=[26]int{}
	for m:=0;m<len(curr);m++{
		occ[curr[m]-'a'] =1
	}
	save[occ]++
}
count :=0

fmt.Println(save)
for _,value:=range save{
		fac := value*(value-1)/2
		count  = count+fac
	fmt.Println(count)
	}
return count
}
