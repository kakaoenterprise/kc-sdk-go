# CreateVolumeSnapshotModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 스냅샷의 이름 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsIncremental** | **bool** | 증분 스냅샷 여부 &lt;br/&gt; - 💡 스냅샷 일정을 통해 전체 스냅샷을 보유하고 있더라도, 수동으로 스냅샷을 생성할 경우에는 최초 1번의 전체 스냅샷 생성이 필요합니다. | 

## Methods

### NewCreateVolumeSnapshotModel

`func NewCreateVolumeSnapshotModel(name string, isIncremental bool, ) *CreateVolumeSnapshotModel`

NewCreateVolumeSnapshotModel instantiates a new CreateVolumeSnapshotModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVolumeSnapshotModelWithDefaults

`func NewCreateVolumeSnapshotModelWithDefaults() *CreateVolumeSnapshotModel`

NewCreateVolumeSnapshotModelWithDefaults instantiates a new CreateVolumeSnapshotModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateVolumeSnapshotModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVolumeSnapshotModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVolumeSnapshotModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateVolumeSnapshotModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateVolumeSnapshotModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateVolumeSnapshotModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateVolumeSnapshotModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateVolumeSnapshotModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateVolumeSnapshotModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsIncremental

`func (o *CreateVolumeSnapshotModel) GetIsIncremental() bool`

GetIsIncremental returns the IsIncremental field if non-nil, zero value otherwise.

### GetIsIncrementalOk

`func (o *CreateVolumeSnapshotModel) GetIsIncrementalOk() (*bool, bool)`

GetIsIncrementalOk returns a tuple with the IsIncremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIncremental

`func (o *CreateVolumeSnapshotModel) SetIsIncremental(v bool)`

SetIsIncremental sets IsIncremental field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


