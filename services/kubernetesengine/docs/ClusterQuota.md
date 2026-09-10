# ClusterQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BmNodeQuota** | **int32** | 클러스터에서 생성할 수 있는 베어메탈 노드 수 | 
**BmNodeUsage** | Pointer to [**[]NodeUsage**](NodeUsage.md) |  | [optional] 
**NodeQuota** | **int32** | 클러스터에서 생성할 수 있는 노드 수 | 
**NodeUsage** | Pointer to [**[]NodeUsage**](NodeUsage.md) |  | [optional] 
**NodepoolQuota** | **int32** | 클러스터에서 생성할 수 있는 노드 풀 수 | 
**NodepoolUsage** | **int32** | 클러스터에서 사용 중인 노드 풀 수 | 

## Methods

### NewClusterQuota

`func NewClusterQuota(bmNodeQuota int32, nodeQuota int32, nodepoolQuota int32, nodepoolUsage int32, ) *ClusterQuota`

NewClusterQuota instantiates a new ClusterQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterQuotaWithDefaults

`func NewClusterQuotaWithDefaults() *ClusterQuota`

NewClusterQuotaWithDefaults instantiates a new ClusterQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBmNodeQuota

`func (o *ClusterQuota) GetBmNodeQuota() int32`

GetBmNodeQuota returns the BmNodeQuota field if non-nil, zero value otherwise.

### GetBmNodeQuotaOk

`func (o *ClusterQuota) GetBmNodeQuotaOk() (*int32, bool)`

GetBmNodeQuotaOk returns a tuple with the BmNodeQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmNodeQuota

`func (o *ClusterQuota) SetBmNodeQuota(v int32)`

SetBmNodeQuota sets BmNodeQuota field to given value.


### GetBmNodeUsage

`func (o *ClusterQuota) GetBmNodeUsage() []NodeUsage`

GetBmNodeUsage returns the BmNodeUsage field if non-nil, zero value otherwise.

### GetBmNodeUsageOk

`func (o *ClusterQuota) GetBmNodeUsageOk() (*[]NodeUsage, bool)`

GetBmNodeUsageOk returns a tuple with the BmNodeUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmNodeUsage

`func (o *ClusterQuota) SetBmNodeUsage(v []NodeUsage)`

SetBmNodeUsage sets BmNodeUsage field to given value.

### HasBmNodeUsage

`func (o *ClusterQuota) HasBmNodeUsage() bool`

HasBmNodeUsage returns a boolean if a field has been set.

### SetBmNodeUsageNil

`func (o *ClusterQuota) SetBmNodeUsageNil(b bool)`

 SetBmNodeUsageNil sets the value for BmNodeUsage to be an explicit nil

### UnsetBmNodeUsage
`func (o *ClusterQuota) UnsetBmNodeUsage()`

UnsetBmNodeUsage ensures that no value is present for BmNodeUsage, not even an explicit nil
### GetNodeQuota

`func (o *ClusterQuota) GetNodeQuota() int32`

GetNodeQuota returns the NodeQuota field if non-nil, zero value otherwise.

### GetNodeQuotaOk

`func (o *ClusterQuota) GetNodeQuotaOk() (*int32, bool)`

GetNodeQuotaOk returns a tuple with the NodeQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeQuota

`func (o *ClusterQuota) SetNodeQuota(v int32)`

SetNodeQuota sets NodeQuota field to given value.


### GetNodeUsage

`func (o *ClusterQuota) GetNodeUsage() []NodeUsage`

GetNodeUsage returns the NodeUsage field if non-nil, zero value otherwise.

### GetNodeUsageOk

`func (o *ClusterQuota) GetNodeUsageOk() (*[]NodeUsage, bool)`

GetNodeUsageOk returns a tuple with the NodeUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeUsage

`func (o *ClusterQuota) SetNodeUsage(v []NodeUsage)`

SetNodeUsage sets NodeUsage field to given value.

### HasNodeUsage

`func (o *ClusterQuota) HasNodeUsage() bool`

HasNodeUsage returns a boolean if a field has been set.

### SetNodeUsageNil

`func (o *ClusterQuota) SetNodeUsageNil(b bool)`

 SetNodeUsageNil sets the value for NodeUsage to be an explicit nil

### UnsetNodeUsage
`func (o *ClusterQuota) UnsetNodeUsage()`

UnsetNodeUsage ensures that no value is present for NodeUsage, not even an explicit nil
### GetNodepoolQuota

`func (o *ClusterQuota) GetNodepoolQuota() int32`

GetNodepoolQuota returns the NodepoolQuota field if non-nil, zero value otherwise.

### GetNodepoolQuotaOk

`func (o *ClusterQuota) GetNodepoolQuotaOk() (*int32, bool)`

GetNodepoolQuotaOk returns a tuple with the NodepoolQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodepoolQuota

`func (o *ClusterQuota) SetNodepoolQuota(v int32)`

SetNodepoolQuota sets NodepoolQuota field to given value.


### GetNodepoolUsage

`func (o *ClusterQuota) GetNodepoolUsage() int32`

GetNodepoolUsage returns the NodepoolUsage field if non-nil, zero value otherwise.

### GetNodepoolUsageOk

`func (o *ClusterQuota) GetNodepoolUsageOk() (*int32, bool)`

GetNodepoolUsageOk returns a tuple with the NodepoolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodepoolUsage

`func (o *ClusterQuota) SetNodepoolUsage(v int32)`

SetNodepoolUsage sets NodepoolUsage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


