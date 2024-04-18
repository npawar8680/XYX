package main
import "fmt"
type student struct{
name string
age int
grade string
}
func (s*student)show(){
fmt.Printf("name:%s\nage:%d\ngrade:%s\n",s.name,s.age,s.grade)
}
func main(){
student:=&student{
name:"Abhishek samrit",age:20,grade:"A",
}
student.show()
}
