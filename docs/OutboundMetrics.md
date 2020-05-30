# OutboundMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dispatched** | [**TypedMetric**](TypedMetric.md) | Contains internally dispatched (e.g. a Ditto event) metric counts | 
**Filtered** | [**TypedMetric**](TypedMetric.md) | Contains the metric counts for messages which passed the filter (e.g. namespace or RQL filter for events) | 
**Mapped** | [**TypedMetric**](TypedMetric.md) | Contains mapped (payload mapping) messages metric counts | 
**Dropped** | [**TypedMetric**](TypedMetric.md) | Contains dropped (in the payload mapping) messages metric counts | 
**Published** | [**TypedMetric**](TypedMetric.md) | Contains published messages metric counts meaning those messages were published to the external source | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


