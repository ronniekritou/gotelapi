TelAPI Helper
==============

In client_test.go, if you would like to run tests just update these variables correctly.

```sh
var (
var (
	testTelapiSid       = "your telapi sid"
	testTelapiAuthToken = "your telapi auth"
	testNumberTo        = "Destination number of a text"
	testNumberFrom      = "Where the text should send from" //Correlates to telapi sid being used
	testCallSid         = "A random call sid for the call tests" //Also correlates to telapi sid being used
	testRecordingSid    = "A random recording sid on the account"
)
```


Then run go test.
