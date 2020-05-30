# \SolutionsApi

All URIs are relative to *https://things.eu-1.bosch-iot-suite.com/api/2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SolutionsSolutionIdClientsClientIdDelete**](SolutionsApi.md#SolutionsSolutionIdClientsClientIdDelete) | **Delete** /solutions/{solutionId}/clients/{clientId} | Delete a specific authorized client (e.g. OAuth2) of a specific Solution
[**SolutionsSolutionIdClientsClientIdGet**](SolutionsApi.md#SolutionsSolutionIdClientsClientIdGet) | **Get** /solutions/{solutionId}/clients/{clientId} | Retrieve a specific authorized client (e.g. OAuth2) of a specific Solution
[**SolutionsSolutionIdClientsClientIdPut**](SolutionsApi.md#SolutionsSolutionIdClientsClientIdPut) | **Put** /solutions/{solutionId}/clients/{clientId} | Create or update a specific authorized client (e.g. OAuth2) of a specific Solution
[**SolutionsSolutionIdClientsDelete**](SolutionsApi.md#SolutionsSolutionIdClientsDelete) | **Delete** /solutions/{solutionId}/clients | Delete all authorized authorized clients (e.g. OAuth2) of a specific Solution
[**SolutionsSolutionIdClientsGet**](SolutionsApi.md#SolutionsSolutionIdClientsGet) | **Get** /solutions/{solutionId}/clients | Retrieve the authorized authorized clients (e.g. OAuth2) of a specific Solution.
[**SolutionsSolutionIdClientsPut**](SolutionsApi.md#SolutionsSolutionIdClientsPut) | **Put** /solutions/{solutionId}/clients | Create or update all authorized authorized clients (e.g. OAuth2) of a specific Solution
[**SolutionsSolutionIdConnectionsConnectionIdCommandPost**](SolutionsApi.md#SolutionsSolutionIdConnectionsConnectionIdCommandPost) | **Post** /solutions/{solutionId}/connections/{connectionId}/command | Send a command to a specific Connection.
[**SolutionsSolutionIdConnectionsConnectionIdDelete**](SolutionsApi.md#SolutionsSolutionIdConnectionsConnectionIdDelete) | **Delete** /solutions/{solutionId}/connections/{connectionId} | Delete a specific Connection registered for a specific Solution
[**SolutionsSolutionIdConnectionsConnectionIdGet**](SolutionsApi.md#SolutionsSolutionIdConnectionsConnectionIdGet) | **Get** /solutions/{solutionId}/connections/{connectionId} | Retrieve a specific Connection registered for a specific Solution
[**SolutionsSolutionIdConnectionsConnectionIdLogsGet**](SolutionsApi.md#SolutionsSolutionIdConnectionsConnectionIdLogsGet) | **Get** /solutions/{solutionId}/connections/{connectionId}/logs | Retrieve logs of a specific Connection
[**SolutionsSolutionIdConnectionsConnectionIdMetricsGet**](SolutionsApi.md#SolutionsSolutionIdConnectionsConnectionIdMetricsGet) | **Get** /solutions/{solutionId}/connections/{connectionId}/metrics | Retrieve metrics of a specific Connection
[**SolutionsSolutionIdConnectionsConnectionIdPut**](SolutionsApi.md#SolutionsSolutionIdConnectionsConnectionIdPut) | **Put** /solutions/{solutionId}/connections/{connectionId} | Update a specific Connection registered for a specific Solution
[**SolutionsSolutionIdConnectionsConnectionIdStatusGet**](SolutionsApi.md#SolutionsSolutionIdConnectionsConnectionIdStatusGet) | **Get** /solutions/{solutionId}/connections/{connectionId}/status | Retrieve status of a specific Connection
[**SolutionsSolutionIdConnectionsGet**](SolutionsApi.md#SolutionsSolutionIdConnectionsGet) | **Get** /solutions/{solutionId}/connections | Retrieve all Connections registered for a specific Solution
[**SolutionsSolutionIdConnectionsPost**](SolutionsApi.md#SolutionsSolutionIdConnectionsPost) | **Post** /solutions/{solutionId}/connections | Create a new Connection for a specific Solution
[**SolutionsSolutionIdGet**](SolutionsApi.md#SolutionsSolutionIdGet) | **Get** /solutions/{solutionId} | Retrieve a specific Solution
[**SolutionsSolutionIdKeyPut**](SolutionsApi.md#SolutionsSolutionIdKeyPut) | **Put** /solutions/{solutionId}/key | Update the public key of a specific Solution
[**SolutionsSolutionIdNamespacesGet**](SolutionsApi.md#SolutionsSolutionIdNamespacesGet) | **Get** /solutions/{solutionId}/namespaces | Retrieve all Namespaces registered for a specific Solution
[**SolutionsSolutionIdNamespacesNamespaceIdDelete**](SolutionsApi.md#SolutionsSolutionIdNamespacesNamespaceIdDelete) | **Delete** /solutions/{solutionId}/namespaces/{namespaceId} | Delete a specific Namespace - and all its associated entities - registered for a specific Solution
[**SolutionsSolutionIdNamespacesNamespaceIdDeletionGet**](SolutionsApi.md#SolutionsSolutionIdNamespacesNamespaceIdDeletionGet) | **Get** /solutions/{solutionId}/namespaces/{namespaceId}/deletion | Retrieve the deletion status report of a specific Namespace registered for a specific Solution
[**SolutionsSolutionIdNamespacesNamespaceIdGet**](SolutionsApi.md#SolutionsSolutionIdNamespacesNamespaceIdGet) | **Get** /solutions/{solutionId}/namespaces/{namespaceId} | Retrieve a specific Namespace registered for a specific Solution
[**SolutionsSolutionIdNamespacesNamespaceIdPut**](SolutionsApi.md#SolutionsSolutionIdNamespacesNamespaceIdPut) | **Put** /solutions/{solutionId}/namespaces/{namespaceId} | Create or update a specific Namespace registered for a specific Solution
[**SolutionsSolutionIdUsageGet**](SolutionsApi.md#SolutionsSolutionIdUsageGet) | **Get** /solutions/{solutionId}/usage | Retrieve the last calculated usage of a specific Solution



## SolutionsSolutionIdClientsClientIdDelete

> SolutionsSolutionIdClientsClientIdDelete(ctx, solutionId, clientId)

Delete a specific authorized client (e.g. OAuth2) of a specific Solution

Deletes an authorized client (e.g. OAuth2) specified by the `clientId` path parameter, of the solution identified by the `solutionId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 
**clientId** | **string**| The ID of the client | 

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


## SolutionsSolutionIdClientsClientIdGet

> Client SolutionsSolutionIdClientsClientIdGet(ctx, solutionId, clientId)

Retrieve a specific authorized client (e.g. OAuth2) of a specific Solution

Returns the authorized client (e.g. OAuth2) specified by the `clientId` path parameter, of the solution identified by the `solutionId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 
**clientId** | **string**| The ID of the client | 

### Return type

[**Client**](Client.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolutionsSolutionIdClientsClientIdPut

> SolutionsSolutionIdClientsClientIdPut(ctx, solutionId, clientId, client)

Create or update a specific authorized client (e.g. OAuth2) of a specific Solution

Create or update an authorized client (e.g. OAuth2) specified by the `clientId` path parameter, of the solution identified by the `solutionId` path parameter.  The client will be authorized to make requests in the context of the solution. The request body must be a JSON object containing the entries for one client.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 
**clientId** | **string**| The ID of the client | 
**client** | [**Client**](Client.md)| JSON representation of the client to be modified. | 

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


## SolutionsSolutionIdClientsDelete

> SolutionsSolutionIdClientsDelete(ctx, solutionId)

Delete all authorized authorized clients (e.g. OAuth2) of a specific Solution

Deletes all authorized clients (e.g. OAuth2) of the solution identified by the `solutionId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 

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


## SolutionsSolutionIdClientsGet

> Clients SolutionsSolutionIdClientsGet(ctx, solutionId)

Retrieve the authorized authorized clients (e.g. OAuth2) of a specific Solution.

Returns all authorized clients (e.g. OAuth2) of the solution.  Clients registered here are authorized to make requests in the context of the solution identified by the `solutionId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 

### Return type

[**Clients**](Clients.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolutionsSolutionIdClientsPut

> SolutionsSolutionIdClientsPut(ctx, solutionId, clients)

Create or update all authorized authorized clients (e.g. OAuth2) of a specific Solution

Create or update all authorized clients (e.g. OAuth2) at once.  These clients - specified in the request body - will be authorized to make requests in the context of the solution identified by the `solutionId` path parameter.  The request body must be a JSON object containing one or more clients.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 
**clients** | [**Clients**](Clients.md)| JSON representation of the clients to be modified. | 

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


## SolutionsSolutionIdConnectionsConnectionIdCommandPost

> SolutionsSolutionIdConnectionsConnectionIdCommandPost(ctx, solutionId, connectionId, body)

Send a command to a specific Connection.

Sends the command specified in the body to the connection identified by the `solutionId` and `connectionId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 
**connectionId** | **string**| The ID of the connection | 
**body** | **string**| The command to send. Supported commands are   * &#x60;connectivity.commands:openConnection&#x60;   * &#x60;connectivity.commands:closeConnection&#x60;   * &#x60;connectivity.commands:resetConnectionMetrics&#x60;   * &#x60;connectivity.commands:enableConnectionLogs&#x60;   * &#x60;connectivity.commands:resetConnectionLogs&#x60; | 

### Return type

 (empty response body)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolutionsSolutionIdConnectionsConnectionIdDelete

> SolutionsSolutionIdConnectionsConnectionIdDelete(ctx, solutionId, connectionId)

Delete a specific Connection registered for a specific Solution

Delete the connection identified by the `solutionId` and `connectionId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 
**connectionId** | **string**| The ID of the connection | 

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


## SolutionsSolutionIdConnectionsConnectionIdGet

> Connection SolutionsSolutionIdConnectionsConnectionIdGet(ctx, solutionId, connectionId)

Retrieve a specific Connection registered for a specific Solution

Returns the connection identified by the `solutionId` and `connectionId` path parameter.  Tip: If you subscribe to the 'Bosch IoT Suite for Asset Communication' package, the connection between Hub and things is created automatically.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 
**connectionId** | **string**| The ID of the connection | 

### Return type

[**Connection**](Connection.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolutionsSolutionIdConnectionsConnectionIdLogsGet

> ConnectionLogs SolutionsSolutionIdConnectionsConnectionIdLogsGet(ctx, solutionId, connectionId)

Retrieve logs of a specific Connection

Returns the logs of the connection identified by the `solutionId` and `connectionId` path parameter.  **Before** log entries are generated and returned, logging needs be enabled with the `command` `connectivity.commands:enableConnectionLogs`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 
**connectionId** | **string**| The ID of the connection | 

### Return type

[**ConnectionLogs**](ConnectionLogs.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolutionsSolutionIdConnectionsConnectionIdMetricsGet

> ConnectionMetrics SolutionsSolutionIdConnectionsConnectionIdMetricsGet(ctx, solutionId, connectionId)

Retrieve metrics of a specific Connection

Returns the metrics of the connection identified by the `solutionId` and `connectionId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 
**connectionId** | **string**| The ID of the connection | 

### Return type

[**ConnectionMetrics**](ConnectionMetrics.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolutionsSolutionIdConnectionsConnectionIdPut

> SolutionsSolutionIdConnectionsConnectionIdPut(ctx, solutionId, connectionId, connection)

Update a specific Connection registered for a specific Solution

Update the connection identified by the `solutionId` and `connectionId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 
**connectionId** | **string**| The ID of the connection | 
**connection** | [**Connection**](Connection.md)|  | 

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


## SolutionsSolutionIdConnectionsConnectionIdStatusGet

> ConnectionStatus SolutionsSolutionIdConnectionsConnectionIdStatusGet(ctx, solutionId, connectionId)

Retrieve status of a specific Connection

Returns the status of the connection identified by the `solutionId` and `connectionId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 
**connectionId** | **string**| The ID of the connection | 

### Return type

[**ConnectionStatus**](ConnectionStatus.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolutionsSolutionIdConnectionsGet

> []Connection SolutionsSolutionIdConnectionsGet(ctx, solutionId)

Retrieve all Connections registered for a specific Solution

Returns all connections of a solution specified by the `solutionId` parameter.  Tip: If you subscribe to the 'Bosch IoT Suite for Asset Communication' package, the connection between Hub and things is created automatically.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 

### Return type

[**[]Connection**](Connection.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolutionsSolutionIdConnectionsPost

> SolutionsSolutionIdConnectionsPost(ctx, solutionId, newConnection, optional)

Create a new Connection for a specific Solution

Creates the connection defined in the JSON body for the solution specified by the `solutionId` parameter.  The ID of the connection will be generated by the backend. Any `ID` specified in the request body is therefore ignored.  All other IDs can be looked up from the credentials section available after a successful service subscription see https://accounts.bosch-iot-suite.com/subscriptions.  Supported connection types are `amqp-091`, `amqp-10`, `mqtt`, `kafka` and `http`.  The connection defines a `ConnectionStatus` mapping and references it by its ID `status` in the telemetry/event source.  Tip: If you subscribe to the 'Bosch IoT Suite for Asset Communication' package, the connection between Hub and things is created automatically.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 
**newConnection** | [**NewConnection**](NewConnection.md)| The example below shows a connection to Bosch IoT Hub. | 
 **optional** | ***SolutionsSolutionIdConnectionsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SolutionsSolutionIdConnectionsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dryRun** | **optional.Bool**| When set to true, the request will not try to create the connection, but only try to connect it. You can use this parameter to verify that the given connection is able to communicate with your external system. | 

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


## SolutionsSolutionIdGet

> Solution SolutionsSolutionIdGet(ctx, solutionId)

Retrieve a specific Solution

Returns the complete solution entity identified by the `solutionId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 

### Return type

[**Solution**](Solution.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolutionsSolutionIdKeyPut

> SolutionsSolutionIdKeyPut(ctx, solutionId, body)

Update the public key of a specific Solution

Updates the public key of the solution identified by the `solutionId` path parameter.  A public key can be used by a technical client to communicate with the Bosch IoT Things backend.  The request body must be a JSON string beginning and ending with double quotes.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 
**body** | **string**| The public key file in PKCS#8 format or a X.509 certificate as JSON string representation. We only accept keys that were created using the elliptic curve algorithm \&quot;EC\&quot; at the moment. | 

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


## SolutionsSolutionIdNamespacesGet

> Namespaces SolutionsSolutionIdNamespacesGet(ctx, solutionId)

Retrieve all Namespaces registered for a specific Solution

Returns all namespaces registered for the solution identified by the `solutionId` path parameter.  A namespace is necessary to link things and policies to the specific solution.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 

### Return type

[**Namespaces**](Namespaces.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolutionsSolutionIdNamespacesNamespaceIdDelete

> SolutionsSolutionIdNamespacesNamespaceIdDelete(ctx, solutionId, namespaceId)

Delete a specific Namespace - and all its associated entities - registered for a specific Solution

Deletes the namespace identified by the `namespaceId` path parameter registered for the Solution identified by the `solutionId` path parameter.  ### Note: By deleting a namespace you automatically delete multiple entities Entities like things and policies, that are part of the namespace, are **deleted permanently** as part of the deletion process.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 
**namespaceId** | **string**| The ID of the Namespace | 

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


## SolutionsSolutionIdNamespacesNamespaceIdDeletionGet

> NamespacePurgingState SolutionsSolutionIdNamespacesNamespaceIdDeletionGet(ctx, solutionId, namespaceId)

Retrieve the deletion status report of a specific Namespace registered for a specific Solution

Returns a status report of the namespace deletion process including the current step number and the amount of total steps as indication of progress.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 
**namespaceId** | **string**| The ID of the Namespace | 

### Return type

[**NamespacePurgingState**](NamespacePurgingState.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolutionsSolutionIdNamespacesNamespaceIdGet

> NamespaceEntry SolutionsSolutionIdNamespacesNamespaceIdGet(ctx, solutionId, namespaceId)

Retrieve a specific Namespace registered for a specific Solution

Returns the namespace identified by the `namespaceId` path parameter registered for the solution identified by the `solutionId` path parameter.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 
**namespaceId** | **string**| The ID of the Namespace | 

### Return type

[**NamespaceEntry**](NamespaceEntry.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SolutionsSolutionIdNamespacesNamespaceIdPut

> SolutionsSolutionIdNamespacesNamespaceIdPut(ctx, solutionId, namespaceId, namespaceEntry)

Create or update a specific Namespace registered for a specific Solution

Create or update a namespace specified by the `namespaceId` path parameter, and the JSON body. * If you set a new `namespaceId` in the path, a namespace will be created. * If you set an existing `namespaceId` in the path, the namespace will be updated.  A namespace is necessary to link things and policies to the specific solution.  The request body must be a JSON object containing a namespace entry.  The namespace entry contains the information if the namespaceId is used as default namespace.  ```{\"default\": true}``` Note: Only one namespace can be the default.  **Restriction**: If your solution is based on a Free and Starter plan only **one namespace** is allowed.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 
**namespaceId** | **string**| The ID of the Namespace | 
**namespaceEntry** | [**NamespaceEntry**](NamespaceEntry.md)| JSON representation of the namespace to be modified. | 

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


## SolutionsSolutionIdUsageGet

> Usage SolutionsSolutionIdUsageGet(ctx, solutionId)

Retrieve the last calculated usage of a specific Solution

Returns the last calculated usage of the solution identified by the `solutionId` path parameter.  The numbers might differ from the effective number reported for billing purposes. Please note however, that quota usage requests and all responses will also count as transactions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**solutionId** | **string**| The ID of the Solution | 

### Return type

[**Usage**](Usage.md)

### Authorization

[BoschID](../README.md#BoschID), [basicAuth](../README.md#basicAuth), [bearerAuth](../README.md#bearerAuth), [thingsApiToken](../README.md#thingsApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

