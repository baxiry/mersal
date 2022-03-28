package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/bashery/im"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var mt sync.Mutex

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//if r.Header.Get("Origin")!="http://"+r.Host {http.Error(w,"Origin not allowed",-1);return}

	conn, err := websocket.Upgrade(w, r, w.Header(), 2, 2) //1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", 404)
	}
	go im.ServeMessages(conn)
}

func main() {
	db = setdb()
	defer db.Close()

	http.HandleFunc("/ws", wsHandler)
	go func() {
		panic(http.ListenAndServe(":8080", nil))
	}()

	e := echo.New()
	// middleware
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	// routers
	e.POST("/signup", signup)

	e.POST("/login", login)

	e.Logger.Fatal(e.Start(":1323"))
}

func simple(c echo.Context) error {

	fmt.Println(c.Request().Body)
	username := c.FormValue("username")
	password := c.FormValue("password")
	fmt.Println("username :", username)
	fmt.Println("password :", password)

	return c.String(http.StatusOK, "ok")
}

// login if user info is correct
func Login(c echo.Context) error {

	femail := c.FormValue("email")
	fpass := c.FormValue("password")
	userid, username, email, pass := getUsername(femail)

	if pass == fpass && femail == email {
		//userSession[email] = username
		setSession(c, username, userid)
		return c.Redirect(http.StatusSeeOther, "/") // 303 code
		// TODO redirect to latest page
	}

	data := make(map[string]interface{}, 2)
	data["userid"] = nil
	data["error"] = "user information is not correct"
	return c.Render(200, "login.html", data)
}
