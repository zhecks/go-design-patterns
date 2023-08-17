package templatemethod

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type LoginMode struct {
	LoginID  string
	Password string
}

type Loginer interface {
	Login(lm LoginMode) bool
}

type implement interface {
	FindLoginUser(loginID string) *LoginMode
	Match(lm LoginMode, dbLm LoginMode) bool
}
type LoginTemplate struct {
	implement
}

func NewTemplate(impl implement) Loginer {
	return &LoginTemplate{
		implement: impl,
	}
}

func (lt *LoginTemplate) Login(lm LoginMode) bool {
	dbLm := lt.implement.FindLoginUser(lm.LoginID)
	if dbLm == nil {
		return false
	}
	return lt.implement.Match(lm, *dbLm)
}

func (lt *LoginTemplate) FindLoginUser(_ string) *LoginMode {
	fmt.Println("default find login user")
	return nil
}

func (lt *LoginTemplate) EncrypPwd(pwd string) string {
	hash := sha256.New()
	hash.Write([]byte(pwd))
	return hex.EncodeToString(hash.Sum(nil))
}

func (lt *LoginTemplate) Match(lm LoginMode, dbLm LoginMode) bool {
	if lm.LoginID == dbLm.LoginID && lm.Password == dbLm.Password {
		return true
	}
	return false
}

type WokerLogin struct {
	LoginTemplate
}

func (wl *WokerLogin) FindLoginUser(loginID string) *LoginMode {
	return &LoginMode{
		LoginID:  loginID,
		Password: wl.EncrypPwd("123"),
	}
}

func (wl *WokerLogin) Match(lm, dbLm LoginMode) bool {
	if lm.LoginID == dbLm.LoginID && wl.EncrypPwd(lm.Password) == dbLm.Password {
		return true
	}
	return false
}

type NormalLogin struct {
	LoginTemplate
}

func (nl *NormalLogin) FindLoginUser(loginID string) *LoginMode {
	return &LoginMode{
		LoginID:  loginID,
		Password: "123",
	}
}
