package main
import "fmt"
func main(){
var num,rem int
var rev int=0
fmt.Printf("Enter number to check pallindrome")
fmt.Scanln(&num)
var temp int=num
for{
 rem = num%10
 rev = rev*10+rem
 num = num/10
if(num==0){
 break
}
}
if(rev==temp){
fmt.Println("Given number is pallindrome")
}else{
fmt.Println("Given number is not pallindrome")
}
}
