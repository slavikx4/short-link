package services

import "errors"

var (
	ErrorNoLink       = errors.New("ошибка: такая ссылка не зарегистрирована")
	ErrorInvalidValue = errors.New("вы ввели некорректное значение")
)
