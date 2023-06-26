package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	defaultSchemas = []string{
		"http://",
		"https://",
	}

	AccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	AccessControlExposeHeaders    = "Access-Control-Expose-Headers"
	AccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	AccessControlAllowMethods     = "Access-Control-Allow-Methods"
	AccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	AccessControlRequestMethod    = "Access-Control-Request-Method"
	AccessControlRequestHeaders   = "Access-Control-Request-Headers"

	Vary   = "Vary"
	Origin = "Origin"
	All    = "*"
	Allow  = "true"
)

type CorsConfig struct {
	AllowOrigins     []string
	AllowHeaders     []string
	AllowMethods     []string
	AllowCredentials bool
	ExposeHeaders    []string
}

func DefaultCorsConfig() CorsConfig {
	return CorsConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Accept-CH", "Accept-Charset", "Accept-Datetime", "Accept-Encoding", "Accept-Ext", "Accept-Features", "Accept-Language", "Accept-Params", "Accept-Ranges", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Origin", "Access-Control-Expose-Headers", "Access-Control-Max-Age", "Access-Control-Request-Headers", "Access-Control-Request-Method", "Age", "Allow", "Alternates", "Authentication-Info", "Authorization", "C-Ext", "C-Man", "C-Opt", "C-PEP", "C-PEP-Info", "CONNECT", "Cache-Control", "Compliance", "Connection", "Content-Base", "Content-Disposition", "Content-Encoding", "Content-ID", "Content-Language", "Content-Length", "Content-Location", "Content-MD5", "Content-Range", "Content-Script-Type", "Content-Security-Policy", "Content-Style-Type", "Content-Transfer-Encoding", "Content-Type", "Content-Version", "Cookie", "Cost", "DAV", "DELETE", "DNT", "DPR", "Date", "Default-Style", "Delta-Base", "Depth", "Derived-From", "Destination", "Differential-ID", "Digest", "ETag", "Expect", "Expires", "Ext", "From", "GET", "GetProfile", "HEAD", "HTTP-date", "Host", "IM", "If", "If-Match", "If-Modified-Since", "If-None-Match", "If-Range", "If-Unmodified-Since", "Keep-Alive", "Label", "Last-Event-ID", "Last-Modified", "Link", "Location", "Lock-Token", "MIME-Version", "Man", "Max-Forwards", "Media-Range", "Message-ID", "Meter", "Negotiate", "Non-Compliance", "OPTION", "OPTIONS", "OWS", "Opt", "Optional", "Ordering-Type", "Origin", "Overwrite", "P3P", "PEP", "PICS-Label", "POST", "PUT", "Pep-Info", "Permanent", "Position", "Pragma", "ProfileObject", "Protocol", "Protocol-Query", "Protocol-Request", "Proxy-Authenticate", "Proxy-Authentication-Info", "Proxy-Authorization", "Proxy-Features", "Proxy-Instruction", "Public", "RWS", "Range", "Referer", "Refresh", "Resolution-Hint", "Resolver-Location", "Retry-After", "Safe", "Sec-Websocket-Extensions", "Sec-Websocket-Key", "Sec-Websocket-Origin", "Sec-Websocket-Protocol", "Sec-Websocket-Version", "Security-Scheme", "Server", "Set-Cookie", "Set-Cookie2", "SetProfile", "SoapAction", "Status", "Status-URI", "Strict-Transport-Security", "SubOK", "Subst", "Surrogate-Capability", "Surrogate-Control", "TCN", "TE", "TRACE", "Timeout", "Title", "Trailer", "Transfer-Encoding", "UA-Color", "UA-Media", "UA-Pixels", "UA-Resolution", "UA-Windowpixels", "URI", "Upgrade", "User-Agent", "Variant-Vary", "Vary", "Version", "Via", "Viewport-Width", "WWW-Authenticate", "Want-Digest", "Warning", "Width", "X-Content-Duration", "X-Content-Security-Policy", "X-Content-Type-Options", "X-CustomHeader", "X-DNSPrefetch-Control", "X-Forwarded-For", "X-Forwarded-Port", "X-Forwarded-Proto", "X-Frame-Options", "X-Modified", "X-OTHER", "X-PING", "X-PINGOTHER", "X-Powered-By", "X-Requested-With"},
		AllowCredentials: false,
	}
}

func (c CorsConfig) validateAllowedSchemas(origin string) bool {
	for _, schema := range defaultSchemas {
		if strings.HasPrefix(origin, schema) {
			return true
		}
	}
	return false
}

func (corsConfig CorsConfig) Validate() error {
	for _, origin := range corsConfig.AllowOrigins {
		if !strings.Contains(origin, "*") && !corsConfig.validateAllowedSchemas(origin) {
			return errors.New("bad origin: origins must contain '*' or include " + strings.Join(defaultSchemas, ","))
		}
	}
	return nil
}

