package clock

import (
	"time"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func DurationToPB(in time.Duration) *duration.Duration {
	// the duration as an integer nanosecond count
	nsTotal := in.Nanoseconds()

	var seconds int64
	var nanos int32

	if time.Duration(nsTotal) > time.Second {
		nanos = int32(nsTotal % int64(time.Second))
		seconds = int64(time.Duration(nsTotal-int64(nanos)) / time.Second)
	} else if time.Duration(nsTotal) == time.Second {
		seconds = 1
	} else {
		nanos = int32(nsTotal)
	}

	return &duration.Duration{
		Seconds: seconds,
		Nanos:   nanos,
	}
}

func DurationFromPB(in *duration.Duration) time.Duration {
	sec := time.Duration(in.Seconds) * time.Second
	ns := time.Duration(in.Nanos) * time.Nanosecond
	return sec + ns
}

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
