package main

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"
	rest "zhonggu-drive/rest"
)

type restService struct {
	songs map[string]map[string]rest.SongDetail
	id    int64
	mux   sync.Mutex
}

func (r *restService) InfoGet(ctx context.Context, params rest.InfoGetParams) (rest.InfoGetRes, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	//var response rest.InfoGetRes
	//song := params.Song
	//group := params.Group
	//pet, ok := r.pets[params.Song]
	songDetail, ok := r.songs[params.Group][params.Song]
	if !ok {
		return &rest.InfoGetBadRequest{}, nil
	}

	return &songDetail, nil
}

func main() {
	// Create service instance.

	songData := map[string]map[string]rest.SongDetail{
		"Muse": {
			"Supermassive Black Hole": {
				ReleaseDate: time.Date(2006, 7, 16, 0, 0, 0, 0, time.UTC).Format("2006-01-02"),
				Text:        "Ooh baby, don't you know I suffer?\n...", // ...оставшийся текст
				Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
			},
			"Time Is Running Out": {
				ReleaseDate: time.Date(2003, 9, 15, 0, 0, 0, 0, time.UTC).Format("2006-01-02"),
				Text:        "The pressure's building, the time is running out...",
				Link:        "https://www.youtube.com/watch?v=YZd-J5vL65U",
			},
		},
		"Queen": {
			"Bohemian Rhapsody": {
				ReleaseDate: time.Date(1975, 10, 31, 0, 0, 0, 0, time.UTC).Format("2006-01-02"),
				Text:        "Is this the real life? Is this just fantasy? ...",
				Link:        "https://www.youtube.com/watch?v=fJ9rEULg_l0",
			},
		},
	}

	service := &restService{
		songs: songData,
	}

	// Create generated server.
	srv, err := rest.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
