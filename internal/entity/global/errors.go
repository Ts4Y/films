package global

import "errors"

var (
	ErrIncorrectParams = errors.New("неверные параметры")
	ErrInternalServerError = errors.New("ошибка на сервере")
	ErrNoData = errors.New("нет данных")
)