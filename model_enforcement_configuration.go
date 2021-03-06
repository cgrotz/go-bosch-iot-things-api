/*
 * Bosch IoT Things HTTP API
 *
 * Bosch IoT Things enables applications to manage digital twins of IoT device assets in a simple, convenient, robust, and secure way.  These descriptions focus on the JSON-based, REST-like **HTTP API 2** of the Bosch IoT Things service.  Find details in our [documentation](https://docs.bosch-iot-suite.com/things/).
 *
 * API version: 2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package things
// EnforcementConfiguration Defines an enforcement for this source to make sure that a device can only access its associated Thing.
type EnforcementConfiguration struct {
	// The input value of the enforcement that should identify the origin of the message (e.g. a device id). You can use placeholders within this field depending on the connection type. E.g. for AMQP 1.0 connections you can use `{{ header:[any-header-name] }}` to resolve the value from a message header.
	Input string `json:"input"`
	// An array of filters. One of the defined filters must match the input value from the message otherwise the message is rejected.
	Filters []string `json:"filters"`
}
