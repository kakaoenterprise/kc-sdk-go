# DefaultParameterGroupInstanceGroupResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullableInstanceGroupStatus**](InstanceGroupStatus.md) |  | [optional] 
**EngineVersion** | Pointer to **NullableString** |  | [optional] 
**FlavorId** | Pointer to **NullableString** |  | [optional] 
**ParameterGroupStatus** | Pointer to [**NullableParameterGroupStatus**](ParameterGroupStatus.md) |  | [optional] 
**InstanceGroupType** | Pointer to [**NullableClusterType**](ClusterType.md) |  | [optional] 
**IsMultiAz** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewDefaultParameterGroupInstanceGroupResponseModel

`func NewDefaultParameterGroupInstanceGroupResponseModel() *DefaultParameterGroupInstanceGroupResponseModel`

NewDefaultParameterGroupInstanceGroupResponseModel instantiates a new DefaultParameterGroupInstanceGroupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultParameterGroupInstanceGroupResponseModelWithDefaults

`func NewDefaultParameterGroupInstanceGroupResponseModelWithDefaults() *DefaultParameterGroupInstanceGroupResponseModel`

NewDefaultParameterGroupInstanceGroupResponseModelWithDefaults instantiates a new DefaultParameterGroupInstanceGroupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DefaultParameterGroupInstanceGroupResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DefaultParameterGroupInstanceGroupResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DefaultParameterGroupInstanceGroupResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DefaultParameterGroupInstanceGroupResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DefaultParameterGroupInstanceGroupResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DefaultParameterGroupInstanceGroupResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *DefaultParameterGroupInstanceGroupResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DefaultParameterGroupInstanceGroupResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DefaultParameterGroupInstanceGroupResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DefaultParameterGroupInstanceGroupResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DefaultParameterGroupInstanceGroupResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DefaultParameterGroupInstanceGroupResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *DefaultParameterGroupInstanceGroupResponseModel) GetStatus() InstanceGroupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DefaultParameterGroupInstanceGroupResponseModel) GetStatusOk() (*InstanceGroupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DefaultParameterGroupInstanceGroupResponseModel) SetStatus(v InstanceGroupStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DefaultParameterGroupInstanceGroupResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *DefaultParameterGroupInstanceGroupResponseModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DefaultParameterGroupInstanceGroupResponseModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetEngineVersion

`func (o *DefaultParameterGroupInstanceGroupResponseModel) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *DefaultParameterGroupInstanceGroupResponseModel) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *DefaultParameterGroupInstanceGroupResponseModel) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.

### HasEngineVersion

`func (o *DefaultParameterGroupInstanceGroupResponseModel) HasEngineVersion() bool`

HasEngineVersion returns a boolean if a field has been set.

### SetEngineVersionNil

`func (o *DefaultParameterGroupInstanceGroupResponseModel) SetEngineVersionNil(b bool)`

 SetEngineVersionNil sets the value for EngineVersion to be an explicit nil

### UnsetEngineVersion
`func (o *DefaultParameterGroupInstanceGroupResponseModel) UnsetEngineVersion()`

UnsetEngineVersion ensures that no value is present for EngineVersion, not even an explicit nil
### GetFlavorId

`func (o *DefaultParameterGroupInstanceGroupResponseModel) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *DefaultParameterGroupInstanceGroupResponseModel) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *DefaultParameterGroupInstanceGroupResponseModel) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.

### HasFlavorId

`func (o *DefaultParameterGroupInstanceGroupResponseModel) HasFlavorId() bool`

HasFlavorId returns a boolean if a field has been set.

### SetFlavorIdNil

`func (o *DefaultParameterGroupInstanceGroupResponseModel) SetFlavorIdNil(b bool)`

 SetFlavorIdNil sets the value for FlavorId to be an explicit nil

### UnsetFlavorId
`func (o *DefaultParameterGroupInstanceGroupResponseModel) UnsetFlavorId()`

UnsetFlavorId ensures that no value is present for FlavorId, not even an explicit nil
### GetParameterGroupStatus

