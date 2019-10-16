package mxj

import (
    "fmt"
    "testing"
    "github.com/ssomagani/mxj"
)

var authorDoc = []byte(`
<biblio>
    <geography country="UK"></geography>
	<author>
		<name>William Gaddis</name>
		<books>
			<book>
				<title>The Recognitions</title>
				<date>1955</date>
				<review>A novel that changed the face of American literature.</review>
			</book>
			<book>
				<title>JR</title>
				<date>1975</date>
				<review>Winner of National Book Award for Fiction.</review>
			</book>
		</books>
	</author>
	<author>
		<name>John Hawkes</name>
		<books>
			<book>
				<title>The Cannibal</title>
				<date>1949</date>
				<review>A novel on his experiences in WWII.</review>
			</book>
			<book>
				<title>The Beetle Leg</title>
				<date>1951</date>
				<review>A lyrical novel about the construction of Ft. Peck Dam in Montana.</review>
			</book>
			<book>
				<title>The Blood Oranges</title>
				<date>1970</date>
				<review>Where everyone wants to vacation.</review>
			</book>
		</books>
        <residence>
            <state>MA</state>
            <city>Boston</city>
        </residence>
        <residence>
            <state>NY</state>
            <city>NY</city>
        </residence>
	</author>
</biblio>`)

func TestUpdateCreate(t *testing.T) {
	m, merr := mxj.NewMapXml(authorDoc)
	if merr != nil {
		t.Fatal("merr:", merr.Error())
	}
//	fmt.Println(m.StringIndent())

	m.UpdateValuesForPath("comment:Whatever", "biblio.author.books.book", "title:JR")
    m.UpdateValuesForPath("-country:usa", "biblio.geography")
    m.UpdateValuesForPath("-country:usa", "biblio.author.residence")
    
    // Don't use wild cards. It'll go to shit
    //m.UpdateValuesForPath("comment:not bad", "*.*.*.*", "title:The Blood Oranges")
    
    fmt.Println("_________________RESULT________________________")
    fmt.Println(m.StringIndent())
//    xmlStr, _ := m.XmlIndent("", "    ")
//    fmt.Println(string(xmlStr))
}
