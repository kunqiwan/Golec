package main


func minimumRecolors(blocks string, k int) int {
	//go语言，string不需要转为array就可以调用
	left:=0
	right:=0
	w_ocunt :=0
	max := 99999999
	for right<len(blocks) {
		if string(blocks[right]) == "W" {
			w_ocunt++
		}
		right++
		if right-left == k {
			if w_ocunt < max {
				max = w_ocunt
			}
			if string(blocks[left]) == "W" {
				w_ocunt--
			}
			left++
		}
	}
	return max
}
