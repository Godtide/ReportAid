package xml

import (
	"encoding/xml"

	"github.com/ReportAid/app/src/server/internal/configs"
)

// home - homepage XML
type home struct {
	XMLName   	xml.Name `xml:"home"`
	Text	    string `xml:"text"`
}

// Home - get total activities
func Home () ([]byte) {

	homeXML := &home{Text: configs.TextHome}
	thisXML, _ := xml.MarshalIndent(homeXML, "", "  ")
	return thisXML
}
