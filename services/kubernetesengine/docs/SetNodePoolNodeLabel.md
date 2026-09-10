# SetNodePoolNodeLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddOrUpdateLabels** | Pointer to [**[]LabelRequest**](LabelRequest.md) |  | [optional] 
**RemoveLabelKeys** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSetNodePoolNodeLabel

`func NewSetNodePoolNodeLabel() *SetNodePoolNodeLabel`

NewSetNodePoolNodeLabel instantiates a new SetNodePoolNodeLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetNodePoolNodeLabelWithDefaults

`func NewSetNodePoolNodeLabelWithDefaults() *SetNodePoolNodeLabel`

NewSetNodePoolNodeLabelWithDefaults instantiates a new SetNodePoolNodeLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddOrUpdateLabels

`func (o *SetNodePoolNodeLabel) GetAddOrUpdateLabels() []LabelRequest`

GetAddOrUpdateLabels returns the AddOrUpdateLabels field if non-nil, zero value otherwise.

### GetAddOrUpdateLabelsOk

`func (o *SetNodePoolNodeLabel) GetAddOrUpdateLabelsOk() (*[]LabelRequest, bool)`

GetAddOrUpdateLabelsOk returns a tuple with the AddOrUpdateLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOrUpdateLabels

`func (o *SetNodePoolNodeLabel) SetAddOrUpdateLabels(v []LabelRequest)`

SetAddOrUpdateLabels sets AddOrUpdateLabels field to given value.

### HasAddOrUpdateLabels

`func (o *SetNodePoolNodeLabel) HasAddOrUpdateLabels() bool`

HasAddOrUpdateLabels returns a boolean if a field has been set.

### SetAddOrUpdateLabelsNil

`func (o *SetNodePoolNodeLabel) SetAddOrUpdateLabelsNil(b bool)`

 SetAddOrUpdateLabelsNil sets the value for AddOrUpdateLabels to be an explicit nil

### UnsetAddOrUpdateLabels
`func (o *SetNodePoolNodeLabel) UnsetAddOrUpdateLabels()`

UnsetAddOrUpdateLabels ensures that no value is present for AddOrUpdateLabels, not even an explicit nil
### GetRemoveLabelKeys

`func (o *SetNodePoolNodeLabel) GetRemoveLabelKeys() []string`

GetRemoveLabelKeys returns the RemoveLabelKeys field if non-nil, zero value otherwise.

### GetRemoveLabelKeysOk

`func (o *SetNodePoolNodeLabel) GetRemoveLabelKeysOk() (*[]string, bool)`

GetRemoveLabelKeysOk returns a tuple with the RemoveLabelKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveLabelKeys

`func (o *SetNodePoolNodeLabel) SetRemoveLabelKeys(v []string)`

SetRemoveLabelKeys sets RemoveLabelKeys field to given value.

### HasRemoveLabelKeys

`func (o *SetNodePoolNodeLabel) HasRemoveLabelKeys() bool`

HasRemoveLabelKeys returns a boolean if a field has been set.

### SetRemoveLabelKeysNil

`func (o *SetNodePoolNodeLabel) SetRemoveLabelKeysNil(b bool)`

 SetRemoveLabelKeysNil sets the value for RemoveLabelKeys to be an explicit nil

### UnsetRemoveLabelKeys
`func (o *SetNodePoolNodeLabel) UnsetRemoveLabelKeys()`

UnsetRemoveLabelKeys ensures that no value is present for RemoveLabelKeys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


