package constants

const (
	MethodGET  = `GET`
	MethodPOST = `POST`
	MethodPUT  = "PUT"

	HeaderContentType = "Content-Type"
	HeaderApiKey = "apikey"
	//JSONType request api options type for json request
	JSONType        = "application/json"

	// ISOTimeLayout ISO standard time layout without timezone (with timezone use time.RFC3339 instead)
	ISOTimeLayout = "2006-01-02T15:04:05"
)