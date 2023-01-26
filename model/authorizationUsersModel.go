package model

import (
	"fmt"
	"httpTable/userData"
	"httpTable/user_session"
)

// TableModel труктура используется для конструктора контроллер
type AuthorizationUsersModel struct {
	module *TableModel
}

// NewTableModel конструктор контроллера, возращающий экземпляр структуры TableModel
// со свойством model контроллера модели
func NewAuthorizationUsersModel() *AuthorizationUsersModel {
	return &AuthorizationUsersModel{
		NewTableModel(),
	}
}

// GetAuthorizationUsers метод модели по входу на сайт
func (aut *AuthorizationUsersModel) GetAuthorizationUsers(name string, password string) (string, error) {
	name, err := userData.Get(name, password)
	if err != nil {
		fmt.Println("Ошибка в имени или пароле пользователя метода GetAuthorizationUsers модели", err)
	}
	return name, err
}

// AddNewUser метод модели по добавлению нового пользователя
func (aut *AuthorizationUsersModel) AddNewUser(name string, password string, token string) (string, error) {
	var err error
	if token == `` {
		name, err := userData.Add(name, password)
		if err != nil {
			fmt.Println("Ошибка в имени или пароле пользователя метода AddNewUser модели", err)
		}
		return name, err
	} else {
		user_session.SetFirstNameToken(name, token)
		user_session.Set(token, 0, 0)
	}
	return name, err
}

// DeleteUser метод модели по удалению пользователя, после ввода пароля и имени пользователя
func (aut *AuthorizationUsersModel) DeleteUser(name string, password string) (string, error) {
	name, err := userData.Delete(name, password)
	if err != nil {
		fmt.Println("Ошибка в имени или пароле пользователя метода AddNewUser модели", err)
	}
	return name, err
}

// LogOutUser метод модели по удалению текущей сессии пользователя
func (aut *AuthorizationUsersModel) LogOutUser(token string) {
	user_session.LogOutToken(token)

}
