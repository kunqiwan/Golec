package main



func countGoodSubstrings(s string) int {
left :=0
right :=0
count :=0
for right < len(s){
	right ++
	if right -left ==3{
		if s[right-1]!=s[right-2] &&s[right-1]!=s[right-3] && s[right-3]!=s[right-2]{
			count++
		}
		left++
	}
}
return count
}
