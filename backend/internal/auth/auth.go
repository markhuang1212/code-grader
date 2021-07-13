package auth

import (
	"bufio"
	"errors"
	"io"
	"strings"
	"sync"
)

type AuthController struct {
	accounts map[string]string
	lock     sync.RWMutex
}

func (ac *AuthController) Add(userName string, password string) {
	ac.lock.Lock()
	defer ac.lock.Unlock()

	ac.accounts[userName] = password
}

func (ac *AuthController) Auth(userName string, password string) bool {
	ac.lock.RLock()
	defer ac.lock.RUnlock()

	psd, ok := ac.accounts[userName]
	return ok && psd == password
}

func NewAuthController(accountInfo io.Reader) (*AuthController, error) {
	ret := AuthController{
		accounts: make(map[string]string),
	}
	scanner := bufio.NewScanner(accountInfo)

	for {
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()
		token := strings.Split(text, " ")
		if len(token) != 2 {
			return nil, errors.New("wrong format")
		}
		ret.Add(token[0], token[1])
	}

	return &ret, nil

}
