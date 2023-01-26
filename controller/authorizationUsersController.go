package controller

import (
	"fmt"
	"httpTable/model"
	"httpTable/user_session"
	"net/http"
)

// AuthorizationUsersController cтруктура используется для конструктора контроллер
type AuthorizationUsersController struct {
	modelUser  *model.AuthorizationUsersModel
	modelTable *model.TableModel
}

// NewAuthorizationUsersController конструктор контроллера, возращающий экземпляр структуры AuthorizationUsersController
// со свойствами modelTable и modelUser контроллера модели
func NewAuthorizationUsersController() *AuthorizationUsersController {
	return &AuthorizationUsersController{
		modelUser:  model.NewAuthorizationUsersModel(),
		modelTable: model.NewTableModel(),
	}
}

//readURL приватная функция считывающая имя и пароль пользователя
func readURL(res http.ResponseWriter, req *http.Request) (string, string) {
	req.ParseForm()
	name := req.FormValue("username")
	password := req.FormValue("password")
	return name, password
}

// GetAuthorizationUsers метод контроллера при авторизации
func (aut *AuthorizationUsersController) GetAuthorizationUsers(res http.ResponseWriter, req *http.Request) int {
	//получение имени и пароля пользователя из URL строки
	name, password := readURL(res, req)
	//если пароль и пользователь пустой выводит сообщение
	if name == `` && password == `` {

		messageUser := `<html>
		<head>
		<meta charset="UTF-8">
		<title>Введите свои данные</title>
		</head>
		<body>
		<h2>Введите свои данные</h2>
		<form method="GET" action="auth">
		<label>Имя</label><br>
		<input type="text" name="username" /><br><br>
		<label>Пароль</label><br>
		<input type="password" maxlength="25" size="40" name="password" /><br><br>
		<input type="submit" value="Отправить" />
 <a href="../register">РЕГИСТРАЦИЯ</a>
		</form>
		</body>
		</html>`
		html := []byte(messageUser)
		res.Write(html)
		return 0
	}
	name, err := aut.modelUser.GetAuthorizationUsers(name, password)
	//если пароль и имя пользователя не совпадает выводит сообщение
	if err != nil {
		messageUserError :=
			`<html>
		<head>
		<meta charset="UTF-8">
		<title>Извините, пароль или имя с ошибкой, введите еще раз или пройдите <a href="../register">РЕГИСТРАЦИЮ</a> </title>
		</head>
		<body>
		<h2>Извините, пароль или имя с ошибкой, введите еще раз или пройдите <a href="../register">РЕГИСТРАЦИЮ</a></h2>
		<form method="GET" action="auth">
		<label>Имя</label><br>
		<input type="text" name="username" /><br><br>
		<label>Пароль</label><br>
		<input type="password" maxlength="25" size="40" name="password" /><br><br>
		<input type="submit" value="Отправить" />
 <a href="../register">РЕГИСТРАЦИЯ</a>
		</form>
		</body>
		</html>`

		Errhtml := []byte(messageUserError)
		res.Write(Errhtml)
		return 0
	}
	token, err := TokenGet(res, req)
	if err != nil {
		fmt.Printf("Ошибка в полученном токене", err)
	}
	user_session.SetFirstNameToken(name, token)
	user_session.Set(token, 0, 0)
	return 1

}

