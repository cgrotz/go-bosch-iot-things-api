# \MessagesApi

All URIs are relative to *https://things.eu-1.bosch-iot-suite.com/api/2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThingsThingIdFeaturesFeatureIdInboxMessagesMessageSubjectPost**](MessagesApi.md#ThingsThingIdFeaturesFeatureIdInboxMessagesMessageSubjectPost) | **Post** /things/{thingId}/features/{featureId}/inbox/messages/{messageSubject} | Send a message TO a specific feature of a specific thing
[**ThingsThingIdFeaturesFeatureIdOutboxMessagesMessageSubjectPost**](MessagesApi.md#ThingsThingIdFeaturesFeatureIdOutboxMessagesMessageSubjectPost) | **Post** /things/{thingId}/features/{featureId}/outbox/messages/{messageSubject} | Send a message FROM a specific feature of a specific thing
[**ThingsThingIdInboxClaimPost**](MessagesApi.md#ThingsThingIdInboxClaimPost) | **Post** /things/{thingId}/inbox/claim | Initiates claiming a specific thing in order to gain access
[**ThingsThingIdInboxMessagesMessageSubjectPost**](MessagesApi.md#ThingsThingIdInboxMessagesMessageSubjectPost) | **Post** /things/{thingId}/inbox/messages/{messageSubject} | Send a message TO a specific thing
[**ThingsThingIdOutboxMessagesMessageSubjectPost**](MessagesApi.md#ThingsThingIdOutboxMessagesMessageSubjectPost) | **Post** /things/{thingId}/outbox/messages/{messageSubject} | Send a message FROM a specific thing



## ThingsThingIdFeaturesFeatureIdInboxMessagesMessageSubjectPost

> ThingsThingIdFeaturesFeatureIdInboxMessagesMessageSubjectPost(ctx, thingId, featureId, messageSubject, optional)

Send a message TO a specific feature of a specific thing

Send a message with the subject `messageSubject` **to** the feature specified by the `featureId` and `thingId` path parameter. The request body contains the message payload and the `Content-Type` header defines its type. The API does not provide any kind of acknowledgement that the message was received by the feature.  The HTTP request blocks until a response to the message is available or until the `timeout` is expired. If many clients respond to the issued message, the first response will complete the HTTP request.  In order to handle the message in a fire and forget manner, add a query-parameter `timeout=0` to the request.  ### Who You will need `WRITE` permission on the root \"message:/\" resource, or at least the resource `message:/features/featureId/inbox/messages/messageSubject`. The receiving device needs `READ` permission on the resource. Such permission is managed  within the policy which controls the access on the thing.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**featureId** | **string**| The ID of the feature - has to conform to RFC-3986 (URI) | 
**messageSubject** | **string**| The subject of the Message - has to conform to RFC-3986 (URI) | 
 **optional** | ***ThingsThingIdFeaturesFeatureIdInboxMessagesMessageSubjectPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdFeaturesFeatureIdInboxMessagesMessageSubjectPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **timeout** | **optional.Int32**| Contains an optional timeout (in seconds) of how long to wait for the message response and therefore block the HTTP request. Default value (if omitted): 10 seconds. Maximum value: 60 seconds. A value of 0 seconds applies fire and forget semantics for the message. | 
 **body** | **optional.String**| Payload of the message with max size of 250 kB. It can be any HTTP supported content, including binary content. | 

### Return type

 (empty response body)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/octet-stream, text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsThingIdFeaturesFeatureIdOutboxMessagesMessageSubjectPost

> ThingsThingIdFeaturesFeatureIdOutboxMessagesMessageSubjectPost(ctx, thingId, featureId, messageSubject, optional)

Send a message FROM a specific feature of a specific thing

Send a message with the subject `messageSubject` **from** the feature specified by the `featureId` and `thingId` path parameter. The request body contains the message payload and the `Content-Type` header defines its type.  The HTTP request blocks until a response to the message is available or until the `timeout` is expired. If many clients respond to the issued message, the first response will complete the HTTP request.  In order to handle the message in a fire and forget manner, add a query-parameter `timeout=0` to the request.  ### Who You will need `WRITE` permission on the root \"message:/\" resource, or at least the resource `message:/features/featureId/outbox/messages/messageSubject`. Such permission is managed  within the policy which controls the access on the thing.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**featureId** | **string**| The ID of the feature - has to conform to RFC-3986 (URI) | 
**messageSubject** | **string**| The subject of the Message - has to conform to RFC-3986 (URI) | 
 **optional** | ***ThingsThingIdFeaturesFeatureIdOutboxMessagesMessageSubjectPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdFeaturesFeatureIdOutboxMessagesMessageSubjectPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **timeout** | **optional.Int32**| Contains an optional timeout (in seconds) of how long to wait for the message response and therefore block the HTTP request. Default value (if omitted): 10 seconds. Maximum value: 60 seconds. A value of 0 seconds applies fire and forget semantics for the message. | 
 **body** | **optional.String**| Payload of the message with max size of 250 kB. It can be any HTTP supported content, including binary content. | 

### Return type

 (empty response body)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/octet-stream, text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsThingIdInboxClaimPost

> ThingsThingIdInboxClaimPost(ctx, thingId, optional)

Initiates claiming a specific thing in order to gain access

### Why A claiming process may enable an end-user to claim things and proof ownership thereof. Such a process is initially triggered via a claim message. This message can be sent to the things service with the HTTP API or the things-client.  ### How At this resource you can send a \"claim\" message to the thing identified by the `thingId` path parameter in order to gain access to it. The \"claim\" message is forwarded together with the request body and `Content-Type` header to client(s) which registered for Claim messages of the specific thing.  The decision whether to grant access (by setting permissions) is completely up to the client(s) which handle the \"claim\" message.  The HTTP request blocks until a response to the issued \"claim\" message is available or until the `timeout` is expired. If many clients respond to the issued message, the first response will complete the HTTP request.  ### Who No special permission is required to issue a claim message.  ### Example See [Claiming](https://docs.bosch-iot-suite.com/things/basic-concepts/auth/claim/) concept in detail and example in GitHub. However, in that scenario, the policy should grant you READ and WRITE permission on the \"message:/\" resource in order to be able to send the message and read the response. Further, the things-client which handles the \"claim\" message, needs permission to change the policy itself (i.e. READ and WRITE permission on the \"policy:/\" resource).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
 **optional** | ***ThingsThingIdInboxClaimPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdInboxClaimPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeout** | **optional.Int32**| Contains an optional timeout (in seconds) of how long to wait for the Claim response and therefore block the HTTP request. Default value (if omitted): 60 seconds. Maximum value: 600 seconds. A value of 0 seconds applies fire and forget semantics for the message. | 
 **body** | **optional.String**| Payload of the message with max size of 250 kB. It can be any HTTP supported content, including binary content. | 

### Return type

 (empty response body)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/octet-stream, text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsThingIdInboxMessagesMessageSubjectPost

> ThingsThingIdInboxMessagesMessageSubjectPost(ctx, thingId, messageSubject, optional)

Send a message TO a specific thing

### Why A message can be sent to a thing or one of its features in order to invoke an operation on the device.  ### How Send a message with a `messageSubject` **to** the thing identified by the `thingId` path parameter. The request body contains the message payload and the `Content-Type` header defines its type.  The API does not provide any kind of acknowledgement that the thing has received the message.  The HTTP request blocks until a response to the message is available or until the `timeout` is expired. If many clients respond to the issued message, the first response will complete the HTTP request.  In order to handle the message in a fire and forget manner, add a query-parameter `timeout=0` to the request.  ### Who You will need `WRITE` permission on the root \"message:/\" resource, or at least the resource `message:/inbox/messages/messageSubject`. The receiving device needs `READ` permission on the resource. Such permission is managed within the policy which controls the access on the thing.  ### Example Given you have a \"coffemaker\" thing as shown in the examples for the `things` resources. The `messageSubject` understood by such a device would be \"makeCoffee\".  Further, as in our example the \"brewed-coffees\" counter would increase as a response, you would need `WRITE` permission for the things resource, at least at the respective path  `/things/{thingId}/features/coffee-brewer/properties/brewed-coffees`

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**messageSubject** | **string**| The subject of the Message - has to conform to RFC-3986 (URI) | 
 **optional** | ***ThingsThingIdInboxMessagesMessageSubjectPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdInboxMessagesMessageSubjectPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeout** | **optional.Int32**| Contains an optional timeout (in seconds) of how long to wait for the message response and therefore block the HTTP request. Default value (if omitted): 10 seconds. Maximum value: 60 seconds. A value of 0 seconds applies fire and forget semantics for the message. | 
 **body** | **optional.String**| Payload of the message with max size of 250 kB. It can be any HTTP supported content, including binary content. | 

### Return type

 (empty response body)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/octet-stream, text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThingsThingIdOutboxMessagesMessageSubjectPost

> ThingsThingIdOutboxMessagesMessageSubjectPost(ctx, thingId, messageSubject, optional)

Send a message FROM a specific thing

Send a message with the subject `messageSubject` **from** the thing identified by the `thingId` path parameter. The request body contains the message payload and the `Content-Type` header defines its type.  The HTTP request blocks until a response to the message is available or until the `timeout` is expired. If many clients respond to the issued message, the first response will complete the HTTP request.  In order to handle the message in a fire and forget manner, add a query-parameter `timeout=0` to the request.  ### Who You will need `WRITE` permission on the root \"message:/\" resource, or at least the resource `message:/outbox/messages/messageSubject`. Such permission is managed  within the policy which controls the access on the thing.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**thingId** | **string**| The ID of a thing needs to follow the namespaced entity ID notation (see [Ditto documentation on namespaced entity IDs](https://www.eclipse.org/ditto/basic-namespaces-and-names.html#namespaced-id)).  The namespace needs to be registered for your solution | 
**messageSubject** | **string**| The subject of the Message - has to conform to RFC-3986 (URI) | 
 **optional** | ***ThingsThingIdOutboxMessagesMessageSubjectPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ThingsThingIdOutboxMessagesMessageSubjectPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeout** | **optional.Int32**| Contains an optional timeout (in seconds) of how long to wait for the message response and therefore block the HTTP request. Default value (if omitted): 10 seconds. Maximum value: 60 seconds. A value of 0 seconds applies fire and forget semantics for the message. | 
 **body** | **optional.String**| Payload of the message with max size of 250 kB. It can be any HTTP supported content, including binary content. | 

### Return type

 (empty response body)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: application/json, application/octet-stream, text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

