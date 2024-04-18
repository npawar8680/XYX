package main
import "fmt"
func remove(s []int,index int) []int{
 return append(s[:index], s[index+1:]...) }
func main() {
s1:= []int {1,2,3,4}
fmt.Println("slice 1:",s1)
var s2 []int
fmt.Println("appending elements at the end of slice")
s1=append(s1,9,8)
fmt.Println("slice after appending:")
fmt.Println(s1)
fmt.Println("removing element from slice:")
s1=remove(s1,3)
fmt.Println("slice after removing an element:",s1)
fmt.Println("copying elements of slice in another slice:")
s2=s1
fmt.Println("copied slice:",s2)
}
