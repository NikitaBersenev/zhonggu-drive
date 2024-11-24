package main

import (
	"fmt"
	"zhonggu-drive/internal/config"
)

func main() {
	//curl -X "GET" -H "Content-Type: application/json" "http://localhost:8080/info?group=Muse&song=Supermassive+Black+Hole"
	cfg := config.MustLoad()

	//fmt.Println("rest music api")
	//
	//err := http.ListenAndServe(`:8080`, http.HandlerFunc(mainPage))
	//if err != nil {
	//	panic(err)
	//}
	/*
					 Выставить rest методы
					-
					Получение данных библиотеки с фильтрацией по всем полям и
					пагинацией
					-
					Получение текста песни с пагинацией по куплетам
					-
					Удаление песни
					-
					Изменение данных песни
					-
					Добавление новой песни в формате
				{
				"group": "Muse",
				"song": "Supermassive Black Hole"
				}
			При добавлении сделать запрос в АПИ, описанного сваггером

		получение из бд
		пагинация
		фильтрация
		удаление песни
		изменение песни
		добавление песни
	*/

	/*
		3. Обогащенную информацию положить в БД postgres (структура БД должна
		быть создана путем миграций при старте сервиса)
	*/

	/*
		4. Покрыть код debug- и info-логами
	*/

	/*
		5. Вынести конфигурационные данные в ..env-файл
	*/

	/*
		6. Сгенерировать сваггер на реализованное АПИ
	*/

	fmt.Println(cfg)
}
