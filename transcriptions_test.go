package telapi

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTranscribeRecording(t *testing.T) {

	var (
		err           error
		telapi_helper TelapiHelper
		resp          *Transcription
	)

	Convey("Tests when TranscribeRecording method called ", t, func() {

		Convey("Should not have an error, bc correct credentials", func() {
			telapi_helper, err = CreateClient(testTelapiSid, testTelapiAuthToken)

			So(err, ShouldBeNil)
		})

		Convey("Should blow up because no voicemail sid", func() {
			resp, err := telapi_helper.TranscribeRecording("", "", "")

			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})

		Convey("Should have no errors", func() {
			resp, err = telapi_helper.TranscribeRecording(testRecordingSid, "", "")

			So(err, ShouldBeNil)
			So(resp.TranscriptionText, ShouldNotBeNil)
		})

	})
}

func TestTranscribeAudioUrl(t *testing.T) {

	var (
		err           error
		telapi_helper TelapiHelper
		resp          *Transcription
	)

	Convey("Tests when TranscribeAudioUrl method called ", t, func() {

		Convey("Should not have an error, bc correct credentials", func() {
			telapi_helper, err = CreateClient(testTelapiSid, testTelapiAuthToken)

			So(err, ShouldBeNil)
		})

		Convey("Should blow up because no url", func() {
			resp, err := telapi_helper.TranscribeAudioUrl("", "", "")

			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})

		Convey("Should have no errors", func() {
			resp, err = telapi_helper.TranscribeAudioUrl("http://recordings.telapi.com/RB47f8c862ab4a478095e5c07252d99f84/REbb8890842aaee99deed04aef836a0679.mp3", "", "")

			So(err, ShouldBeNil)
			So(resp.TranscriptionText, ShouldNotBeNil)
		})

	})
}

func TestViewTranscription(t *testing.T) {

	var (
		err           error
		telapi_helper TelapiHelper
		resp          *Transcription
	)

	Convey("Tests when TranscribeAudioUrl method called ", t, func() {

		Convey("Should not have an error, bc correct credentials", func() {
			telapi_helper, err = CreateClient(testTelapiSid, testTelapiAuthToken)

			So(err, ShouldBeNil)
		})

		Convey("Should blow up because no url", func() {
			resp, err := telapi_helper.ViewTranscription("")

			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})

		Convey("Should have no errors", func() {
			resp, err = telapi_helper.ViewTranscription("TR3c889084986054ad121d40c383db39ec")

			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
		})

	})
}
