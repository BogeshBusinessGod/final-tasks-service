package config

type Logger struct {
	Level string `envconfig:"LEVEL" default:"debug"`
}

// МИЖГАН Я ЗНАЮ ЧТО У ТЕБЯ ДРУГОЙ ЛОГГЕР НО ТЫ ГОВОРИЛ ЧТО НА ВОРКЕ БУДУТ ЗА МЕНЯ ЕГО ПИСАТЬ
//ТАК ЧТО Я ПОКА ОСТАВИЛ ТАКОЙ ОН МАТЬ ТВОЮ РАБОЧИЙ
