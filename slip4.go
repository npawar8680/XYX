package main
import "fmt"
func main(){
var num,rem int
var sum int=0
fmt.Printf("Enter number ")
fmt.Scanln(&num)
for{
 rem =num%10
 sum =sum+rem
 num=num/10
if(num==0){
 break
}
}
fmt.Println("Recursive Sum:",sum)
}
