# BcsInstanceV1ApiUpdateInstanceModelInstanceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 인스턴스의 ID | 
**Name** | **string** | 인스턴스 이름 | 
**ProjectId** | **string** | 인스턴스가 속한 프로젝트의 ID | 
**UserId** | **string** | 인스턴스를 생성한 사용자 계정의 ID | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Flavor** | [**BcsInstanceV1ApiUpdateInstanceModelInstanceFlavorModel**](BcsInstanceV1ApiUpdateInstanceModelInstanceFlavorModel.md) | 인스턴스의 사양을 정의하는 정보 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**Addresses** | Pointer to **[]map[string]string** |  | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**KeyName** | Pointer to **NullableString** |  | [optional] 
**TerminatedAt** | Pointer to **NullableString** |  | [optional] 
**SecurityGroups** | Pointer to [**[]BcsInstanceV1ApiUpdateInstanceModelInstanceSecurityGroupModel**](BcsInstanceV1ApiUpdateInstanceModelInstanceSecurityGroupModel.md) |  | [optional] 
**TaskState** | Pointer to **NullableString** |  | [optional] 
**VmState** | Pointer to **NullableString** |  | [optional] 
**PowerState** | Pointer to [**NullablePowerState**](PowerState.md) |  | [optional] 
**AttachedVolumes** | Pointer to [**[]BcsInstanceV1ApiUpdateInstanceModelInstanceAttachedVolumeModel**](BcsInstanceV1ApiUpdateInstanceModelInstanceAttachedVolumeModel.md) |  | [optional] 
**Description** | **string** | 인스턴스에 대한 설명 | 
**Hostname** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBcsInstanceV1ApiUpdateInstanceModelInstanceModel

`func NewBcsInstanceV1ApiUpdateInstanceModelInstanceModel(id string, name string, projectId string, userId string, flavor BcsInstanceV1ApiUpdateInstanceModelInstanceFlavorModel, createdAt time.Time, updatedAt time.Time, description string, ) *BcsInstanceV1ApiUpdateInstanceModelInstanceModel`

NewBcsInstanceV1ApiUpdateInstanceModelInstanceModel instantiates a new BcsInstanceV1ApiUpdateInstanceModelInstanceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsInstanceV1ApiUpdateInstanceModelInstanceModelWithDefaults

`func NewBcsInstanceV1ApiUpdateInstanceModelInstanceModelWithDefaults() *BcsInstanceV1ApiUpdateInstanceModelInstanceModel`

NewBcsInstanceV1ApiUpdateInstanceModelInstanceModelWithDefaults instantiates a new BcsInstanceV1ApiUpdateInstanceModelInstanceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetUserId

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetMetadata

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetFlavor

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetFlavor() BcsInstanceV1ApiUpdateInstanceModelInstanceFlavorModel`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetFlavorOk() (*BcsInstanceV1ApiUpdateInstanceModelInstanceFlavorModel, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetFlavor(v BcsInstanceV1ApiUpdateInstanceModelInstanceFlavorModel)`

SetFlavor sets Flavor field to given value.


