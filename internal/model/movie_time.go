package model

import "time"

// MovieTime ...
type MovieTime struct {
	StartTs int64 // Seconds
	EndTs   int64 // Seconds
}

// StartTime string of StartTs
func (mt *MovieTime) StartTime() string {
	return time.Unix(mt.StartTs, 0).Local().Format("2006-01-02 15:04:05")
}

// EndTime string of EndTs
func (mt *MovieTime) EndTime() string {
	return time.Unix(mt.EndTs, 0).Local().Format("2006-01-02 15:04:05")
}
