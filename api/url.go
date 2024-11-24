package api

import (
	_ "effective_mobile/docs"
	logger "effective_mobile/logs"
	"net/http"
)

var urls = make(map[string]*Handle)

func createHandlers(logger logger.LogInterface) {
	newHandle(http.MethodGet, "/songs", getSongs, logger)
	newHandle(http.MethodPost, "/songs", createSong, logger)
	newHandle(http.MethodGet, "/songs/{id}", getSong, logger)
	newHandle(http.MethodDelete, "/delete/{id}", deleteSong, logger)
	newHandle(http.MethodPut, "/update/{id}", updateSong, logger)
}
