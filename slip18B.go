package calculator
import "fmt"
func Add (a,b float64) float64 {
return a+b
}
func Subtract(a,b float64) float64 {
return a-b
}
func Multiply(a,b float64) float64 {
return a*b
}
main.go

package main
import(
	"fmt"
	"calculator"
)
func main(){
	var choice int
	var a,b floa64
	
	fmt.Println("Simple Calculator")
	fmt.Println("1.Add")
	fmt.Println("2.Subtract")
	fmt.Println("3.Multiply")
	fmt.Println("4.Divide")
	
	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)
	
	switch choice {
	case 1:
	result := calculator.Add(a,b)
	fmt.Println("Result :"result)
	case 2:
	result := calculator.Subtract(a,b)
	fmt.Println("Result :"result)
	case 3:
	result := calculator.Multiply(a,b)
	fmt.Println("Result :"result)
	case 4:
	result,err := calculator.Divided(a,b)
	if err !=nil{
		fmt.Println("Result:",result)
	}else{
		fmt.Println("Result:",result)
		}
	default:
	fmt.Println("Invalid choice")
	}
}
func Divide (a,b float64) (float64,error) {
if b == 0 {
return 0
fmt.Errorf("division by zero")
}
return a/b,nil
}	
