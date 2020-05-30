/*
 * Bosch IoT Things HTTP API
 *
 * Bosch IoT Things enables applications to manage digital twins of IoT device assets in a simple, convenient, robust, and secure way.  These descriptions focus on the JSON-based, REST-like **HTTP API 2** of the Bosch IoT Things service.  Find details in our [documentation](https://docs.bosch-iot-suite.com/things/).
 *
 * API version: 2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SourceMetrics Source metrics of the connection
type SourceMetrics struct {
	// Contains \"inbound\" from external sources consumed metric counts
	AddressMetrics map[string]InboundMetrics `json:"addressMetrics"`
}
