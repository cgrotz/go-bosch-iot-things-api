/*
 * Bosch IoT Things HTTP API
 *
 * Bosch IoT Things enables applications to manage digital twins of IoT device assets in a simple, convenient, robust, and secure way.  These descriptions focus on the JSON-based, REST-like **HTTP API 2** of the Bosch IoT Things service.  Find details in our [documentation](https://docs.bosch-iot-suite.com/things/).
 *
 * API version: 2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Connection struct for Connection
type Connection struct {
	// The generated unique identifier of the connection
	Id string `json:"id,omitempty"`
	// The name of the connection
	Name string `json:"name,omitempty"`
	ConnectionType ConnectionType `json:"connectionType"`
	ConnectionStatus ConnectionStatus `json:"connectionStatus"`
	// The URI of the connection
	Uri string `json:"uri"`
	// The subscription sources of this connection
	Sources []Source `json:"sources"`
	// The publish targets of this connection
	Targets []Target `json:"targets"`
	// Configuration which is only applicable for a specific connection type
	SpecificConfig map[string]interface{} `json:"specificConfig,omitempty"`
	// How many clients on different cluster nodes should establish the connection
	ClientCount float32 `json:"clientCount,omitempty"`
	// Whether or not failover is enabled for this connection
	FailoverEnabled bool `json:"failoverEnabled,omitempty"`
	// Whether or not to validate server certificates on connection establishment
	ValidateCertificates bool `json:"validateCertificates,omitempty"`
	// List of mapping definitions where the key represents the ID of each mapping that can be used in sources and targets to reference a mapping.
	MappingDefinitions map[string]PayloadMappingDefinition `json:"mappingDefinitions,omitempty"`
	MappingContext MappingContext `json:"mappingContext,omitempty"`
	// The tags of the connection
	Tags []string `json:"tags,omitempty"`
}
