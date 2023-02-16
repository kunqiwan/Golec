package main

func main() {

}
func numJewelsInStones(jewels string, stones string) int {
count:=0
	save :=map[rune]struct{}{}
	for m:=0;m<len(jewels);m++{
		save[rune(jewels[m])] = struct{}{}
	}
for n:=0;n<len(stones);n++{
	_,has:=save[rune(stones[n])]
	if has{
		count++
	}
}
return count
}
