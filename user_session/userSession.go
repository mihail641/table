package user_session

import (
	"httpTable/table_metadata"
	"sync"
)

// UserSession структура состаящая из map значение токена-ключ, количество строк и столбцов структура
type UserSession struct {
	userInfo map[string]table_metadata.TableMetaData
	lock     sync.RWMutex
}

//переменная для доступа к структуре UserSession
var session UserSession

// UserName структура хранящая Имя пользователя-токены с которых пользователь заходит
type UserName struct {
	infoName map[string][]string
	lock     sync.RWMutex
}

//переменная для доступа к структуре UserName
var userName UserName

// TokenName структура хранящая Токен-имя пользователя
type TokenName struct {
	tokenName map[string]string
	lock      sync.RWMutex
}

//переменная для доступа к структуре TokenName
var tokenName TokenName

// Token структура токена, для инициализации уникального токена
type Token struct {
	id   int
	key  string
	lock sync.RWMutex
}

//переменная для доступа к структуре Token
var token Token

const key = "my new token for table api"
