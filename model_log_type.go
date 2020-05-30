/*
 * Bosch IoT Things HTTP API
 *
 * Bosch IoT Things enables applications to manage digital twins of IoT device assets in a simple, convenient, robust, and secure way.  These descriptions focus on the JSON-based, REST-like **HTTP API 2** of the Bosch IoT Things service.  Find details in our [documentation](https://docs.bosch-iot-suite.com/things/).
 *
 * API version: 2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// LogType The type of a log entry describing during what kind of activity the entry was created.
type LogType string

// List of LogType
const (
	CONSUMED LogType = "consumed"
	DISPATCHED LogType = "dispatched"
	FILTERED LogType = "filtered"
	MAPPED LogType = "mapped"
	DROPPED LogType = "dropped"
	ENFORCED LogType = "enforced"
	PUBLISHED LogType = "published"
	OTHER LogType = "other"
)