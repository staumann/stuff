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
	_ = json.Unmarshal(recorder.Body.Bytes(), responseObj)
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

}
