package telapi

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFindConferenceByFriendlyName(t *testing.T) {

	var (
		err           error
		telapi_helper TelapiHelper
		conference    *Conference
	)

	Convey("Tests when findConference method has been called", t, func() {

		Convey("Should not have an error, bc correct credentials", func() {

			telapi_helper, err = CreateClient(testing_telapi_sid, testing_telapi_auth_token)

			So(err, ShouldBeNil)
		})

		Convey("Should blow up because no  sid", func() {
			conference, err = telapi_helper.FindConferenceByFriendlyName("")

			So(err, ShouldNotBeNil)
			So(conference, ShouldBeNil)
		})

		Convey("Should have no errors", func() {
			conference, err = telapi_helper.FindConferenceByFriendlyName("CF4d889084a46cc36f46794904a87a86d8")
			So(err, ShouldBeNil)
			So(conference, ShouldNotBeNil)
			So(conference.ActiveParticipantsCount, ShouldEqual, 0)
		})

	})
}
