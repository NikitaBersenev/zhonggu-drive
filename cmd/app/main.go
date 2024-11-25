package main

import (
	"log/slog"
	"os"
	"zhonggu-drive/internal/config"
	"zhonggu-drive/internal/lib/logger/sl"
	"zhonggu-drive/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//curl -X "GET" -H "Content-Type: application/json" "http://localhost:8080/info?group=Muse&song=Supermassive+Black+Hole"
	cfg := config.MustLoad()
	//fmt.Println(cfg.Env)

	log := setupLogger(cfg.Env)
	log.Info("starting music api")
	log.Info("environment type", slog.String("env", cfg.Env))
	log.Debug("debugging messages are enabled")

	log.Debug("storage path", cfg.StoragePath)
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err)) // как без костылей?
		os.Exit(1)
	}

	id, err := storage.SaveMusic("group1", "music1")
	if err != nil {
		log.Error("failed to save music", sl.Err(err))
	}

	log.Info("music id is", id)
	_ = storage

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

	//fmt.Println(cfg)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
