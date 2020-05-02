package clock

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_TimestampConversion(t *testing.T) {
	Convey("Given some non-zero time", t, func(c C) {
		now := time.Now()

		Convey("When we convert it to a Timestamp protobuf, it should not be nil", func() {
			pb := TimeToTimestamp(now)
			So(pb, ShouldNotBeNil)

			Convey("And when we convert the protobuf back into a time, we should get our original value", func() {
				res := TimestampToTime(pb)
				So(now, ShouldEqual, res)
			})
		})
	})

	Convey("Given some zero time", t, func(c C) {
		now := time.Time{}

		Convey("When we convert it to a Timestamp protobuf, it should be nil", func() {
			pb := TimeToTimestamp(now)
			So(pb, ShouldBeNil)

			Convey("And when we convert the protobuf back into a time, we should get our original value", func() {
				res := TimestampToTime(pb)
				So(now, ShouldEqual, res)
			})
		})
	})
}

func Test_timeToNanosecondSecondFraction(t *testing.T) {
	Convey("Given some time", t, func(c C) {
		var fraction int32

		fraction = timeToNanosecondSecondFraction(time.Unix(1000000, 0))
		So(fraction, ShouldEqual, 0)

		fraction = timeToNanosecondSecondFraction(time.Unix(1000000, 12345678))
		So(fraction, ShouldEqual, 12345678)

		fraction = timeToNanosecondSecondFraction(time.Unix(1000000, 999999999))
		So(fraction, ShouldEqual, 999999999)
	})
}
