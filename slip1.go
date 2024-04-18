package main
import "fmt"
func main(){
     var num1 int
     var num2 int
     var ch int
     fmt.Println("Enter the number 1")
     fmt.Scanln(&num1)
     fmt.Println("Enter the number of 2")
     fmt.Scanln(&num2)
     fmt.Println("1.Addition\n2.Subtraction\n3.Multiplication\n4.Division\n enter your choice:")
     fmt.Scanln(&ch)
     switch ch{
     case 1:fmt.Printf("\nAddition of %d=",num1+num2);
     case 2:fmt.Printf("\nSubtraction of %d=",num1-num2);
     case 3:fmt.Printf("\nMultiplication of %d=",num1*num2);
     case 4:fmt.Printf("\nDivision of %d=",num1/num2);
     }
}     
