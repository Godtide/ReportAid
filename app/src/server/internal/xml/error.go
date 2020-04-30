package xml

import (
	"log"
	"encoding/xml"
)

// Error - create conditions for reporting back errors
type Error struct {
	XMLName    xml.Name `xml:"errors"`
    Errors []*Errors
}

// Errors - make specific error
type Errors struct {
	Error      string `xml:"error>narrative"`
    ID         error `xml:"error>error"`
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
