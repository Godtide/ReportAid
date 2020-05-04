package xml

import (
	"log"
	"encoding/xml"

	"github.com/ReportAid/app/src/server/internal/configs"
)

// Error - create conditions for reporting back errors
type Error struct {
	XMLName    xml.Name `xml:"reportaid:error"`
	Namespace  string   `xml:"xmlns:reportaid,attr"`
    Errors 	   []*Errors `xml:"reportaid:error-list"`
}

// Errors - make specific error
type Errors struct {
	Error      string 	`xml:"reportaid:error>narrative"`
    ID         error 	`xml:"reportaid:error>error"`
}

// notFound - 404 XML
type notFound struct {
	XMLName   	xml.Name `xml:"reportaid:error"`
	Namespace   string   `xml:"xmlns:reportaid,attr"`
	Text	    string 	 `xml:"reportaid:not-found"`
}

// disallowed - POST XML
type disallowed struct {
	XMLName   	xml.Name 	`xml:"reportaid:error"`
	Namespace  string   	`xml:"xmlns:reportaid,attr"`
	Text	    string 		`xml:"reportaid:method-disallowed"`
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

// NotFound - 404
func NotFound () ([]byte) {

	notFound := &notFound{Text: configs.ErrorNotFound}
	notFoundXML, _ := xml.MarshalIndent(notFound, "", "")
	return notFoundXML
}

// PostDisallowed - handle POSt requests
func PostDisallowed () ([]byte) {

	disallowed := &disallowed{Text: configs.ErrorPostDisallowed}
	disallowedXML, _ := xml.MarshalIndent(disallowed, "", "")
	return disallowedXML
}
