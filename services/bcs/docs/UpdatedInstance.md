# UpdatedInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 인스턴스의 ID | 
**Name** | **string** | 인스턴스 이름 | 
**ProjectId** | **string** | 인스턴스가 속한 프로젝트의 ID | 
**UserId** | **string** | 인스턴스를 생성한 사용자 계정의 ID | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Flavor** | [**UpdatedFlavor**](UpdatedFlavor.md) | 인스턴스의 사양을 정의하는 정보 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**KeyName** | Pointer to **NullableString** |  | [optional] 
**TerminatedAt** | Pointer to **NullableTime** |  | [optional] 
**SecurityGroups** | Pointer to [**[]SecurityGroup**](SecurityGroup.md) |  | [optional] 
**TaskState** | Pointer to **NullableString** |  | [optional] 
**VmState** | Pointer to **NullableString** |  | [optional] 
**PowerState** | Pointer to [**NullablePowerState**](PowerState.md) |  | [optional] 
**AttachedVolumes** | Pointer to [**[]UpdatedAttachedVolume**](UpdatedAttachedVolume.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Hostname** | Pointer to **NullableString** |  | [optional] 
**Addresses** | Pointer to [**[]InstanceAddress**](InstanceAddress.md) |  | [optional] 

## Methods

### NewUpdatedInstance

`func NewUpdatedInstance(id string, name string, projectId string, userId string, flavor UpdatedFlavor, createdAt time.Time, updatedAt time.Time, ) *UpdatedInstance`

NewUpdatedInstance instantiates a new UpdatedInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatedInstanceWithDefaults

`func NewUpdatedInstanceWithDefaults() *UpdatedInstance`

NewUpdatedInstanceWithDefaults instantiates a new UpdatedInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdatedInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatedInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatedInstance) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdatedInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatedInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatedInstance) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *UpdatedInstance) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdatedInstance) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdatedInstance) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetUserId

`func (o *UpdatedInstance) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdatedInstance) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdatedInstance) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetMetadata

`func (o *UpdatedInstance) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdatedInstance) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdatedInstance) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdatedInstance) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *UpdatedInstance) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *UpdatedInstance) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetFlavor

`func (o *UpdatedInstance) GetFlavor() UpdatedFlavor`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *UpdatedInstance) GetFlavorOk() (*UpdatedFlavor, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *UpdatedInstance) SetFlavor(v UpdatedFlavor)`

SetFlavor sets Flavor field to given value.


### GetCreatedAt

`func (o *UpdatedInstance) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdatedInstance) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdatedInstance) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UpdatedInstance) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdatedInstance) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdatedInstance) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAvailabilityZone

`func (o *UpdatedInstance) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *UpdatedInstance) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *UpdatedInstance) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *UpdatedInstance) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *UpdatedInstance) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *UpdatedInstance) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetKeyName

`func (o *UpdatedInstance) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *UpdatedInstance) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *UpdatedInstance) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *UpdatedInstance) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### SetKeyNameNil

