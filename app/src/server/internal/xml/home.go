package xml

import (
    "net/http"
	"encoding/xml"

	"github.com/ReportAid/app/src/server/internal/configs"
)

// home - homepage XML
type home struct {
	XMLName   	xml.Name `xml:"home"`
	Text	    string `xml:"text"`
}


// Home - get total activities
func Home (w http.ResponseWriter, r *http.Request) {

	homeXML := &home{Text: configs.TextHome}
	thisXML, _ := xml.MarshalIndent(homeXML, "", " ")
	w.Write(thisXML)
}
