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
			conference, err = telapi_helper.FindConferenceByFriendlyName("347249516")
			So(err, ShouldBeNil)
			So(conference, ShouldBeNil)
			// So(conference.FriendlyName, ShouldEqual, "347249516")
			So(conference.Sid, ShouldEqual, "CF4d8890840765528ab03a476598bc15fa")
			So(conference.ActiveParticipantsCount, ShouldEqual, 0)
		})

	})
}
