package main
import "fmt"
func Demo(n1,n2 int)(int,int){
return n1+n2,n1-n2
}
func main(){
var n1,n2 int
fmt.Printf("Enter Two` Number :")
fmt.Scanf("%d%d",&n1,&n2)
add,Subtract:=Demo(n1,n2)
fmt.Printf("Add :%d\nSubtraction :%d\n",add,Subtract)
}
