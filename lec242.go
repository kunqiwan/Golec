package main


func isAnagram(s string, t string) bool {
s1 :=occ(s)
t1 :=occ(t)
if s1 ==t1{
	return true
} else {
	return false
}
}
func occ(target string) [26]int{
	save :=[26]int{}
	for n:=0;n<len(target);n++{
		curr :=target[n]
		save[curr-'a']++
	}
	return  save
}