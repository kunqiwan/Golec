package main

import "fmt"

func main() {
 s := "lEeTcOdE"
 fmt.Println(greatestLetter(s))
}
func greatestLetter(s string) string {
save :=[58]int{}
for n:=0;n<len(s);n++{
curr :=rune(s[n])
save[curr-'A']++
}
 fmt.Println(save)
max :=-1
for i:=32;i<58;i++{
if save[i] >0 && save[i-32]>0{
if i>max{
 max =i
}
}
 fmt.Println(max)
}

if max!=-1{
return string(max+33)
}
return ""
}