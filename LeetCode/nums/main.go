package main



func twoNums(nums []int,target int) []int {
	/*for i:=0;i< len(nums)-1;i++{
		for j:=i+1;j<len(nums);j++{
			if nums[i]+nums[j]==target {
				return []int{i,j}
			}
			return []int{}
		}
	}
	return []int{}*/

	/*m:=make(map[int]int)
	for i,x:=range nums {
		_,prs:=m[x]
		if prs {
			return []int{m[x],i}
		}else {
			m[target-x]=i
		}
	}
	return nil*/

	//其实就是在 m 中存下遇到的数的位置，然后对于新的 num, 看 target-num 有没有保存，有的话那就找到这个数对了，返回就行
	//注意这里肯定是先找再存
	marks := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		numTarget := target - num
		index, ok := marks[numTarget]
		if ok {
			return []int{i, index}
		}

		marks[num] = i
	}
}

func reverseList(head *ListNode) *ListNode {
	var newh ListNode
	for head!=nil {
		newh.Next,head.Next,head=head,newh.Next,head.Next
	}
	return newh.Next

}
func main(){
	
}
