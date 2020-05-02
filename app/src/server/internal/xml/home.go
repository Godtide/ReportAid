package xml

import (
	"encoding/xml"

	"github.com/ReportAid/app/src/server/internal/configs"
)

// home - homepage XML
type home struct {
	XMLName   	xml.Name 	`xml:"reportaid:home"`
    Namespace   string   	`xml:"xmlns:reportaid,attr"`
	Text	    string 		`xml:"reportaid:text"`
}

// Home - get total activities
func Home () ([]byte) {

	homeXML := &home{Namespace: configs.URLReportAidNamespace, Text: configs.TextHome}
	thisXML, _ := xml.MarshalIndent(homeXML, "", "  ")
	return thisXML
}
