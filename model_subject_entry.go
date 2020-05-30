/*
 * Bosch IoT Things HTTP API
 *
 * Bosch IoT Things enables applications to manage digital twins of IoT device assets in a simple, convenient, robust, and secure way.  These descriptions focus on the JSON-based, REST-like **HTTP API 2** of the Bosch IoT Things service.  Find details in our [documentation](https://docs.bosch-iot-suite.com/things/).
 *
 * API version: 2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SubjectEntry Single (Authorization) Subject entry holding its type.
type SubjectEntry struct {
	// The type is offered only for documentation purposes. You are not restricted to any specific types, but we recommend to use it to specify the kind of the subject as shown in our examples.
	Type string `json:"type"`
}