package handlers

import (
	"fmt"
	"net/http"
)

func (h *AppHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		writeJson(w, http.StatusBadRequest, Map{"message": "failed to parse form"})
	}

	files := r.MultipartForm.File

	for _, value := range files {
		for i := range value {
			fmt.Println(value[i].Filename)
		}
	}
}

func (h *AppHandler) DeleteFile(w http.ResponseWriter, r *http.Request) {

}

func (h *AppHandler) GetUserFiles(w http.ResponseWriter, r *http.Request) {

}
