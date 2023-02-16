package main

func main() {

}
func repeatedCharacter(s string) byte {
occ := [26]int{}
for n:=0;n<len(s);n++{
	curr := s[n]
	occ[curr-'a']++
	if occ[curr-'a'] == 2{
		return byte(curr)
	}
}
return 'z'
}
