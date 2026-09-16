# DefaultParameterGroupInstanceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | MySQL 인스턴스 그룹 ID | [optional] 
**Name** | Pointer to **NullableString** | MySQL 인스턴스 그룹 이름 | [optional] 
**Status** | Pointer to [**NullableInstanceGroupStatus**](InstanceGroupStatus.md) | MySQL 인스턴스 그룹의 현재 상태 | [optional] 
**EngineVersion** | Pointer to **NullableString** | MySQL 엔진 버전 | [optional] 
**FlavorId** | Pointer to **NullableString** | MySQL 인스턴스 그룹에 적용된 Flavor ID | [optional] 
**ParameterGroupStatus** | Pointer to [**NullableParameterGroupStatus**](ParameterGroupStatus.md) | 파라미터 그룹 적용 상태 | [optional] 
**InstanceGroupType** | Pointer to [**NullableClusterType**](ClusterType.md) | MySQL 인스턴스 그룹 유형 | [optional] 
**IsMultiAz** | Pointer to **NullableBool** | 다중 가용 영역(Multi-AZ) 구성 여부 | [optional] 

## Methods

### NewDefaultParameterGroupInstanceGroup

`func NewDefaultParameterGroupInstanceGroup() *DefaultParameterGroupInstanceGroup`

NewDefaultParameterGroupInstanceGroup instantiates a new DefaultParameterGroupInstanceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultParameterGroupInstanceGroupWithDefaults

`func NewDefaultParameterGroupInstanceGroupWithDefaults() *DefaultParameterGroupInstanceGroup`

NewDefaultParameterGroupInstanceGroupWithDefaults instantiates a new DefaultParameterGroupInstanceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DefaultParameterGroupInstanceGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DefaultParameterGroupInstanceGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DefaultParameterGroupInstanceGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DefaultParameterGroupInstanceGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DefaultParameterGroupInstanceGroup) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DefaultParameterGroupInstanceGroup) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *DefaultParameterGroupInstanceGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DefaultParameterGroupInstanceGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DefaultParameterGroupInstanceGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DefaultParameterGroupInstanceGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DefaultParameterGroupInstanceGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DefaultParameterGroupInstanceGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *DefaultParameterGroupInstanceGroup) GetStatus() InstanceGroupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DefaultParameterGroupInstanceGroup) GetStatusOk() (*InstanceGroupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DefaultParameterGroupInstanceGroup) SetStatus(v InstanceGroupStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DefaultParameterGroupInstanceGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *DefaultParameterGroupInstanceGroup) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DefaultParameterGroupInstanceGroup) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetEngineVersion

`func (o *DefaultParameterGroupInstanceGroup) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *DefaultParameterGroupInstanceGroup) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *DefaultParameterGroupInstanceGroup) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.

### HasEngineVersion

`func (o *DefaultParameterGroupInstanceGroup) HasEngineVersion() bool`

HasEngineVersion returns a boolean if a field has been set.

### SetEngineVersionNil

`func (o *DefaultParameterGroupInstanceGroup) SetEngineVersionNil(b bool)`

 SetEngineVersionNil sets the value for EngineVersion to be an explicit nil

### UnsetEngineVersion
`func (o *DefaultParameterGroupInstanceGroup) UnsetEngineVersion()`

UnsetEngineVersion ensures that no value is present for EngineVersion, not even an explicit nil
### GetFlavorId

`func (o *DefaultParameterGroupInstanceGroup) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *DefaultParameterGroupInstanceGroup) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *DefaultParameterGroupInstanceGroup) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.

### HasFlavorId

`func (o *DefaultParameterGroupInstanceGroup) HasFlavorId() bool`

HasFlavorId returns a boolean if a field has been set.

### SetFlavorIdNil

`func (o *DefaultParameterGroupInstanceGroup) SetFlavorIdNil(b bool)`

 SetFlavorIdNil sets the value for FlavorId to be an explicit nil

### UnsetFlavorId
`func (o *DefaultParameterGroupInstanceGroup) UnsetFlavorId()`

