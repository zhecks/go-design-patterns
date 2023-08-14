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
type LoginTemplate struct {
}

func (lt *LoginTemplate) Login(lm LoginMode) {

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

func (wl *WokerLogin) Match(lm, dbLm LoginMode) bool {
	if lm.LoginID == dbLm.LoginID && wl.EncrypPwd(lm.Password) == dbLm.Password {
		return true
	}
	return false
}

type NormalLogin struct {
	LoginTemplate
}
