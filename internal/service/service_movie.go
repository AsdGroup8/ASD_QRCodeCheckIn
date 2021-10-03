package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/AsdGroup8/ASD_QRCodeCheckIn/conf"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/model"
)

var (
	_inTheaterMovie = ImdbInTheaterResp{}
	_movieMtx       = sync.Mutex{}
	_lastGetMovieTs time.Time
)

// ImdbInTheaterResp ...
type ImdbInTheaterResp struct {
	Items []model.Movie `json:"items"`
}

// GetInTheatersMovies ...
func GetInTheatersMovies() ([]model.Movie, error) {
	now := time.Now()
	if _lastGetMovieTs.Add(conf.MovieUpdateGap).Before(now) {
		req, err := http.NewRequest(http.MethodGet, conf.ImdbAPI, nil)
		if err != nil {
			return nil, fmt.Errorf("fail to retrive movie data. %v", err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("fail to retrive movie data. %v", err)
		}
		defer resp.Body.Close()
		tmpData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("fail to retrive movie data. %v", err)
		}
		_movieMtx.Lock()
		defer _movieMtx.Unlock()

		if err := json.Unmarshal(tmpData, &_inTheaterMovie); err != nil {
			return nil, fmt.Errorf("fail to unmarshal movie data. %v", err)
		}
		_lastGetMovieTs = now
	}
	return _inTheaterMovie.Items, nil
}
