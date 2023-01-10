

package main

import "fmt"

// 如果定义全局变量，只有前三种可以
var f int = 900

func main() {
	ex1 := "A"
	t := "AA"
	res := minWindow(ex1, t)
	fmt.Println(res)
}

func minWindow(s string, t string) string {
	//循环t，得到需要的字符个数
	needs := make(map[string]int, len(t))
	windows := make(map[string]int, len(t))
	for i := 0; i < len(t); i++ {
		curr := string(t[i])
		val, ok := needs[curr]
		if ok {
			needs[curr] = val + 1
		} else {
			needs[curr] = 1
		}
	}
	fmt.Println("看一下循环后的needs", needs)
	//开始遍历目标字符，设立左右指针，进行初步划窗
	left := 0
	right := 0
	valid := 0
	slen := len(s)
	//滑窗的右指针一直向右，直到碰到满足needs中的条件，每满足一个字符条件，就可以让valid+1
	//start用来记录符合条件的最小子串的起始位置
	start := 0
	max := 100000
	for right < slen {
		now := string(s[right])
		right++
		_, ok := needs[now]
		if ok {
			windows[now] = windows[now] + 1
			if windows[now] == needs[now] {
				valid++
			}
			//满足了最低条件,开始左滑
			for valid == len(needs) {
				if right-left < max {
					max = right - left
					start = left
					fmt.Println("max:", max)
				}
				//开始移动符合条件的子串中的子串，将左指针向右移动，减少无关项，直到不满足needs条件
				remove := string(s[left])
				left++
				_, ok := needs[remove]
				if ok {
					if windows[remove] == needs[remove] {
						//判断是否正好满足条件，如果是，就需要减少valid，让其跳出循环，已经减多了
						valid--
					}
					windows[remove]--
				}
			}
		}

	}
	if max == 100000 {
		return ""
	} else {
		return s[start : start+max]
	}
}
