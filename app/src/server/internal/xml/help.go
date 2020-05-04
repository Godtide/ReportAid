package xml

import (
	"encoding/xml"

	"github.com/ReportAid/app/src/server/internal/configs"
)

// home - homepage XML
type help struct {
	XMLName   	xml.Name 	`xml:"reportaid:help"`
    Namespace   string   	`xml:"xmlns:reportaid,attr"`
	Text	    string 		`xml:"reportaid:text"`
}

// Help - get total activities
func Help () ([]byte) {

	helpXML := &home{Namespace: configs.URLReportAidNamespace, Text: configs.TextHelp}
	thisXML, _ := xml.MarshalIndent(helpXML, "", "  ")
	return thisXML
}
