package service

import (
	"github.com/darkweak/souin/cache/providers"
	"github.com/darkweak/souin/cache/types"
	"github.com/darkweak/souin/helpers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeResponse(t *testing.T) {
	c := MockConfiguration()
	prs := providers.InitializeProvider(c)
	regexpUrls := helpers.InitializeRegexp(c)
	retriever := &types.RetrieverResponseProperties{
		Configuration: c,
		Provider:      prs,
		MatchedURL:    getMatchedURL(PATH),
		RegexpUrls:    regexpUrls,
	}
	r := httptest.NewRequest("GET", "http://"+DOMAIN+PATH, nil)
	w := httptest.NewRecorder()
	ServeResponse(
		w,
		r,
		retriever,
		func(rw http.ResponseWriter, rq *http.Request, r types.RetrieverResponsePropertiesInterface, key string) {},
	)
}