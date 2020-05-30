# PayloadMappingDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MappingEngine** | **string** | The mapping engine used to process incoming and outgoing messages. Available mapping engines are &#x60;JavaScript&#x60;, &#x60;Normalized&#x60;, &#x60;ConnectionStatus&#x60;, &#x60;Ditto&#x60;. | 
**Options** | [**map[string]interface{}**](.md) | Configuration options specific to the used mapping engine:  #### JavaScript   * &#x60;incomingScript&#x60; (&#x60;string&#x60;, required): The mapping script for incoming messages   * &#x60;outgoingScript&#x60; (&#x60;string&#x60;, required): The mapping script for outgoing messages   * &#x60;loadBytebufferJS&#x60; (&#x60;boolean&#x60;, optional): Whether or not ByteBufferJS library should be included     (default: &#x60;false&#x60;)   * &#x60;loadLongJS&#x60; (&#x60;boolean&#x60;, optional): Whether or not LongJS library should be included (default: &#x60;false&#x60;)  #### Normalized   * &#x60;fields&#x60; (&#x60;string&#x60;, optional): Comma separated list of fields included in the normalized message     (default: all fields included)  #### ConnectionStatus   * &#x60;thingId&#x60; (&#x60;string&#x60;, required): The ID of the thing   * &#x60;featureId&#x60; (&#x60;string&#x60;, optional): The ID of the modified feature (default: &#x60;ConnectionStatus&#x60;)  #### Ditto   * no options required | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


