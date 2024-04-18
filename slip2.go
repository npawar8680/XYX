package main
import "fmt"
func main() {
var num1 int=0
var num2 int=1
var next int=num2
var n int
fmt.Printf("Enter number of terms :")
fmt.Scanln(&n)
fmt.Printf("%d ",num1)
for i:=2; i<=n; i++ {
fmt.Printf("%d ",next)
next=num1+num2
num1=num2
num2=next
}
}
