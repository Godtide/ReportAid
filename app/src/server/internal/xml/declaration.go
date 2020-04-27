package xml

import (
    "net/http"
)

// First - outputs the first line of XML
func First(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/xml")
    w.Write([]byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n"))
}
