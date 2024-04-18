//slip 3,5 B same
package main
import "fmt"
type employee struct{
eno int
esal float64
ename string
}
func main(){
var e1[10] employee
var n,k int
var max float64
fmt.Println("enter no of employee you want:")
fmt.Scanf("%d",&n)
for i:=0;i<n;i++{
fmt.Println("enter employee no")
fmt.Scanf("%d",&e1[i].eno)
fmt.Println("enter employee name:")
fmt.Scanf("%s",&e1[i].ename)
fmt.Println("enter employee salary")
fmt.Scanf("%f",&e1[i].esal)
}
max=e1[0].esal
for i:=0;i<n;i++{
if(e1[i].esal>max){
max=e1[i].esal
k=i
}
}
fmt.Println("employee having maximun salary...")
fmt.Println("employee no:",e1[k].eno)
fmt.Println("employee name:",e1[k].ename)
fmt.Println("employee salary:",e1[k].esal)
}
