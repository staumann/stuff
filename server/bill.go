package server

import (
	"encoding/json"
	"fmt"
	"github.com/staumann/caluclation/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func createBillHandler(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Content-Type", contentTypeJson)
	bts, _ := ioutil.ReadAll(r.Body)
	bill := new(model.Bill)
	err := json.Unmarshal(bts, bill)

	if err != nil {
		log.Printf("could not marshal bill from request: %s", err.Error())
	}

	if err = billRepository.SaveBill(bill); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		bts, _ = json.Marshal(model.ErrorResponse{Message: err.Error()})
	} else {
		writer.WriteHeader(http.StatusOK)
		bts, _ = json.Marshal(*bill)
	}

	_, _ = writer.Write(bts)
}

func getBillHandler(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Content-Type", contentTypeJson)
	idString := r.URL.Query().Get("id")
	var bts []byte
	if idString == "" {
		writer.WriteHeader(http.StatusBadRequest)
		bts, _ = json.Marshal(model.ErrorResponse{Message: "no id passed"})
	} else {
		id, _ := strconv.ParseInt(idString, 10, 64)

		if obj := billRepository.GetBillByID(id); obj != nil {
			writer.WriteHeader(http.StatusOK)
			bts, _ = json.Marshal(*obj)
		} else {
			writer.WriteHeader(http.StatusNotFound)
			bts, _ = json.Marshal(model.ErrorResponse{Message: fmt.Sprintf("error bill with id %d not found", id)})
		}
	}
	_, _ = writer.Write(bts)
}

func updateBillHandler(writer http.ResponseWriter, r *http.Request) {
	bts, _ := ioutil.ReadAll(r.Body)
	bill := new(model.Bill)
	_ = json.Unmarshal(bts, bill)
	writer.Header().Set("Content-Type", contentTypeJson)

	err := billRepository.UpdateBill(bill)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		bts, _ = json.Marshal(model.ErrorResponse{Message: err.Error()})
	} else {
		writer.WriteHeader(http.StatusOK)
		bts, _ = json.Marshal(*bill)
	}
	_, _ = writer.Write(bts)
}

func deleteBillHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", contentTypeJson)
	idString := request.URL.Query().Get("id")
	var bts []byte
	if idString == "" {
		writer.WriteHeader(http.StatusBadRequest)
		bts, _ = json.Marshal(model.ErrorResponse{Message: "error no id given"})
	} else {
		id, _ := strconv.ParseInt(idString, 10, 64)
		if err := billRepository.DeleteBillByID(id); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			bts, _ = json.Marshal(model.ErrorResponse{Message: err.Error()})
		} else {
			writer.WriteHeader(http.StatusOK)
			bts, _ = json.Marshal(model.DeleteResponse{ID: id})
		}
	}
	_, _ = writer.Write(bts)
}
