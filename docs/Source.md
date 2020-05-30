# Source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** | The source addresses this connection consumes messages from | [optional] 
**ConsumerCount** | **int32** | The number of consumers that should be attached to each source address | [optional] [default to 1]
**AuthorizationContext** | **[]string** | The authorization context defines all authorization subjects associated for this source | [optional] 
**Enforcement** | [**EnforcementConfiguration**](Enforcement_configuration.md) |  | [optional] 
**PayloadMapping** | **[]string** | A list of payload mappings that are applied to messages received via this source. If no payload mapping is specified the standard Ditto mapping is used as default. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


