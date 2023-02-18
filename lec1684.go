package main

func main() {

}
func countConsistentStrings(allowed string, words []string) int {
count :=0
allow :=map[rune]struct{}{}
for z:=0;z<len(allowed);z++{
	_,has :=allow[rune(allowed[z])]
	if has{
	} else{
		allow[rune(allowed[z])] = struct{}{}
	}

}
for n:=0;n<len(words);n++{
	curr :=words[n]
	length :=len(curr)
	bad :=0
	for m:=0;m<length;m++{
		cha :=rune(curr[m])
		_,have := allow[cha]
		if have{

		} else {
			bad =1
			break
		}
	}
	if bad == 0{
		count++
	}
}
return count
}