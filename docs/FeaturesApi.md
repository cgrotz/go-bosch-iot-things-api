# \FeaturesApi

All URIs are relative to *https://things.eu-1.bosch-iot-suite.com/api/2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThingsThingIdFeaturesDelete**](FeaturesApi.md#ThingsThingIdFeaturesDelete) | **Delete** /things/{thingId}/features | Delete all features of a specific thing
[**ThingsThingIdFeaturesFeatureIdDefinitionDelete**](FeaturesApi.md#ThingsThingIdFeaturesFeatureIdDefinitionDelete) | **Delete** /things/{thingId}/features/{featureId}/definition | Delete the definition of a feature
[**ThingsThingIdFeaturesFeatureIdDefinitionGet**](FeaturesApi.md#ThingsThingIdFeaturesFeatureIdDefinitionGet) | **Get** /things/{thingId}/features/{featureId}/definition | List the definition of a feature
[**ThingsThingIdFeaturesFeatureIdDefinitionPut**](FeaturesApi.md#ThingsThingIdFeaturesFeatureIdDefinitionPut) | **Put** /things/{thingId}/features/{featureId}/definition | Create or update the definition of a feature
[**ThingsThingIdFeaturesFeatureIdDelete**](FeaturesApi.md#ThingsThingIdFeaturesFeatureIdDelete) | **Delete** /things/{thingId}/features/{featureId} | Delete a specific feature of a specific thing
[**ThingsThingIdFeaturesFeatureIdGet**](FeaturesApi.md#ThingsThingIdFeaturesFeatureIdGet) | **Get** /things/{thingId}/features/{featureId} | Retrieve a specific feature of a specific thing
[**ThingsThingIdFeaturesFeatureIdPropertiesDelete**](FeaturesApi.md#ThingsThingIdFeaturesFeatureIdPropertiesDelete) | **Delete** /things/{thingId}/features/{featureId}/properties | Delete all properties of a feature
[**ThingsThingIdFeaturesFeatureIdPropertiesGet**](FeaturesApi.md#ThingsThingIdFeaturesFeatureIdPropertiesGet) | **Get** /things/{thingId}/features/{featureId}/properties | List all properties of a feature
[**ThingsThingIdFeaturesFeatureIdPropertiesPropertyPathDelete**](FeaturesApi.md#ThingsThingIdFeaturesFeatureIdPropertiesPropertyPathDelete) | **Delete** /things/{thingId}/features/{featureId}/properties/{propertyPath} | Delete a specific property of a feature
[**ThingsThingIdFeaturesFeatureIdPropertiesPropertyPathGet**](FeaturesApi.md#ThingsThingIdFeaturesFeatureIdPropertiesPropertyPathGet) | **Get** /things/{thingId}/features/{featureId}/properties/{propertyPath} | Retrieve a specific property of a feature
[**ThingsThingIdFeaturesFeatureIdPropertiesPropertyPathPut**](FeaturesApi.md#ThingsThingIdFeaturesFeatureIdPropertiesPropertyPathPut) | **Put** /things/{thingId}/features/{featureId}/properties/{propertyPath} | Create or update a specific property of a feature
[**ThingsThingIdFeaturesFeatureIdPropertiesPut**](FeaturesApi.md#ThingsThingIdFeaturesFeatureIdPropertiesPut) | **Put** /things/{thingId}/features/{featureId}/properties | Create or update all properties of a feature at once
[**ThingsThingIdFeaturesFeatureIdPut**](FeaturesApi.md#ThingsThingIdFeaturesFeatureIdPut) | **Put** /things/{thingId}/features/{featureId} | Create or modify a specific feature of a specific thing
[**ThingsThingIdFeaturesGet**](FeaturesApi.md#ThingsThingIdFeaturesGet) | **Get** /things/{thingId}/features | List all features of a specific thing
[**ThingsThingIdFeaturesPut**](FeaturesApi.md#ThingsThingIdFeaturesPut) | **Put** /things/{thingId}/features | Create or modify all features of a specific thing at once



## ThingsThingIdFeaturesDelete

> ThingsThingIdFeaturesDelete(ctx, thingId, optional)

Delete all features of a specific thing

Deletes all features of the thing identified by the `thingId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
 **optional** | ***ThingsThingIdFeaturesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdFeaturesDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **requestedAcks** | **optional.String**| Contains the \&quot;requested acknowledgements\&quot; for this modifying request as comma separated list. The HTTP call will block until all requested acknowledgements were aggregated or will time out based on the specified &#x60;timeout&#x60; parameter.  The default (if omitted) requested acks is &#x60;requested-acks&#x3D;twin-persisted&#x60; which will block the HTTP call until the change was persited to the twin. | 
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


## ThingsThingIdFeaturesFeatureIdDefinitionDelete

> ThingsThingIdFeaturesFeatureIdDefinitionDelete(ctx, thingId, featureId, optional)

Delete the definition of a feature

Deletes the complete definition of the feature identified by the `thingId` and `featureId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**featureId** | **string**| The ID of the feature - has to conform to RFC-3986 (URI) | 
 **optional** | ***ThingsThingIdFeaturesFeatureIdDefinitionDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdFeaturesFeatureIdDefinitionDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **requestedAcks** | **optional.String**| Contains the \&quot;requested acknowledgements\&quot; for this modifying request as comma separated list. The HTTP call will block until all requested acknowledgements were aggregated or will time out based on the specified &#x60;timeout&#x60; parameter.  The default (if omitted) requested acks is &#x60;requested-acks&#x3D;twin-persisted&#x60; which will block the HTTP call until the change was persited to the twin. | 
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


## ThingsThingIdFeaturesFeatureIdDefinitionGet

> []string ThingsThingIdFeaturesFeatureIdDefinitionGet(ctx, thingId, featureId, optional)

List the definition of a feature

Returns the complete definition field of the feature identified by the `thingId` and `featureId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**featureId** | **string**| The ID of the feature - has to conform to RFC-3986 (URI) | 
 **optional** | ***ThingsThingIdFeaturesFeatureIdDefinitionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdFeaturesFeatureIdDefinitionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 

### Return type

**[]string**

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsThingIdFeaturesFeatureIdDefinitionPut

> []string ThingsThingIdFeaturesFeatureIdDefinitionPut(ctx, thingId, featureId, requestBody, optional)

Create or update the definition of a feature

Create or update the complete definition of a feature identified by the `thingId` and `featureId` path parameter.  The definition field will be overwritten with the JSON array set in the request body

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**featureId** | **string**| The ID of the feature - has to conform to RFC-3986 (URI) | 
**requestBody** | [**[]string**](string.md)| JSON array of the complete definition to be updated. Consider that the value has to be a JSON array or &#x60;null&#x60;.  The content of the JSON array are strings in the format &#x60;\&quot;namespace:name:version\&quot;&#x60; which is enforced. | 
 **optional** | ***ThingsThingIdFeaturesFeatureIdDefinitionPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdFeaturesFeatureIdDefinitionPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **requestedAcks** | **optional.String**| Contains the \&quot;requested acknowledgements\&quot; for this modifying request as comma separated list. The HTTP call will block until all requested acknowledgements were aggregated or will time out based on the specified &#x60;timeout&#x60; parameter.  The default (if omitted) requested acks is &#x60;requested-acks&#x3D;twin-persisted&#x60; which will block the HTTP call until the change was persited to the twin. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 

### Return type

**[]string**

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsThingIdFeaturesFeatureIdDelete

> ThingsThingIdFeaturesFeatureIdDelete(ctx, thingId, featureId, optional)

Delete a specific feature of a specific thing

Deletes a specific feature identified by the `featureId` path parameter of the thing identified by the `thingId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**featureId** | **string**| The ID of the feature - has to conform to RFC-3986 (URI) | 
 **optional** | ***ThingsThingIdFeaturesFeatureIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdFeaturesFeatureIdDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **requestedAcks** | **optional.String**| Contains the \&quot;requested acknowledgements\&quot; for this modifying request as comma separated list. The HTTP call will block until all requested acknowledgements were aggregated or will time out based on the specified &#x60;timeout&#x60; parameter.  The default (if omitted) requested acks is &#x60;requested-acks&#x3D;twin-persisted&#x60; which will block the HTTP call until the change was persited to the twin. | 
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


## ThingsThingIdFeaturesFeatureIdGet

> Feature ThingsThingIdFeaturesFeatureIdGet(ctx, thingId, featureId, optional)

Retrieve a specific feature of a specific thing

Returns a specific feature identified by the `featureId` path parameter of the thing identified by the `thingId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**featureId** | **string**| The ID of the feature - has to conform to RFC-3986 (URI) | 
 **optional** | ***ThingsThingIdFeaturesFeatureIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdFeaturesFeatureIdGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Contains a comma-separated list of fields from the selected feature to be included in the returned JSON.  #### Selectable fields  * &#x60;properties&#x60;    Supports selecting arbitrary sub-fields by using a comma-separated list:     * several properties paths can be passed as a comma-separated list of JSON pointers (RFC-6901)        For example:         * &#x60;?fields&#x3D;properties/color&#x60; would select only &#x60;color&#x60; property value (if present)         * &#x60;?fields&#x3D;properties/color,properties/brightness&#x60; would select only &#x60;color&#x60; and &#x60;brightness&#x60; property values (if present)    Supports selecting arbitrary sub-fields of objects by wrapping sub-fields inside parentheses &#x60;( )&#x60;:     * a comma-separated list of sub-fields (a sub-field is a JSON pointer (RFC-6901) separated with &#x60;/&#x60;) to select     * sub-selectors can be used to request only specific sub-fields by placing expressions in parentheses &#x60;( )&#x60; after a selected subfield        For example:        * &#x60;?fields&#x3D;properties(color,brightness)&#x60; would select only &#x60;color&#x60; and &#x60;brightness&#x60; property values (if present)        * &#x60;?fields&#x3D;properties(location/longitude)&#x60; would select the &#x60;longitude&#x60; value inside the &#x60;location&#x60; object  #### Examples  * &#x60;?fields&#x3D;properties(color,brightness)&#x60; | 
 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 

### Return type

[**Feature**](Feature.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsThingIdFeaturesFeatureIdPropertiesDelete

> ThingsThingIdFeaturesFeatureIdPropertiesDelete(ctx, thingId, featureId, optional)

Delete all properties of a feature

Deletes all properties of the feature identified by the `thingId` and `featureId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**featureId** | **string**| The ID of the feature - has to conform to RFC-3986 (URI) | 
 **optional** | ***ThingsThingIdFeaturesFeatureIdPropertiesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdFeaturesFeatureIdPropertiesDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **requestedAcks** | **optional.String**| Contains the \&quot;requested acknowledgements\&quot; for this modifying request as comma separated list. The HTTP call will block until all requested acknowledgements were aggregated or will time out based on the specified &#x60;timeout&#x60; parameter.  The default (if omitted) requested acks is &#x60;requested-acks&#x3D;twin-persisted&#x60; which will block the HTTP call until the change was persited to the twin. | 
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


## ThingsThingIdFeaturesFeatureIdPropertiesGet

> map[string]interface{} ThingsThingIdFeaturesFeatureIdPropertiesGet(ctx, thingId, featureId, optional)

List all properties of a feature

Returns all properties of the feature identified by the `thingId` and `featureId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**featureId** | **string**| The ID of the feature - has to conform to RFC-3986 (URI) | 
 **optional** | ***ThingsThingIdFeaturesFeatureIdPropertiesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdFeaturesFeatureIdPropertiesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| Contains a comma-separated list of fields from the properties to be included in the returned JSON.  #### Selectable fields  Supports selecting arbitrary sub-fields as defined in the properties by using a comma-separated list:   * several properties paths can be passed as a comma-separated list of JSON pointers (RFC-6901)      For example:       * &#x60;?fields&#x3D;temperature&#x60; would select only &#x60;temperature&#x60; property value (if present)       * &#x60;?fields&#x3D;temperature,humidity&#x60; would select only &#x60;temperature&#x60; and &#x60;humidity&#x60; property values (if present)  Supports selecting arbitrary sub-fields of objects by wrapping sub-fields inside parentheses &#x60;( )&#x60;:   * a comma-separated list of sub-fields (a sub-field is a JSON pointer (RFC-6901) separated with &#x60;/&#x60;) to select   * sub-selectors can be used to request only specific sub-fields by placing expressions in parentheses &#x60;( )&#x60; after a selected subfield      For example:      * &#x60;?fields&#x3D;location(longitude,latitude)&#x60; would select the &#x60;longitude&#x60; and &#x60;latitude&#x60; value inside the &#x60;location&#x60; property  #### Examples  * &#x60;?fields&#x3D;temperature,humidity,location(longitude,latitude)&#x60;  * &#x60;?fields&#x3D;configuration,status(powerConsumption/watts)&#x60; | 
 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 

### Return type

**map[string]interface{}**

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsThingIdFeaturesFeatureIdPropertiesPropertyPathDelete

> ThingsThingIdFeaturesFeatureIdPropertiesPropertyPathDelete(ctx, thingId, featureId, propertyPath, optional)

Delete a specific property of a feature

Deletes a specific property of the feature identified by the `thingId` and `featureId` path parameter.  The property (JSON) can be referenced hierarchically, by applying JSON Pointer notation (RFC-6901)  ### Example To delete the value of the brewingTemp in the water-tank of our coffeemaker example the full path is:  `/things/{thingId}/features/water-tank/properties/configuration/brewingTemp`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**featureId** | **string**| The ID of the feature - has to conform to RFC-3986 (URI) | 
**propertyPath** | **string**| The path to the property | 
 **optional** | ***ThingsThingIdFeaturesFeatureIdPropertiesPropertyPathDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdFeaturesFeatureIdPropertiesPropertyPathDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **requestedAcks** | **optional.String**| Contains the \&quot;requested acknowledgements\&quot; for this modifying request as comma separated list. The HTTP call will block until all requested acknowledgements were aggregated or will time out based on the specified &#x60;timeout&#x60; parameter.  The default (if omitted) requested acks is &#x60;requested-acks&#x3D;twin-persisted&#x60; which will block the HTTP call until the change was persited to the twin. | 
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


## ThingsThingIdFeaturesFeatureIdPropertiesPropertyPathGet

> ThingsThingIdFeaturesFeatureIdPropertiesPropertyPathGet(ctx, thingId, featureId, propertyPath, optional)

Retrieve a specific property of a feature

Returns the a specific property path of the feature identified by the `thingId` and `featureId` path parameter.  The property (JSON) can be referenced hierarchically, by applying JSON Pointer notation (RFC-6901)  ### Example To retrieve the value of the `brewingTemp` in the `water-tank` of our coffeemaker example the full path is:  `/things/{thingId}/features/water-tank/properties/configuration/brewingTemp`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**featureId** | **string**| The ID of the feature - has to conform to RFC-3986 (URI) | 
**propertyPath** | **string**| The path to the property | 
 **optional** | ***ThingsThingIdFeaturesFeatureIdPropertiesPropertyPathGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdFeaturesFeatureIdPropertiesPropertyPathGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 

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


## ThingsThingIdFeaturesFeatureIdPropertiesPropertyPathPut

> ThingsThingIdFeaturesFeatureIdPropertiesPropertyPathPut(ctx, thingId, featureId, propertyPath, body, optional)

Create or update a specific property of a feature

Create or update a specific property of a feature identified by the `thingId` and `featureId` path parameter.  The property will be created if it doesn't exist or else updated.  The property (JSON) can be referenced hierarchically, by applying JSON Pointer notation (RFC-6901),  ### Example To set the value of the brewingTemp in the water-tank of our coffeemaker example the full path is:  `/things/{thingId}/features/water-tank/properties/configuration/brewingTemp`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**featureId** | **string**| The ID of the feature - has to conform to RFC-3986 (URI) | 
**propertyPath** | **string**| The path to the property | 
**body** | **map[string]interface{}**| JSON representation of the value to be created/updated. This may be as well &#x60;null&#x60; or an empty object.  Consider that the value has to be a JSON value, examples:    * for a number, the JSON value is the number: &#x60;42&#x60;    * for a string, the JSON value must be quoted: &#x60;\&quot;aString\&quot;&#x60;    * for a boolean, the JSON value is the boolean: &#x60;true&#x60;    * for an object, the JSON value is the object: &#x60;{ \&quot;key\&quot;: \&quot;value\&quot;}&#x60; -} We strongly recommend to use a restricted set of characters for the key (identifier). Currently these identifiers should follow the pattern: [_a-zA-Z][_a-zA-Z0-9\\-]*    * for an list, the JSON value is the list: &#x60;[ 1,2,3 ]&#x60; | 
 **optional** | ***ThingsThingIdFeaturesFeatureIdPropertiesPropertyPathPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdFeaturesFeatureIdPropertiesPropertyPathPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **requestedAcks** | **optional.String**| Contains the \&quot;requested acknowledgements\&quot; for this modifying request as comma separated list. The HTTP call will block until all requested acknowledgements were aggregated or will time out based on the specified &#x60;timeout&#x60; parameter.  The default (if omitted) requested acks is &#x60;requested-acks&#x3D;twin-persisted&#x60; which will block the HTTP call until the change was persited to the twin. | 
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


## ThingsThingIdFeaturesFeatureIdPropertiesPut

> map[string]interface{} ThingsThingIdFeaturesFeatureIdPropertiesPut(ctx, thingId, featureId, body, optional)

Create or update all properties of a feature at once

Create or update the properties of a feature identified by the `thingId` and `featureId` path parameter.  The properties will be overwritten with the JSON content from the request body.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**featureId** | **string**| The ID of the feature - has to conform to RFC-3986 (URI) | 
**body** | **map[string]interface{}**| JSON object of all properties to be updated at once.  Consider that the value has to be a JSON object or &#x60;null&#x60;. We strongly recommend to use a restricted set of characters for the key (identifier).  Currently these identifiers should follow the pattern: [_a-zA-Z][_a-zA-Z0-9\\-]* | 
 **optional** | ***ThingsThingIdFeaturesFeatureIdPropertiesPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdFeaturesFeatureIdPropertiesPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **requestedAcks** | **optional.String**| Contains the \&quot;requested acknowledgements\&quot; for this modifying request as comma separated list. The HTTP call will block until all requested acknowledgements were aggregated or will time out based on the specified &#x60;timeout&#x60; parameter.  The default (if omitted) requested acks is &#x60;requested-acks&#x3D;twin-persisted&#x60; which will block the HTTP call until the change was persited to the twin. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 

### Return type

**map[string]interface{}**

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsThingIdFeaturesFeatureIdPut

> Feature ThingsThingIdFeaturesFeatureIdPut(ctx, thingId, featureId, feature, optional)

Create or modify a specific feature of a specific thing

Create or modify a specific feature identified by the `featureId` path parameter of the thing identified by the `thingId` path parameter.  ### Create feature If the feature ID is new, the feature and all content from the JSON body will be created  ### Update feature If the feature ID is used already in this thing, the feature will be overwrittern with the content from the JSON body.  Example: Set the `featureId` to **coffee-brewer** and all properties in the body part.  ``` {   \"definition\": [\"com.acme:coffeebrewer:0.1.0\"],   \"properties\": {     \"brewed-coffees\": 42   } } ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**featureId** | **string**| The ID of the feature - has to conform to RFC-3986 (URI) | 
**feature** | [**Feature**](Feature.md)| JSON representation of the feature to be created/modified. Consider that the value has to be a JSON object or null.  Examples: * an empty object: {} - would just create the featureID but would delete all content of the feature * a nested object with multiple model definitions and multiple properties as shown in the example value field | 
 **optional** | ***ThingsThingIdFeaturesFeatureIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdFeaturesFeatureIdPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **requestedAcks** | **optional.String**| Contains the \&quot;requested acknowledgements\&quot; for this modifying request as comma separated list. The HTTP call will block until all requested acknowledgements were aggregated or will time out based on the specified &#x60;timeout&#x60; parameter.  The default (if omitted) requested acks is &#x60;requested-acks&#x3D;twin-persisted&#x60; which will block the HTTP call until the change was persited to the twin. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 

### Return type

[**Feature**](Feature.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsThingIdFeaturesGet

> map[string]Feature ThingsThingIdFeaturesGet(ctx, thingId, optional)

List all features of a specific thing

Returns all features of the thing identified by the `thingId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
 **optional** | ***ThingsThingIdFeaturesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdFeaturesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Contains a comma-separated list of fields from one or more features to be included in the returned JSON.  #### Selectable fields  * &#x60;{featureId}&#x60; The ID of the feature to select properties in   * &#x60;properties&#x60;     Supports selecting arbitrary sub-fields by using a comma-separated list:       * several properties paths can be passed as a comma-separated list of JSON pointers (RFC-6901)         For example:           * &#x60;?fields&#x3D;{featureId}/properties/color&#x60; would select only &#x60;color&#x60; property value (if present) of the feature identified with &#x60;{featureId}&#x60;           * &#x60;?fields&#x3D;{featureId}/properties/color,properties/brightness&#x60; would select only &#x60;color&#x60; and &#x60;brightness&#x60; property values (if present) of the feature identified with &#x60;{featureId}&#x60;     Supports selecting arbitrary sub-fields of objects by wrapping sub-fields inside parentheses &#x60;( )&#x60;:       * a comma-separated list of sub-fields (a sub-field is a JSON pointer (RFC-6901) separated with &#x60;/&#x60;) to select       * sub-selectors can be used to request only specific sub-fields by placing expressions in parentheses &#x60;( )&#x60; after a selected subfield         For example:          * &#x60;?fields&#x3D;{featureId}/properties(color,brightness)&#x60; would select only &#x60;color&#x60; and &#x60;brightness&#x60; property values (if present) of the feature identified with &#x60;{featureId}&#x60;          * &#x60;?fields&#x3D;{featureId}/properties(location/longitude)&#x60; would select the &#x60;longitude&#x60; value inside the &#x60;location&#x60; object of the feature identified with &#x60;{featureId}&#x60;   #### Examples * &#x60;?fields&#x3D;EnvironmentScanner/properties(temperature,humidity)&#x60; * &#x60;?fields&#x3D;EnvironmentScanner/properties(temperature,humidity),Vehicle/properties/configuration&#x60; | 
 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 

### Return type

[**map[string]Feature**](Feature.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsThingIdFeaturesPut

> map[string]Feature ThingsThingIdFeaturesPut(ctx, thingId, requestBody, optional)

Create or modify all features of a specific thing at once

Create or modify all features of a thing identified by the `thingId` path parameter. ### Create all features at once In case at the initial creation of your thing you have not specified any features, these can be created here.  ### Update all features at once To update all features at once prepare the JSON body accordingly.  Note: In contrast to the \"PUT thing\" request, a partial update is not supported here, but the content will be **overwritten**. If you need to update single features or their paths, please use the sub-resources instead.  Example:  ``` {      \"coffee-brewer\": {        \"definition\": [\"com.acme:coffeebrewer:0.1.0\"],        \"properties\": {          \"brewed-coffees\": 0        }      },      \"water-tank\": {        \"properties\": {          \"configuration\": {            \"smartMode\": true,            \"brewingTemp\": 87,            \"tempToHold\": 44,            \"timeoutSeconds\": 6000          },          \"status\": {            \"waterAmount\": 731,            \"temperature\": 44          }        }      } } ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**requestBody** | [**map[string]Feature**](Feature.md)| JSON object of all features to be modified at once. Consider that the value has to be a JSON object or null.  Examples: * an empty object: {} - would just delete all old features * an empty feature: { \&quot;featureId\&quot;: {} } - We strongly recommend to use a restricted set of characters   for the &#x60;featureId&#x60;, as it might be needed for the (URL) path later.    Currently these identifiers should follow the pattern: [_a-zA-Z][_a-zA-Z0-9-]*  * a nested object with multiple features as shown in the example value field | 
 **optional** | ***ThingsThingIdFeaturesPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdFeaturesPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **requestedAcks** | **optional.String**| Contains the \&quot;requested acknowledgements\&quot; for this modifying request as comma separated list. The HTTP call will block until all requested acknowledgements were aggregated or will time out based on the specified &#x60;timeout&#x60; parameter.  The default (if omitted) requested acks is &#x60;requested-acks&#x3D;twin-persisted&#x60; which will block the HTTP call until the change was persited to the twin. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 

### Return type

[**map[string]Feature**](Feature.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

