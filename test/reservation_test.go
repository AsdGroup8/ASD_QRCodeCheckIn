package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/model"
)

var cus = &model.Customer{
	Name:     "xiao zhongzhong",
	Gender:   "male",
	DOB:      time.Now().Unix(),
	Email:    strconv.FormatInt(time.Now().Unix(), 36) + "@gmail.com",
	Password: "123aaa",
	Phone:    strconv.FormatInt(time.Now().Unix(), 10),
	Address:  "earth",
}

func Test_CreateReservation(t *testing.T) {
	payload := map[string]interface{}{
		"movie_id":  "tt7097896",
		"starttime": time.Now().Unix(),
	}
	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.Post(_baseAPI+"/customer/reserv", "json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
}

func Test_GetReservation(t *testing.T) {
	resp, err := http.Get(_baseAPI + "/customer/reserv/history")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
}
