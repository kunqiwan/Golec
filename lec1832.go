package main


func checkIfPangram(sentence string) bool {
letter :=[26]int{}
for n:=0;n<len(sentence);n++{
	curr :=rune(sentence[n])
	letter[curr-'a']++
}
for m:=0;m<len(letter);m++{
if letter[m] ==0{
return false
}

}
return true
}
