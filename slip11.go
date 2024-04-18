package main
import "fmt"
func main(){
     var num int
     fmt.Println("Enter number :")
     fmt.Scanln(&num) 
     if(num/10==0){
     fmt.Println("number of single digit")
     } else {
     fmt.Println("number is two digit")
     }
} 
