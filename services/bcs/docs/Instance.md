# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 인스턴스의 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Flavor** | Pointer to [**NullableInstanceFlavor**](InstanceFlavor.md) |  | [optional] 
**Addresses** | Pointer to [**[]InstanceAddress**](InstanceAddress.md) |  | [optional] 
**IsHyperThreading** | Pointer to **NullableBool** |  | [optional] 
**IsHadoop** | Pointer to **NullableBool** |  | [optional] 
**IsK8se** | Pointer to **NullableBool** |  | [optional] 
**Image** | Pointer to [**NullableImage**](Image.md) |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**KeyName** | Pointer to **NullableString** |  | [optional] 
**VmState** | Pointer to **NullableString** |  | [optional] 
**PowerState** | Pointer to [**NullablePowerState**](PowerState.md) |  | [optional] 
**Hostname** | Pointer to **NullableString** |  | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**TaskState** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**AttachedVolumes** | Pointer to [**[]AttachedVolume**](AttachedVolume.md) |  | [optional] 
**AttachedVolumeCount** | Pointer to **NullableInt64** |  | [optional] 
**SecurityGroups** | Pointer to [**[]InstanceSg**](InstanceSg.md) |  | [optional] 
**SecurityGroupCount** | Pointer to **NullableInt64** |  | [optional] 
**InstanceType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInstance

`func NewInstance(id string, ) *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Instance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Instance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Instance) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Instance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Instance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Instance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Instance) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Instance) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Instance) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *Instance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Instance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Instance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Instance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Instance) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Instance) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFlavor

`func (o *Instance) GetFlavor() InstanceFlavor`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *Instance) GetFlavorOk() (*InstanceFlavor, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *Instance) SetFlavor(v InstanceFlavor)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *Instance) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### SetFlavorNil

`func (o *Instance) SetFlavorNil(b bool)`

 SetFlavorNil sets the value for Flavor to be an explicit nil

### UnsetFlavor
`func (o *Instance) UnsetFlavor()`

UnsetFlavor ensures that no value is present for Flavor, not even an explicit nil
### GetAddresses

`func (o *Instance) GetAddresses() []InstanceAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Instance) GetAddressesOk() (*[]InstanceAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Instance) SetAddresses(v []InstanceAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *Instance) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### SetAddressesNil

`func (o *Instance) SetAddressesNil(b bool)`

 SetAddressesNil sets the value for Addresses to be an explicit nil

### UnsetAddresses
`func (o *Instance) UnsetAddresses()`

UnsetAddresses ensures that no value is present for Addresses, not even an explicit nil
### GetIsHyperThreading

`func (o *Instance) GetIsHyperThreading() bool`

GetIsHyperThreading returns the IsHyperThreading field if non-nil, zero value otherwise.

### GetIsHyperThreadingOk

`func (o *Instance) GetIsHyperThreadingOk() (*bool, bool)`

GetIsHyperThreadingOk returns a tuple with the IsHyperThreading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyperThreading

`func (o *Instance) SetIsHyperThreading(v bool)`

SetIsHyperThreading sets IsHyperThreading field to given value.

### HasIsHyperThreading

`func (o *Instance) HasIsHyperThreading() bool`

HasIsHyperThreading returns a boolean if a field has been set.

### SetIsHyperThreadingNil

`func (o *Instance) SetIsHyperThreadingNil(b bool)`

 SetIsHyperThreadingNil sets the value for IsHyperThreading to be an explicit nil

### UnsetIsHyperThreading
`func (o *Instance) UnsetIsHyperThreading()`

UnsetIsHyperThreading ensures that no value is present for IsHyperThreading, not even an explicit nil
### GetIsHadoop

`func (o *Instance) GetIsHadoop() bool`

GetIsHadoop returns the IsHadoop field if non-nil, zero value otherwise.

### GetIsHadoopOk

`func (o *Instance) GetIsHadoopOk() (*bool, bool)`

GetIsHadoopOk returns a tuple with the IsHadoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHadoop

`func (o *Instance) SetIsHadoop(v bool)`

SetIsHadoop sets IsHadoop field to given value.

### HasIsHadoop

`func (o *Instance) HasIsHadoop() bool`

HasIsHadoop returns a boolean if a field has been set.

