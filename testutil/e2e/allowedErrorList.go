package e2e

// This is a map of error strings with the description why they are allowed
// Detect extra text to be sure it is the allowed error
var allowedErrors = map[string]string{
	"getSupportedApi":        "This error is allowed because the Tendermint URI tests have a test that checks if the error is caught.",
	"No pairings available.": "This error is allowed because when the network is just booted up and pairings are not yet done this would happen. If after a few seconds the pairings are still not available the e2e would fail because the initial check if the provider is responsive would time out.",
	`error connecting to provider error="context deadline exceeded"`: "This error is allowed because it is caused by the initial bootup, continuous failure would be caught by the e2e so we can allowed this error.",
	"purging provider after all endpoints are disabled provider":     "This error is allowed because it is caused by the initial bootup, continuous failure would be caught by the e2e so we can allowed this error.",
	"Provider Side Failed Sending Message, Reason: Unavailable":      "This error is allowed because it is caused by the lavad restart to turn on emergency mode",
}
