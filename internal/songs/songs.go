package songs

import (
	"effective_mobile/pkg/dto"
	"effective_mobile/pkg/storage/db"
	"effective_mobile/pkg/storage/db/postgres"
	"errors"
	"strings"
)

var (
	ErrNoSuchCouplet = errors.New("no such couplet in song")
)

var ApiClient interface {
	GetSongInfo(group, song string) (*dto.SongDetail, error)
}

func Delete(id int) error {
	err := postgres.Connection.DeleteSong(id)
	return err
}

func Create(data dto.SongRequest) error {
	err := postgres.Connection.CreateSong(data.ReleaseDate, data.Text, data.Link, data.Song, data.Group)
	ApiClient.GetSongInfo(data.Group, data.Song) // нужно исключительно только как имитация запроса???
	return err
}

func GetAll(page int, filter db.SongFilter) ([]db.Song, error) {
	songsList, err := postgres.Connection.GetSongs(page, filter)
	return songsList, err
}

func GetSong(page, id int) (db.Song, error) {
	//поидее можно сделать кеширование для оптимизации
	song, err := postgres.Connection.GetSong(id)
	if err != nil {
		return db.Song{}, err
	}

	couplets := strings.Split(song.Words, "\\n\\n")

	if page > len(couplets) || page < 1 {
		return db.Song{}, ErrNoSuchCouplet
	}
	song.Words = couplets[page-1]

	return song, nil
}

func Update(id int, params dto.SongRequest) error {
	err := postgres.Connection.UpdateSong(id, params.Group, params.ReleaseDate, params.Text, params.Link)
	return err
}

// 1. Получать песни
//
//
//
//
//
