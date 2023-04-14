package api

import (
	"time"

	"github.com/chmike/securecookie"
)

const CookieSessionName = "session_cookie"
const CookieLifeTime = int(time.Hour) * 24

var Location string
var CookieHashKey []byte
var CookieIssued *securecookie.Obj
