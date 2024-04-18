<data>
<person>
	<name>John Doe</name>
	<age>30</age>
	<city>New Yourk</city>
</person>
<person>
	<name>Jane Smith</name>
	<age>25</age>
	<city>Los Angeles</city>
</person>
</data>

package main
import (
	"encoding/xml"
	"fmt"
	"os"
)
type Person struct {
	Name string 'xml:"name"'
	Age int      'xml:"age"'
	City string   'xml:"city"
}
type Data struct {
      People []Person 'xml:"person"'
}
func main() {
	file,err := os.Open("data.xml")
	if err ! = nil {
		fmt.Println(Error:",err)
		return
	}
	fmt.Println("Data from XML file:")
	forr_, person := range data.People {
	fmt.Println("Name:",person.Name)
	fmt.Println("Age:",person.Age)
	fmt.Println("City:",person.City)
	fmt.Println()
	}
	
