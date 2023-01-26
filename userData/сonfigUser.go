package userData

import (
	"errors"
	"fmt"
	"httpTable/user_data"
	"httpTable/user_session"
)

// UsersInitializationData запускается при запуске приложения создает map Имя пользователя-пароль
func UsersInitializationData() {
	user.userVerif = make(map[string]user_data.UserDataVerification)
}

// Get  функция проверяет соответсвие Имени-пользователя -пароля
func Get(name string, password string) (string, error) {
	var err error
	user.lock.Lock()
	defer user.lock.Unlock()
	if valueMap, ok := user.userVerif[name]; ok {
		if password == valueMap.Password {
			return name, nil
			fmt.Printf("Совпало полностью")
		} else {
			fmt.Printf("Необходимо создать новый токен")
			err = errors.New("Такого пароля нет в базе данных ошибка метода Get userData")
			return "", err
		}
	} else {
		err = errors.New("Такого пользователя нет в базе данных ошибка метода Get userData")
		return "", err

	}
	return name, err
}

// Add функция добавляет нового пользователя и его пароль
func Add(name string, password string) (string, error) {
	var err error
	user.lock.Lock()
	defer user.lock.Unlock()
	if _, ok := user.userVerif[name]; ok {
		err := errors.New("Пользователь с таким именем уже существует, укажите другое имя")
		return "", err
	} else {
		user.userVerif[name] = user_data.UserDataVerification{password}
		return name, nil
	}

	return name, err
}

// Delete функция удаляет пользователя и его пароль из БД
func Delete(name string, password string) (string, error) {
	var err error
	user.lock.Lock()
	defer user.lock.Unlock()
	if valueMap, ok := user.userVerif[name]; ok {
		if password == valueMap.Password {
			delete(user.userVerif, name)
			user_session.DeleteToken(name)
			fmt.Printf("Совпало полностью")
			return name, nil

		} else {
			err = errors.New("Такого пароля нет в базе данных ошибка метода Delete userData")
			return "", nil
		}
	} else {
		err := errors.New("Пользователь с таким именем уже существует, укажите другое имя")
		return name, err
	}
	return name, err
}
