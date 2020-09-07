package req

import (
	"io"
	"net/http"
)

// Generate Request for the client
func GenReq(reqType, url string, reader io.Reader) (req *http.Request, error)  {
    return http.NewRequest(reqType, url, reader)
}

// Build client
func Build() {

}
