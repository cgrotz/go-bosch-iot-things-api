# EnforcementConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | **string** | The input value of the enforcement that should identify the origin of the message (e.g. a device id). You can use placeholders within this field depending on the connection type. E.g. for AMQP 1.0 connections you can use &#x60;{{ header:[any-header-name] }}&#x60; to resolve the value from a message header. | 
**Filters** | **[]string** | An array of filters. One of the defined filters must match the input value from the message otherwise the message is rejected. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


