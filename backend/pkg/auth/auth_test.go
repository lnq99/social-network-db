//go:build unit

package auth

import (
	"strings"
	"testing"
)

var authManager = InitManager("id", "pw")

func TestAuthToken(t *testing.T) {
	token, err := authManager.CreateToken("3")
	//t.Log(token)
	if err != nil || len(strings.Split(token, ".")) != 3 {
		t.Error(err)
	}
}

func TestComparePassword(t *testing.T) {
	testcases := [][]string{
		{"easypass", "q%3@QYw#", "DLWhDfQXMMLHaitFIE7v3XpCbgg="},
	}
	for _, v := range testcases {
		if !authManager.ComparePassword(v[0], v[1], v[2]) {
			t.Error("Error compare password")
		}
	}
}

func TestGeneratePassword(t *testing.T) {
	password := "1212321"
	manager := InitManager("dsfsdaf", "secret")
	s, h := manager.GetHashSalt(password)
	//t.Log(s, h)
	if !manager.ComparePassword(password, s, h) {
		t.Error("Error generate password")
	}
}

//func TestGenerate(t *testing.T) {
//	for _, p := range []string{"pw1", "pw2", "pw3"} {
//		println(authManager.GetHashSalt(p))
//	}
//}
