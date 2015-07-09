package telapi

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var (
	testTelapiSid       = ""
	testTelapiAuthToken = ""
	testNumberTo        = ""
	testNumberFrom      = "" //Correlates to telapi sid being used
	testCallSid         = "" //Also correlates to telapi sid being used
	testRecordingSid    = ""
)

func TestCreateClient(t *testing.T) {

	Convey("Tests when creating a client", t, func() {

		Convey("calling method CreateClient with empty strings", func() {
			telapi_helper, err := CreateClient("", "")

			Convey("Should bubble an error due to empty strings", func() {
				So(err, ShouldNotBeNil)
			})
			Convey("Since empty strings, should return a bad helper", func() {
				So(telapi_helper.Sid, ShouldEqual, "")
			})
		})

		Convey("calling method CreateClient with bad credentials", func() {
			telapi_helper, err := CreateClient("test1", "test1")

			Convey("Should bubble an error due to invalid credentials on telapi end", func() {
				So(err, ShouldNotBeNil)
			})
			Convey("Since error bubbled should return a bad helper", func() {
				So(telapi_helper.Sid, ShouldEqual, "")
			})
		})

		Convey("calling method CreateClient with good credentials", func() {
			telapi_helper, err := CreateClient(testTelapiSid, testTelapiAuthToken)

			Convey("Should be no error, due to good credentials", func() {
				So(err, ShouldBeNil)
			})
			Convey("The sid should be correctly set.", func() {
				So(telapi_helper.Sid, ShouldEqual, testTelapiSid)
			})
		})

	})

}
