package main //import "github.com/sneakstarberry/web1"

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"

	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
)

const (
	// 애플리케이션에서 사용할 세션의 키 정보
	sessionKey    = "simple_chat_session"
	sessionSecret = "simple_chat_session_secret"
)

var renderer *render.Render

func init() {
	renderer = render.New()
}

func main() {
	r := httprouter.New()

	r.GET("/", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		renderer.HTML(w, http.StatusOK, "index", map[string]string{"title": "Simple"})
	})

	r.GET("/login", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		renderer.HTML(w, http.StatusOK, "login", nil)
	})

	r.GET("/logout", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		sessions.GetSession(req).Delete(currentUserKey)
		http.Redirect(w, req, "/login", http.StatusFound)
	})

	n := negroni.Classic()
	store := cookiestore.New([]byte(sessionSecret))
	n.Use(sessions.Sessions(sessionKey, store))
	n.UseHandler(r)

	n.Run(":5000")
}

////////////////////////////////////////////////

const (
	currentUserKey  = "oauth2_current_user" // 세션에 저장되는 CurrentUser의 키
	sessionDuration = time.Hour             // 로그인 세션 유지 시간
)

type User struct {
	Uid       string    `json:"uid"`
	Name      string    `json:"name"`
	Email     string    `json:"user"`
	AvatarUrl string    `json:"avatar_url"`
	Expired   time.Time `json:"expired"`
}

func (u *User) Valid() bool {
	// 현재 시간 기준으로 만료 시간 확인
	return u.Expired.Sub(time.Now()) > 0
}

func (u *User) Refresh() {
	// 만료 시간 시간 연장
	u.Expired = time.Now().Add(sessionDuration)
}

func GetCurrentUser(r *http.Request) *User {
	// 세션에서 CurrentUser 정보를 가져옴
	s := sessions.GetSession(r)

	if s.Get(currentUserKey) == nil {
		return nil
	}

	data := s.Get(currentUserKey).([]byte)
	var u User
	json.Unmarshal(data, &u)
	return &u
}

func SetCurrentUser(r *http.Request, u *User) {
	if u != nil {
		// CurrentUser 만료 시간 갱신
		u.Refresh()
	}

	// 세션에 CurrentUser 정보를 json으로 저장
	s := sessions.GetSession(r)
	val, _ := json.Marshal(u)
	s.Set(currentUserKey, val)
}
