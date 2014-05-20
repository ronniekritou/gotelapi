package telapi

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPostRequest(t *testing.T) {

	var (
		telapi_helper TelapiHelper
		err           error
	)

	Convey("Tests when PostRequest has been made ", t, func() {

		Convey("Should not have an error, bc correct credentials", func() {
			telapi_helper, err = CreateClient(testing_telapi_sid, testing_telapi_auth_token)

			So(err, ShouldEqual, nil)
		})

		data := map[string]string{
			"To":   testing_number_to,
			"From": testing_number_from,
			"Body": "This is a test",
		}

		Convey("Should blow up due to bad endpoint", func() {
			_, err = telapi_helper.PostRequest("/SMS/Messges", data)

			So(err, ShouldNotEqual, nil)
		})

		_, err = telapi_helper.PostRequest("/SMS/Messages", data)

		Convey("Should have no errors", func() {
			So(err, ShouldEqual, nil)
		})

	})

}

func TestGetRequest(t *testing.T) {

	var (
		telapi_helper TelapiHelper
		err           error
	)

	Convey("Tests when GetRequest has been made ", t, func() {

		Convey("Should not have an error, bc correct credentials", func() {
			telapi_helper, err = CreateClient(testing_telapi_sid, testing_telapi_auth_token)

			So(err, ShouldEqual, nil)
		})

		data := map[string]string{}

		Convey("Should blow up due to bad endpoint", func() {
			_, err = telapi_helper.GetRequest("Transcriptns", data)

			So(err, ShouldNotEqual, nil)
		})

		Convey("Should have no errors", func() {
			_, err = telapi_helper.GetRequest("/Transcriptions", data)

			So(err, ShouldEqual, nil)
		})

	})

}
