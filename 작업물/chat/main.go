package main // import "github.com/sneakstarberry/chat"

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"gopkg.in/mgo.v2"
)

const socketBufferSize = 1024

var (
	renderer     *render.Render
	mongoSession *mgo.Session
)

var (
	upgrader = &websocket.Upgrader{
		ReadBufferSize:  socketBufferSize,
		WriteBufferSize: socketBufferSize,
	}
)

const (
	sessionKey    = "simple_chat_session"
	sessionSecret = "simple_chat_session_secret"
)

func init() {
	// 렌더러 생성
	renderer = render.New()

	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	mongoSession = s
}

func main() {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		renderer.HTML(w, http.StatusOK, "index", map[string]string{"title": "Simple Chat"})
	})

	router.GET("/login", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		renderer.HTML(w, http.StatusOK, "login", nil)
	})

	router.GET("/logout", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		sessions.GetSession(r).Delete(currentUserKey)
		http.Redirect(w, r, "/login", http.StatusFound)
	})

	router.GET("/auth/:action/:provider", loginHandler)
	router.POST("/rooms", createRoom)
	router.GET("/rooms", retrieveRooms)

	router.GET("/rooms/:id/messages", retrieveMessages)

	router.GET("/ws/:room_id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		socket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal("ServeHTTP:", err)
			return
		}
		newClient(socket, ps.ByName("room_id"), GetCurrentUser(r))
	})

	n := negroni.Classic()
	store := cookiestore.New([]byte(sessionSecret))
	n.Use(sessions.Sessions(sessionKey, store))

	n.Use(LoginRequired("/login", "/auth"))
	n.UseHandler(router)

	n.Run(":3000")
}
