TelAPI Helper
==============

In client_test.go, if you would like to run tests just update these variables correctly.

```sh
var (
	testing_telapi_sid        = "telapi_sid"
	testing_telapi_auth_token = "your_telapi_auth_token"
	testing_number_to         = "your_telapi_number"
	testing_number_from       = "other_testing_number" //Correlates to telapi sid being used
)
```


Then run go test.
