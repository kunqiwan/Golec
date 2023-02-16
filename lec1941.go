package main

func main() {

}
func areOccurrencesEqual(s string) bool {
save :=make(map[rune]int)
for n:=0;n<len(s);n++{
	curr:=rune(s[n])
	save[curr]++
}
first :=save[rune(s[0])]
for _,value := range save{
			if value != first{
				return false
			}
}
return true
}