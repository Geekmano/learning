package sort

import "testing"

var nums=[...]int{1,5,4,2,3,6}

func TestBubbleSort(t *testing.T) {
	for i:=0;i<len(nums);i++{
		sort:=false
		for j:=0;j<len(nums)-i-1;j++{
			if nums[j]>nums[j+1]{
				nums[j],nums[j+1]=nums[j+1],nums[j]
				sort=true
			}
		}
		if !sort{
			break
		}
	}
	t.Log(nums)
}

func TestQUicSort(t *testing.T) {
	var nums=[]int{1,6,3,5,4,2}
	quickSort(nums,0,len(nums)-1)
	t.Log(nums)

}
func quickSort(nums []int,left,right int)  {
	if left<right{
		index:=partition(nums,left,right)
		quickSort(nums,left,index-1)
		quickSort(nums,index+1,right)
	}
}
func partition(nums []int,left,right int) int {
	pivotVal:=nums[left]
	for left<right{
		for right>left&&nums[right]>=pivotVal{
			right--
		}
		nums[left]=nums[right]
		for left<right&&nums[left]<=pivotVal{
			left++
		}
		nums[right]=nums[left]
	}
	nums[left]=pivotVal
	return left

}