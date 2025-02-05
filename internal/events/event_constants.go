package events

const (
	// SummaryEventsSchemaVersion is the minimum event schema that supports summary events.
	SummaryEventsSchemaVersion = 3

	// CurrentEventsSchemaVersion is the latest event schema version.
	CurrentEventsSchemaVersion = 4

	// EventSchemaHeader is an HTTP header that describes the schema version for event requests.
	EventSchemaHeader = "X-LaunchDarkly-Event-Schema"

	// TagsHeader is an HTTP header that may be sent by SDKs that support application metadata.
	// We copy the value of this header when proxying events.
	TagsHeader = "X-LaunchDarkly-Tags"
)
