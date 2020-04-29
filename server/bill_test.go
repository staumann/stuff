package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/staumann/caluclation/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_CreateBillHandler_Success(t *testing.T) {
	adapter = &AdapterSpy{
		saveHandler: func(bill *model.Bill) error {
			assert.Equal(t, int64(7), bill.UserID)
			assert.Equal(t, int64(5), bill.ShopID)
			assert.Equal(t, 10.55, bill.TotalDiscount)

			return nil
		},
	}
	bts, _ := json.Marshal(model.Bill{
		UserID:        7,
		ShopID:        5,
		TotalDiscount: 10.55,
	})
	recorder := httptest.NewRecorder()

	var buffer bytes.Buffer
	buffer.Write(bts)
	request := httptest.NewRequest(http.MethodPost, "/api/bill/create", &buffer)

	createBillHandler(recorder, request)

	result := recorder.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	responseObj := new(model.Bill)

	assert.Nil(t, json.Unmarshal(recorder.Body.Bytes(), responseObj))
	assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
	assert.Equal(t, int64(7), responseObj.UserID)
	assert.Equal(t, int64(5), responseObj.ShopID)
	assert.Equal(t, 10.55, responseObj.TotalDiscount)
}

func Test_CreateBillHandler_Failure(t *testing.T) {
	adapter = &AdapterSpy{
		saveHandler: func(bill *model.Bill) error {
			return errors.New("test error")
		},
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/api/bill/create", nil)

	createBillHandler(recorder, request)

	assert.Equal(t, http.StatusInternalServerError, recorder.Result().StatusCode)
	assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
	assert.True(t, strings.Contains(recorder.Body.String(), "test error"))
}

func Test_GetBillHandler_Success(t *testing.T) {
	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "/api/bill/get?id=5", nil)

	adapter = &AdapterSpy{
		getHandler: func(i int64) *model.Bill {
			assert.Equal(t, int64(5), i)
			return &model.Bill{
				UserID:        15,
				TotalDiscount: 22.22,
			}
		},
	}

	getBillHandler(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	responseObj := new(model.Bill)

	assert.Nil(t, json.Unmarshal(recorder.Body.Bytes(), responseObj))
	assert.Equal(t, int64(15), responseObj.UserID)
	assert.Equal(t, 22.22, responseObj.TotalDiscount)
}

func Test_GetBillHandler_MissingID(t *testing.T) {
	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "/api/bill/get?id=", nil)

	getBillHandler(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	responseObj := new(model.ErrorResponse)

	assert.Nil(t, json.Unmarshal(recorder.Body.Bytes(), responseObj))
	assert.Equal(t, "no id passed", responseObj.Message)
}

func Test_GetBillHandler_NotFound(t *testing.T) {
	adapter = &AdapterSpy{
		getHandler: func(i int64) *model.Bill {
			return nil
		},
	}
	recorder := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodGet, "/api/bill/get?id=5", nil)

	getBillHandler(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	responseObj := new(model.ErrorResponse)

	assert.Nil(t, json.Unmarshal(recorder.Body.Bytes(), responseObj))
	assert.Equal(t, "error bill with id 5 not found", responseObj.Message)
}

func Test_UpdateBillHandler_Success(t *testing.T) {
	adapter = &AdapterSpy{
		updateHandler: func(bill *model.Bill) error {
			assert.Equal(t, int64(5), bill.ID)
			assert.Equal(t, 25.25, bill.TotalDiscount)
			return nil
		},
	}
	bts, _ := json.Marshal(model.Bill{
		ID:            5,
		UserID:        7,
		ShopID:        5,
		TotalDiscount: 25.25,
	})
	recorder := httptest.NewRecorder()

	var buffer bytes.Buffer
	buffer.Write(bts)
	request := httptest.NewRequest(http.MethodPost, "/api/bill/update", &buffer)

	updateBillHandler(recorder, request)

	result := recorder.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	responseObj := new(model.Bill)

	assert.Nil(t, json.Unmarshal(recorder.Body.Bytes(), responseObj))
	assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
	assert.Equal(t, int64(7), responseObj.UserID)
	assert.Equal(t, int64(5), responseObj.ShopID)
	assert.Equal(t, 25.25, responseObj.TotalDiscount)
}

func Test_UpdateBillHandler_Failure(t *testing.T) {
	adapter = &AdapterSpy{
		updateHandler: func(bill *model.Bill) error {
			return errors.New("test error")
		},
	}
	bts, _ := json.Marshal(model.Bill{
		ID:            5,
		UserID:        7,
		ShopID:        5,
		TotalDiscount: 25.25,
	})
	recorder := httptest.NewRecorder()

	var buffer bytes.Buffer
	buffer.Write(bts)
	request := httptest.NewRequest(http.MethodPost, "/api/bill/update", &buffer)

	updateBillHandler(recorder, request)

	result := recorder.Result()
	assert.Equal(t, http.StatusInternalServerError, result.StatusCode)

	responseObj := new(model.ErrorResponse)

	assert.Nil(t, json.Unmarshal(recorder.Body.Bytes(), responseObj))
	assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
	assert.Equal(t, "test error", responseObj.Message)
}

func Test_DeleteHandler_Success(t *testing.T) {
	adapter = &AdapterSpy{
		deleteHandlerById: func(i int64) error {
			assert.Equal(t, int64(5), i)
			return nil
		},
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/bill/delete?id=5", nil)

	deleteBillHandler(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
	responseObject := new(model.DeleteResponse)
	_ = json.Unmarshal(recorder.Body.Bytes(), responseObject)
	assert.Equal(t, int64(5), responseObject.ID)
}

func Test_DeleteHandler_Failure(t *testing.T) {
	adapter = &AdapterSpy{
		deleteHandlerById: func(i int64) error {
			assert.Equal(t, int64(5), i)
			return errors.New("test error")
		},
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/bill/delete?id=5", nil)

	deleteBillHandler(recorder, request)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
	responseObject := new(model.ErrorResponse)
	_ = json.Unmarshal(recorder.Body.Bytes(), responseObject)
	assert.Equal(t, "test error", responseObject.Message)
}

func Test_DeleteHandler_MissingID(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/bill/delete", nil)

	deleteBillHandler(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
	responseObject := new(model.ErrorResponse)
	_ = json.Unmarshal(recorder.Body.Bytes(), responseObject)
	assert.Equal(t, "error no id given", responseObject.Message)
}
