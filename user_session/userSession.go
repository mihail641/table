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

// Token структура токена
type Token struct {
	id   int
	key  string
	lock sync.RWMutex
}

var token Token

const key = "my new token for table api"
