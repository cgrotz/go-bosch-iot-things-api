/*
 * Bosch IoT Things HTTP API
 *
 * Bosch IoT Things enables applications to manage digital twins of IoT device assets in a simple, convenient, robust, and secure way.  These descriptions focus on the JSON-based, REST-like **HTTP API 2** of the Bosch IoT Things service.  Find details in our [documentation](https://docs.bosch-iot-suite.com/things/).
 *
 * API version: 2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package things
// SolutionCustomer struct for SolutionCustomer
type SolutionCustomer struct {
	// The name of the customer
	Name string `json:"name"`
	// The email address of the customer
	Email string `json:"email"`
	// Additional information about the customer
	Info string `json:"info"`
}
