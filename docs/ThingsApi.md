# \ThingsApi

All URIs are relative to *https://things.eu-1.bosch-iot-suite.com/api/2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThingsGet**](ThingsApi.md#ThingsGet) | **Get** /things | Retrieve multiple things with specified IDs
[**ThingsPost**](ThingsApi.md#ThingsPost) | **Post** /things | Create a new thing
[**ThingsThingIdAttributesAttributePathDelete**](ThingsApi.md#ThingsThingIdAttributesAttributePathDelete) | **Delete** /things/{thingId}/attributes/{attributePath} | Delete a specific attribute of a specific thing
[**ThingsThingIdAttributesAttributePathGet**](ThingsApi.md#ThingsThingIdAttributesAttributePathGet) | **Get** /things/{thingId}/attributes/{attributePath} | Retrieve a specific attribute of a specific thing
[**ThingsThingIdAttributesAttributePathPut**](ThingsApi.md#ThingsThingIdAttributesAttributePathPut) | **Put** /things/{thingId}/attributes/{attributePath} | Create or update a specific attribute of a specific thing
[**ThingsThingIdAttributesDelete**](ThingsApi.md#ThingsThingIdAttributesDelete) | **Delete** /things/{thingId}/attributes | Delete all attributes of a specific thing at once
[**ThingsThingIdAttributesGet**](ThingsApi.md#ThingsThingIdAttributesGet) | **Get** /things/{thingId}/attributes | List all attributes of a specific thing
[**ThingsThingIdAttributesPut**](ThingsApi.md#ThingsThingIdAttributesPut) | **Put** /things/{thingId}/attributes | Create or update all attributes of a specific thing at once
[**ThingsThingIdDefinitionDelete**](ThingsApi.md#ThingsThingIdDefinitionDelete) | **Delete** /things/{thingId}/definition | Delete the definition of a specific thing
[**ThingsThingIdDefinitionGet**](ThingsApi.md#ThingsThingIdDefinitionGet) | **Get** /things/{thingId}/definition | Retrieve the definition of a specific thing
[**ThingsThingIdDefinitionPut**](ThingsApi.md#ThingsThingIdDefinitionPut) | **Put** /things/{thingId}/definition | Create or update the definition of a specific thing
[**ThingsThingIdDelete**](ThingsApi.md#ThingsThingIdDelete) | **Delete** /things/{thingId} | Delete a specific thing
[**ThingsThingIdGet**](ThingsApi.md#ThingsThingIdGet) | **Get** /things/{thingId} | Retrieve a specific thing
[**ThingsThingIdPolicyIdGet**](ThingsApi.md#ThingsThingIdPolicyIdGet) | **Get** /things/{thingId}/policyId | Retrieve the policy ID of a thing
[**ThingsThingIdPolicyIdPut**](ThingsApi.md#ThingsThingIdPolicyIdPut) | **Put** /things/{thingId}/policyId | Create or update the policy ID of a thing
[**ThingsThingIdPut**](ThingsApi.md#ThingsThingIdPut) | **Put** /things/{thingId} | Create or update a thing with a specified ID



## ThingsGet

> []Thing ThingsGet(ctx, ids, optional)

Retrieve multiple things with specified IDs

Returns all things passed in by the required parameter `ids`, which you (the authorized subject) are allowed to read.  Optionally, if you want to retrieve only some of the thing's fields, you can use the specific field selectors (see parameter `fields`) .  Tip: If you don't know the thing IDs, start with the search resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ids** | **string**| Contains a comma-separated list of &#x60;thingId&#x60;s to retrieve in one single request. | 
 **optional** | ***ThingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Contains a comma-separated list of fields to be included in the returned JSON. attributes can be selected in the same manner.  #### Selectable fields  * &#x60;thingId&#x60; * &#x60;policyId&#x60; * &#x60;definition&#x60; * &#x60;attributes&#x60;     Supports selecting arbitrary sub-fields by using a comma-separated list:     * several attribute paths can be passed as a comma-separated list of JSON pointers (RFC-6901)        For example:         * &#x60;?fields&#x3D;attributes/model&#x60; would select only &#x60;model&#x60; attribute value (if present)         * &#x60;?fields&#x3D;attributes/model,attributes/location&#x60; would select only &#x60;model&#x60; and            &#x60;location&#x60; attribute values (if present)    Supports selecting arbitrary sub-fields of objects by wrapping sub-fields inside parentheses &#x60;( )&#x60;:     * a comma-separated list of sub-fields (a sub-field is a JSON pointer (RFC-6901)       separated with &#x60;/&#x60;) to select      * sub-selectors can be used to request only specific sub-fields by placing expressions       in parentheses &#x60;( )&#x60; after a selected subfield        For example:        * &#x60;?fields&#x3D;attributes(model,location)&#x60; would select only &#x60;model&#x60;           and &#x60;location&#x60; attribute values (if present)        * &#x60;?fields&#x3D;attributes(coffeemaker/serialno)&#x60; would select the &#x60;serialno&#x60; value           inside the &#x60;coffeemaker&#x60; object        * &#x60;?fields&#x3D;attributes/address/postal(city,street)&#x60; would select the &#x60;city&#x60; and           &#x60;street&#x60; values inside the &#x60;postal&#x60; object inside the &#x60;address&#x60; object  * &#x60;features&#x60;    Supports selecting arbitrary fields in features similar to &#x60;attributes&#x60; (see also features documentation for more details)  * &#x60;_namespace&#x60;    Specifically selects the namespace also contained in the &#x60;thingId&#x60;  * &#x60;_revision&#x60;    Specifically selects the revision of the thing. The revision is a counter, which is incremented on each modification of a thing.  * &#x60;_modified&#x60;    Specifically selects the modified timestamp of the thing in ISO-8601 UTC format. The timestamp is set on each modification of a thing.  * &#x60;_policy&#x60;    Specifically selects the content of the policy associated to the thing. (By default, only the policyId is returned.)  #### Examples  * &#x60;?fields&#x3D;thingId,attributes,features&#x60; * &#x60;?fields&#x3D;attributes(model,manufacturer),features&#x60; | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 

### Return type

[**[]Thing**](Thing.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsPost

> Thing ThingsPost(ctx, optional)

Create a new thing

Creates a thing with a default `thingId` and a default `policyId`.  The thing will be empty, i.e. no features, definition, attributes etc. by default.  The default `thingId` consists of your default namespace and a UUID. Make sure a default namespace is registered for your solution.  The default `policyId` is identical with the default `thingId`, and allows the currently authorized subject all permissions.  In case you need to create a thing with a specific ID, use a *PUT* request instead, as any `thingId` specified in the request body will be ignored.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ThingsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestedAcks** | **optional.String**| Contains the \&quot;requested acknowledgements\&quot; for this modifying request as comma separated list. The HTTP call will block until all requested acknowledgements were aggregated or will time out based on the specified &#x60;timeout&#x60; parameter.  The default (if omitted) requested acks is &#x60;requested-acks&#x3D;twin-persisted&#x60; which will block the HTTP call until the change was persited to the twin. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 
 **newThing** | [**optional.Interface of NewThing**](NewThing.md)| JSON representation of the thing to be created. Use { } to create an empty thing with a default policy. | 

### Return type

[**Thing**](Thing.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsThingIdAttributesAttributePathDelete

> ThingsThingIdAttributesAttributePathDelete(ctx, thingId, attributePath, optional)

Delete a specific attribute of a specific thing

Deletes a specific attribute of the thing identified by the `thingId` path parameter.  The attribute (JSON) can be referenced hierarchically, by applying JSON Pointer notation (RFC-6901).  Example: In order to delete the `name` field of an `manufacturer` attribute, the full path would be         `/things/{thingId}/attributes/manufacturer/name`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**attributePath** | **string**| The path to the attribute, e.g. **manufacturer/name** | 
 **optional** | ***ThingsThingIdAttributesAttributePathDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdAttributesAttributePathDeleteOpts struct


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


## ThingsThingIdAttributesAttributePathGet

> ThingsThingIdAttributesAttributePathGet(ctx, thingId, attributePath, optional)

Retrieve a specific attribute of a specific thing

Returns a specific attribute of the thing identified by the `thingId` path parameter.  The attribute (JSON) can be referenced hierarchically, by applying JSON Pointer notation (RFC-6901).  Example: In order to retrieve the `name` field of an `manufacturer` attribute, the full path would be `/things/{thingId}/attributes/manufacturer/name`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**attributePath** | **string**| The path to the attribute, e.g. **manufacturer/name** | 
 **optional** | ***ThingsThingIdAttributesAttributePathGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdAttributesAttributePathGetOpts struct


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


## ThingsThingIdAttributesAttributePathPut

> ThingsThingIdAttributesAttributePathPut(ctx, thingId, attributePath, body, optional)

Create or update a specific attribute of a specific thing

Create or update a specific attribute of the thing identified by the `thingId` path parameter.  * If you specify a new attribute path, this will be created * If you specify an existing attribute path, this will be updated  The attribute (JSON) can be referenced hierarchically, by applying JSON Pointer notation (RFC-6901).  Example: In order to put the `name` field of an `manufacturer` attribute, the full path would be `/things/{thingId}/attributes/manufacturer/name`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**attributePath** | **string**| The path to the attribute, e.g. **manufacturer/name** | 
**body** | **map[string]interface{}**| JSON representation of the value to be created/updated. This may be as well &#x60;null&#x60; or an empty object.  Consider that the value has to be a JSON value, examples:    * for a number, the JSON value is the number: &#x60;42&#x60;    * for a string, the JSON value must be quoted: &#x60;\&quot;aString\&quot;&#x60;    * for a boolean, the JSON value is the boolean: &#x60;true&#x60;    * for an object, the JSON value is the object: &#x60;{ \&quot;key\&quot;: \&quot;value\&quot;}&#x60; -} We strongly recommend to use a restricted set of characters for the key (identifier). Currently these identifiers should follow the pattern: [_a-zA-Z][_a-zA-Z0-9\\-]*    * for an list, the JSON value is the list: &#x60;[ 1,2,3 ]&#x60; | 
 **optional** | ***ThingsThingIdAttributesAttributePathPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdAttributesAttributePathPutOpts struct


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


## ThingsThingIdAttributesDelete

> ThingsThingIdAttributesDelete(ctx, thingId, optional)

Delete all attributes of a specific thing at once

Deletes all attributes of the thing identified by the `thingId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
 **optional** | ***ThingsThingIdAttributesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdAttributesDeleteOpts struct


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


## ThingsThingIdAttributesGet

> map[string]interface{} ThingsThingIdAttributesGet(ctx, thingId, optional)

List all attributes of a specific thing

Returns all attributes of the thing identified by the `thingId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
 **optional** | ***ThingsThingIdAttributesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdAttributesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Contains a comma-separated list of fields from the attributes to be included in the returned JSON.  #### Selectable fields  Supports selecting arbitrary sub-fields as defined in the attributes by using a comma-separated list:   * several properties paths can be passed as a comma-separated list of JSON pointers (RFC-6901)      For example:       * &#x60;?fields&#x3D;model&#x60; would select only &#x60;model&#x60; attribute value (if present)       * &#x60;?fields&#x3D;model,make&#x60; would select &#x60;model&#x60; and &#x60;make&#x60; attribute values (if present)  Supports selecting arbitrary sub-fields of objects by wrapping sub-fields inside parentheses &#x60;( )&#x60;:   * a comma-separated list of sub-fields (a sub-field is a JSON pointer (RFC-6901) separated with &#x60;/&#x60;) to select   * sub-selectors can be used to request only specific sub-fields by placing expressions in parentheses &#x60;( )&#x60; after a selected subfield      For example:      * &#x60;?fields&#x3D;location(longitude,latitude)&#x60; would select the &#x60;longitude&#x60; and &#x60;latitude&#x60; value inside the &#x60;location&#x60; attribute  #### Examples  * &#x60;?fields&#x3D;model,make,location(longitude,latitude)&#x60;  * &#x60;?fields&#x3D;listOfAddresses/postal(city,street))&#x60; | 
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


## ThingsThingIdAttributesPut

> map[string]interface{} ThingsThingIdAttributesPut(ctx, thingId, body, optional)

Create or update all attributes of a specific thing at once

Create or update the attributes of a thing identified by the `thingId` path parameter. The attributes will be overwritten - all at once - with the content (JSON) set in the request body.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**body** | **map[string]interface{}**| JSON object of all attributes to be modified at once. Consider that the value has to be a JSON object or &#x60;null&#x60;.  Examples:  * an empty object: &#x60;{}&#x60; - would just delete all old attributes  * a simple object: &#x60;{ \&quot;key\&quot;: \&quot;value\&quot;}&#x60; - We strongly recommend to use a restricted set of characters for the key (identifier), as the key might be needed for the (URL) path later.&lt;br&gt; Currently these identifiers should follow the pattern: [_a-zA-Z][_a-zA-Z0-9\\-]*  * a nested object as shown in the example value | 
 **optional** | ***ThingsThingIdAttributesPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdAttributesPutOpts struct


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


## ThingsThingIdDefinitionDelete

> ThingsThingIdDefinitionDelete(ctx, thingId, optional)

Delete the definition of a specific thing

Deletes the definition of the thing identified by the `thingId`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
 **optional** | ***ThingsThingIdDefinitionDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdDefinitionDeleteOpts struct


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


## ThingsThingIdDefinitionGet

> string ThingsThingIdDefinitionGet(ctx, thingId, optional)

Retrieve the definition of a specific thing

Returns the definition of the thing identified by the `thingId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
 **optional** | ***ThingsThingIdDefinitionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdDefinitionGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 

### Return type

**string**

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsThingIdDefinitionPut

> string ThingsThingIdDefinitionPut(ctx, thingId, optional)

Create or update the definition of a specific thing

* If the thing does not have a definition yet, this request will create it. * If the thing already has a definition you can assign it to a new one by setting the new definition in the request body.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
 **optional** | ***ThingsThingIdDefinitionPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdDefinitionPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **requestedAcks** | **optional.String**| Contains the \&quot;requested acknowledgements\&quot; for this modifying request as comma separated list. The HTTP call will block until all requested acknowledgements were aggregated or will time out based on the specified &#x60;timeout&#x60; parameter.  The default (if omitted) requested acks is &#x60;requested-acks&#x3D;twin-persisted&#x60; which will block the HTTP call until the change was persited to the twin. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 
 **body** | **optional.String**| JSON string of the definition to be modified. Consider that the value has to be a JSON string or &#x60;null&#x60;, examples:    * an string: &#x60;{ \&quot;\&quot;value\&quot;}&#x60; -}. Currently the definition should follow the pattern: [_a-zA-Z0-9\\-]:[_a-zA-Z0-9\\-]:[_a-zA-Z0-9\\-]   * an empty string: &#x60;\&quot;\&quot;&#x60; | 

### Return type

**string**

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsThingIdDelete

> ThingsThingIdDelete(ctx, thingId, optional)

Delete a specific thing

Deletes the thing identified by the `thingId` path parameter.  This will not delete the policy, which is used for controlling access to this thing.  You can delete the policy afterwards via DELETE `/policies/{policyId}` if you don't need it for other things.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
 **optional** | ***ThingsThingIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous GET response, e.g. &#x60;If-Match: \&quot;rev:4711\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
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


## ThingsThingIdGet

> Thing ThingsThingIdGet(ctx, thingId, optional)

Retrieve a specific thing

Returns the thing identified by the `thingId` path parameter. The response includes details about the thing, including the `policyId`, attributes, definition and features.  Optionally, you can use the field selectors (see parameter `fields`) to only get specific fields, which you are interested in.  Example: Use the field selector `_policy` to retrieve the content of the policy.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
 **optional** | ***ThingsThingIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Contains a comma-separated list of fields to be included in the returned JSON. attributes can be selected in the same manner.  #### Selectable fields  * &#x60;thingId&#x60; * &#x60;policyId&#x60; * &#x60;definition&#x60; * &#x60;attributes&#x60;     Supports selecting arbitrary sub-fields by using a comma-separated list:     * several attribute paths can be passed as a comma-separated list of JSON pointers (RFC-6901)        For example:         * &#x60;?fields&#x3D;attributes/model&#x60; would select only &#x60;model&#x60; attribute value (if present)         * &#x60;?fields&#x3D;attributes/model,attributes/location&#x60; would select only &#x60;model&#x60; and            &#x60;location&#x60; attribute values (if present)    Supports selecting arbitrary sub-fields of objects by wrapping sub-fields inside parentheses &#x60;( )&#x60;:     * a comma-separated list of sub-fields (a sub-field is a JSON pointer (RFC-6901)       separated with &#x60;/&#x60;) to select      * sub-selectors can be used to request only specific sub-fields by placing expressions       in parentheses &#x60;( )&#x60; after a selected subfield        For example:        * &#x60;?fields&#x3D;attributes(model,location)&#x60; would select only &#x60;model&#x60;           and &#x60;location&#x60; attribute values (if present)        * &#x60;?fields&#x3D;attributes(coffeemaker/serialno)&#x60; would select the &#x60;serialno&#x60; value           inside the &#x60;coffeemaker&#x60; object        * &#x60;?fields&#x3D;attributes/address/postal(city,street)&#x60; would select the &#x60;city&#x60; and           &#x60;street&#x60; values inside the &#x60;postal&#x60; object inside the &#x60;address&#x60; object  * &#x60;features&#x60;    Supports selecting arbitrary fields in features similar to &#x60;attributes&#x60; (see also features documentation for more details)  * &#x60;_namespace&#x60;    Specifically selects the namespace also contained in the &#x60;thingId&#x60;  * &#x60;_revision&#x60;    Specifically selects the revision of the thing. The revision is a counter, which is incremented on each modification of a thing.  * &#x60;_modified&#x60;    Specifically selects the modified timestamp of the thing in ISO-8601 UTC format. The timestamp is set on each modification of a thing.  * &#x60;_policy&#x60;    Specifically selects the content of the policy associated to the thing. (By default, only the policyId is returned.)  #### Examples  * &#x60;?fields&#x3D;thingId,attributes,features&#x60; * &#x60;?fields&#x3D;attributes(model,manufacturer),features&#x60; | 
 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous GET response, e.g. &#x60;If-Match: \&quot;rev:4711\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 

### Return type

[**Thing**](Thing.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsThingIdPolicyIdGet

> string ThingsThingIdPolicyIdGet(ctx, thingId, optional)

Retrieve the policy ID of a thing

Returns the policy ID of the thing identified by the `thingId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
 **optional** | ***ThingsThingIdPolicyIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdPolicyIdGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 

### Return type

**string**

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsThingIdPolicyIdPut

> string ThingsThingIdPolicyIdPut(ctx, thingId, body, optional)

Create or update the policy ID of a thing

Create or update the policy ID of the thing identified by the `thingId` path parameter.  ### Create If the thing does not have a `policyId` yet, it is considered to have been created via API version 1, where the access control list `acl` mechanism is used. In that case, this request will create the `policyId`.  Note: You will need to create the policy content separately.  ### Update If the thing already has a `policyId` you can assign it to an existing policy by setting the new `policyId` in the request body.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**body** | **string**| The policy is used for controlling access to this thing. It is managed by resource &#x60;/policies/{policyId}&#x60;.  The ID of a policy needs to conform to the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
 **optional** | ***ThingsThingIdPolicyIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdPolicyIdPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous HTTP response, e.g. &#x60;If-Match: \&quot;hash:a75ece4e\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **requestedAcks** | **optional.String**| Contains the \&quot;requested acknowledgements\&quot; for this modifying request as comma separated list. The HTTP call will block until all requested acknowledgements were aggregated or will time out based on the specified &#x60;timeout&#x60; parameter.  The default (if omitted) requested acks is &#x60;requested-acks&#x3D;twin-persisted&#x60; which will block the HTTP call until the change was persited to the twin. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 

### Return type

**string**

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsThingIdPut

> Thing ThingsThingIdPut(ctx, thingId, optional)

Create or update a thing with a specified ID

Create or update the thing specified by the `thingId` path parameter and the optional JSON body.  * If you set a new `thingId` in the path, a thing will be created. * If you set an existing `thingId` in the path, the thing will be updated.   ### Create a new thing At the initial creation of a thing, only a valid `thingId` is required. However, you can create a full-fledged thing all at once.  Example: To create a coffee maker thing, set the `thingId` in the path, e.g. to \"com.acme.coffeemaker:BE-42\" and the body part, like in the following snippet.  ```  {    \"definition\": \"com.acme:coffeebrewer:0.1.0\",    \"attributes\": {      \"manufacturer\": \"ACME demo corp.\",      \"location\": \"Berlin, main floor\",      \"serialno\": \"42\",      \"model\": \"Speaking coffee machine\"    },    \"features\": {      \"coffee-brewer\": {        \"definition\": [ \"com.acme:coffeebrewer:0.1.0\" ],        \"properties\": {          \"brewed-coffees\": 0        }      },      \"water-tank\": {        \"properties\": {          \"configuration\": {            \"smartMode\": true,            \"brewingTemp\": 87,            \"tempToHold\": 44,            \"timeoutSeconds\": 6000          },          \"status\": {            \"waterAmount\": 731,            \"temperature\": 44          }        }      }    }   }  ``` As the example does not set a policy in the request body, but the thing concept requires one, the service will create a default policy. The default policy, has the exactly same id as the thing, and grants ALL permissions to the authorized subject.  In case you need to associate the new thing to an already existing policy you can additionally set a policy e.g. \"policyId\": \"com.acme.coffeemaker:policy-1\" as the first element in the body part. Keep in mind, that you can also change the assignment to another policy anytime, with a request on the sub-resource \"PUT /things/{thingId}/policyId\"  ### Update an existing thing  For updating an existing thing, the authorized subject needs **WRITE** permission on the thing's root resource.  The ID of a thing cannot be changed after creation. Any `thingId` specified in the request body is therefore ignored.  ### Partially update an existing thing  When updating an existing thing, which already contains `attributes`, `definition` `policyId` or `features`, the existing fields must not explicitly be provided again. For this \"PUT thing\" request - and only for this top-level update on the thing - the top-level field to update is **merged** with the existing top-level fields of the thing.  ### Example for a partial update  Given, a thing already exists with this content:  ``` {   \"thingId\": \"namespace:thing-name\",   \"policyId\": \"namespace:policy-name\",   \"definition\": \"namespace:model:version\",   \"attributes\": {     \"foo\": 1   },   \"features\": {...} } ```  The thing's `attributes` may be modified without having to pass the `policyId` or the `features` in again. For updating the `attributes`, following request body would be sufficient :  ``` {   \"attributes\": {     \"foo\": 2,     \"bar\": false   } } ```  The `policyId` and `features` of the thing will not be overwritten. The thing will be merged into:  ``` {   \"thingId\": \"namespace:thing-name\",   \"policyId\": \"namespace:policy-name\",   \"definition\": \"namespace:model:version\",   \"attributes\": {     \"foo\": 2,     \"bar\": false   },   \"features\": {...} } ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
 **optional** | ***ThingsThingIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **optional.String**| The &#x60;If-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). Common usages are:   * optimistic locking by specifying the &#x60;ETag&#x60; from a previous GET response, e.g. &#x60;If-Match: \&quot;rev:4711\&quot;&#x60;   * retrieving or modifying a resource only if it already exists, e.g. &#x60;If-Match: *&#x60; | 
 **ifNoneMatch** | **optional.String**| The &#x60;If-None-Match&#x60; header, which has to conform to RFC-7232 (Conditional Requests). A common usage scenario is to modify a resource only if it does not yet exist, thus to create it, by specifying &#x60;If-None-Match: *&#x60;. | 
 **requestedAcks** | **optional.String**| Contains the \&quot;requested acknowledgements\&quot; for this modifying request as comma separated list. The HTTP call will block until all requested acknowledgements were aggregated or will time out based on the specified &#x60;timeout&#x60; parameter.  The default (if omitted) requested acks is &#x60;requested-acks&#x3D;twin-persisted&#x60; which will block the HTTP call until the change was persited to the twin. | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **responseRequired** | **optional.Bool**| Defines whether a response is required to the API call or not - if set to &#x60;false&#x60; the response will directly sent back with a status code of &#x60;202&#x60; (Accepted).  The default (if ommited) response is &#x60;true&#x60;. | 
 **newThing** | [**optional.Interface of NewThing**](NewThing.md)| JSON representation of the thing to be modified. | 

### Return type

[**Thing**](Thing.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

