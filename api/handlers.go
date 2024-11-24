package api

import (
	"effective_mobile/internal/songs"
	"effective_mobile/pkg/dto"
	"effective_mobile/pkg/storage/db"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GETALLSONGS Получение всех песен
// @Summary Список всех песен
// @Description Возвращает все песни с фильтрацией
// @Tags Songs
// @Produce json
// @Param group query string false "Название группы"
// @Param release_date_start query string false "Начальная дата (yyyy-mm-dd)"
// @Param release_date_end query string false "Конечная дата (yyyy-mm-dd)"
// @Param page query int false "Номер страницы пагинации"
// @Param song query string false "Название песни"
// @Success 200 {array} dto.SongRequest
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /songs [get]
func getSongs(r *http.Request) ([]byte, int, error) {
	query := r.URL.Query()

	pg := query.Get("page")
	if pg == "" {
		pg = "1"
	}
	page, err := strconv.Atoi(pg)
	if err != nil {
		res := fmt.Sprintf(`{"err": "%s"}`, err.Error())
		return []byte(res), http.StatusInternalServerError, err
	}

	filter := db.SongFilter{
		Group:       query.Get("group"),
		Song:        query.Get("song"),
		ReleaseDate: []string{query.Get("release_date_start"), query.Get("release_date_end")},
	}

	songsList, err := songs.GetAll(page, filter)
	if err != nil {
		res := fmt.Sprintf(`{"err": "%s"}`, err.Error())
		return []byte(res), http.StatusInternalServerError, err
	}

	res, _ := json.Marshal(songsList)
	return res, http.StatusOK, nil
}

// @Summary Получить песню по ID
// @Description Возвращает песню по ID
// @Tags Songs
// @Produce json
// @Param id path string true "ID песни"
// @Success 200 {object} dto.SongRequest
// @Failure 404 {object} map[string]string "Песня не найдена"
// @Router /songs/{id} [get]
func getSong(r *http.Request) ([]byte, int, error) {
	query := r.URL.Query()

	cp := query.Get("couplet")
	if cp == "" {
		cp = "1"
	}

	couplet, err := strconv.Atoi(cp)
	if err != nil {
		res := fmt.Sprintf(`{"err": "%s"}`, err.Error())
		return []byte(res), http.StatusInternalServerError, err
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		res := fmt.Sprintf(`{"err": "%s"}`, err.Error())
		return []byte(res), http.StatusNotFound, err
	}

	song, err := songs.GetSong(couplet, id)
	if err != nil {
		res := fmt.Sprintf(`{"err": "%s"}`, err.Error())
		return []byte(res), http.StatusInternalServerError, err
	}

	res, _ := json.Marshal(song)
	return res, http.StatusOK, nil
}

// @Summary Создать новую песню
// @Description Добавляет песню в базу данных
// @Tags Songs
// @Accept json
// @Produce json
// @Param song body dto.SongRequest true "Информация о песне"
// @Success 201 {object} map[string]string "Успешное создание"
// @Failure 400 {object} map[string]string "Ошибка запроса"
// @Router /songs [post]
func createSong(r *http.Request) ([]byte, int, error) {
	var data dto.SongRequest

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res := fmt.Sprintf(`{"err": "%s"}`, err.Error())
		return []byte(res), http.StatusNotAcceptable, err
	}

	err = songs.Create(data)
	if err != nil {
		res := fmt.Sprintf(`{"err": "%s"}`, err.Error())
		return []byte(res), http.StatusInternalServerError, err
	}

	return []byte(`{"message": "song created successfully"}`), http.StatusCreated, nil
}

// @Summary Удалить песню
// @Description Удаляет песню из базы данных
// @Tags Songs
// @Param id path string true "ID песни"
// @Success 200 {object} map[string]string "Успешное удаление"
// @Failure 404 {object} map[string]string "Песня не найдена"
// @Router /delete/{id} [delete]
func deleteSong(r *http.Request) ([]byte, int, error) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		res := fmt.Sprintf(`{"err": "%s"}`, err.Error())
		return []byte(res), http.StatusInternalServerError, err
	}

	err = songs.Delete(id)
	if err != nil {
		res := fmt.Sprintf(`{"err": "%s"}`, err.Error())
		return []byte(res), http.StatusInternalServerError, err
	}

	return []byte(`{"message": "song deleted successfully"}`), http.StatusOK, nil
}

// @Summary Обновить песню
// @Description Обновляет данные о песне в базе
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path string true "ID песни"
// @Param song body dto.SongRequest true "Обновленная информация о песне"
// @Success 200 {object} map[string]string "Успешное обновление"
// @Failure 404 {object} map[string]string "Песня не найдена"
// @Router /update/{id} [put]
func updateSong(r *http.Request) ([]byte, int, error) {
	var data dto.SongRequest
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		res := fmt.Sprintf(`{"err": "%s"}`, err.Error())
		return []byte(res), http.StatusNotFound, err
	}

	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res := fmt.Sprintf(`{"err": "%s"}`, err.Error())
		return []byte(res), http.StatusNotAcceptable, err
	}

	err = songs.Update(id, data)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return []byte(`{"message": "song updated successfully"}`), http.StatusOK, nil
}
