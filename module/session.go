package module

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/sha3"
)

type Session struct {
	Uid     int64
	Session int64
	Time    time.Time
	User    *User
}

var (
	salt = `` // TODO 写配置里

	CookieKey = `tango_sesssion`

	patternCookie, _ = regexp.Compile(`[a-f0-9]{64},[0-9]{10},[0-9]{1,20},[0-9]{1,20}`)

	exampleCookie = `c0150be3447a23ddfbd0be5a8936131037355738be3c0a634cf9a0a713178501,1489645249,123,321`

	cookieExpire = time.Unix(2147483647, 0)
)

func (this *Session) GetUser() *User {
	if this.User == nil {
		this.User = &User{
			Id: this.Uid,
		}
	}
	return this.User
}

func (this *Session) MakeCookie() (c *http.Cookie) {
	time := time.Now()
	hash := SessionMakeHash(this.Uid, this.Session, time)
	if this.Uid < 1 {
		return
	}

	c = &http.Cookie{
		Name:     CookieKey,
		Value:    fmt.Sprintf(`%x,%d,%d,%d`, hash, time.Unix(), this.Uid, this.Session),
		Path:     `/`,
		Domain:   `goweb.funplus.io`, // TODO 写配置里
		Expires:  cookieExpire,
		Secure:   true,
		HttpOnly: true,
	}
	return
}

func (this *Session) Valid() bool {
	return this.Uid > 0
}

func SetSessionSalt(s string) {
	salt = s
}

func SessionParse(w http.ResponseWriter, r *http.Request) (s *Session) {
	c, err := r.Cookie(CookieKey)
	if err != nil {
		return
	}
	s, _ = SessionImport(c.Value)

	if s != nil {
		dur := s.Time.Sub(time.Now()).Minutes()
		if dur > 5 {
			http.SetCookie(w, s.MakeCookie())
		}
	}
	return
}

func SessionImport(sCookie string) (user *Session, err error) {
	//user = nil
	if !patternCookie.MatchString(sCookie) {
		errors.New(`valid cookie format`)
		return
	}
	s := strings.Split(sCookie, `,`)

	iTime, _ := strconv.ParseInt(s[1], 10, 64)
	time := time.Unix(iTime, 0)

	uid, _ := strconv.ParseInt(s[2], 10, 64)
	session, _ := strconv.ParseInt(s[3], 10, 64)

	// fmt.Println(time, uid, session)

	hash := SessionMakeHash(uid, session, time)
	// fmt.Printf("%x\n", hash)

	hashCheck, _ := hex.DecodeString(s[0])

	if bytes.Equal(hashCheck, hash) {
		user = &Session{}
		user.Uid = uid
		user.Session = session
		user.Time = time
		return
	}
	return
}

func SessionMakeHash(uid int64, session int64, time time.Time) []byte {

	// fmt.Println(`input`, uid, time, session)

	hash := sha3.New256()
	binary.Write(hash, binary.BigEndian, uid)
	binary.Write(hash, binary.BigEndian, session)
	binary.Write(hash, binary.BigEndian, time.Unix())

	hash.Write([]byte(salt))

	return hash.Sum(nil)
}
