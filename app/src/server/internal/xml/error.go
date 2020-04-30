package xml

import (
	"log"
	"encoding/xml"

	"github.com/ReportAid/app/src/server/internal/configs"
)

// Error - create conditions for reporting back errors
type Error struct {
	XMLName    xml.Name `xml:"error"`
    Errors []*Errors `xml:"error-list"`
}

// Errors - make specific error
type Errors struct {
	Error      string `xml:"error>narrative"`
    ID         error `xml:"error>error"`
}

// home - homepage XML
type notFound struct {
	XMLName   	xml.Name `xml:"error"`
	Text	    string `xml:"not-found"`
}

// LogInit =
func LogInit () (*Error) {
    return &Error{}
}

// Log - return error and log
func Log (e string, id error) (*Errors) {

    thisError := &Errors{Error: e, ID: id}
    if id == nil {
        log.Printf("%s\n", e)
    } else{
        log.Printf("%s\n%v\n", e, id)
    }
    return thisError
}

// NotFound - get total activities
func NotFound () ([]byte) {

	notFound := &notFound{Text: configs.ErrorNotFound}
	notFoundXML, _ := xml.MarshalIndent(notFound, "", "")
	return notFoundXML
}