`func (o *UpdatedInstance) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *UpdatedInstance) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetTerminatedAt

`func (o *UpdatedInstance) GetTerminatedAt() time.Time`

GetTerminatedAt returns the TerminatedAt field if non-nil, zero value otherwise.

### GetTerminatedAtOk

`func (o *UpdatedInstance) GetTerminatedAtOk() (*time.Time, bool)`

GetTerminatedAtOk returns a tuple with the TerminatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedAt

`func (o *UpdatedInstance) SetTerminatedAt(v time.Time)`

SetTerminatedAt sets TerminatedAt field to given value.

### HasTerminatedAt

`func (o *UpdatedInstance) HasTerminatedAt() bool`

HasTerminatedAt returns a boolean if a field has been set.

### SetTerminatedAtNil

`func (o *UpdatedInstance) SetTerminatedAtNil(b bool)`

 SetTerminatedAtNil sets the value for TerminatedAt to be an explicit nil

### UnsetTerminatedAt
`func (o *UpdatedInstance) UnsetTerminatedAt()`

UnsetTerminatedAt ensures that no value is present for TerminatedAt, not even an explicit nil
### GetSecurityGroups

`func (o *UpdatedInstance) GetSecurityGroups() []SecurityGroup`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *UpdatedInstance) GetSecurityGroupsOk() (*[]SecurityGroup, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *UpdatedInstance) SetSecurityGroups(v []SecurityGroup)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *UpdatedInstance) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *UpdatedInstance) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *UpdatedInstance) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetTaskState

`func (o *UpdatedInstance) GetTaskState() string`

GetTaskState returns the TaskState field if non-nil, zero value otherwise.

### GetTaskStateOk

`func (o *UpdatedInstance) GetTaskStateOk() (*string, bool)`

GetTaskStateOk returns a tuple with the TaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskState

`func (o *UpdatedInstance) SetTaskState(v string)`

SetTaskState sets TaskState field to given value.

### HasTaskState

`func (o *UpdatedInstance) HasTaskState() bool`

HasTaskState returns a boolean if a field has been set.

### SetTaskStateNil

`func (o *UpdatedInstance) SetTaskStateNil(b bool)`

 SetTaskStateNil sets the value for TaskState to be an explicit nil

### UnsetTaskState
`func (o *UpdatedInstance) UnsetTaskState()`

UnsetTaskState ensures that no value is present for TaskState, not even an explicit nil
### GetVmState

`func (o *UpdatedInstance) GetVmState() string`

GetVmState returns the VmState field if non-nil, zero value otherwise.

### GetVmStateOk

`func (o *UpdatedInstance) GetVmStateOk() (*string, bool)`

GetVmStateOk returns a tuple with the VmState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmState

`func (o *UpdatedInstance) SetVmState(v string)`

SetVmState sets VmState field to given value.

### HasVmState

`func (o *UpdatedInstance) HasVmState() bool`

HasVmState returns a boolean if a field has been set.

### SetVmStateNil

`func (o *UpdatedInstance) SetVmStateNil(b bool)`

 SetVmStateNil sets the value for VmState to be an explicit nil

### UnsetVmState
`func (o *UpdatedInstance) UnsetVmState()`

UnsetVmState ensures that no value is present for VmState, not even an explicit nil
### GetPowerState

`func (o *UpdatedInstance) GetPowerState() PowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *UpdatedInstance) GetPowerStateOk() (*PowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *UpdatedInstance) SetPowerState(v PowerState)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *UpdatedInstance) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### SetPowerStateNil

`func (o *UpdatedInstance) SetPowerStateNil(b bool)`

 SetPowerStateNil sets the value for PowerState to be an explicit nil

### UnsetPowerState
`func (o *UpdatedInstance) UnsetPowerState()`

UnsetPowerState ensures that no value is present for PowerState, not even an explicit nil
### GetAttachedVolumes

`func (o *UpdatedInstance) GetAttachedVolumes() []UpdatedAttachedVolume`

GetAttachedVolumes returns the AttachedVolumes field if non-nil, zero value otherwise.

### GetAttachedVolumesOk

`func (o *UpdatedInstance) GetAttachedVolumesOk() (*[]UpdatedAttachedVolume, bool)`

GetAttachedVolumesOk returns a tuple with the AttachedVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedVolumes

`func (o *UpdatedInstance) SetAttachedVolumes(v []UpdatedAttachedVolume)`

SetAttachedVolumes sets AttachedVolumes field to given value.

### HasAttachedVolumes

`func (o *UpdatedInstance) HasAttachedVolumes() bool`

HasAttachedVolumes returns a boolean if a field has been set.

### SetAttachedVolumesNil

`func (o *UpdatedInstance) SetAttachedVolumesNil(b bool)`

 SetAttachedVolumesNil sets the value for AttachedVolumes to be an explicit nil

### UnsetAttachedVolumes
`func (o *UpdatedInstance) UnsetAttachedVolumes()`

UnsetAttachedVolumes ensures that no value is present for AttachedVolumes, not even an explicit nil
### GetDescription

`func (o *UpdatedInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatedInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatedInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatedInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdatedInstance) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdatedInstance) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHostname

`func (o *UpdatedInstance) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *UpdatedInstance) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *UpdatedInstance) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *UpdatedInstance) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *UpdatedInstance) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *UpdatedInstance) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetAddresses

`func (o *UpdatedInstance) GetAddresses() []InstanceAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *UpdatedInstance) GetAddressesOk() (*[]InstanceAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *UpdatedInstance) SetAddresses(v []InstanceAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *UpdatedInstance) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### SetAddressesNil

`func (o *UpdatedInstance) SetAddressesNil(b bool)`

 SetAddressesNil sets the value for Addresses to be an explicit nil

### UnsetAddresses
`func (o *UpdatedInstance) UnsetAddresses()`

UnsetAddresses ensures that no value is present for Addresses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


