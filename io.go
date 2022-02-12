package leetcodeIO

import (
	"fmt"
	"strconv"
	"strings"
)

func GetLinearArray()[]int{
	var input string
	fmt.Scan(&input)
	input = input[1:len(input)-1]
	strArr := strings.Split(input,",")
	ans := make([]int, len(strArr))
	for i:=0;i< len(strArr);i++{
		ans[i],_ = strconv.Atoi(strArr[i])
	}
	return ans
}
type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

func Get2DArray()[][]int{
	var input string
	fmt.Scan(&input)
	input = input[2:len(input)-2]
	strArr := strings.Split(input,"],[")
	ans := make([][]int, len(strArr))
	for i,str1 := range strArr{
		if str1 == ""{
			continue
		}
		strArr1 := strings.Split(str1,",")
		ans[i] = make([]int, len(strArr1))
		for j,str2 := range strArr1{
			ans[i][j], _ = strconv.Atoi(str2)
		}
	}
	return ans
}

func GetTreeNode()*TreeNode{
	var input string
	fmt.Scan(&input)
	input = input[1:len(input)-1]
	strArr := strings.Split(input,",")
	if len(strArr) == 0 || strArr[0] == "null"{
		return nil
	}
	val,_ := strconv.Atoi(strArr[0])
	root := TreeNode{
		Val:	val,
	}
	var q []*TreeNode
	q = append(q, &root)
	index1,index2,l := 1,2, len(strArr)
	for index1 < l{
		curN := q[0]
		q = q[1:]
		if strArr[index1] == "null"{
			curN.Left = nil
		}else{
			val,_ := strconv.Atoi(strArr[index1])
			curN.Left = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
			q = append(q, curN.Left)
		}
		if index2 >= l{
			break
		}
		if strArr[index2] == "null"{
			curN.Right = nil
		}else{
			val,_ := strconv.Atoi(strArr[index2])
			curN.Right = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
			q = append(q, curN.Right)
		}
		index1 += 2
		index2 += 2
	}
	return &root
}

func GetString()string{
	var in string
	fmt.Scan(&in)
	in = in[1:len(in)-1]
	return in
}
func GetStringArr()[]string{
	var in string
	fmt.Scan(&in)
	in = in[2:len(in)-2]
	ans := strings.Split(in,"\",\"")
	return ans
}
func Get2DStringArray()[][]string{
	var input string
	fmt.Scan(&input)
	input = input[2:len(input)-2]
	strArr := strings.Split(input,"],[")
	ans := make([][]string, len(strArr))
	for i,str1 := range strArr{
		if str1 == ""{
			continue
		}
		ans[i] = strings.Split(str1,",")
	}
	return ans
}