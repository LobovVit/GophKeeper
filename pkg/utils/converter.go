package utils

import (
	"time"

	timestamp "google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertTimeToTimestamp(t time.Time) *timestamp.Timestamp {
	return timestamp.New(t)
}

func ConvertTimestampToTime(ts *timestamp.Timestamp) (time.Time, error) {
	err := ts.CheckValid()
	if err != nil {
		return time.Time{}, err
	}
	return ts.AsTime(), nil
}
