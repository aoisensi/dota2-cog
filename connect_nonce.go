package main

import (
	"time"

	"github.com/google/uuid"
)

var nonces = make(map[int64]ConnectNonce)

// ConnectNonce is connect nonce
type ConnectNonce struct {
	DiscordID int64
	Token     string
	Expire    time.Time
}

func findNonce(token string) int64 {
	if token == "" {
		return 0
	}
	for _, nonce := range nonces {
		if token == nonce.Token {
			return nonce.DiscordID
		}
	}
	return 0
}

func makeConnectNonce(id int64) string {
	token := uuid.New().String()
	expire := time.Now().Add(time.Minute * 5)
	nonces[id] = ConnectNonce{
		DiscordID: id,
		Token:     token,
		Expire:    expire,
	}
	return token
}

func deleteNonce(id int64) {
	delete(nonces, id)
}

func deleteExpiredNonce() {
	now := time.Now()
	for id, nonce := range nonces {
		if now.Before(nonce.Expire) {
			deleteNonce(id)
		}
	}
}
