# NodePoolLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | [**[]Label**](Label.md) | 노드 풀에 적용된 레이블 목록 | 
**Name** | **string** | 노드 풀 이름 | 
**Id** | **string** | 노드 풀의 고유 ID | 

## Methods

### NewNodePoolLabel

`func NewNodePoolLabel(labels []Label, name string, id string, ) *NodePoolLabel`

NewNodePoolLabel instantiates a new NodePoolLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePoolLabelWithDefaults

`func NewNodePoolLabelWithDefaults() *NodePoolLabel`

NewNodePoolLabelWithDefaults instantiates a new NodePoolLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *NodePoolLabel) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NodePoolLabel) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NodePoolLabel) SetLabels(v []Label)`

SetLabels sets Labels field to given value.


### GetName

`func (o *NodePoolLabel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodePoolLabel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodePoolLabel) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *NodePoolLabel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodePoolLabel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodePoolLabel) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