type corsConfig struct {
	allowOrigins     []string
	allowCredentials bool
	allowAllOrigins  bool
	normalHeaders    http.Header
	preflightHeaders http.Header
}

func normalizeToLower(values []string) []string {
	if values == nil {
		return nil
	}
	distinctKeyValues := make(map[string]bool, len(values))
	normalizedStrings := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		value = strings.ToLower(value)
		if _, distinctedValue := distinctKeyValues[value]; !distinctedValue {
			normalizedStrings = append(normalizedStrings, value)
			distinctKeyValues[value] = true
		}
	}
	return normalizedStrings
}

func normalizeToUpper(values []string) []string {
	if values == nil {
		return nil
	}
	distinctKeyValues := make(map[string]bool, len(values))
	normalizedStrings := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		value = strings.ToUpper(value)
		if _, distinctedValue := distinctKeyValues[value]; !distinctedValue {
			normalizedStrings = append(normalizedStrings, value)
			distinctKeyValues[value] = true
		}
	}
	return normalizedStrings
}

func generateNormalHeaders(c CorsConfig, allowAllOrigins bool) http.Header {
	headers := make(http.Header)
	if c.AllowCredentials {
		headers.Set(AccessControlAllowCredentials, Allow)
	}
	if len(c.ExposeHeaders) > 0 {
		headers.Set(AccessControlExposeHeaders, strings.Join(normalizeToLower(c.ExposeHeaders), ","))
	}
	if allowAllOrigins {
		headers.Set(AccessControlAllowOrigin, All)
	} else {
		headers.Set(Vary, Origin)
	}
	return headers
}

func generatePreflightHeaders(c CorsConfig, allowAllOrigins bool) http.Header {
	headers := make(http.Header)
	if c.AllowCredentials {
		headers.Set(AccessControlAllowCredentials, Allow)
	}
	if len(c.AllowMethods) > 0 {
		headers.Set(AccessControlAllowMethods, strings.Join(normalizeToUpper(c.AllowMethods), ","))
	}
	if len(c.AllowHeaders) > 0 {
		allowHeaders := normalizeToLower(c.AllowHeaders)
		headers.Set(AccessControlAllowHeaders, strings.Join(allowHeaders, ","))
	}
	if allowAllOrigins {
		headers.Set(AccessControlAllowOrigin, All)
	} else {
		headers.Add(Vary, Origin)
		headers.Add(Vary, AccessControlRequestMethod)
		headers.Add(Vary, AccessControlRequestHeaders)
	}
	return headers
}

func initCors(config CorsConfig) *corsConfig {
	if err := config.Validate(); err != nil {
		panic(err.Error())
	}

	allowAllOrigins := false
	for _, origin := range config.AllowOrigins {
		if origin == "*" {
			allowAllOrigins = true
		}
	}

	return &corsConfig{
		allowAllOrigins:  allowAllOrigins,
		allowCredentials: config.AllowCredentials,
		allowOrigins:     normalizeToLower(config.AllowOrigins),
		normalHeaders:    generateNormalHeaders(config, allowAllOrigins),
		preflightHeaders: generatePreflightHeaders(config, allowAllOrigins),
	}
}

func CORSMiddleware(corsConfig CorsConfig) gin.HandlerFunc {
	corsSetup := initCors(corsConfig)
	return func(c *gin.Context) {
		origin := c.Request.Header.Get(Origin)
		if len(origin) == 0 {
			// not apply if is not a CORS request
			return
		}
		host := c.Request.Host

		if origin == "http://"+host || origin == "https://"+host {
			return
		}

		if !corsSetup.validateOrigin(origin) {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		if c.Request.Method == http.MethodOptions {
			corsSetup.preflightRequest(c)
			defer c.AbortWithStatus(http.StatusNoContent)
		} else {
			corsSetup.normalRequest(c)
		}

		if !corsSetup.allowAllOrigins {
			c.Header(AccessControlAllowOrigin, origin)
		}
	}
}

func (cors *corsConfig) validateOrigin(origin string) bool {
	if cors.allowAllOrigins {
		return true
	}
	for _, value := range cors.allowOrigins {
		if value == origin {
			return true
		}
	}
	return false
}

func (cors *corsConfig) preflightRequest(c *gin.Context) {
	header := c.Writer.Header()
	for key, value := range cors.preflightHeaders {
		header[key] = value
	}
}

func (cors *corsConfig) normalRequest(c *gin.Context) {
	header := c.Writer.Header()
	for key, value := range cors.normalHeaders {
		header[key] = value
	}
}
