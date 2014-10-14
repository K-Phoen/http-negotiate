package negotiate

import (
	"github.com/K-Phoen/negotiation"
	"net/http"
)

type formatNegotiator struct {
	acceptedFormats []string
}

func FormatNegotiator(acceptedFormats []string) *formatNegotiator {
	return &formatNegotiator{
		acceptedFormats: acceptedFormats,
	}
}

func (negotiator *formatNegotiator) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	// no Accept header found
	if len(req.Header["Accept"]) == 0 {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	format, err := negotiation.NegotiateAccept(req.Header["Accept"][0], negotiator.acceptedFormats)

	// the negotiation failed
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// store the negotiated Content-Type in a header
	w.Header().Set("Content-Type", format.Value)

	// and call the other middlewares
	next(w, req)
}
