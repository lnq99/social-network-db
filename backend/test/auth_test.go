package test

import (
	"testing"
)

func TestAuthToken(t *testing.T) {
	token, err := authManager.CreateToken("3")
	t.Log(token)
	if err != nil {
		t.Error(err)
	}
}

func TestAuthPassword(t *testing.T) {
	testcases := [][]string{
		{"easypass", "q%3@QYw#", "DLWhDfQXMMLHaitFIE7v3XpCbgg="},
	}
	for _, v := range testcases {
		if !authManager.ComparePassword(v[0], v[1], v[2]) {
			t.Error("Error compare password")
		}
	}
}
