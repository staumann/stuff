package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/staumann/caluclation/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CreateUserHandler_Success(t *testing.T) {
	adapter = &AdapterSpy{
		saveUserHandler: func(user *model.User) error {
			assert.Equal(t, "TestName", user.Name)
			user.ID = 5
			return nil
		},
	}
	user := model.User{Name: "TestName"}
	bts, _ := json.Marshal(user)
	var buffer bytes.Buffer

	buffer.Write(bts)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/api/user/create", &buffer)

	createUserHandler(recorder, request)
	responseObj := new(model.User)
	assert.Nil(t, json.Unmarshal(recorder.Body.Bytes(), responseObj))
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, int64(5), responseObj.ID)
	assert.Equal(t, "TestName", responseObj.Name)
}

func Test_CreateUserHandler_Failure(t *testing.T) {
	adapter = &AdapterSpy{
		saveUserHandler: func(user *model.User) error {
			return errors.New("test error")
		},
	}
	user := model.User{Name: "TestName"}
	bts, _ := json.Marshal(user)
	var buffer bytes.Buffer

	buffer.Write(bts)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/api/user/create", &buffer)

	createUserHandler(recorder, request)
	responseObj := new(model.ErrorResponse)
	_ = json.Unmarshal(recorder.Body.Bytes(), responseObj)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, "test error", responseObj.Message)
}

func Test_GetUserHandler_Success(t *testing.T) {
	adapter = &AdapterSpy{
		getUserHandler: func(i int64) *model.User {
			assert.Equal(t, int64(5), i)
			return &model.User{Name: "TestName", ID: int64(5)}
		},
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/user/get?id=5", nil)

	getUserHandler(recorder, request)

	responseObj := new(model.User)
	assert.Nil(t, json.Unmarshal(recorder.Body.Bytes(), responseObj))
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "TestName", responseObj.Name)
	assert.Equal(t, int64(5), responseObj.ID)
}

func Test_GetUserHandler_NotFound(t *testing.T) {
	adapter = &AdapterSpy{
		getUserHandler: func(i int64) *model.User {
			return nil
		},
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/user/get?id=5", nil)

	getUserHandler(recorder, request)

	responseObj := new(model.ErrorResponse)
	assert.Nil(t, json.Unmarshal(recorder.Body.Bytes(), responseObj))
	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.Equal(t, "user with id 5 not found", responseObj.Message)
}

func Test_GetUserHandler_MissingID(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/user/get?id=", nil)

	getUserHandler(recorder, request)

	responseObj := new(model.ErrorResponse)
	assert.Nil(t, json.Unmarshal(recorder.Body.Bytes(), responseObj))
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "error no id given", responseObj.Message)
}

func Test_UpdateUserHandler_Success(t *testing.T) {
	adapter = &AdapterSpy{
		updateUserHandler: func(user *model.User) error {
			assert.Equal(t, "NewName", user.Name)
			return nil
		},
	}

	var buffer bytes.Buffer

	bts, _ := json.Marshal(model.User{
		ID:   25,
		Name: "NewName",
	})

	buffer.Write(bts)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/api/user/update", &buffer)

	updateUserHandler(recorder, request)
	responseObj := new(model.User)
	_ = json.Unmarshal(recorder.Body.Bytes(), responseObj)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "NewName", responseObj.Name)
}

func Test_UpdateUserHandler_Failure(t *testing.T) {
	adapter = &AdapterSpy{
		updateUserHandler: func(user *model.User) error {
			return errors.New("test error")
		},
	}

	var buffer bytes.Buffer

	bts, _ := json.Marshal(model.User{
		ID:   25,
		Name: "NewName",
	})

	buffer.Write(bts)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/api/user/update", &buffer)

	updateUserHandler(recorder, request)
	responseObj := new(model.ErrorResponse)
	_ = json.Unmarshal(recorder.Body.Bytes(), responseObj)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, "test error", responseObj.Message)
}
