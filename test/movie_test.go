package test

import (
	"net/http"
	"testing"
)

func Test_GetInTheaterMovie(t *testing.T) {
	resp, err := http.Get(_baseAPI + "/app/v1")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	defer resp.Body.Close()
}
