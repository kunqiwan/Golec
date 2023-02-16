package main

func main() {

}
func countNegatives(grid [][]int) int {
	count :=0
for m:=0;m<len(grid);m++{
//curr :=grid[m]
l :=0
r :=len(grid[m])-1
for l<=r{
	mid :=l+(r-l)/2
	if grid[m][mid] <0{
		r--
	} else {
		l++
	}
}
count = count +len(grid[m])-1-r
}
return count
}

