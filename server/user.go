package server

import (
	"encoding/json"
	"fmt"
	"github.com/staumann/caluclation/model"
	"io/ioutil"
	"net/http"
	"strconv"
)

func createUserHandler(writer http.ResponseWriter, request *http.Request) {
	bts, _ := ioutil.ReadAll(request.Body)
	writer.Header().Set("Content-Type", contentTypeJson)
	user := new(model.User)
	_ = json.Unmarshal(bts, user)

	if e := adapter.SaveUser(user); e != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		bts, _ = json.Marshal(model.ErrorResponse{Message: e.Error()})
	} else {
		writer.WriteHeader(http.StatusOK)
		bts, _ = json.Marshal(*user)
	}

	_, _ = writer.Write(bts)
}

func getUserHandler(writer http.ResponseWriter, request *http.Request) {
	idString := request.URL.Query().Get("id")
	writer.Header().Set("Content-Type", contentTypeJson)
	var bts []byte
	if idString == "" {
		writer.WriteHeader(http.StatusBadRequest)
		bts, _ = json.Marshal(model.ErrorResponse{Message: "error no id given"})
	} else {
		id, _ := strconv.ParseInt(idString, 10, 64)
		if user := adapter.GetUserByID(id); user != nil {
			writer.WriteHeader(http.StatusOK)
			bts, _ = json.Marshal(*user)
		} else {
			writer.WriteHeader(http.StatusNotFound)
			bts, _ = json.Marshal(model.ErrorResponse{Message: fmt.Sprintf("user with id %d not found", id)})
		}
	}

	_, _ = writer.Write(bts)
}
