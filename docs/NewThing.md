# NewThing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | [**Policy**](Policy.md) | The initial policy to create for this thing. This will create a separate policy entity managed by resource &#x60;/policies/{thingId}&#x60;.   Use the placeholder &#x60;{{ request:subjectId }}&#x60; in order to let the backend insert the authenticated subjectId of the HTTP request. | [optional] 
**CopyPolicyFrom** | **string** | This field may contain the policy ID of an existing policy. The policy is copied and used for this newly created thing. This field may also contain a placeholder reference to a thing in the format &#x60;{{ ref:things/[thingId]/policyId }}&#x60; where you need to replace &#x60;[thingId]&#x60; with a valid thing ID. The newly created thing will then obtain a copy of the policy of the referenced thing. In the case of using a reference, the caller needs to have READ access to both the thing and the policy of the thing. In the case of using an explicit policy id to copy from, the caller needs to have READ access to the policy. If you want to specify a policy ID for the copied policy, use the policyId field. This field must not be used together with the field &#x60;_policy&#x60;. If you specify both &#x60;_policy&#x60; and &#x60;_copyPolicyFrom&#x60; this will lead to an error response. | [optional] 
**PolicyId** | **string** | The policy ID used for controlling access to this thing. Managed by resource &#x60;/policies/{policyId}&#x60;. | [optional] 
**Definition** | **string** | A single fully qualified identifier of a definition in the form &#39;namespace:name:version&#39; | [optional] 
**Attributes** | [**map[string]interface{}**](.md) | An arbitrary JSON object describing the attributes of a thing. | [optional] 
**Features** | [**map[string]Feature**](Feature.md) | List of features where the key represents the &#x60;featureId&#x60; of each feature. The &#x60;featureId&#x60; key must be unique in the list. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


