package api

import (
	"time"

	"github.com/chmike/securecookie"
)

const CookieSessionName = "session_cookie"
const CookieLifeTime = int(time.Hour) * 24
const Location = "localhost"

var CookieHashKey = []byte("ABcAsdfawe1241!1q2142@fgsv.,h/@!")
var CookieIssued *securecookie.Obj