// AddNewUser метод контроллера по добавлению новых пользователей
func (aut *AuthorizationUsersController) AddNewUser(res http.ResponseWriter, req *http.Request) int {
	//получение имени и пароля пользователя из URL строки
	name, password := readURL(res, req)
	//если пароль и имя не заданы выводит сообщение
	if name == `` && password == `` {
		messageUser :=
			`<html>
		<head>
		<meta charset="UTF-8">
		<title>Для регистрации,введите свои данные</title>
		</head>
		<body>
		<h2>Для регистрации,введите свои данные</h2>
		<form method="GET" action="register">
		<label>Имя</label><br>
		<input type="text" name="username" /><br><br>
		<label>Пароль</label><br>
		<input type="password" maxlength="25" size="40" name="password" /><br><br>
        <label>Повторите пароль</label><br>
		<input type="password" maxlength="25" size="40" name="repeatpassword" /><br><br>
		<input type="submit" value="Отправить" />
		</form>
		</body>
		</html>`

		html := []byte(messageUser)
		res.Write(html)

		return 0
	}
	req.ParseForm()
	repeatPassword := req.FormValue("repeatpassword")
	if repeatPassword != password {
		messageUser :=
			`<html>
		<head>
		<meta charset="UTF-8">
		<title>Пароль, не совпадает с паролем подтверждения введите свои данные повторно</title>
		</head>
		<body>
		<h2>Пароль, не совпадает с паролем подтверждения введите свои данные повторно</h2>
		<form method="GET" action="register">
		<label>Имя</label><br>
		<input type="text" name="username" /><br><br>
		<label>Пароль</label><br>
		<input type="password" maxlength="25" size="40" name="password" /><br><br>
		<label>Повторите пароль</label><br>
		<input type="password" maxlength="25" size="40" name="repeatpassword" /><br><br>
		<input type="submit" value="Отправить" />
		</form>
		</body>
		</html>`

		html := []byte(messageUser)
		res.Write(html)
		return 0
	}
	//проверка имени на отсутсвие повторений
	name, err := aut.modelUser.AddNewUser(name, password, ``)
	if err != nil {
		messageUserError := `<html>
		<head>
		<meta charset="UTF-8">
		<title> Извините такой пользователь уже существует, введите свои данные снова</title>
		</head>
		<body>
		<h2>Извините такой пользователь уже существует, введите свои данные снова</h2>
		<form method="GET" action="register">
		<label>Имя</label><br>
		<input type="text" name="username" /><br><br>
		<label>Пароль</label><br>
		<input type="password" maxlength="25" size="40" name="password" /><br><br>
		<label>Повторите пароль</label><br>
		<input type="password" maxlength="25" size="40" name="repeatpassword" /><br><br>
		<input type="submit" value="Отправить" />
		</form>
		</body>
		</html>`

		Errhtml := []byte(messageUserError)
		res.Write(Errhtml)
		return 0
	}
	//создание и отправка cookie файлов
	token, err := TokenGet(res, req)
	if err != nil {
		fmt.Printf("Ошибка в полученном токене", err)
	}
	_, err = aut.modelUser.AddNewUser(name, password, token)
	return 1
}

// DeleteUser удаляет пользователя из БД при вводе правильных: пароля и имени
func (aut *AuthorizationUsersController) DeleteUser(res http.ResponseWriter, req *http.Request) {
	//получение имени и пароля пользователя из URL строки
	name, password := readURL(res, req)
	//если пароль и имя пользователя пустые выводит сообщение
	if name == `` && password == `` {
		messageUser := `<html>
		<head>
		<meta charset="UTF-8">
		<title> Если Вы, хотите удалить пользователя, введите свои данные </title>
		</head>
		<body>
		<h2>Если Вы, хотите удалить пользователя, введите свои данные</h2>
		<form method="GET" action="delete">
		<label>Имя</label><br>
		<input type="text" name="username" /><br><br>
		<label>Пароль</label><br>
		<input type="password" maxlength="25" size="40" name="password" /><br><br>
		<input type="submit" value="Отправить" />
		</form>
		</body>
		</html>`
		html := []byte(messageUser)
		res.Write(html)
		return
	}
	//проверка имени и пароля, в случае ошибки, вывод сообщения
	name, err := aut.modelUser.DeleteUser(name, password)
	if err != nil {
		messageUserError :=
			`<html>
		<head>
		<meta charset="UTF-8">
		<title> Извините, пароль или имя с ошибкой, введите свои данные снова</title>
		</head>
		<body>
		<h2>Извините, пароль или имя с ошибкой, введите свои данные снова</h2>
		<form method="GET" action="delete">
		<label>Имя</label><br>
		<input type="text" name="username" /><br><br>
		<label>Пароль</label><br>
		<input type="password" maxlength="25" size="40" name="password" /><br><br>
		<input type="submit" value="Отправить" />
		</form>
		</body>
		</html>`
		Errhtml := []byte(messageUserError)
		res.Write(Errhtml)
		return
	}
	messageUserError := `<html lang="ru">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	Ваш пользователь удален с сайта <a href="../">САЙТ</a>`
	Errhtml := []byte(messageUserError)
	res.Write(Errhtml)
}

// LogOutUser метод контроллера по сбросу куки файлов и выходу из текущей сессии
func (aut *AuthorizationUsersController) LogOutUser(res http.ResponseWriter, req *http.Request) {
	token, err := TokenVerification(res, req)
	fmt.Printf("Токен из GetAuthorizationUsers контроллера", token)
	if err != nil {
		fmt.Printf("Ошибка такого токена нет в БД")
	}
	aut.modelUser.LogOutUser(token)
	messageUserError := `<html lang="ru">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	Вы успешно вышли из текущей сессии, для продолжения работы прошу авторизоваться <a href="../auth">АВТОРИЗАЦИЯ</a>`
	Errhtml := []byte(messageUserError)
	res.Write(Errhtml)
}
