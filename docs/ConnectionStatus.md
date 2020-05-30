# ConnectionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | The connection ID | 
**ConnectionStatus** | [**ConnectivityStatus**](ConnectivityStatus.md) | The desired/target status of the connection | 
**LiveStatus** | [**ConnectivityStatus**](ConnectivityStatus.md) | The current/actual status of the connection | 
**ConnectedSince** | **string** | The timestamp since when the connection is connected | 
**ClientStatus** | [**[]ResourceStatus**](ResourceStatus.md) | The client states of the of the connection | 
**SourceStatus** | [**[]ResourceStatus**](ResourceStatus.md) | The states of the sources the of the connection | 
**TargetStatus** | [**[]ResourceStatus**](ResourceStatus.md) | The states of the targets the of the connection | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


