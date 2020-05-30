/*
 * Bosch IoT Things HTTP API
 *
 * Bosch IoT Things enables applications to manage digital twins of IoT device assets in a simple, convenient, robust, and secure way.  These descriptions focus on the JSON-based, REST-like **HTTP API 2** of the Bosch IoT Things service.  Find details in our [documentation](https://docs.bosch-iot-suite.com/things/).
 *
 * API version: 2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package things
// Source A subscription source subscribed by this connection
type Source struct {
	// The source addresses this connection consumes messages from
	Addresses []string `json:"addresses,omitempty"`
	// The number of consumers that should be attached to each source address
	ConsumerCount int32 `json:"consumerCount,omitempty"`
	// The authorization context defines all authorization subjects associated for this source
	AuthorizationContext []string `json:"authorizationContext,omitempty"`
	Enforcement EnforcementConfiguration `json:"enforcement,omitempty"`
	// A list of payload mappings that are applied to messages received via this source. If no payload mapping is specified the standard Ditto mapping is used as default.
	PayloadMapping []string `json:"payloadMapping,omitempty"`
}
