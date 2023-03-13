package main

func isHappy(n int) bool {
	//create a map to stroe every sum of every digit to  determine whether they appear
	m :=make(map[int]bool)
	for n!=1 && !m[n]{
		//if not appear before, save it in the map and change it to true
		n,m[n] = digit_sum(n),true
	}

return n==1
}
func digit_sum(n int)int{
	sum:=0
	for n>0{
		// if we want ot get the last digit number , we could use  a%10 to get it
		sum = sum+(n%10)*(n%10)
		n=n/10
	}
	return sum
}
//func isHappy(n int) bool {
//	if n==1 {
//		return false
//	}
//	numStr := strconv.Itoa(n) // Convert the number to a string
//	numDigits := len(numStr)
//	result:=0
//	for numDigits>0{
//		result=0
//		for i := 0; i < len(numStr); i++ {
//			numInt, err := strconv.Atoi(string(numStr[i]))
//			if err != nil {
//				// Handle error
//			}
//			power :=numInt*numInt
//			result =result+int(power)
//		}
//
//		numStr = strconv.Itoa(result) // Convert the number to a string
//		numDigits = len(numStr)
//	}
//	if result == 1{
//		return true
//	} else {
//		return false
//	}
//}
