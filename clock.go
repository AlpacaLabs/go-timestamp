package clock

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

// TimestampToTime converts a Timestamp (from Google's well-known Protocol Buffer types) to a time.Time.
func TimestampToTime(pb *timestamp.Timestamp) time.Time {
	if pb != nil {
		if pb.Seconds != 0 {
			return time.Unix(pb.Seconds, int64(pb.Nanos))
		}
	}
	return time.Time{}
}

// TimeToTimestamp converts a time.Time into Timestamp (from Google's well-known Protocol Buffer types).
func TimeToTimestamp(t time.Time) *timestamp.Timestamp {
	if t.IsZero() {
		return nil
	}
	return &timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   timeToNanosecondSecondFraction(t),
	}
}

func timeToNanosecondSecondFraction(in time.Time) int32 {
	return int32(in.UnixNano() % 1000000000)
}
