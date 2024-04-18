package main
import "fmt"
func main(){
var a,b int
fmt.Printf("enter two number\n")
fmt.Scanf("%d%d",&a,&b)
fmt.Printf("before swap value of a=%d b=%d\n",a,b)
swap(&a,&b)
fmt.Printf("after swap value of a=%d b=%d\n",a,b)
}
func swap(x*int,y*int){
var temp int
temp=*x
*x=*y
*y=temp
}