### GetCreatedAt

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAddresses

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetAddresses() []map[string]string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetAddressesOk() (*[]map[string]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetAddresses(v []map[string]string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### SetAddressesNil

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetAddressesNil(b bool)`

 SetAddressesNil sets the value for Addresses to be an explicit nil

### UnsetAddresses
`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) UnsetAddresses()`

UnsetAddresses ensures that no value is present for Addresses, not even an explicit nil
### GetAvailabilityZone

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetKeyName

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### SetKeyNameNil

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetTerminatedAt

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetTerminatedAt() string`

GetTerminatedAt returns the TerminatedAt field if non-nil, zero value otherwise.

### GetTerminatedAtOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetTerminatedAtOk() (*string, bool)`

GetTerminatedAtOk returns a tuple with the TerminatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedAt

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetTerminatedAt(v string)`

SetTerminatedAt sets TerminatedAt field to given value.

### HasTerminatedAt

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) HasTerminatedAt() bool`

HasTerminatedAt returns a boolean if a field has been set.

### SetTerminatedAtNil

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetTerminatedAtNil(b bool)`

 SetTerminatedAtNil sets the value for TerminatedAt to be an explicit nil

### UnsetTerminatedAt
`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) UnsetTerminatedAt()`

UnsetTerminatedAt ensures that no value is present for TerminatedAt, not even an explicit nil
### GetSecurityGroups

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetSecurityGroups() []BcsInstanceV1ApiUpdateInstanceModelInstanceSecurityGroupModel`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetSecurityGroupsOk() (*[]BcsInstanceV1ApiUpdateInstanceModelInstanceSecurityGroupModel, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetSecurityGroups(v []BcsInstanceV1ApiUpdateInstanceModelInstanceSecurityGroupModel)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetTaskState

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetTaskState() string`

GetTaskState returns the TaskState field if non-nil, zero value otherwise.

### GetTaskStateOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetTaskStateOk() (*string, bool)`

GetTaskStateOk returns a tuple with the TaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskState

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetTaskState(v string)`

SetTaskState sets TaskState field to given value.

### HasTaskState

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) HasTaskState() bool`

HasTaskState returns a boolean if a field has been set.

### SetTaskStateNil

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetTaskStateNil(b bool)`

 SetTaskStateNil sets the value for TaskState to be an explicit nil

### UnsetTaskState
`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) UnsetTaskState()`

UnsetTaskState ensures that no value is present for TaskState, not even an explicit nil
### GetVmState

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetVmState() string`

GetVmState returns the VmState field if non-nil, zero value otherwise.

### GetVmStateOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetVmStateOk() (*string, bool)`

GetVmStateOk returns a tuple with the VmState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmState

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetVmState(v string)`

SetVmState sets VmState field to given value.

### HasVmState

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) HasVmState() bool`

HasVmState returns a boolean if a field has been set.

### SetVmStateNil

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetVmStateNil(b bool)`

 SetVmStateNil sets the value for VmState to be an explicit nil

### UnsetVmState
`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) UnsetVmState()`

UnsetVmState ensures that no value is present for VmState, not even an explicit nil
### GetPowerState

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetPowerState() PowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetPowerStateOk() (*PowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetPowerState(v PowerState)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### SetPowerStateNil

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetPowerStateNil(b bool)`

 SetPowerStateNil sets the value for PowerState to be an explicit nil

### UnsetPowerState
`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) UnsetPowerState()`

UnsetPowerState ensures that no value is present for PowerState, not even an explicit nil
### GetAttachedVolumes

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetAttachedVolumes() []BcsInstanceV1ApiUpdateInstanceModelInstanceAttachedVolumeModel`

GetAttachedVolumes returns the AttachedVolumes field if non-nil, zero value otherwise.

### GetAttachedVolumesOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetAttachedVolumesOk() (*[]BcsInstanceV1ApiUpdateInstanceModelInstanceAttachedVolumeModel, bool)`

GetAttachedVolumesOk returns a tuple with the AttachedVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedVolumes

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetAttachedVolumes(v []BcsInstanceV1ApiUpdateInstanceModelInstanceAttachedVolumeModel)`

SetAttachedVolumes sets AttachedVolumes field to given value.

### HasAttachedVolumes

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) HasAttachedVolumes() bool`

HasAttachedVolumes returns a boolean if a field has been set.

### SetAttachedVolumesNil

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetAttachedVolumesNil(b bool)`

 SetAttachedVolumesNil sets the value for AttachedVolumes to be an explicit nil

### UnsetAttachedVolumes
`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) UnsetAttachedVolumes()`

UnsetAttachedVolumes ensures that no value is present for AttachedVolumes, not even an explicit nil
### GetDescription

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHostname

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *BcsInstanceV1ApiUpdateInstanceModelInstanceModel) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


