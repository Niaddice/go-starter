package xml

import (
	"fmt"
	"strings"
	"testing"
)

func TestReading(t *testing.T) {
	doc := NewDocument()
	if err := doc.ReadFromFile("./test.xml"); err != nil {
		panic(err)
	}

	for _, t := range doc.FindElements(fmt.Sprintf("//sql[@id='%s']", "find1")) {
		fmt.Println(strings.Trim(t.Text(), " \n\t"))
	}
}
