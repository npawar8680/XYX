package main
import "fmt"
func main(){

var matrix1 [2][2]int
var matrix2 [2][2]int
var matrix3 [2][2]int
fmt.Printf("enter matrix1 elements:\n")
for i:=0;i<2;i++{
for j:=0;j<2;j++{
fmt.Printf("elements :matrix[%d][%d]:",i,j)
fmt.Scanf("%d ",&matrix1[i][j])
}
}
fmt.Printf("enter matrix2 elements:\n")
for i:=0;i<2;i++{
for j:=0;j<2;j++{
fmt.Printf("elements matrix2[%d][%d]:",i,j)
fmt.Scanf("%d ",&matrix2[i][j])
}
}
for i:=0;i<2;i++{
for j:=0;j<2;j++{
matrix3[i][j]= matrix1[i][j]*matrix2[i][j]

}
}
fmt.Printf("matrix1:\n")
for i:=0;i<2;i++{
for j:=0;j<2;j++{
fmt.Printf("%d ",matrix1[i][j])
}
fmt.Printf("\n")
}
fmt.Printf("matrix2:\n")
for i:=0;i<2;i++{
for j:=0;j<2;j++{
fmt.Printf("%d ",matrix2[i][j])
}
fmt.Printf("\n")
}
fmt.Printf("multiplication of matrix1 & matrix2:\n")
for i:=0;i<2;i++{
for j:=0;j<2;j++{
fmt.Printf("%d ",matrix3[i][j])
}
fmt.Printf("\n")
}
}
