# ProjectNodeUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **string** | 클러스터 이름 | 
**NodeCount** | **int32** | 클러스터에서 사용 중인 노드 수 | 

## Methods

### NewProjectNodeUsage

`func NewProjectNodeUsage(clusterName string, nodeCount int32, ) *ProjectNodeUsage`

NewProjectNodeUsage instantiates a new ProjectNodeUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectNodeUsageWithDefaults

`func NewProjectNodeUsageWithDefaults() *ProjectNodeUsage`

NewProjectNodeUsageWithDefaults instantiates a new ProjectNodeUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *ProjectNodeUsage) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ProjectNodeUsage) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ProjectNodeUsage) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetNodeCount

`func (o *ProjectNodeUsage) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ProjectNodeUsage) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ProjectNodeUsage) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


