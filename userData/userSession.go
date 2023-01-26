package userData

import (
	"httpTable/user_data"
	"sync"
)

// UserVerification структура хранения Имени пользователя-пароля
type UserVerification struct {
	userVerif map[string]user_data.UserDataVerification
	lock      sync.RWMutex
}

//переменная доступа к структуре UserVerification
var user UserVerification
