package service

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/AsdGroup8/ASD_QRCodeCheckIn/conf"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/log"
)

var (
	_movieData      string
	_lastGetMovieTs time.Time
)

// GetInTheatersMovies ...
func GetInTheatersMovies() string {
	now := time.Now()
	if _lastGetMovieTs.Add(conf.MovieUpdateGap).Before(now) {
		req, err := http.NewRequest(http.MethodGet, conf.ImdbAPI, nil)
		if err != nil {
			log.Errorf("fail to retrive movie data. %v", err)
			return ""
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Errorf("fail to retrive movie data. %v", err)
			return ""
		}
		defer resp.Body.Close()
		tmpData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Errorf("fail to retrive movie data. %v", err)
			return ""
		}
		_movieData = string(tmpData)
		_lastGetMovieTs = now
	}
	return _movieData
}
