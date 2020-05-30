# \PoliciesApi

All URIs are relative to *https://things.eu-1.bosch-iot-suite.com/api/2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PoliciesPolicyIdDelete**](PoliciesApi.md#PoliciesPolicyIdDelete) | **Delete** /policies/{policyId} | Delete a specific policy
[**PoliciesPolicyIdEntriesGet**](PoliciesApi.md#PoliciesPolicyIdEntriesGet) | **Get** /policies/{policyId}/entries | Retrieve the entries of a specific policy
[**PoliciesPolicyIdEntriesLabelDelete**](PoliciesApi.md#PoliciesPolicyIdEntriesLabelDelete) | **Delete** /policies/{policyId}/entries/{label} | Delete the entries of a specific Label of a specific policy
[**PoliciesPolicyIdEntriesLabelGet**](PoliciesApi.md#PoliciesPolicyIdEntriesLabelGet) | **Get** /policies/{policyId}/entries/{label} | Retrieve the entries of a specific Label of a specific policy
[**PoliciesPolicyIdEntriesLabelPut**](PoliciesApi.md#PoliciesPolicyIdEntriesLabelPut) | **Put** /policies/{policyId}/entries/{label} | Create or modify the entries of a specific Label of a specific policy
[**PoliciesPolicyIdEntriesLabelResourcesGet**](PoliciesApi.md#PoliciesPolicyIdEntriesLabelResourcesGet) | **Get** /policies/{policyId}/entries/{label}/resources | Retrieve all Resources for a specific Label of a specific policy
[**PoliciesPolicyIdEntriesLabelResourcesPut**](PoliciesApi.md#PoliciesPolicyIdEntriesLabelResourcesPut) | **Put** /policies/{policyId}/entries/{label}/resources | Create or modify all Resources for a specific Label of a specific policy
[**PoliciesPolicyIdEntriesLabelResourcesResourcePathDelete**](PoliciesApi.md#PoliciesPolicyIdEntriesLabelResourcesResourcePathDelete) | **Delete** /policies/{policyId}/entries/{label}/resources/{resourcePath} | Delete one specific Resource for a specific Label of a specific policy
[**PoliciesPolicyIdEntriesLabelResourcesResourcePathGet**](PoliciesApi.md#PoliciesPolicyIdEntriesLabelResourcesResourcePathGet) | **Get** /policies/{policyId}/entries/{label}/resources/{resourcePath} | Retrieve one specific Resource for a specific Label of a specific policy
[**PoliciesPolicyIdEntriesLabelResourcesResourcePathPut**](PoliciesApi.md#PoliciesPolicyIdEntriesLabelResourcesResourcePathPut) | **Put** /policies/{policyId}/entries/{label}/resources/{resourcePath} | Create or modify one specific Resource for a specific Label of a specific policy
[**PoliciesPolicyIdEntriesLabelSubjectsGet**](PoliciesApi.md#PoliciesPolicyIdEntriesLabelSubjectsGet) | **Get** /policies/{policyId}/entries/{label}/subjects | Retrieve all Subjects for a specific Label of a specific policy
[**PoliciesPolicyIdEntriesLabelSubjectsPut**](PoliciesApi.md#PoliciesPolicyIdEntriesLabelSubjectsPut) | **Put** /policies/{policyId}/entries/{label}/subjects | Create or modify all Subjects for a specific Label of a specific policy
[**PoliciesPolicyIdEntriesLabelSubjectsSubjectIdDelete**](PoliciesApi.md#PoliciesPolicyIdEntriesLabelSubjectsSubjectIdDelete) | **Delete** /policies/{policyId}/entries/{label}/subjects/{subjectId} | Delete one specific Subject for a specific Label of a specific policy
[**PoliciesPolicyIdEntriesLabelSubjectsSubjectIdGet**](PoliciesApi.md#PoliciesPolicyIdEntriesLabelSubjectsSubjectIdGet) | **Get** /policies/{policyId}/entries/{label}/subjects/{subjectId} | Retrieve one specific Subject for a specific Label of a specific policy
[**PoliciesPolicyIdEntriesLabelSubjectsSubjectIdPut**](PoliciesApi.md#PoliciesPolicyIdEntriesLabelSubjectsSubjectIdPut) | **Put** /policies/{policyId}/entries/{label}/subjects/{subjectId} | Create or modify one specific Subject for a specific Label of a specific policy
[**PoliciesPolicyIdEntriesPut**](PoliciesApi.md#PoliciesPolicyIdEntriesPut) | **Put** /policies/{policyId}/entries | Modify the entries of a specific policy
[**PoliciesPolicyIdGet**](PoliciesApi.md#PoliciesPolicyIdGet) | **Get** /policies/{policyId} | Retrieve a specific policy
[**PoliciesPolicyIdPut**](PoliciesApi.md#PoliciesPolicyIdPut) | **Put** /policies/{policyId} | Create or update a policy with a specified ID



## PoliciesPolicyIdDelete

> PoliciesPolicyIdDelete(ctx, policyId, optional)

Delete a specific policy

Deletes the policy identified by the `policyId` path parameter. Deleting a policy does not implicitly delete other entities (e.g. things) which use this policy.  Note: Delete the respective things **before** deleting the policy, otherwise nobody has permission to read, update, or delete the things. If you accidentally run into such a scenario, re-create the policy via PUT `/policies/{policyId}`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**| The ID of the policy needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to: * conform to the reverse domain name notation * be registered for your solution | 
 **optional** | ***PoliciesPolicyIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyIdDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous GET response, e.g. &#x60;If-Match: \&quot;rev:4711\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 

### Return type

 (empty response body)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyIdEntriesGet

> PolicyEntries PoliciesPolicyIdEntriesGet(ctx, policyId, optional)

Retrieve the entries of a specific policy

Returns all policy entries of the policy identified by the `policyId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**| The ID of the policy needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to: * conform to the reverse domain name notation * be registered for your solution | 
 **optional** | ***PoliciesPolicyIdEntriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyIdEntriesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 

### Return type

[**PolicyEntries**](PolicyEntries.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyIdEntriesLabelDelete

> PoliciesPolicyIdEntriesLabelDelete(ctx, policyId, label, optional)

Delete the entries of a specific Label of a specific policy

Deletes the entry of the policy identified by the `policyId` path parameter and with the label identified by the `label` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**| The ID of the policy needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to: * conform to the reverse domain name notation * be registered for your solution | 
**label** | **string**| The label of a policy entry | 
 **optional** | ***PoliciesPolicyIdEntriesLabelDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyIdEntriesLabelDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 

### Return type

 (empty response body)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyIdEntriesLabelGet

> PolicyEntry PoliciesPolicyIdEntriesLabelGet(ctx, policyId, label, optional)

Retrieve the entries of a specific Label of a specific policy

Returns all entries (subjects, resources, etc.) of the policy identified by the `policyId` path parameter, and by the `label` path parameter. Example label: DEFAULT.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**| The ID of the policy needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to: * conform to the reverse domain name notation * be registered for your solution | 
**label** | **string**| The label of a policy entry | 
 **optional** | ***PoliciesPolicyIdEntriesLabelGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyIdEntriesLabelGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 

### Return type

[**PolicyEntry**](PolicyEntry.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyIdEntriesLabelPut

> PolicyEntry PoliciesPolicyIdEntriesLabelPut(ctx, policyId, label, policyEntry, optional)

Create or modify the entries of a specific Label of a specific policy

Create or modify the policy entry identified by the `policyId` path parameter and with the label identified by the `label` path parameter. * If you specify a new label, the respective policy entry will be created * If you specify an existig label, the respective policy entry will be updated

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**| The ID of the policy needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to: * conform to the reverse domain name notation * be registered for your solution | 
**label** | **string**| The label of a policy entry | 
**policyEntry** | [**PolicyEntry**](PolicyEntry.md)| JSON representation of the policy entry. Use the placeholder &#x60;{{ request:subjectId }}&#x60; in order to let the backend insert the authenticated subjectId of the HTTP request.  ### Example Given your policy \&quot;com.acme.coffeemaker:policy-01\&quot; only has the DEFAULT entry, and you want to add a \&quot;Consumer\&quot; section which additionally allows USER-01 (managed within a Bosch IoT Permissions service instance) to *read* the thing and to trigger a \&quot;makeCoffee\&quot; operation (i.e. POST such a message - see POST /things/{thingId}/inbox/messages/{messageSubject}).  Set the label value to **Consumer** and the following request body: &#x60;&#x60;&#x60; {    \&quot;subjects\&quot;: {      \&quot;iot-permissions:USER-01\&quot;: {        \&quot;type\&quot;: \&quot;iot-permissions-userid\&quot;      },    },    \&quot;resources\&quot;: {      \&quot;thing:/\&quot;: {        \&quot;grant\&quot;: [          \&quot;READ\&quot;        ],        \&quot;revoke\&quot;: []      },      \&quot;message:/\&quot;: {        \&quot;grant\&quot;: [          \&quot;WRITE\&quot;        ],        \&quot;revoke\&quot;: []      }    }  } &#x60;&#x60;&#x60; | 
 **optional** | ***PoliciesPolicyIdEntriesLabelPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyIdEntriesLabelPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 

### Return type

[**PolicyEntry**](PolicyEntry.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyIdEntriesLabelResourcesGet

> Resources PoliciesPolicyIdEntriesLabelResourcesGet(ctx, policyId, label, optional)

Retrieve all Resources for a specific Label of a specific policy

Returns all resource entries of the policy identified by the `policyId` path parameter, and by the `label` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**| The ID of the policy needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to: * conform to the reverse domain name notation * be registered for your solution | 
**label** | **string**| The label of a policy entry | 
 **optional** | ***PoliciesPolicyIdEntriesLabelResourcesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyIdEntriesLabelResourcesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 

### Return type

[**Resources**](Resources.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyIdEntriesLabelResourcesPut

> PoliciesPolicyIdEntriesLabelResourcesPut(ctx, policyId, label, resources, optional)

Create or modify all Resources for a specific Label of a specific policy

Create or modify all resources of the policy identified by the `policyId` path parameter, and by the `label` path parameter.  ### Delete all resource entries  Set the empty body part, if you need to delete all resource entries: { }  ### Set max permissions on all ressources ``` {           \"policy:/\": {                   \"grant\": [                     \"READ\",                     \"WRITE\"                   ],                   \"revoke\": []                 },          \"thing:/\": {                   \"grant\": [                     \"READ\",                     \"WRITE\"                   ],                   \"revoke\": []                 },          \"message:/\": {                   \"grant\": [                     \"READ\",                     \"WRITE\"                   ],                   \"revoke\": []                 }               }   }   ```   ### Allow to read all parts of a thing except the \"confidential\" feature   ```   {           \"thing:/\": {                     \"grant\": [ \"READ\" ],                     \"revoke\": []           },           \"things:/{thingId}/features/confidential:/\": {                     \"grant\": [],                     \"revoke\": [ \"READ\"]           },   }   ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**| The ID of the policy needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to: * conform to the reverse domain name notation * be registered for your solution | 
**label** | **string**| The label of a policy entry | 
**resources** | [**Resources**](Resources.md)| JSON representation of the Resources | 
 **optional** | ***PoliciesPolicyIdEntriesLabelResourcesPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyIdEntriesLabelResourcesPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 

### Return type

 (empty response body)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyIdEntriesLabelResourcesResourcePathDelete

> PoliciesPolicyIdEntriesLabelResourcesResourcePathDelete(ctx, policyId, label, resourcePath, optional)

Delete one specific Resource for a specific Label of a specific policy

Deletes the resource with path `resourcePath` from the policy identified by the the `policyId` path parameter, and by the `label` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**| The ID of the policy needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to: * conform to the reverse domain name notation * be registered for your solution | 
**label** | **string**| The label of a policy entry | 
**resourcePath** | **string**| The path of an (Authorization) Resource | 
 **optional** | ***PoliciesPolicyIdEntriesLabelResourcesResourcePathDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyIdEntriesLabelResourcesResourcePathDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 

### Return type

 (empty response body)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyIdEntriesLabelResourcesResourcePathGet

> ResourceEntry PoliciesPolicyIdEntriesLabelResourcesResourcePathGet(ctx, policyId, label, resourcePath, optional)

Retrieve one specific Resource for a specific Label of a specific policy

Returns the resource with path `resourcePath` of the policy identified by the `policyId` path parameter, and by the `label` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**| The ID of the policy needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to: * conform to the reverse domain name notation * be registered for your solution | 
**label** | **string**| The label of a policy entry | 
**resourcePath** | **string**| The path of an (Authorization) Resource | 
 **optional** | ***PoliciesPolicyIdEntriesLabelResourcesResourcePathGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyIdEntriesLabelResourcesResourcePathGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 

### Return type

[**ResourceEntry**](ResourceEntry.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyIdEntriesLabelResourcesResourcePathPut

> ResourceEntry PoliciesPolicyIdEntriesLabelResourcesResourcePathPut(ctx, policyId, label, resourcePath, resourceEntry, optional)

Create or modify one specific Resource for a specific Label of a specific policy

Create or modify the Resource with path `resourcePath` of the policy entry identified by the `label` path parameter belonging to the policy identified by the `policyId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**| The ID of the policy needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to: * conform to the reverse domain name notation * be registered for your solution | 
**label** | **string**| The label of a policy entry | 
**resourcePath** | **string**| The path of an (Authorization) Resource | 
**resourceEntry** | [**ResourceEntry**](ResourceEntry.md)| JSON representation of the Resource | 
 **optional** | ***PoliciesPolicyIdEntriesLabelResourcesResourcePathPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyIdEntriesLabelResourcesResourcePathPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 

### Return type

[**ResourceEntry**](ResourceEntry.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyIdEntriesLabelSubjectsGet

> Subjects PoliciesPolicyIdEntriesLabelSubjectsGet(ctx, policyId, label, optional)

Retrieve all Subjects for a specific Label of a specific policy

Returns all subject entries of the policy identified by the `policyId` path parameter, and by the `label` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**| The ID of the policy needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to: * conform to the reverse domain name notation * be registered for your solution | 
**label** | **string**| The label of a policy entry | 
 **optional** | ***PoliciesPolicyIdEntriesLabelSubjectsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyIdEntriesLabelSubjectsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 

### Return type

[**Subjects**](Subjects.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyIdEntriesLabelSubjectsPut

> PoliciesPolicyIdEntriesLabelSubjectsPut(ctx, policyId, label, subjects, optional)

Create or modify all Subjects for a specific Label of a specific policy

Create or modify at once ALL subjects of the policy entry identified by the `policyId` path parameter, and by the `label` path parameter.  ### Example - delete all subjects To delete all subjects set an empty body { }  ### Example - entities of Bosch IoT Permissions To add a user, a goup, and a role managed within Bosch IoT Permissions  ``` {   \"iot-permissions:ID-user\": {     \"type\": \"iot-permissions-userid\"   },   \"iot-permissions:ID-group\": {     \"type\": \"iot-permissions-groupid\"   },   \"iot-permissions:ID-role\": {     \"type\": \"iot-permissions-roleid\"   } } ```  ### Example - technical clients To add a technical client ``` {    \"iot-things:Your-Things-Solution-ID-user:your-postfix\": {      \"type\": \"iot-things-clientid\"    },    \"iot-integration:your-Things-solution-ID-user:bosch-iot-hub\": {      \"type\": \"hub-amqp-clientid\"    },    \"iot-suite:service:iot-things-eu-1:Your-Things-Solution-ID_things/full-access\": {      \"type\": \"suite-auth\"    } } ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**| The ID of the policy needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to: * conform to the reverse domain name notation * be registered for your solution | 
**label** | **string**| The label of a policy entry | 
**subjects** | [**Subjects**](Subjects.md)| JSON representation of the Subjects.   Use the placeholder &#x60;{{ request:subjectId }}&#x60; in order to let the backend insert the authenticated subjectId of the HTTP request. | 
 **optional** | ***PoliciesPolicyIdEntriesLabelSubjectsPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyIdEntriesLabelSubjectsPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 

### Return type

 (empty response body)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyIdEntriesLabelSubjectsSubjectIdDelete

> PoliciesPolicyIdEntriesLabelSubjectsSubjectIdDelete(ctx, policyId, label, subjectId, optional)

Delete one specific Subject for a specific Label of a specific policy

Deletes the subject with ID `subjectId` from the policy identified by the `policyId` path parameter and by the `label` path parameter.  Note: If the subject is used in other labels, it will not be deleted there, i.e. it will not lose those permissions, but only the permissions defined in the label specified at this path.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**| The ID of the policy needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to: * conform to the reverse domain name notation * be registered for your solution | 
**label** | **string**| The label of a policy entry | 
**subjectId** | **string**| The ID of an (Authorization) Subject | 
 **optional** | ***PoliciesPolicyIdEntriesLabelSubjectsSubjectIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyIdEntriesLabelSubjectsSubjectIdDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 

### Return type

 (empty response body)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyIdEntriesLabelSubjectsSubjectIdGet

> SubjectEntry PoliciesPolicyIdEntriesLabelSubjectsSubjectIdGet(ctx, policyId, label, subjectId, optional)

Retrieve one specific Subject for a specific Label of a specific policy

Returns the subject with ID `subjectId` of the policy entry identified by the `policyId` path parameter, and by the `label` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**| The ID of the policy needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to: * conform to the reverse domain name notation * be registered for your solution | 
**label** | **string**| The label of a policy entry | 
**subjectId** | **string**| The ID of an (Authorization) Subject | 
 **optional** | ***PoliciesPolicyIdEntriesLabelSubjectsSubjectIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyIdEntriesLabelSubjectsSubjectIdGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 

### Return type

[**SubjectEntry**](SubjectEntry.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyIdEntriesLabelSubjectsSubjectIdPut

> SubjectEntry PoliciesPolicyIdEntriesLabelSubjectsSubjectIdPut(ctx, policyId, label, subjectId, subjectEntry, optional)

Create or modify one specific Subject for a specific Label of a specific policy

Create or modify the subject with ID `subjectId` of the policy identified by the `policyId` path parameter, and by the `label` path parameter.  ### Example -  add user managed within Bosch IoT Permissions To add a user managed within Bosch IoT Permissions, set the subjectId path to **iot-permissions:ID-user** and the body to  { \"type\": \"iot-permissions-userid\" }.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**| The ID of the policy needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to: * conform to the reverse domain name notation * be registered for your solution | 
**label** | **string**| The label of a policy entry | 
**subjectId** | **string**| The ID of an (Authorization) Subject | 
**subjectEntry** | [**SubjectEntry**](SubjectEntry.md)| JSON representation of the Subject | 
 **optional** | ***PoliciesPolicyIdEntriesLabelSubjectsSubjectIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyIdEntriesLabelSubjectsSubjectIdPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 

### Return type

[**SubjectEntry**](SubjectEntry.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyIdEntriesPut

> PoliciesPolicyIdEntriesPut(ctx, policyId, policyEntries, optional)

Modify the entries of a specific policy

Modify the policy entries of the policy identified by the `policyId` path parameter.  Note: Take care to not lock yourself out. Use the placeholder {{ request:subjectId }} in order to let the backend insert the authenticated subjectId of the HTTP request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**| The ID of the policy needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to: * conform to the reverse domain name notation * be registered for your solution | 
**policyEntries** | [**PolicyEntries**](PolicyEntries.md)| JSON representation of the policy entries. Use the placeholder &#x60;{{ request:subjectId }}&#x60; in order to let the backend insert the authenticated subjectId of the HTTP request. | 
 **optional** | ***PoliciesPolicyIdEntriesPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyIdEntriesPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 

### Return type

 (empty response body)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyIdGet

> Policy PoliciesPolicyIdGet(ctx, policyId, optional)

Retrieve a specific policy

Returns the complete policy identified by the `policyId` path parameter. The response contains the policy as JSON object.  Tip: If you don't know the policy ID, request it via GET ​/things​/{thingId}.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**| The ID of the policy needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to: * conform to the reverse domain name notation * be registered for your solution | 
 **optional** | ***PoliciesPolicyIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyIdGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous GET response, e.g. &#x60;If-Match: \&quot;rev:4711\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 

### Return type

[**Policy**](Policy.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PoliciesPolicyIdPut

> Policy PoliciesPolicyIdPut(ctx, policyId, policy, optional)

Create or update a policy with a specified ID

Create or update the policy specified by the policyId path parameter. * If you set a new policyId in the path, a new policy will be created. * If you set an existing policyId in the path, the policy will be updated.  ### Create a new policy At the initial creation of a policy, at least one valid entry is required. However, you can create a full-fledged policy all at once.  Example: To create a policy for multiple coffee maker things, which gives **yourself** all permissions on all resources, set the policyId in the path, e.g. to \"com.acme.coffeemaker:policy-01\" and the body part, like in the following snippet.  ``` {   \"entries\": {     \"DEFAULT\": {       \"subjects\": {         \"{{ request:subjectId }}\": {           \"type\": \"bosch-id\"         }       },       \"resources\": {         \"policy:/\": {           \"grant\": [             \"READ\",             \"WRITE\"           ],           \"revoke\": []         },         \"thing:/\": {           \"grant\": [             \"READ\",             \"WRITE\"           ],           \"revoke\": []         },         \"message:/\": {           \"grant\": [             \"READ\",             \"WRITE\"           ],           \"revoke\": []         }       }     }   } } ```  ### Update an existing policy For updating an existing policy, the authorized subject needs WRITE permission on the policy's root resource.  The ID of a policy cannot be changed after creation. Any `policyId` specified in the request body is therefore ignored.  ### Partially update an existing policy Partial updates are not supported.  If you need to create or update a specific label, resource, or subject, please use the respective sub-resources.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string**| The ID of the policy needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to: * conform to the reverse domain name notation * be registered for your solution | 
**policy** | [**Policy**](Policy.md)| JSON representation of the policy. Use the placeholder &#x60;{{ request:subjectId }}&#x60; in order to let the backend insert the authenticated subjectId of the HTTP request. | 
 **optional** | ***PoliciesPolicyIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PoliciesPolicyIdPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous GET response, e.g. &#x60;If-Match: \&quot;rev:4711\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 

### Return type

[**Policy**](Policy.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

