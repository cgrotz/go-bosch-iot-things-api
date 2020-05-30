/*
 * Bosch IoT Things HTTP API
 *
 * Bosch IoT Things enables applications to manage digital twins of IoT device assets in a simple, convenient, robust, and secure way.  These descriptions focus on the JSON-based, REST-like **HTTP API 2** of the Bosch IoT Things service.  Find details in our [documentation](https://docs.bosch-iot-suite.com/things/).
 *
 * API version: 2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Client An OAuth 2.0 client which is authorized to perform requests on behalf of the solution.
type Client struct {
	// The ID of the OAuth client. The ID must be provided in the \"azp\" claim of a JWT.
	ClientId string `json:"clientId,omitempty"`
	// Issuer for authorization subjects which identifies the associated OAuth 2.0 authorization server.
	SubjectIssuer string `json:"subjectIssuer,omitempty"`
}