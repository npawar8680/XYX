package main
import (
	"fmt"
	"os"
)
func main(){
file,err := os.OpenFile("example.txt",os.O_APPEND|os.O_WRONLY|os.O_CREATE,0644)
	if err !=nil {
		fmt.Println("Error :",err)
		return
		}
		defer file.Close()
		if err !=nil{
		fmt.Println("Error :",err)
		return
	}
	fmt.Println("content appent to the file Success .")
	}
