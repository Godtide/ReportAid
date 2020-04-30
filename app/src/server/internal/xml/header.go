package xml

import (
    "net/http"
)

// Header - XML header
func Header(w http.ResponseWriter) {

    w.Header().Set("Content-Type", "application/xml")
    w.Write([]byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n"))
}
