package main


func checkInclusion(s1 string, s2 string) bool {
 needs :=make(map[string]int,len(s1))
 windows:=make(map[string]int,len(s1))
 for i:=0;i<len(s1);i++{
	 target :=string(s1[i])
	 _,ok := needs[target]
	 if ok {
		 needs[target] = needs[target]+1
	 } else {
		 needs[target] = 1
	 }
 }
 left :=0
 right :=0
 valid :=0
 for right <len(s2){
	 curr :=string(s2[right])
	 right++
	 _,ok := needs[curr]
	 if ok{
		 windows[curr]++
		 if windows[curr] == needs[curr]{
			 valid++
		 }
		 for  valid == len(needs) {
			 if right-left == len(s1){
					 return true}
				 now := string(s2[left])
			     left++
				 _,ok := needs[now]
				 if ok{
					 if windows[now] == needs[now]{
						 valid--
					 }
					 windows[now]--
				 }
		 }

	 }
 }
return false
}
