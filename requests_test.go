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
			telapi_helper, err = CreateClient(testTelapiSid, testTelapiAuthToken)

			So(err, ShouldBeNil)
			So(telapi_helper, ShouldNotBeNil)
		})

		Convey("Should blow up due to bad endpoint", func() {
			_, err = telapi_helper.PostRequest("/SMS/Messges", nil)
			So(err, ShouldNotBeNil)
		})

		Convey("Should have no errors", func() {
			telapi_helper, err = CreateClient(testTelapiSid, testTelapiAuthToken)
			So(err, ShouldBeNil)
			So(telapi_helper, ShouldNotBeNil)

			sms, err := telapi_helper.SendSMS(testNumberTo, testNumberFrom, `TrapCall New Transcription
Cell Phone   NJ
(848) 210-6084
NEW JERSEY NJ 
(06/30/14 4:05 PM)
I make phone calls because Ronnie talk to you soon.
http://v.trapcall.com/445jccko

fklnfdlkdnldskdfd

fdkldsnskln
`)

			So(err, ShouldBeNil)
			So(sms, ShouldNotBeNil)

			So(sms.Sid, ShouldNotEqual, "")
			So(sms.Body, ShouldNotEqual, "")
			So(sms.To, ShouldNotEqual, "")
			So(sms.From, ShouldNotEqual, "")
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
			telapi_helper, err = CreateClient(testTelapiSid, testTelapiAuthToken)

			So(err, ShouldBeNil)
		})

		data := map[string]string{}

		Convey("Should blow up due to bad endpoint", func() {
			_, err = telapi_helper.GetRequest("Transcriptns", data)

			So(err, ShouldNotBeNil)
		})

		Convey("Should have no errors", func() {
			_, err = telapi_helper.GetRequest("/Transcriptions", data)

			So(err, ShouldBeNil)
		})

	})

}
