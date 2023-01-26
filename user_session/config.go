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

// InitializationSession  -функция запускающаяся когда первый раз запустили приложение для инициализации карт
//Имени пользователя-пароля и Имени пользователя-токенов, Токен-имя пользователя
func InitializationSession() {
	session.userInfo = make(map[string]table_metadata.TableMetaData)
	userName.infoName = make(map[string][]string)
	tokenName.tokenName = make(map[string]string)
}

// SetFirstNameToken функция для осуществления первого входа пользователя после регистрации
func SetFirstNameToken(name string, token string) {
	tokenName.lock.Lock()
	defer tokenName.lock.Unlock()
	if _, ok := tokenName.tokenName[token]; ok {
		fmt.Printf("Такой токен есть")
	} else {
		tokenName.tokenName[token] = name
	}
	userName.lock.Lock()
	defer userName.lock.Unlock()
	if _, ok := userName.infoName[name]; ok {
		userName.infoName[name] = append(userName.infoName[name], token)
	} else {
		userName.infoName[name] = append(userName.infoName[name], token)
	}
}

// LogOutToken функция сброса текущей сессии удаляет запись в map
func LogOutToken(token string) {
	session.lock.Lock()
	defer session.lock.Unlock()

	if _, ok := session.userInfo[token]; ok {
		delete(session.userInfo, token)
	}
}

// DeleteToken функция удаления сессии пользователя
func DeleteToken(name string) {
	userName.lock.Lock()
	defer userName.lock.Unlock()
	if valueMap, ok := userName.infoName[name]; ok {
		delete(userName.infoName, name)
		for _, value := range valueMap {
			fmt.Printf("value", value)
			if _, ok := userName.infoName[value]; ok {
				fmt.Printf("userName.infoName[value]", userName.infoName[value])
				delete(userName.infoName, value)
			}
			session.lock.Lock()
			defer session.lock.Unlock()
			if _, ok := session.userInfo[value]; ok {
				delete(session.userInfo, value)
			}
		}
	}
}

//функция получения имени пользователя по токену
func GetName(value string) (string, error) {
	var err error
	if valueToken, ok := tokenName.tokenName[value]; ok {
		name := valueToken
		return name, nil
	} else {
		err = errors.New("Токена нет в базе данных ошибка метода GetName config")
	}
	return "", err
}

// Get функция получение строк и колонок из map по Токену
func Get(value string) (int, int, string, error) {
	var rows int
	var columns int
	var err error
	session.lock.Lock()
	defer session.lock.Unlock()
	if valueMap, ok := session.userInfo[value]; ok {
		rows = valueMap.Row
		columns = valueMap.Column
	} else {
		fmt.Printf("Токена нет в базе данных ошибка метода Get config")
		err := errors.New("Токена нет в базе данных ошибка метода Get config")
		return 0, 0, ``, err
	}
	return rows, columns, value, err
}

//полусение нового уникального id для генерации токена под нового пользователя
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
	fmt.Printf("session.userInfo из config", session.userInfo)
}

//createJWTToken-функция создания нового Токена для нового пользователя
func createToken(id int) string {
	idString := strconv.Itoa(id)
	tokenString := key + idString
	tokenByte := md5.Sum([]byte(tokenString))
	token := hex.EncodeToString(tokenByte[:])
	return token
}
