package xml

import "fmt"

func Sql(id string) string {
	doc := NewDocument()
	if err := doc.ReadFromFile("./test.xml"); err != nil {
		panic(err)
	}

	for _, t := range doc.FindElements(fmt.Sprintf("//sql[@id='%s']", id)) {
		return t.Text()
	}
	return ""
}
