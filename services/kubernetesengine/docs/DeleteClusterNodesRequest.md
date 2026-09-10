# DeleteClusterNodesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**DeleteClusterNodes**](DeleteClusterNodes.md) | 대상 클러스터 정보 | 

## Methods

### NewDeleteClusterNodesRequest

`func NewDeleteClusterNodesRequest(cluster DeleteClusterNodes, ) *DeleteClusterNodesRequest`

NewDeleteClusterNodesRequest instantiates a new DeleteClusterNodesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteClusterNodesRequestWithDefaults

`func NewDeleteClusterNodesRequestWithDefaults() *DeleteClusterNodesRequest`

NewDeleteClusterNodesRequestWithDefaults instantiates a new DeleteClusterNodesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *DeleteClusterNodesRequest) GetCluster() DeleteClusterNodes`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DeleteClusterNodesRequest) GetClusterOk() (*DeleteClusterNodes, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DeleteClusterNodesRequest) SetCluster(v DeleteClusterNodes)`

SetCluster sets Cluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


