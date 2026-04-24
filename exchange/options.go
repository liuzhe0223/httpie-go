package exchange

import (
	"net/http"
	"time"

	"github.com/liuzhe0223/vasign"
)

type Options struct {
	Timeout         time.Duration
	FollowRedirects bool
	Auth            AuthOptions
	SkipVerify      bool
	ForceHTTP1      bool
	CheckStatus     bool
	Transport       http.RoundTripper
	VASigner        *vasign.Signer
}

type AuthOptions struct {
	Enabled  bool
	UserName string
	Password string
}
