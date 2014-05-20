package telapi

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCarrierLookup(t *testing.T) {

	var (
		err           error
		telapi_helper TelapiHelper
		resp          string
	)

	Convey("Tests when CarrierLookup method called ", t, func() {

		Convey("Should not have an error, bc correct credentials", func() {
			telapi_helper, err = CreateClient(testing_telapi_sid, testing_telapi_auth_token)

			So(err, ShouldEqual, nil)
		})

		Convey("Should blow up because no phone number", func() {
			resp, err = telapi_helper.CarrierLookup("")

			So(err, ShouldNotEqual, nil)
			So(resp, ShouldEqual, "")
		})

		Convey("Should have no errors", func() {
			resp, err = telapi_helper.CarrierLookup("+13134333244")

			So(err, ShouldEqual, nil)

			So(resp, ShouldContainSubstring, "Sprint")
		})

	})
}