### SetIsHadoopNil

`func (o *Instance) SetIsHadoopNil(b bool)`

 SetIsHadoopNil sets the value for IsHadoop to be an explicit nil

### UnsetIsHadoop
`func (o *Instance) UnsetIsHadoop()`

UnsetIsHadoop ensures that no value is present for IsHadoop, not even an explicit nil
### GetIsK8se

`func (o *Instance) GetIsK8se() bool`

GetIsK8se returns the IsK8se field if non-nil, zero value otherwise.

### GetIsK8seOk

`func (o *Instance) GetIsK8seOk() (*bool, bool)`

GetIsK8seOk returns a tuple with the IsK8se field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsK8se

`func (o *Instance) SetIsK8se(v bool)`

SetIsK8se sets IsK8se field to given value.

### HasIsK8se

`func (o *Instance) HasIsK8se() bool`

HasIsK8se returns a boolean if a field has been set.

### SetIsK8seNil

`func (o *Instance) SetIsK8seNil(b bool)`

 SetIsK8seNil sets the value for IsK8se to be an explicit nil

### UnsetIsK8se
`func (o *Instance) UnsetIsK8se()`

UnsetIsK8se ensures that no value is present for IsK8se, not even an explicit nil
### GetImage

`func (o *Instance) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Instance) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Instance) SetImage(v Image)`

SetImage sets Image field to given value.

### HasImage

`func (o *Instance) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *Instance) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *Instance) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetStatus

`func (o *Instance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Instance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Instance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Instance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Instance) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Instance) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUserId

`func (o *Instance) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Instance) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Instance) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Instance) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *Instance) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *Instance) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetProjectId

`func (o *Instance) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Instance) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Instance) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Instance) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *Instance) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *Instance) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetKeyName

`func (o *Instance) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *Instance) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *Instance) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *Instance) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### SetKeyNameNil

`func (o *Instance) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *Instance) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetVmState

`func (o *Instance) GetVmState() string`

GetVmState returns the VmState field if non-nil, zero value otherwise.

### GetVmStateOk

`func (o *Instance) GetVmStateOk() (*string, bool)`

GetVmStateOk returns a tuple with the VmState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmState

`func (o *Instance) SetVmState(v string)`

SetVmState sets VmState field to given value.

### HasVmState

`func (o *Instance) HasVmState() bool`

HasVmState returns a boolean if a field has been set.

### SetVmStateNil

`func (o *Instance) SetVmStateNil(b bool)`

 SetVmStateNil sets the value for VmState to be an explicit nil

### UnsetVmState
`func (o *Instance) UnsetVmState()`

UnsetVmState ensures that no value is present for VmState, not even an explicit nil
### GetPowerState

`func (o *Instance) GetPowerState() PowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *Instance) GetPowerStateOk() (*PowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *Instance) SetPowerState(v PowerState)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *Instance) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### SetPowerStateNil

`func (o *Instance) SetPowerStateNil(b bool)`

 SetPowerStateNil sets the value for PowerState to be an explicit nil

### UnsetPowerState
`func (o *Instance) UnsetPowerState()`

UnsetPowerState ensures that no value is present for PowerState, not even an explicit nil
### GetHostname

`func (o *Instance) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Instance) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Instance) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Instance) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *Instance) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *Instance) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetAvailabilityZone

`func (o *Instance) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *Instance) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *Instance) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *Instance) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *Instance) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *Instance) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetTaskState

`func (o *Instance) GetTaskState() string`

GetTaskState returns the TaskState field if non-nil, zero value otherwise.

### GetTaskStateOk

`func (o *Instance) GetTaskStateOk() (*string, bool)`

GetTaskStateOk returns a tuple with the TaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskState

`func (o *Instance) SetTaskState(v string)`

SetTaskState sets TaskState field to given value.

### HasTaskState

`func (o *Instance) HasTaskState() bool`

HasTaskState returns a boolean if a field has been set.

### SetTaskStateNil

`func (o *Instance) SetTaskStateNil(b bool)`

 SetTaskStateNil sets the value for TaskState to be an explicit nil

### UnsetTaskState
`func (o *Instance) UnsetTaskState()`

UnsetTaskState ensures that no value is present for TaskState, not even an explicit nil
### GetCreatedAt

`func (o *Instance) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Instance) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Instance) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Instance) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Instance) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Instance) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Instance) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Instance) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Instance) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Instance) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Instance) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Instance) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetAttachedVolumes

