# NodeLabelsRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddOrUpdateLabels** | Pointer to [**[]LabelRequestModel**](LabelRequestModel.md) |  | [optional] 
**RemoveLabelKeys** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNodeLabelsRequestModel

`func NewNodeLabelsRequestModel() *NodeLabelsRequestModel`

NewNodeLabelsRequestModel instantiates a new NodeLabelsRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeLabelsRequestModelWithDefaults

`func NewNodeLabelsRequestModelWithDefaults() *NodeLabelsRequestModel`

NewNodeLabelsRequestModelWithDefaults instantiates a new NodeLabelsRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddOrUpdateLabels

`func (o *NodeLabelsRequestModel) GetAddOrUpdateLabels() []LabelRequestModel`

GetAddOrUpdateLabels returns the AddOrUpdateLabels field if non-nil, zero value otherwise.

### GetAddOrUpdateLabelsOk

`func (o *NodeLabelsRequestModel) GetAddOrUpdateLabelsOk() (*[]LabelRequestModel, bool)`

GetAddOrUpdateLabelsOk returns a tuple with the AddOrUpdateLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOrUpdateLabels

`func (o *NodeLabelsRequestModel) SetAddOrUpdateLabels(v []LabelRequestModel)`

SetAddOrUpdateLabels sets AddOrUpdateLabels field to given value.

### HasAddOrUpdateLabels

`func (o *NodeLabelsRequestModel) HasAddOrUpdateLabels() bool`

HasAddOrUpdateLabels returns a boolean if a field has been set.

### SetAddOrUpdateLabelsNil

`func (o *NodeLabelsRequestModel) SetAddOrUpdateLabelsNil(b bool)`

 SetAddOrUpdateLabelsNil sets the value for AddOrUpdateLabels to be an explicit nil

### UnsetAddOrUpdateLabels
`func (o *NodeLabelsRequestModel) UnsetAddOrUpdateLabels()`

UnsetAddOrUpdateLabels ensures that no value is present for AddOrUpdateLabels, not even an explicit nil
### GetRemoveLabelKeys

`func (o *NodeLabelsRequestModel) GetRemoveLabelKeys() []string`

GetRemoveLabelKeys returns the RemoveLabelKeys field if non-nil, zero value otherwise.

### GetRemoveLabelKeysOk

`func (o *NodeLabelsRequestModel) GetRemoveLabelKeysOk() (*[]string, bool)`

GetRemoveLabelKeysOk returns a tuple with the RemoveLabelKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveLabelKeys

`func (o *NodeLabelsRequestModel) SetRemoveLabelKeys(v []string)`

SetRemoveLabelKeys sets RemoveLabelKeys field to given value.

### HasRemoveLabelKeys

`func (o *NodeLabelsRequestModel) HasRemoveLabelKeys() bool`

HasRemoveLabelKeys returns a boolean if a field has been set.

### SetRemoveLabelKeysNil

`func (o *NodeLabelsRequestModel) SetRemoveLabelKeysNil(b bool)`

 SetRemoveLabelKeysNil sets the value for RemoveLabelKeys to be an explicit nil

### UnsetRemoveLabelKeys
`func (o *NodeLabelsRequestModel) UnsetRemoveLabelKeys()`

UnsetRemoveLabelKeys ensures that no value is present for RemoveLabelKeys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


