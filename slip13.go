package main
import "fmt"
func main(){
	evensum:=0
	oddsum:=0
	for i:=1;i<=99;i++{
		if i%2==0{
			evensum+=i
		}else{
			oddsum+=i
		}
		}
		fmt.Printf("sum of even number :%d\n",evensum)
		fmt.Printf("sum of odd number :%d\n",oddsum)
}
