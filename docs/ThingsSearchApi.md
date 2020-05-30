# \ThingsSearchApi

All URIs are relative to *https://things.eu-1.bosch-iot-suite.com/api/2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchThingsCountGet**](ThingsSearchApi.md#SearchThingsCountGet) | **Get** /search/things/count | Count things
[**SearchThingsGet**](ThingsSearchApi.md#SearchThingsGet) | **Get** /search/things | Search for things



## SearchThingsCountGet

> int32 SearchThingsCountGet(ctx, optional)

Count things

This resource can be used to count things.  The query parameter `filter` is not mandatory. If it is not set there is returned the total amount of things which the logged in user is allowed to read.  To search for nested properties, we use JSON Pointer notation (RFC-6901). See the following example how to search for the sub property `location` of the parent property `attributes` with a forward slash as separator:  ```eq(attributes/location,\"kitchen\")```  The search result is limited to the things within the namespace (or namespaces) the solution is configured for. You can explicitly search in other namespaces by including them in the query via the `namespaces` parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchThingsCountGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchThingsCountGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**|  #### Filter predicates:  * &#x60;&#x60;&#x60;eq({property},{value})&#x60;&#x60;&#x60;  (i.e. equal to the given value)  * &#x60;&#x60;&#x60;ne({property},{value})&#x60;&#x60;&#x60;  (i.e. not equal to the given value)  * &#x60;&#x60;&#x60;gt({property},{value})&#x60;&#x60;&#x60;  (i.e. greater than the given value)  * &#x60;&#x60;&#x60;ge({property},{value})&#x60;&#x60;&#x60;  (i.e. equal to the given value or greater than it)  * &#x60;&#x60;&#x60;lt({property},{value})&#x60;&#x60;&#x60;  (i.e. lower than the given value or equal to it)  * &#x60;&#x60;&#x60;le({property},{value})&#x60;&#x60;&#x60;  (i.e. lower than the given value)  * &#x60;&#x60;&#x60;in({property},{value},{value},...)&#x60;&#x60;&#x60;  (i.e. contains at least one of the values listed)  * &#x60;&#x60;&#x60;like({property},{value})&#x60;&#x60;&#x60;  (i.e. contains values similar to the expressions listed)  * &#x60;&#x60;&#x60;exists({property})&#x60;&#x60;&#x60;  (i.e. all things in which the given path exists)   Note: When using filter operations, only things with the specified properties are returned. For example, the filter &#x60;ne(attributes/owner, \&quot;SID123\&quot;)&#x60; will only return things that do have the &#x60;owner&#x60; attribute.   #### Logical operations:   * &#x60;&#x60;&#x60;and({query},{query},...)&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;or({query},{query},...)&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;not({query})&#x60;&#x60;&#x60;   #### Examples:  * &#x60;&#x60;&#x60;eq(attributes/location,\&quot;kitchen\&quot;)&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;ge(thingId,\&quot;myThing1\&quot;)&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;exists(features/featureId)&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;and(eq(attributes/location,\&quot;kitchen\&quot;),eq(attributes/color,\&quot;red\&quot;))&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;or(eq(attributes/location,\&quot;kitchen\&quot;),eq(attributes/location,\&quot;living-room\&quot;))&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;like(attributes/key1,\&quot;known-chars-at-start*\&quot;)&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;like(attributes/key1,\&quot;*known-chars-at-end\&quot;)&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;like(attributes/key1,\&quot;*known-chars-in-between*\&quot;)&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;like(attributes/key1,\&quot;just-som?-char?-unkn?wn\&quot;)&#x60;&#x60;&#x60;  The &#x60;like&#x60; filters with the wildcard &#x60;*&#x60; at the beginning can slow down your search request. | 
 **namespaces** | **optional.String**| A comma-separated list of namespaces. This list is used to limit the query to things in the given namespaces only. If this parameter is omitted, all registered namespaces of your solution will be queried. The solution is determined by the API token sent with the request.   #### Examples:  * &#x60;?namespaces&#x3D;com.example.namespace&#x60;  * &#x60;?namespaces&#x3D;com.example.namespace1,com.example.namespace2&#x60; | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 

### Return type

**int32**

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchThingsGet

> SearchResultThings SearchThingsGet(ctx, optional)

Search for things

This resource can be used to search for things.  * The query parameter `filter` is not mandatory. If it is not set, the   result contains all things which the logged in user is allowed to read.  * The search is case sensitive. In case you don't know how exactly the   spelling of the namespace, name, attribute, feature etc. is, use the *like*   notation for filtering  * The search result is limited to the things within the namespace (or namespaces)   the solution is configured for. You can explicitly search in specific namespaces   by including them in the query via the *namespaces* parameter.  * The resource supports sorting and paging. If paging is not explicitly   specified by means of the `size` option, a default count of `25`   documents is returned.  * The internal search index is \"eventually consistent\".  Consistency with the latest   thing updates should recover within milliseconds.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchThingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchThingsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**|  #### Filter predicates:  * &#x60;&#x60;&#x60;eq({property},{value})&#x60;&#x60;&#x60;  (i.e. equal to the given value)  * &#x60;&#x60;&#x60;ne({property},{value})&#x60;&#x60;&#x60;  (i.e. not equal to the given value)  * &#x60;&#x60;&#x60;gt({property},{value})&#x60;&#x60;&#x60;  (i.e. greater than the given value)  * &#x60;&#x60;&#x60;ge({property},{value})&#x60;&#x60;&#x60;  (i.e. equal to the given value or greater than it)  * &#x60;&#x60;&#x60;lt({property},{value})&#x60;&#x60;&#x60;  (i.e. lower than the given value or equal to it)  * &#x60;&#x60;&#x60;le({property},{value})&#x60;&#x60;&#x60;  (i.e. lower than the given value)  * &#x60;&#x60;&#x60;in({property},{value},{value},...)&#x60;&#x60;&#x60;  (i.e. contains at least one of the values listed)  * &#x60;&#x60;&#x60;like({property},{value})&#x60;&#x60;&#x60;  (i.e. contains values similar to the expressions listed)  * &#x60;&#x60;&#x60;exists({property})&#x60;&#x60;&#x60;  (i.e. all things in which the given path exists)   Note: When using filter operations, only things with the specified properties are returned. For example, the filter &#x60;ne(attributes/owner, \&quot;SID123\&quot;)&#x60; will only return things that do have the &#x60;owner&#x60; attribute.   #### Logical operations:   * &#x60;&#x60;&#x60;and({query},{query},...)&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;or({query},{query},...)&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;not({query})&#x60;&#x60;&#x60;   #### Examples:  * &#x60;&#x60;&#x60;eq(attributes/location,\&quot;kitchen\&quot;)&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;ge(thingId,\&quot;myThing1\&quot;)&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;exists(features/featureId)&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;and(eq(attributes/location,\&quot;kitchen\&quot;),eq(attributes/color,\&quot;red\&quot;))&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;or(eq(attributes/location,\&quot;kitchen\&quot;),eq(attributes/location,\&quot;living-room\&quot;))&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;like(attributes/key1,\&quot;known-chars-at-start*\&quot;)&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;like(attributes/key1,\&quot;*known-chars-at-end\&quot;)&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;like(attributes/key1,\&quot;*known-chars-in-between*\&quot;)&#x60;&#x60;&#x60;  * &#x60;&#x60;&#x60;like(attributes/key1,\&quot;just-som?-char?-unkn?wn\&quot;)&#x60;&#x60;&#x60;  The &#x60;like&#x60; filters with the wildcard &#x60;*&#x60; at the beginning can slow down your search request. | 
 **namespaces** | **optional.String**| A comma-separated list of namespaces. This list is used to limit the query to things in the given namespaces only. If this parameter is omitted, all registered namespaces of your solution will be queried. The solution is determined by the API token sent with the request.   #### Examples:  * &#x60;?namespaces&#x3D;com.example.namespace&#x60;  * &#x60;?namespaces&#x3D;com.example.namespace1,com.example.namespace2&#x60; | 
 **fields** | **optional.String**| Contains a comma-separated list of fields to be included in the returned JSON. attributes can be selected in the same manner.  #### Selectable fields  * &#x60;thingId&#x60; * &#x60;policyId&#x60; * &#x60;definition&#x60; * &#x60;attributes&#x60;     Supports selecting arbitrary sub-fields by using a comma-separated list:     * several attribute paths can be passed as a comma-separated list of JSON pointers (RFC-6901)        For example:         * &#x60;?fields&#x3D;attributes/model&#x60; would select only &#x60;model&#x60; attribute value (if present)         * &#x60;?fields&#x3D;attributes/model,attributes/location&#x60; would select only &#x60;model&#x60; and            &#x60;location&#x60; attribute values (if present)    Supports selecting arbitrary sub-fields of objects by wrapping sub-fields inside parentheses &#x60;( )&#x60;:     * a comma-separated list of sub-fields (a sub-field is a JSON pointer (RFC-6901)       separated with &#x60;/&#x60;) to select      * sub-selectors can be used to request only specific sub-fields by placing expressions       in parentheses &#x60;( )&#x60; after a selected subfield        For example:        * &#x60;?fields&#x3D;attributes(model,location)&#x60; would select only &#x60;model&#x60;           and &#x60;location&#x60; attribute values (if present)        * &#x60;?fields&#x3D;attributes(coffeemaker/serialno)&#x60; would select the &#x60;serialno&#x60; value           inside the &#x60;coffeemaker&#x60; object        * &#x60;?fields&#x3D;attributes/address/postal(city,street)&#x60; would select the &#x60;city&#x60; and           &#x60;street&#x60; values inside the &#x60;postal&#x60; object inside the &#x60;address&#x60; object  * &#x60;features&#x60;    Supports selecting arbitrary fields in features similar to &#x60;attributes&#x60; (see also features documentation for more details)  * &#x60;_namespace&#x60;    Specifically selects the namespace also contained in the &#x60;thingId&#x60;  * &#x60;_revision&#x60;    Specifically selects the revision of the thing. The revision is a counter, which is incremented on each modification of a thing.  * &#x60;_modified&#x60;    Specifically selects the modified timestamp of the thing in ISO-8601 UTC format. The timestamp is set on each modification of a thing.  * &#x60;_policy&#x60;    Specifically selects the content of the policy associated to the thing. (By default, only the policyId is returned.)  #### Examples  * &#x60;?fields&#x3D;thingId,attributes,features&#x60; * &#x60;?fields&#x3D;attributes(model,manufacturer),features&#x60; | 
 **timeout** | **optional.String**| Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the &#x60;requested-acks&#x60; param. Can be specified without unit (then seconds are assumed) or together with &#x60;s&#x60;, &#x60;ms&#x60; or &#x60;m&#x60; unit. Example: &#x60;42s&#x60;, &#x60;1m&#x60;.  The default (if omitted) timeout is &#x60;10s&#x60;. Maximum value: &#x60;60s&#x60;.  A value of &#x60;0&#x60; applies fire and forget semantics for the command resulting in setting &#x60;response-required&#x3D;false&#x60;. | 
 **option** | **optional.String**| Possible values for the parameter:  #### Sort operations  * &#x60;&#x60;&#x60;sort([+|-]{property})&#x60;&#x60;&#x60; * &#x60;&#x60;&#x60;sort([+|-]{property},[+|-]{property},...)&#x60;&#x60;&#x60;  #### Paging operations  * &#x60;&#x60;&#x60;size({page-size})&#x60;&#x60;&#x60;  Maximum allowed page size is &#x60;200&#x60;. Default page size is &#x60;25&#x60;. * &#x60;&#x60;&#x60;cursor({cursor-id})&#x60;&#x60;&#x60;  Start the search from the cursor location. Specify the cursor ID without quotation marks. Cursor IDs are given in search responses and mark the position after the last entry of the previous search. The meaning of cursor IDs is unspecified and may change without notice.  The paging option &#x60;limit({offset},{count})&#x60; is deprecated. It may result in slow queries or timeouts and will be removed eventually.  #### Examples:  * &#x60;&#x60;&#x60;sort(+thingId)&#x60;&#x60;&#x60; * &#x60;&#x60;&#x60;sort(-attributes/manufacturer)&#x60;&#x60;&#x60; * &#x60;&#x60;&#x60;sort(+thingId,-attributes/manufacturer)&#x60;&#x60;&#x60; * &#x60;&#x60;&#x60;size(10)&#x60;&#x60;&#x60; return 10 results * &#x60;&#x60;&#x60;cursor(LOREMIPSUM)&#x60;&#x60;&#x60;  return results after the position of the cursor &#x60;LOREMIPSUM&#x60;.  #### Combine:  If you need to specify multiple options, when using the swagger UI just write each option in a new line. When using the plain REST API programmatically, you will need to separate the options using a comma (,) character.  &#x60;&#x60;&#x60;size(200),cursor(LOREMIPSUM)&#x60;&#x60;&#x60;  The deprecated paging option &#x60;limit&#x60; may not be combined with the other paging options &#x60;size&#x60; and &#x60;cursor&#x60;. | 

### Return type

[**SearchResultThings**](SearchResultThings.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

