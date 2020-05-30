# NamespacePurgingState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the status report. | [optional] [default to solutions.responses:retrieveNamespacePurgingState]
**Status** | **int32** | HTTP status code of the status report. | [optional] [default to 200]
**NamespaceName** | **string** | The namespace being deleted permanently. | [optional] 
**CurrentStep** | **int32** | Ordinal number of the current step in the namespace deletion process. | [optional] 
**TotalSteps** | **int32** | Amount of steps in the namespace deletion process. | [optional] 
**Description** | **string** | Description of the current step in the namespace deletion process. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


