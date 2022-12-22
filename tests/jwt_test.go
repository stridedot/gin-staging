package tests

import (
	"go_code/gintest/pkg/jwt"
	"testing"
)

func testGenJWT(t *testing.T) {
	token := jwt.GenToken(1, "test")
	t.Log(token)
}

func TestParseToken(t *testing.T) {
	claims, _ := jwt.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJpc3MiOiLpmoblhqzlvLoiLCJleHAiOjE2Njk2ODA5Njl9.VyCKS-xFDJ3Nu9ezSFA5I3JqXTQ2HQ6ESUFVc9V49sI")
	t.Log(claims)
}
