/*
 * Bosch IoT Things HTTP API
 *
 * Bosch IoT Things enables applications to manage digital twins of IoT device assets in a simple, convenient, robust, and secure way.  These descriptions focus on the JSON-based, REST-like **HTTP API 2** of the Bosch IoT Things service.  Find details in our [documentation](https://docs.bosch-iot-suite.com/things/).
 *
 * API version: 2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Subjects A SubjectEntry defines who is addressed. Example: iot-permissions:some-user-id. See all supported types at [policy > Who can be addressed](https://docs.bosch-iot-suite.com/things/basic-concepts/auth/auth_policy/#who-can-be-addressed)
type Subjects struct {
	IotPermissionssubjectId1 SubjectEntry `json:"iot-permissions:subjectId1,omitempty"`
	IotPermissionssubjectIdN SubjectEntry `json:"iot-permissions:subjectIdN,omitempty"`
}