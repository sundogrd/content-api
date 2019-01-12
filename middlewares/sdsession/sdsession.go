package sdsession

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

// TODO: 驱动器加锁单例

//定义驱动列表
var SessionDrivers = make(map[string]Store)

//驱动注册器
func Register(name string, store Store) {
	if store == nil {
		fmt.Printf("[middlewares/sdsession] 注册存储驱动错误")

	}
	if _, dup := SessionDrivers[name]; dup {
		fmt.Printf("[middlewares/sdsession] 已经注册过了")

	}
	SessionDrivers[name] = store
	fmt.Printf("[middlewares/sdsession] 驱动器注册成功")

}

func Middleware(config string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取或者设置cookie 中的session_id
		session := SessionStart(config, c.Request, c.Writer)
		c.Set("session", session)
		c.Next()
	}
}

type SessionConfig struct {
	CookieName      string `json:"cookieName"`
	EnableSetCookie bool   `json:"enableSetCookie,omitempty"`
	Secure          bool   `json:"secure"`
	CookieLifeTime  int    `json:"cookieLifeTime"`
	Domain          string `json:"domain"`
	StoreDriver     string `json:"storeDriver"`
}

//定义session保存接口
type Store interface {
	Get(k string, sid string) interface{}
	Set(k string, v interface{}, sid string) error
}

type session struct {
	Session_id  string
	LiftTime    int
	storeDriver string
	store       Store
}

func (s *session) Get(k string) interface{} {
	return s.store.Get(k, s.Session_id)
}
func (s *session) Set(k string, v interface{}) {
	s.store.Set(k, v, s.Session_id)
}

// TODO: 雪花算法

func buildId() (string, error) {
	b := make([]byte, 32)
	n, err := rand.Read(b)
	if n != len(b) || err != nil {
		return "", fmt.Errorf("随机数生成失败！")
	}
	return hex.EncodeToString(b), nil
}
func GetSession(c *gin.Context) (s *session) {

	tmp := c.MustGet("session")
	s = tmp.(*session)
	return
}
func SessionStart(config string, r *http.Request, w gin.ResponseWriter) *session {
	var session_id string
	c := new(SessionConfig)
	err := json.Unmarshal([]byte(config), c)
	//根据配置文件获取注册过的存储驱动
	store, ok := SessionDrivers[c.StoreDriver]
	if !ok {
		fmt.Printf("[middlewares/sdsession] 没有找到该注册驱动" + c.StoreDriver)
		return nil
	}
	//根据配置生成session_id

	if err != nil {
		fmt.Println(err)
		return nil
	}
	cookie, err := r.Cookie(c.CookieName)
	if err != nil || cookie.Value == "" {
		session_id, err := buildId()
		if err != nil {
			fmt.Println(err)
			return nil
		}
		cookie = &http.Cookie{Name: c.CookieName,
			Value:    url.QueryEscape(session_id),
			Path:     "/",
			HttpOnly: true,
			Secure:   c.Secure,
			Domain:   c.Domain}
		if c.CookieLifeTime >= 0 {
			cookie.MaxAge = c.CookieLifeTime
		}

		http.SetCookie(w, cookie)
		r.AddCookie(cookie)
	} else {
		session_id, err = url.QueryUnescape(cookie.Value)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		cookie = &http.Cookie{Name: c.CookieName,
			Value:    url.QueryEscape(session_id),
			Path:     "/",
			HttpOnly: true,
			Secure:   c.Secure,
			Domain:   c.Domain}
		if c.CookieLifeTime >= 0 {
			cookie.MaxAge = c.CookieLifeTime
		}

		http.SetCookie(w, cookie)
		r.AddCookie(cookie)

	}
	//传入初始值初始化session
	session := new(session)
	session.Session_id = session_id
	session.LiftTime = c.CookieLifeTime
	session.store = store
	return session
}
