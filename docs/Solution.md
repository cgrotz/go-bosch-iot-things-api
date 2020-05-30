# Solution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SolutionId** | **string** | Unique identifier representing a specific Solution | 
**ApiToken** | **string** | A token to identify requests of users associated with the solution | 
**Plan** | [**PlanType**](PlanType.md) |  | 
**PlanObject** | [**SolutionPlanObject**](Solution_planObject.md) |  | 
**ServiceName** | **string** | The name of the service | 
**Status** | [**SolutionStatus**](Solution_status.md) |  | 
**Customer** | [**SolutionCustomer**](Solution_customer.md) |  | 
**Namespaces** | [**Namespaces**](Namespaces.md) |  | [optional] 
**EffectiveUsageLimits** | [**UsageLimits**](UsageLimits.md) |  | 
**Clients** | [**Clients**](Clients.md) |  | [optional] 
**PolicyId** | **string** | The policy ID used for controlling access to this solution, managed by resource &#x60;/policies/{policyId}&#x60; | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


