# SetClusterNodesCordonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**SetClusterNodesCordon**](SetClusterNodesCordon.md) | 클러스터 이름 | 

## Methods

### NewSetClusterNodesCordonRequest

`func NewSetClusterNodesCordonRequest(cluster SetClusterNodesCordon, ) *SetClusterNodesCordonRequest`

NewSetClusterNodesCordonRequest instantiates a new SetClusterNodesCordonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetClusterNodesCordonRequestWithDefaults

`func NewSetClusterNodesCordonRequestWithDefaults() *SetClusterNodesCordonRequest`

NewSetClusterNodesCordonRequestWithDefaults instantiates a new SetClusterNodesCordonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *SetClusterNodesCordonRequest) GetCluster() SetClusterNodesCordon`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *SetClusterNodesCordonRequest) GetClusterOk() (*SetClusterNodesCordon, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *SetClusterNodesCordonRequest) SetCluster(v SetClusterNodesCordon)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


