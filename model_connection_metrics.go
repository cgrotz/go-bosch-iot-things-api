/*
 * Bosch IoT Things HTTP API
 *
 * Bosch IoT Things enables applications to manage digital twins of IoT device assets in a simple, convenient, robust, and secure way.  These descriptions focus on the JSON-based, REST-like **HTTP API 2** of the Bosch IoT Things service.  Find details in our [documentation](https://docs.bosch-iot-suite.com/things/).
 *
 * API version: 2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ConnectionMetrics Metrics of a connection
type ConnectionMetrics struct {
	// The connection ID
	ConnectionId string `json:"connectionId"`
	// Whether the connection metrics contains any failures
	ContainsFailures bool `json:"containsFailures"`
	ConnectionMetrics OverallConnectionMetrics `json:"connectionMetrics"`
	SourceMetrics SourceMetrics `json:"sourceMetrics"`
	TargetMetrics TargetMetrics `json:"targetMetrics"`
}