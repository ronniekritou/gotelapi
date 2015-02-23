package telapi

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCarrierLookup(t *testing.T) {

	var (
		err           error
		telapi_helper TelapiHelper
		carrier       *CarrierLookup
	)

	Convey("Tests when CarrierLookup method called ", t, func() {

		Convey("Should not have an error, bc correct credentials", func() {
			telapi_helper, err = CreateClient(testing_telapi_sid, testing_telapi_auth_token)

			So(err, ShouldBeNil)
		})

		Convey("Should blow up because no phone number", func() {
			carrier, err = telapi_helper.CarrierLookup("")

			So(err, ShouldNotBeNil)
			So(carrier, ShouldBeNil)
		})

		Convey("Should have no errors", func() {
			carrier, err = telapi_helper.CarrierLookup("+17325807596")

			So(err, ShouldBeNil)
			So(carrier, ShouldNotBeNil)
			So(carrier.CarrierId, ShouldEqual, 20314)
		})

	})
}

func TestBNALookup(t *testing.T) {

	var (
		err           error
		telapi_helper TelapiHelper
		resp          *BnaLookup
	)

	Convey("Tests when BnaLookup method called ", t, func() {

		Convey("Should not have an error, bc correct credentials", func() {
			telapi_helper, err = CreateClient(testing_telapi_sid, testing_telapi_auth_token)

			So(err, ShouldBeNil)
		})

		Convey("Should blow up because no phone number", func() {
			resp, err = telapi_helper.BnaLookup("")

			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})

		Convey("Should have no errors", func() {
			resp, err = telapi_helper.BnaLookup("+13134333244")

			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
		})

	})
}

func TestCnamLookup(t *testing.T) {

	var (
		err           error
		telapi_helper TelapiHelper
		resp          *CnamLookup
	)

	Convey("Tests when BnaLookup method called ", t, func() {

		Convey("Should not have an error, bc correct credentials", func() {
			telapi_helper, err = CreateClient(testing_telapi_sid, testing_telapi_auth_token)

			So(err, ShouldBeNil)
		})

		Convey("Should blow up because no phone number", func() {
			resp, err = telapi_helper.CnamLookup("")

			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})

		Convey("Should have no errors", func() {
			resp, err = telapi_helper.CnamLookup("+13134333244")

			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
		})

	})
}
