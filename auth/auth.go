package auth

import (
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// GetSession return username & userid as session's user
func GetSession(c echo.Context) (string, int, error) {
	sess, _ := session.Get("session", c)
	if sess.Values["userid"] == nil {
		return "", -1, fmt.Errorf("no session")
	}
	return sess.Values["username"].(string), sess.Values["userid"].(int), nil
	//return sess.Values["userid"].(int), nil
}

// newSession creates new session
func NewSession(c echo.Context, userid int) {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 10, // 10 minutes of session,
		HttpOnly: true,    // no websocket or any protocol else
	}
	// sess.Values["username"] = username
	sess.Values["userid"] = userid
	sess.Save(c.Request(), c.Response())
}