`func (o *Instance) GetAttachedVolumes() []AttachedVolume`

GetAttachedVolumes returns the AttachedVolumes field if non-nil, zero value otherwise.

### GetAttachedVolumesOk

`func (o *Instance) GetAttachedVolumesOk() (*[]AttachedVolume, bool)`

GetAttachedVolumesOk returns a tuple with the AttachedVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedVolumes

`func (o *Instance) SetAttachedVolumes(v []AttachedVolume)`

SetAttachedVolumes sets AttachedVolumes field to given value.

### HasAttachedVolumes

`func (o *Instance) HasAttachedVolumes() bool`

HasAttachedVolumes returns a boolean if a field has been set.

### SetAttachedVolumesNil

`func (o *Instance) SetAttachedVolumesNil(b bool)`

 SetAttachedVolumesNil sets the value for AttachedVolumes to be an explicit nil

### UnsetAttachedVolumes
`func (o *Instance) UnsetAttachedVolumes()`

UnsetAttachedVolumes ensures that no value is present for AttachedVolumes, not even an explicit nil
### GetAttachedVolumeCount

`func (o *Instance) GetAttachedVolumeCount() int64`

GetAttachedVolumeCount returns the AttachedVolumeCount field if non-nil, zero value otherwise.

### GetAttachedVolumeCountOk

`func (o *Instance) GetAttachedVolumeCountOk() (*int64, bool)`

GetAttachedVolumeCountOk returns a tuple with the AttachedVolumeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedVolumeCount

`func (o *Instance) SetAttachedVolumeCount(v int64)`

SetAttachedVolumeCount sets AttachedVolumeCount field to given value.

### HasAttachedVolumeCount

`func (o *Instance) HasAttachedVolumeCount() bool`

HasAttachedVolumeCount returns a boolean if a field has been set.

### SetAttachedVolumeCountNil

`func (o *Instance) SetAttachedVolumeCountNil(b bool)`

 SetAttachedVolumeCountNil sets the value for AttachedVolumeCount to be an explicit nil

### UnsetAttachedVolumeCount
`func (o *Instance) UnsetAttachedVolumeCount()`

UnsetAttachedVolumeCount ensures that no value is present for AttachedVolumeCount, not even an explicit nil
### GetSecurityGroups

`func (o *Instance) GetSecurityGroups() []InstanceSg`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *Instance) GetSecurityGroupsOk() (*[]InstanceSg, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *Instance) SetSecurityGroups(v []InstanceSg)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *Instance) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *Instance) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *Instance) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetSecurityGroupCount

`func (o *Instance) GetSecurityGroupCount() int64`

GetSecurityGroupCount returns the SecurityGroupCount field if non-nil, zero value otherwise.

### GetSecurityGroupCountOk

`func (o *Instance) GetSecurityGroupCountOk() (*int64, bool)`

GetSecurityGroupCountOk returns a tuple with the SecurityGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupCount

`func (o *Instance) SetSecurityGroupCount(v int64)`

SetSecurityGroupCount sets SecurityGroupCount field to given value.

### HasSecurityGroupCount

`func (o *Instance) HasSecurityGroupCount() bool`

HasSecurityGroupCount returns a boolean if a field has been set.

### SetSecurityGroupCountNil

`func (o *Instance) SetSecurityGroupCountNil(b bool)`

 SetSecurityGroupCountNil sets the value for SecurityGroupCount to be an explicit nil

### UnsetSecurityGroupCount
`func (o *Instance) UnsetSecurityGroupCount()`

UnsetSecurityGroupCount ensures that no value is present for SecurityGroupCount, not even an explicit nil
### GetInstanceType

`func (o *Instance) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *Instance) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *Instance) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *Instance) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### SetInstanceTypeNil

`func (o *Instance) SetInstanceTypeNil(b bool)`

 SetInstanceTypeNil sets the value for InstanceType to be an explicit nil

### UnsetInstanceType
`func (o *Instance) UnsetInstanceType()`

UnsetInstanceType ensures that no value is present for InstanceType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


