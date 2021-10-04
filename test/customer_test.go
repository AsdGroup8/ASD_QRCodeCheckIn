package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func Test_RegisterCustomer(t *testing.T) {
	payload := map[string]interface{}{
		"name":     "user_" + strconv.FormatInt(time.Now().Unix(), 36),
		"gender":   1,
		"dob":      time.Now().Unix(),
		"email":    "zhongxiao0711@gmail.com",
		"phone":    "13281122813",
		"password": "123aaa",
		"address":  "chengdu, sichuan",
	}
	data, err := json.Marshal(payload)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	resp, err := http.Post(_baseApi+"/customer/reg", "json", bytes.NewReader(data))
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	defer resp.Body.Close()
}

func Test_ModifyCustomer(t *testing.T) {
	payload := map[string]interface{}{
		"name":     "user_" + strconv.FormatInt(time.Now().Unix(), 36),
		"gender":   1,
		"dob":      time.Now().Unix(),
		"email":    "zhongxiao0711123@gmail.com",
		"phone":    "132811123",
		"password": "123aaa",
		"address":  "chengdu, sichuan",
	}
	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPut, _baseApi+"/customer/profile", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	req.Header["Content-Type"] = []string{"application/json"}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
}
