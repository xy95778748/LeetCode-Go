package string

/*
	171. Excel表列序号
	https://leetcode-cn.com/problems/excel-sheet-column-number/submissions/
*/
func titleToNumber(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum = 26*sum + (int(s[i]) - 64)
	}
	return sum
}
