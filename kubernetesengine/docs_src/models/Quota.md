# Quota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterQuota** | **int32** | 프로젝트에서 생성할 수 있는 클러스터 수 | 
**ClusterUsage** | **int32** | 프로젝트에서 사용 중인 클러스터 수 | 
**ProjectNodeUsage** | Pointer to [**[]ProjectNodeUsage**](ProjectNodeUsage.md) | 클러스터별 사용 중인 노드 수 | [optional] 

## Methods

### NewQuota

`func NewQuota(clusterQuota int32, clusterUsage int32, ) *Quota`

NewQuota instantiates a new Quota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaWithDefaults

`func NewQuotaWithDefaults() *Quota`

NewQuotaWithDefaults instantiates a new Quota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterQuota

`func (o *Quota) GetClusterQuota() int32`

GetClusterQuota returns the ClusterQuota field if non-nil, zero value otherwise.

### GetClusterQuotaOk

`func (o *Quota) GetClusterQuotaOk() (*int32, bool)`

GetClusterQuotaOk returns a tuple with the ClusterQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterQuota

`func (o *Quota) SetClusterQuota(v int32)`

SetClusterQuota sets ClusterQuota field to given value.


### GetClusterUsage

`func (o *Quota) GetClusterUsage() int32`

GetClusterUsage returns the ClusterUsage field if non-nil, zero value otherwise.

### GetClusterUsageOk

`func (o *Quota) GetClusterUsageOk() (*int32, bool)`

GetClusterUsageOk returns a tuple with the ClusterUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUsage

`func (o *Quota) SetClusterUsage(v int32)`

SetClusterUsage sets ClusterUsage field to given value.


### GetProjectNodeUsage

`func (o *Quota) GetProjectNodeUsage() []ProjectNodeUsage`

GetProjectNodeUsage returns the ProjectNodeUsage field if non-nil, zero value otherwise.

### GetProjectNodeUsageOk

`func (o *Quota) GetProjectNodeUsageOk() (*[]ProjectNodeUsage, bool)`

GetProjectNodeUsageOk returns a tuple with the ProjectNodeUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectNodeUsage

`func (o *Quota) SetProjectNodeUsage(v []ProjectNodeUsage)`

SetProjectNodeUsage sets ProjectNodeUsage field to given value.

### HasProjectNodeUsage

`func (o *Quota) HasProjectNodeUsage() bool`

HasProjectNodeUsage returns a boolean if a field has been set.

### SetProjectNodeUsageNil

`func (o *Quota) SetProjectNodeUsageNil(b bool)`

 SetProjectNodeUsageNil sets the value for ProjectNodeUsage to be an explicit nil

### UnsetProjectNodeUsage
`func (o *Quota) UnsetProjectNodeUsage()`

UnsetProjectNodeUsage ensures that no value is present for ProjectNodeUsage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


