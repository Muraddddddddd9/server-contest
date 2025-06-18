package constants

const (
	ErrLoadEnv        = "Ошибка загрузки env"
	ErrInternalServer = "Ошибка сервера"
	ErrInputValue     = "Не верный ввод данных"
	ErrUserNotFound   = "Пользователь не найден"
	ErrUserExit       = "Ошибка выхода"
	ErrAlreadyEntry   = "Пользовтаель уже зарегестрирован"
	ErrNotLeader      = "Вы не являетесь лидером группы"
	ErrStageLessonEnd = "Этапы закончились"
	ErrTestNotStart   = "Тест ещё не начался"
	ErrTestEnd        = "Тест закончился"
	ErrNoFullAnser    = "Ответьте на все вопросы"
	ErrAlreadyReplied = "Вы уже ответили на все вопросы"
)

const (
	SuccCloseWS           = "WebSocket закрыт"
	SuccEntry             = "Успешный вход"
	SuccExit              = "Успешный выход"
	SuccUpdateLessonStage = "Этап урока успешно обновился"
	SuccGetAnswer         = "Ответ получен"
	SuccChangeTime        = "Успешное действие со временем"
)
