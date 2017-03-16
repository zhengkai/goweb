package module

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/sha3"
)

type Session struct {
	uid     int64
	session int64
	time    time.Time
}

var (
	salt = ``

	patternCookie, _ = regexp.Compile(`[a-f0-9]{64},[0-9]{10},[0-9]{1,20},[0-9]{1,20}`)

	exampleCookie = `c0150be3447a23ddfbd0be5a8936131037355738be3c0a634cf9a0a713178501,1489645249,123,321`
)

func (this *Session) Auth(s string) {
}

func (this *Session) MakeCookie() string {
	time := time.Now()
	hash := SessionMakeHash(this.uid, this.session, time)
	if this.uid < 1 {
		return ``
	}

	return fmt.Sprintf(`%x,%d,%d,%d`, hash, time.Unix(), this.uid, this.session)
}

func (this *Session) Valid() bool {
	return this.uid > 0
}

func SetSessionSalt(s string) {
	salt = s
}

func SessionImport(sCookie string) (user *Session, err error) {
	user = nil
	if !patternCookie.MatchString(sCookie) {
		errors.New(`valid cookie format`)
		return
	}
	s := strings.Split(sCookie, `,`)

	iTime, _ := strconv.ParseInt(s[1], 10, 64)
	time := time.Unix(iTime, 0)

	uid, _ := strconv.ParseInt(s[2], 10, 64)
	session, _ := strconv.ParseInt(s[3], 10, 64)

	fmt.Println(time, uid, session)

	hash := SessionMakeHash(uid, session, time)
	fmt.Printf("%x\n", hash)

	hashCheck, _ := hex.DecodeString(s[0])

	if bytes.Equal(hashCheck, hash) {
		user = &Session{}
		user.uid = uid
		user.session = session
		user.time = time
		return
	}
	return
}

func SessionMakeHash(uid int64, session int64, time time.Time) []byte {

	fmt.Println(`input`, uid, time, session)

	hash := sha3.New256()
	binary.Write(hash, binary.BigEndian, uid)
	binary.Write(hash, binary.BigEndian, session)
	binary.Write(hash, binary.BigEndian, time.Unix())

	hash.Write([]byte(salt))

	return hash.Sum(nil)
}

func init() {
	fmt.Println(`session`)
	fmt.Println(exampleCookie)
	valid, _ := SessionImport(exampleCookie)
	fmt.Println(valid)
	fmt.Println(valid.MakeCookie())
	valid2, _ := SessionImport(exampleCookie + `a`)
	fmt.Println(valid2)
}
