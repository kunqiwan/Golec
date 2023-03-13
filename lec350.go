package main

func main() {

}
func intersect(nums1 []int, nums2 []int) []int {
lennums1 :=len(nums1)
lennums2 :=len(nums2)
res := []int{}
//find the small size slice
if lennums1>=lennums2{
save :=map[int]int{}
for i:=0;i<len(nums2);i++{
save[nums2[i]]++
}

for n:=0;n<len(nums1);n++{
	curr :=nums1[n]
	v,_ := save[curr]
	if v>0{
		res = append(res,curr)
		save[curr]--
	} else {

	}
}

}

	if lennums1<lennums2{
		save :=map[int]int{}
		for i:=0;i<len(nums1);i++{
			save[nums1[i]]++
		}

		for n:=0;n<len(nums2);n++{
			curr :=nums2[n]
			v,_ := save[curr]
			if v>0{
				res = append(res,curr)
				save[curr]--
			} else {

			}
		}

	}

	return res
}