`func (o *DefaultParameterGroupInstanceGroupResponseModel) GetParameterGroupStatus() ParameterGroupStatus`

GetParameterGroupStatus returns the ParameterGroupStatus field if non-nil, zero value otherwise.

### GetParameterGroupStatusOk

`func (o *DefaultParameterGroupInstanceGroupResponseModel) GetParameterGroupStatusOk() (*ParameterGroupStatus, bool)`

GetParameterGroupStatusOk returns a tuple with the ParameterGroupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterGroupStatus

`func (o *DefaultParameterGroupInstanceGroupResponseModel) SetParameterGroupStatus(v ParameterGroupStatus)`

SetParameterGroupStatus sets ParameterGroupStatus field to given value.

### HasParameterGroupStatus

`func (o *DefaultParameterGroupInstanceGroupResponseModel) HasParameterGroupStatus() bool`

HasParameterGroupStatus returns a boolean if a field has been set.

### SetParameterGroupStatusNil

`func (o *DefaultParameterGroupInstanceGroupResponseModel) SetParameterGroupStatusNil(b bool)`

 SetParameterGroupStatusNil sets the value for ParameterGroupStatus to be an explicit nil

### UnsetParameterGroupStatus
`func (o *DefaultParameterGroupInstanceGroupResponseModel) UnsetParameterGroupStatus()`

UnsetParameterGroupStatus ensures that no value is present for ParameterGroupStatus, not even an explicit nil
### GetInstanceGroupType

`func (o *DefaultParameterGroupInstanceGroupResponseModel) GetInstanceGroupType() ClusterType`

GetInstanceGroupType returns the InstanceGroupType field if non-nil, zero value otherwise.

### GetInstanceGroupTypeOk

`func (o *DefaultParameterGroupInstanceGroupResponseModel) GetInstanceGroupTypeOk() (*ClusterType, bool)`

GetInstanceGroupTypeOk returns a tuple with the InstanceGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupType

`func (o *DefaultParameterGroupInstanceGroupResponseModel) SetInstanceGroupType(v ClusterType)`

SetInstanceGroupType sets InstanceGroupType field to given value.

### HasInstanceGroupType

`func (o *DefaultParameterGroupInstanceGroupResponseModel) HasInstanceGroupType() bool`

HasInstanceGroupType returns a boolean if a field has been set.

### SetInstanceGroupTypeNil

`func (o *DefaultParameterGroupInstanceGroupResponseModel) SetInstanceGroupTypeNil(b bool)`

 SetInstanceGroupTypeNil sets the value for InstanceGroupType to be an explicit nil

### UnsetInstanceGroupType
`func (o *DefaultParameterGroupInstanceGroupResponseModel) UnsetInstanceGroupType()`

UnsetInstanceGroupType ensures that no value is present for InstanceGroupType, not even an explicit nil
### GetIsMultiAz

`func (o *DefaultParameterGroupInstanceGroupResponseModel) GetIsMultiAz() bool`

GetIsMultiAz returns the IsMultiAz field if non-nil, zero value otherwise.

### GetIsMultiAzOk

`func (o *DefaultParameterGroupInstanceGroupResponseModel) GetIsMultiAzOk() (*bool, bool)`

GetIsMultiAzOk returns a tuple with the IsMultiAz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiAz

`func (o *DefaultParameterGroupInstanceGroupResponseModel) SetIsMultiAz(v bool)`

SetIsMultiAz sets IsMultiAz field to given value.

### HasIsMultiAz

`func (o *DefaultParameterGroupInstanceGroupResponseModel) HasIsMultiAz() bool`

HasIsMultiAz returns a boolean if a field has been set.

### SetIsMultiAzNil

`func (o *DefaultParameterGroupInstanceGroupResponseModel) SetIsMultiAzNil(b bool)`

 SetIsMultiAzNil sets the value for IsMultiAz to be an explicit nil

### UnsetIsMultiAz
`func (o *DefaultParameterGroupInstanceGroupResponseModel) UnsetIsMultiAz()`

UnsetIsMultiAz ensures that no value is present for IsMultiAz, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


