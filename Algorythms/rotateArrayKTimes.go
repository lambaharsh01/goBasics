package main
import "fmt"

func main() {
  var nums []int= []int {1,2,3,4,5,6};
  var k int= 2;
  rotate(nums, k);
  
  fmt.Println(nums);
}

func rotate(nums []int, k int){
    
    k = k % len(nums); 

    var newArray []int= nums[(len(nums)-k):];
    newArray= append(newArray, nums[:(len(nums)-k)]...);

    for i:=0; i< len(nums); i++{
        nums[i]=newArray[i];
    }
    
}