package telapi

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetCallData(t *testing.T) {

	var (
		err           error
		telapi_helper TelapiHelper
		resp          map[string]interface{}
	)

	Convey("Tests when GetCallData method has been called", t, func() {

		Convey("Should not have an error, bc correct credentials", func() {
			telapi_helper, err = CreateClient(testing_telapi_sid, testing_telapi_auth_token)

			So(err, ShouldEqual, nil)
		})

		Convey("Should blow up because no voicemail sid", func() {
			resp, err := telapi_helper.GetCallData("")

			So(err, ShouldNotEqual, nil)
			So(resp, ShouldEqual, nil)
		})

		Convey("Should have no errors", func() {
			resp, err = telapi_helper.GetCallData("CA6c88908491f4a2a7b2cd41478a408584")
			So(err, ShouldEqual, nil)
			So(resp["call_status"], ShouldNotEqual, "")
		})

	})
}

/*

CALL INFORMATION MAP PRINTED OUT

map[
	date_updated:Wed, 14 May 2014 11:16:57 -0000
	duration:46
	from:+15135054217
	sid:CA6c88908491f4a2a7b2cd41478a408584
	recordings_count:
	start_time:Wed, 14 May 2014 11:16:11 -0000
	subresource_uris:map[notifications:/v2/Accounts/lksnksdlknfdlknfdlknsfdklreplacementsid/Calls/CA6c88908491f4a2a7b2cd41478a408584/Notifications recordings:/v2/Accounts/AC1d530461c32a4840a1a19183d0a0bb8c/Calls/CA6c88908491f4a2a7b2cd41478a408584/Recordings]
	account_sid:'lksnksdlknfdlknfdlknsfdklreplacementsid'
	to:+17327305402
	answered_by:telapi
	price:0.01
	direction:inbound
	forwarded_from:
	duration_billed:60
	parent_call_sid:
	caller_id_blocked:false
	status:completed
	uri:/v2/Accounts/lksnksdlknfdlknfdlknsfdklreplacementsid/Calls/CA6c88908491f4a2a7b2cd41478a408584
	api_version:v2
	date_created:Wed, 14 May 2014 11:16:11 -0000
	end_time:Wed, 14 May 2014 11:16:57 -0000
	phone_number_sid:DIa0575b3c755911e1824d000c29cca3b8
]


*/
