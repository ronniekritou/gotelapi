package telapi

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetAllIncomingNumbers(t *testing.T) {

	var (
		err           error
		telapi_helper TelapiHelper
		// call          *Call
	)

	Convey("Tests when MakeCall method has been called", t, func() {

		Convey("Should not have an error, bc correct credentials", func() {

			telapi_helper, err = CreateClient(testTelapiSid, testTelapiAuthToken)

			So(err, ShouldBeNil)
		})

		Convey("If we were to use the trapcall telapi credentials", func() {
			numbers, err := telapi_helper.GetAllIncomingNumbers()

			So(err, ShouldBeNil)
			So(numbers, ShouldNotBeNil)
			So(len(*numbers), ShouldBeLessThanOrEqualTo, 375)

		})

	})
}