UnsetFlavorId ensures that no value is present for FlavorId, not even an explicit nil
### GetParameterGroupStatus

`func (o *DefaultParameterGroupInstanceGroup) GetParameterGroupStatus() ParameterGroupStatus`

GetParameterGroupStatus returns the ParameterGroupStatus field if non-nil, zero value otherwise.

### GetParameterGroupStatusOk

`func (o *DefaultParameterGroupInstanceGroup) GetParameterGroupStatusOk() (*ParameterGroupStatus, bool)`

GetParameterGroupStatusOk returns a tuple with the ParameterGroupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterGroupStatus

`func (o *DefaultParameterGroupInstanceGroup) SetParameterGroupStatus(v ParameterGroupStatus)`

SetParameterGroupStatus sets ParameterGroupStatus field to given value.

### HasParameterGroupStatus

`func (o *DefaultParameterGroupInstanceGroup) HasParameterGroupStatus() bool`

HasParameterGroupStatus returns a boolean if a field has been set.

### SetParameterGroupStatusNil

`func (o *DefaultParameterGroupInstanceGroup) SetParameterGroupStatusNil(b bool)`

 SetParameterGroupStatusNil sets the value for ParameterGroupStatus to be an explicit nil

### UnsetParameterGroupStatus
`func (o *DefaultParameterGroupInstanceGroup) UnsetParameterGroupStatus()`

UnsetParameterGroupStatus ensures that no value is present for ParameterGroupStatus, not even an explicit nil
### GetInstanceGroupType

`func (o *DefaultParameterGroupInstanceGroup) GetInstanceGroupType() ClusterType`

GetInstanceGroupType returns the InstanceGroupType field if non-nil, zero value otherwise.

### GetInstanceGroupTypeOk

`func (o *DefaultParameterGroupInstanceGroup) GetInstanceGroupTypeOk() (*ClusterType, bool)`

GetInstanceGroupTypeOk returns a tuple with the InstanceGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupType

`func (o *DefaultParameterGroupInstanceGroup) SetInstanceGroupType(v ClusterType)`

SetInstanceGroupType sets InstanceGroupType field to given value.

### HasInstanceGroupType

`func (o *DefaultParameterGroupInstanceGroup) HasInstanceGroupType() bool`

HasInstanceGroupType returns a boolean if a field has been set.

### SetInstanceGroupTypeNil

`func (o *DefaultParameterGroupInstanceGroup) SetInstanceGroupTypeNil(b bool)`

 SetInstanceGroupTypeNil sets the value for InstanceGroupType to be an explicit nil

### UnsetInstanceGroupType
`func (o *DefaultParameterGroupInstanceGroup) UnsetInstanceGroupType()`

UnsetInstanceGroupType ensures that no value is present for InstanceGroupType, not even an explicit nil
### GetIsMultiAz

`func (o *DefaultParameterGroupInstanceGroup) GetIsMultiAz() bool`

GetIsMultiAz returns the IsMultiAz field if non-nil, zero value otherwise.

### GetIsMultiAzOk

`func (o *DefaultParameterGroupInstanceGroup) GetIsMultiAzOk() (*bool, bool)`

GetIsMultiAzOk returns a tuple with the IsMultiAz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiAz

`func (o *DefaultParameterGroupInstanceGroup) SetIsMultiAz(v bool)`

SetIsMultiAz sets IsMultiAz field to given value.

### HasIsMultiAz

`func (o *DefaultParameterGroupInstanceGroup) HasIsMultiAz() bool`

HasIsMultiAz returns a boolean if a field has been set.

### SetIsMultiAzNil

`func (o *DefaultParameterGroupInstanceGroup) SetIsMultiAzNil(b bool)`

 SetIsMultiAzNil sets the value for IsMultiAz to be an explicit nil

### UnsetIsMultiAz
`func (o *DefaultParameterGroupInstanceGroup) UnsetIsMultiAz()`

UnsetIsMultiAz ensures that no value is present for IsMultiAz, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


