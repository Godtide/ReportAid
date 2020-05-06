package xml

import (
    "net/http"
	"encoding/xml"
)

// Header - XML header
func Header(w http.ResponseWriter) {

    w.Write([]byte(xml.Header))
}
