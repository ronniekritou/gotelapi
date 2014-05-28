package telapi

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSendSMS(t *testing.T) {

	var (
		err           error
		telapi_helper TelapiHelper
	)

	Convey("Tests when SendSMS method called ", t, func() {

		Convey("Should not have an error, bc correct credentials", func() {
			telapi_helper, err = CreateClient(testing_telapi_sid, testing_telapi_auth_token)

			So(err, ShouldEqual, nil)
		})

		Convey("Should blow up to no 'To' ", func() {
			err = telapi_helper.SendSMS("", "hey", "")

			So(err, ShouldNotEqual, nil)
		})

		Convey("Should blow up to no 'From' ", func() {
			err = telapi_helper.SendSMS("+", "", "")

			So(err, ShouldNotEqual, nil)
		})

		Convey("Should blow up to no 'Body' ", func() {
			err = telapi_helper.SendSMS("372", "hey", "")

			So(err, ShouldNotEqual, nil)
		})

		Convey("Should have no errors", func() {
			err = telapi_helper.SendSMS(testing_number_to, testing_number_from, "Testing SendSMS")

			So(err, ShouldEqual, nil)
		})

	})
}
