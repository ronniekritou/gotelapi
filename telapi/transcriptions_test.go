package telapi

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTranscribeRecording(t *testing.T) {

	var (
		err           error
		telapi_helper TelapiHelper
		resp          map[string]interface{}
	)

	Convey("Tests when TranscribeRecording method called ", t, func() {

		Convey("Should not have an error, bc correct credentials", func() {
			telapi_helper, err = CreateClient(testing_telapi_sid, testing_telapi_auth_token)

			So(err, ShouldEqual, nil)
		})

		Convey("Should blow up because no voicemail sid", func() {
			resp, err := telapi_helper.TranscribeRecording("", "")

			So(err, ShouldNotEqual, nil)
			So(resp, ShouldEqual, nil)
		})

		Convey("Should have no errors", func() {
			resp, err = telapi_helper.TranscribeRecording("RE4c88908451eaa184b4514b528142505c", "")

			So(err, ShouldEqual, nil)
			So(resp["transcription_text"], ShouldNotEqual, "")
		})

	})
}

func TestTranscribeAudioUrl(t *testing.T) {

	var (
		err           error
		telapi_helper TelapiHelper
		resp          map[string]interface{}
	)

	Convey("Tests when TranscribeAudioUrl method called ", t, func() {

		Convey("Should not have an error, bc correct credentials", func() {
			telapi_helper, err = CreateClient(testing_telapi_sid, testing_telapi_auth_token)

			So(err, ShouldEqual, nil)
		})

		Convey("Should blow up because no url", func() {
			resp, err := telapi_helper.TranscribeAudioUrl("", "")

			So(err, ShouldNotEqual, nil)
			So(resp, ShouldEqual, nil)
		})

		Convey("Should have no errors", func() {
			resp, err = telapi_helper.TranscribeAudioUrl("http://recordings.telapi.com/RB47f8c862ab4a478095e5c07252d99f84/REbb8890842aaee99deed04aef836a0679.mp3", "")

			So(err, ShouldEqual, nil)
			So(resp["transcription_text"], ShouldNotEqual, "")
		})

	})
}

func TestViewTranscription(t *testing.T) {

	var (
		err           error
		telapi_helper TelapiHelper
		resp          map[string]interface{}
	)

	Convey("Tests when TranscribeAudioUrl method called ", t, func() {

		Convey("Should not have an error, bc correct credentials", func() {
			telapi_helper, err = CreateClient(testing_telapi_sid, testing_telapi_auth_token)

			So(err, ShouldEqual, nil)
		})

		Convey("Should blow up because no url", func() {
			resp, err := telapi_helper.ViewTranscription("")

			So(err, ShouldNotEqual, nil)
			So(resp, ShouldEqual, nil)
		})

		Convey("Should have no errors", func() {
			resp, err = telapi_helper.ViewTranscription("TR3c889084986054ad121d40c383db39ec")

			So(err, ShouldEqual, nil)
			So(resp, ShouldNotEqual, nil)
		})

	})
}
