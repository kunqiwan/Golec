package main

func main() {

}
func groupAnagrams(strs []string) [][]string {
save :=map[[26]int][]string{}
for n:=0;n<len(strs);n++{
	every_ch := [26]int{}
	for m:=0;m<len(strs[n]);m++{
		every_ch[strs[n][m]-'a']++
	}
	save[every_ch] = append(save[every_ch],strs[n])
}
result :=[][]string{}
for _,value:=range save{
	result = append(result,value)
	}
return result
}
