/*
 * Bosch IoT Things HTTP API
 *
 * Bosch IoT Things enables applications to manage digital twins of IoT device assets in a simple, convenient, robust, and secure way.  These descriptions focus on the JSON-based, REST-like **HTTP API 2** of the Bosch IoT Things service.  Find details in our [documentation](https://docs.bosch-iot-suite.com/things/).
 *
 * API version: 2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package things

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// ThingsSearchApiService ThingsSearchApi service
type ThingsSearchApiService service

// SearchThingsCountGetOpts Optional parameters for the method 'SearchThingsCountGet'
type SearchThingsCountGetOpts struct {
    Filter optional.String
    Namespaces optional.String
    Timeout optional.String
}

/*
SearchThingsCountGet Count things
This resource can be used to count things.  The query parameter &#x60;filter&#x60; is not mandatory. If it is not set there is returned the total amount of things which the logged in user is allowed to read.  To search for nested properties, we use JSON Pointer notation (RFC-6901). See the following example how to search for the sub property &#x60;location&#x60; of the parent property &#x60;attributes&#x60; with a forward slash as separator:  &#x60;&#x60;&#x60;eq(attributes/location,\&quot;kitchen\&quot;)&#x60;&#x60;&#x60;  The search result is limited to the things within the namespace (or namespaces) the solution is configured for. You can explicitly search in other namespaces by including them in the query via the &#x60;namespaces&#x60; parameter.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *SearchThingsCountGetOpts - Optional Parameters:
 * @param "Filter" (optional.String) -   #### Filter predicates:  * ```eq({property},{value})```  (i.e. equal to the given value)  * ```ne({property},{value})```  (i.e. not equal to the given value)  * ```gt({property},{value})```  (i.e. greater than the given value)  * ```ge({property},{value})```  (i.e. equal to the given value or greater than it)  * ```lt({property},{value})```  (i.e. lower than the given value or equal to it)  * ```le({property},{value})```  (i.e. lower than the given value)  * ```in({property},{value},{value},...)```  (i.e. contains at least one of the values listed)  * ```like({property},{value})```  (i.e. contains values similar to the expressions listed)  * ```exists({property})```  (i.e. all things in which the given path exists)   Note: When using filter operations, only things with the specified properties are returned. For example, the filter `ne(attributes/owner, \"SID123\")` will only return things that do have the `owner` attribute.   #### Logical operations:   * ```and({query},{query},...)```  * ```or({query},{query},...)```  * ```not({query})```   #### Examples:  * ```eq(attributes/location,\"kitchen\")```  * ```ge(thingId,\"myThing1\")```  * ```exists(features/featureId)```  * ```and(eq(attributes/location,\"kitchen\"),eq(attributes/color,\"red\"))```  * ```or(eq(attributes/location,\"kitchen\"),eq(attributes/location,\"living-room\"))```  * ```like(attributes/key1,\"known-chars-at-start*\")```  * ```like(attributes/key1,\"*known-chars-at-end\")```  * ```like(attributes/key1,\"*known-chars-in-between*\")```  * ```like(attributes/key1,\"just-som?-char?-unkn?wn\")```  The `like` filters with the wildcard `*` at the beginning can slow down your search request.
 * @param "Namespaces" (optional.String) -  A comma-separated list of namespaces. This list is used to limit the query to things in the given namespaces only. If this parameter is omitted, all registered namespaces of your solution will be queried. The solution is determined by the API token sent with the request.   #### Examples:  * `?namespaces=com.example.namespace`  * `?namespaces=com.example.namespace1,com.example.namespace2`
 * @param "Timeout" (optional.String) -  Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the `requested-acks` param. Can be specified without unit (then seconds are assumed) or together with `s`, `ms` or `m` unit. Example: `42s`, `1m`.  The default (if omitted) timeout is `10s`. Maximum value: `60s`.  A value of `0` applies fire and forget semantics for the command resulting in setting `response-required=false`.
@return int32
*/
func (a *ThingsSearchApiService) SearchThingsCountGet(ctx _context.Context, localVarOptionals *SearchThingsCountGetOpts) (int32, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  int32
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/search/things/count"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Filter.IsSet() {
		localVarQueryParams.Add("filter", parameterToString(localVarOptionals.Filter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Namespaces.IsSet() {
		localVarQueryParams.Add("namespaces", parameterToString(localVarOptionals.Namespaces.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Timeout.IsSet() {
		localVarQueryParams.Add("timeout", parameterToString(localVarOptionals.Timeout.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-cr-api-token"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v AdvancedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AdvancedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v AdvancedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AdvancedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 504 {
			var v AdvancedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// SearchThingsGetOpts Optional parameters for the method 'SearchThingsGet'
type SearchThingsGetOpts struct {
    Filter optional.String
    Namespaces optional.String
    Fields optional.String
    Timeout optional.String
    Option optional.String
}

/*
SearchThingsGet Search for things
This resource can be used to search for things.  * The query parameter &#x60;filter&#x60; is not mandatory. If it is not set, the   result contains all things which the logged in user is allowed to read.  * The search is case sensitive. In case you don&#39;t know how exactly the   spelling of the namespace, name, attribute, feature etc. is, use the *like*   notation for filtering  * The search result is limited to the things within the namespace (or namespaces)   the solution is configured for. You can explicitly search in specific namespaces   by including them in the query via the *namespaces* parameter.  * The resource supports sorting and paging. If paging is not explicitly   specified by means of the &#x60;size&#x60; option, a default count of &#x60;25&#x60;   documents is returned.  * The internal search index is \&quot;eventually consistent\&quot;.  Consistency with the latest   thing updates should recover within milliseconds.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *SearchThingsGetOpts - Optional Parameters:
 * @param "Filter" (optional.String) -   #### Filter predicates:  * ```eq({property},{value})```  (i.e. equal to the given value)  * ```ne({property},{value})```  (i.e. not equal to the given value)  * ```gt({property},{value})```  (i.e. greater than the given value)  * ```ge({property},{value})```  (i.e. equal to the given value or greater than it)  * ```lt({property},{value})```  (i.e. lower than the given value or equal to it)  * ```le({property},{value})```  (i.e. lower than the given value)  * ```in({property},{value},{value},...)```  (i.e. contains at least one of the values listed)  * ```like({property},{value})```  (i.e. contains values similar to the expressions listed)  * ```exists({property})```  (i.e. all things in which the given path exists)   Note: When using filter operations, only things with the specified properties are returned. For example, the filter `ne(attributes/owner, \"SID123\")` will only return things that do have the `owner` attribute.   #### Logical operations:   * ```and({query},{query},...)```  * ```or({query},{query},...)```  * ```not({query})```   #### Examples:  * ```eq(attributes/location,\"kitchen\")```  * ```ge(thingId,\"myThing1\")```  * ```exists(features/featureId)```  * ```and(eq(attributes/location,\"kitchen\"),eq(attributes/color,\"red\"))```  * ```or(eq(attributes/location,\"kitchen\"),eq(attributes/location,\"living-room\"))```  * ```like(attributes/key1,\"known-chars-at-start*\")```  * ```like(attributes/key1,\"*known-chars-at-end\")```  * ```like(attributes/key1,\"*known-chars-in-between*\")```  * ```like(attributes/key1,\"just-som?-char?-unkn?wn\")```  The `like` filters with the wildcard `*` at the beginning can slow down your search request.
 * @param "Namespaces" (optional.String) -  A comma-separated list of namespaces. This list is used to limit the query to things in the given namespaces only. If this parameter is omitted, all registered namespaces of your solution will be queried. The solution is determined by the API token sent with the request.   #### Examples:  * `?namespaces=com.example.namespace`  * `?namespaces=com.example.namespace1,com.example.namespace2`
 * @param "Fields" (optional.String) -  Contains a comma-separated list of fields to be included in the returned JSON. attributes can be selected in the same manner.  #### Selectable fields  * `thingId` * `policyId` * `definition` * `attributes`     Supports selecting arbitrary sub-fields by using a comma-separated list:     * several attribute paths can be passed as a comma-separated list of JSON pointers (RFC-6901)        For example:         * `?fields=attributes/model` would select only `model` attribute value (if present)         * `?fields=attributes/model,attributes/location` would select only `model` and            `location` attribute values (if present)    Supports selecting arbitrary sub-fields of objects by wrapping sub-fields inside parentheses `( )`:     * a comma-separated list of sub-fields (a sub-field is a JSON pointer (RFC-6901)       separated with `/`) to select      * sub-selectors can be used to request only specific sub-fields by placing expressions       in parentheses `( )` after a selected subfield        For example:        * `?fields=attributes(model,location)` would select only `model`           and `location` attribute values (if present)        * `?fields=attributes(coffeemaker/serialno)` would select the `serialno` value           inside the `coffeemaker` object        * `?fields=attributes/address/postal(city,street)` would select the `city` and           `street` values inside the `postal` object inside the `address` object  * `features`    Supports selecting arbitrary fields in features similar to `attributes` (see also features documentation for more details)  * `_namespace`    Specifically selects the namespace also contained in the `thingId`  * `_revision`    Specifically selects the revision of the thing. The revision is a counter, which is incremented on each modification of a thing.  * `_modified`    Specifically selects the modified timestamp of the thing in ISO-8601 UTC format. The timestamp is set on each modification of a thing.  * `_policy`    Specifically selects the content of the policy associated to the thing. (By default, only the policyId is returned.)  #### Examples  * `?fields=thingId,attributes,features` * `?fields=attributes(model,manufacturer),features`
 * @param "Timeout" (optional.String) -  Defines how long the backend should wait for completion of the request, e.g. applied when waiting for requested acknowledgements via the `requested-acks` param. Can be specified without unit (then seconds are assumed) or together with `s`, `ms` or `m` unit. Example: `42s`, `1m`.  The default (if omitted) timeout is `10s`. Maximum value: `60s`.  A value of `0` applies fire and forget semantics for the command resulting in setting `response-required=false`.
 * @param "Option" (optional.String) -  Possible values for the parameter:  #### Sort operations  * ```sort([+|-]{property})``` * ```sort([+|-]{property},[+|-]{property},...)```  #### Paging operations  * ```size({page-size})```  Maximum allowed page size is `200`. Default page size is `25`. * ```cursor({cursor-id})```  Start the search from the cursor location. Specify the cursor ID without quotation marks. Cursor IDs are given in search responses and mark the position after the last entry of the previous search. The meaning of cursor IDs is unspecified and may change without notice.  The paging option `limit({offset},{count})` is deprecated. It may result in slow queries or timeouts and will be removed eventually.  #### Examples:  * ```sort(+thingId)``` * ```sort(-attributes/manufacturer)``` * ```sort(+thingId,-attributes/manufacturer)``` * ```size(10)``` return 10 results * ```cursor(LOREMIPSUM)```  return results after the position of the cursor `LOREMIPSUM`.  #### Combine:  If you need to specify multiple options, when using the swagger UI just write each option in a new line. When using the plain REST API programmatically, you will need to separate the options using a comma (,) character.  ```size(200),cursor(LOREMIPSUM)```  The deprecated paging option `limit` may not be combined with the other paging options `size` and `cursor`.
@return SearchResultThings
*/
func (a *ThingsSearchApiService) SearchThingsGet(ctx _context.Context, localVarOptionals *SearchThingsGetOpts) (SearchResultThings, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SearchResultThings
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/search/things"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Filter.IsSet() {
		localVarQueryParams.Add("filter", parameterToString(localVarOptionals.Filter.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Namespaces.IsSet() {
		localVarQueryParams.Add("namespaces", parameterToString(localVarOptionals.Namespaces.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Timeout.IsSet() {
		localVarQueryParams.Add("timeout", parameterToString(localVarOptionals.Timeout.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Option.IsSet() {
		localVarQueryParams.Add("option", parameterToString(localVarOptionals.Option.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-cr-api-token"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v AdvancedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AdvancedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v AdvancedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AdvancedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 504 {
			var v AdvancedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
