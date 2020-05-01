package xml

import (
	"encoding/xml"
)

// Desc - activity description
type Desc struct {
        XMLName   	xml.Name `xml:"description"`
        Type        uint8 `xml:"type,attr"`
        Narrative   string `xml:"narrative"`
}
