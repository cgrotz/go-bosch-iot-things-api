# Thing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThingId** | **string** | Unique identifier representing the thing | 
**PolicyId** | **string** | The ID of the policy which controls the access to this thing. policies are managed by resource &#x60;/policies/{policyId}&#x60; | 
**Definition** | **string** | A single fully qualified identifier of a definition in the form &#39;namespace:name:version&#39; | 
**Attributes** | [**map[string]interface{}**](.md) | An arbitrary JSON object describing the attributes of a thing. | 
**Features** | [**map[string]Feature**](Feature.md) | List of features where the key represents the &#x60;featureId&#x60; of each feature. The &#x60;featureId&#x60; key must be unique in the list. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


