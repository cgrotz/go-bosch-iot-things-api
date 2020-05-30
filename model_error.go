/*
 * Bosch IoT Things HTTP API
 *
 * Bosch IoT Things enables applications to manage digital twins of IoT device assets in a simple, convenient, robust, and secure way.  These descriptions focus on the JSON-based, REST-like **HTTP API 2** of the Bosch IoT Things service.  Find details in our [documentation](https://docs.bosch-iot-suite.com/things/).
 *
 * API version: 2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package things
// Error struct for Error
type Error struct {
	// The HTTP status of the error
	Status int32 `json:"status"`
	// The message of the error - what went wrong
	Message string `json:"message"`
	// A description how to fix the error or more details
	Description string `json:"description,omitempty"`
	// A link to further information about the error and how to fix it
	Href string `json:"href,omitempty"`
}
