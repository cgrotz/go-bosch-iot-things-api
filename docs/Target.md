# Target

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The target address where events, commands and messages are published to. The following placeholders are allowed within the target address:  * Thing ID: &#x60;{{ thing:id }}&#x60;  * Thing Namespace: &#x60;{{ thing:namespace }}&#x60;  * Thing Name: &#x60;{{ thing:name }}&#x60; (the part of the ID without the namespace) | [optional] 
**Topics** | **[]string** | The topics to which this target is registered for | [optional] 
**AuthorizationContext** | **[]string** | The authorization context defines all authorization subjects associated for this target | [optional] 
**PayloadMapping** | **[]string** | A list of payload mappings that are applied to messages sent via this target. If no payload mapping is specified the standard Ditto mapping is used as default. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


