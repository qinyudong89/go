package slice_test

import "testing"

func TestDel(t *testing.T){
  nums := [] int{1,2,3,4,5}
  //从头删除
  nums = nums[2:]
  t.Log(nums)
}

func TestAppend(t *testing.T){
  var nums [] int
  for i := 0; i < 10; i++ {
    nums = append(nums,i)
  }
  t.Log(nums,cap(nums))
}

func TestCap(t *testing.T)  {
  monthos := []int{1,2,3,4,5,6,7,8,9,10,11,12}
  val := monthos[3:5]
  t.Log(cap(val))
  for i := 0; i < len(monthos); i++ {
    t.Log(monthos[i])
  }
}

func TestClear(t *testing.T){
  monthos := []int{1,2,3,4,5,6,7,8,9,10,11,12}
  t.Log(monthos)
  monthos = append(monthos[:1],30)
  t.Log(monthos[:])
}

func TestMake(t *testing.T){
  nums := make([]int, 0, 8)
  slice := []int{10,20,30,40,50}
  nums = append(nums,slice...)
  nums= append(nums[:2], append([]int{99}, nums[2:]...)...)
  t.Log(nums,slice)
}
