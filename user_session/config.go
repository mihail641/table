package user_session

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"httpTable/table_metadata"
	"strconv"
)

// Row начальные значения таблицы
const Row = 2
const Column = 2

// Init -функция запускающаяся когда первый раз запустили приложение
func Init() {
	session.userInfo = make(map[string]table_metadata.TableMetaData)
	fmt.Sprintf("session.userInfo из конфига функция ИНИТ%v\n\t", session.userInfo)
}

// Get функция получение строк и колонок из map по Токену
func Get(value string) (int, int, string, error) {
	var rows int
	var columns int
	var err error
	session.lock.Lock()
	defer session.lock.Unlock()
	if valueMap, ok := session.userInfo[value]; ok {
		rows := valueMap.Row
		columns := valueMap.Column
		return rows, columns, value, err
	} else {
		fmt.Printf("Токена нет в базе данных ошибка метода Get config")
		err := errors.New("Токена нет в базе данных ошибка метода Get config")
		return 0, 0, ``, err
	}
	return rows, columns, value, err
}
func SetCookie() string {
	token.lock.Lock()
	defer token.lock.Unlock()
	token.id++
	token := createToken(token.id)
	return token
}

// Set функция присвоения нового значения map по значению Токена
func Set(value string, row int, column int) {
	session.lock.Lock()
	defer session.lock.Unlock()
	if _, ok := session.userInfo[value]; ok {
		session.userInfo[value] = table_metadata.TableMetaData{row, column}
	} else {
		session.userInfo[value] = table_metadata.TableMetaData{Row, Column}
	}
}

//createJWTToken-функция создания нового Токена для нового пользователя
func createToken(id int) string {
	idString := strconv.Itoa(id)
	tokenString := key + idString
	tokenByte := md5.Sum([]byte(tokenString))
	token := hex.EncodeToString(tokenByte[:])
	return token
}
