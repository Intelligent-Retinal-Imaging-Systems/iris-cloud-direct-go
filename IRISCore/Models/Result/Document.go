package Result

// Contains the result document payload.
type Document struct {
	// Actual Content, as requested by the caller.  If the caller did not request content, will be null.
	Content string
	// Base64, HL7 or specific encoding scheme.
	Encoding string
	// The type of document contained in the Content property (e.g.: PDF, HTML, ORU).
	Type string
}
