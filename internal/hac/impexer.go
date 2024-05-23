package hac

import "net/http"

type Impexer struct {
	client *http.Client
}

func NewImpexer() Impexer {
	return Impexer{
		client: http.DefaultClient}
}

func (imp Impexer) TestHttp() {

}
