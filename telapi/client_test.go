package telapi

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var (
	testing_telapi_sid        = "AC1d530461c32a4840a1a19183d0a0bb8c"
	testing_telapi_auth_token = "335cfd52bd534c18b55801f4cb2f76fa"
	testing_number_to         = "+17325807596"
	testing_number_from       = "+17327305402" //Correlates to telapi sid being used
)

func TestCreateClient(t *testing.T) {

	Convey("Tests when creating a client", t, func() {

		Convey("calling method CreateClient with empty strings", func() {
			telapi_helper, err := CreateClient("", "")

			Convey("Should bubble an error due to empty strings", func() {
				So(err, ShouldNotEqual, nil)
			})
			Convey("Since empty strings, should return a bad helper", func() {
				So(telapi_helper.sid, ShouldEqual, "")
			})
		})

		Convey("calling method CreateClient with bad credentials", func() {
			telapi_helper, err := CreateClient("test1", "test1")

			Convey("Should bubble an error due to invalid credentials on telapi end", func() {
				So(err, ShouldNotEqual, nil)
			})
			Convey("Since error bubbled should return a bad helper", func() {
				So(telapi_helper.sid, ShouldEqual, "")
			})
		})

		Convey("calling method CreateClient with good credentials", func() {
			telapi_helper, err := CreateClient(testing_telapi_sid, testing_telapi_auth_token)

			Convey("Should be no error, due to good credentials", func() {
				So(err, ShouldEqual, nil)
			})
			Convey("The sid should be correctly set.", func() {
				So(telapi_helper.sid, ShouldEqual, testing_telapi_sid)
			})
		})

	})

}
