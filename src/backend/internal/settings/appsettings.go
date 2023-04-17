package settings

import (
	"time"

	"github.com/chmike/securecookie"
)

const CookieLifeTime = int(time.Hour) * 24 * 7
const CookieSessionName = "session_cookie"

var Location string
var CookieHashKey []byte
var CookieIssued *securecookie.Obj

const TimeStoreInCache = time.Minute * 5
