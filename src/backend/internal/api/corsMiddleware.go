package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type CorsConfig struct {
	AllowOrigin      string
	AllowHeaders     string
	AllowMethods     string
	AllowCredentials bool
}

func DefaultCorsConfig() CorsConfig {
	return CorsConfig{
		AllowOrigin:      "*",
		AllowMethods:     "GET, POST, PATCH, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Accept, Accept-CH, Accept-Charset, Accept-Datetime, Accept-Encoding, Accept-Ext, Accept-Features, Accept-Language, Accept-Params, Accept-Ranges, Access-Control-Allow-Credentials, Access-Control-Allow-Headers, Access-Control-Allow-Methods, Access-Control-Allow-Origin, Access-Control-Expose-Headers, Access-Control-Max-Age, Access-Control-Request-Headers, Access-Control-Request-Method, Age, Allow, Alternates, Authentication-Info, Authorization, C-Ext, C-Man, C-Opt, C-PEP, C-PEP-Info, CONNECT, Cache-Control, Compliance, Connection, Content-Base, Content-Disposition, Content-Encoding, Content-ID, Content-Language, Content-Length, Content-Location, Content-MD5, Content-Range, Content-Script-Type, Content-Security-Policy, Content-Style-Type, Content-Transfer-Encoding, Content-Type, Content-Version, Cookie, Cost, DAV, DELETE, DNT, DPR, Date, Default-Style, Delta-Base, Depth, Derived-From, Destination, Differential-ID, Digest, ETag, Expect, Expires, Ext, From, GET, GetProfile, HEAD, HTTP-date, Host, IM, If, If-Match, If-Modified-Since, If-None-Match, If-Range, If-Unmodified-Since, Keep-Alive, Label, Last-Event-ID, Last-Modified, Link, Location, Lock-Token, MIME-Version, Man, Max-Forwards, Media-Range, Message-ID, Meter, Negotiate, Non-Compliance, OPTION, OPTIONS, OWS, Opt, Optional, Ordering-Type, Origin, Overwrite, P3P, PEP, PICS-Label, POST, PUT, Pep-Info, Permanent, Position, Pragma, ProfileObject, Protocol, Protocol-Query, Protocol-Request, Proxy-Authenticate, Proxy-Authentication-Info, Proxy-Authorization, Proxy-Features, Proxy-Instruction, Public, RWS, Range, Referer, Refresh, Resolution-Hint, Resolver-Location, Retry-After, Safe, Sec-Websocket-Extensions, Sec-Websocket-Key, Sec-Websocket-Origin, Sec-Websocket-Protocol, Sec-Websocket-Version, Security-Scheme, Server, Set-Cookie, Set-Cookie2, SetProfile, SoapAction, Status, Status-URI, Strict-Transport-Security, SubOK, Subst, Surrogate-Capability, Surrogate-Control, TCN, TE, TRACE, Timeout, Title, Trailer, Transfer-Encoding, UA-Color, UA-Media, UA-Pixels, UA-Resolution, UA-Windowpixels, URI, Upgrade, User-Agent, Variant-Vary, Vary, Version, Via, Viewport-Width, WWW-Authenticate, Want-Digest, Warning, Width, X-Content-Duration, X-Content-Security-Policy, X-Content-Type-Options, X-CustomHeader, X-DNSPrefetch-Control, X-Forwarded-For, X-Forwarded-Port, X-Forwarded-Proto, X-Frame-Options, X-Modified, X-OTHER, X-PING, X-PINGOTHER, X-Powered-By, X-Requested-With",
		AllowCredentials: true,
	}
}

func CORSMiddleware(corsConfig CorsConfig) gin.HandlerFunc {

	return func(c *gin.Context) {
		if len(corsConfig.AllowOrigin) > 0 {
			c.Writer.Header().Set("Access-Control-Allow-Origin", corsConfig.AllowOrigin)
		}

		if len(corsConfig.AllowHeaders) > 0 {
			c.Writer.Header().Set("Access-Control-Allow-Headers", corsConfig.AllowHeaders)
		}

		if len(corsConfig.AllowMethods) > 0 {
			c.Writer.Header().Set("Access-Control-Allow-Methods", corsConfig.AllowMethods)
		}

		if corsConfig.AllowCredentials {
			c.Writer.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(corsConfig.AllowCredentials))

		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
