package xml

import (
	"encoding/xml"
	"strings"

	"github.com/ReportAid/app/src/server/internal/configs"
)

// home - homepage XML
type help struct {
	XMLName   	xml.Name 	`xml:"reportaid:help"`
    Namespace   string   	`xml:"xmlns:reportaid,attr"`
	Text	    []*helpText	`xml:"reportaid:text"`
}

type helpText struct {
	Heading		string 		`xml:"reportaid:text>reportaid:heading"`
	Text	    []string 	`xml:"reportaid:text>reportaid:narrative"`
}

// Help - get total activities
func Help () ([]byte) {

	helpTextXML := &helpText{Heading: configs.TextHelpURLSHeading}

	split := strings.Split( configs.TextHelpURLs, "\n")
    for _, line := range split {
		helpTextXML.Text = append(helpTextXML.Text, line)
    }

	helpXML := &help{Namespace: configs.URLReportAidNamespace}
	helpXML.Text = append(helpXML.Text, helpTextXML)
	thisXML, _ := xml.MarshalIndent(helpXML, "", "  ")
	return thisXML
}
