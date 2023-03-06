package main
func findTheDifference(s string, t string) byte {
tocc:=[26]int{}
for n:=0;n<len(t);n++{
	curr :=t[n]
	tocc[rune(curr-'a')]++
}
socc:=[26]int{}
for n:=0;n<len(s);n++{
	curr :=s[n]
		socc[rune(curr-'a')]++
}

for m:=0;m<len(tocc);m++{
	if tocc[m]-socc[m]!=0{
		return byte(rune('a' + m))
	}
	}
return 'a'
}
