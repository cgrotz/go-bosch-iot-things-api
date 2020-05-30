# NewConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the connection | [optional] 
**ConnectionType** | [**ConnectionType**](ConnectionType.md) |  | 
**ConnectionStatus** | [**ConnectionStatus**](ConnectionStatus.md) |  | 
**Uri** | **string** | The URI of the connection | 
**Sources** | [**[]Source**](Source.md) | The subscription sources of this connection | 
**Targets** | [**[]Target**](Target.md) | The publish targets of this connection | 
**SpecificConfig** | [**map[string]interface{}**](.md) | Configuration which is only applicable for a specific connection type | [optional] 
**ClientCount** | **float32** | How many clients on different cluster nodes should establish the connection | [optional] 
**FailoverEnabled** | **bool** | Whether or not failover is enabled for this connection | [optional] 
**ValidateCertificates** | **bool** | Whether or not to validate server certificates on connection establishment | [optional] 
**MappingDefinitions** | [**map[string]PayloadMappingDefinition**](PayloadMappingDefinition.md) | List of mapping definitions where the key represents the ID of each mapping that can be used in sources and targets to reference a mapping. | [optional] 
**MappingContext** | [**MappingContext**](MappingContext.md) |  | [optional] 
**Tags** | **[]string** | The tags of the connection | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


