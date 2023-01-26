package cookies

import (
	"fmt"
	"httpTable/user_session"
	"net/http"
)

// Cookie структура cookie-файлов
type Cookie struct {
	Name  string
	Value string
}

//переменная доступа с структуре Cookie
var cookies *Cookie

// SetCookie функция установка cookie файлов для нового пользователя
func SetCookie(res http.ResponseWriter, cookie http.Cookie) (string, error) {
	cookie.Value = user_session.SetCookie()
	http.SetCookie(res, &cookie)
	return cookie.Value, nil
}

// GetCookie функция по получению значения имени и значения Токена с браузера
func GetCookie(res *http.Request, name string) (string, error) {
	cookie, err := res.Cookie(name)
	if err != nil {
		fmt.Printf("Ошибка в имени cookie-файла в контроллере функции Get ", err)
		return ``, err
	}
	_, _, _, err = user_session.Get(cookie.Value)
	if err != nil {
		fmt.Printf("Ошибка в имени cookie-файла в контроллере функции Get", err)

		return ``, err
	}
	return cookie.Value, nil
}
