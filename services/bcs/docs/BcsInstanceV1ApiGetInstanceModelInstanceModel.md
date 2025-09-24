# BcsInstanceV1ApiGetInstanceModelInstanceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 인스턴스 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Flavor** | Pointer to [**NullableBcsInstanceV1ApiGetInstanceModelInstanceFlavorModel**](BcsInstanceV1ApiGetInstanceModelInstanceFlavorModel.md) |  | [optional] 
**Addresses** | Pointer to [**[]BcsInstanceV1ApiGetInstanceModelInstanceAddressModel**](BcsInstanceV1ApiGetInstanceModelInstanceAddressModel.md) |  | [optional] 
**IsHyperThreading** | Pointer to **NullableBool** |  | [optional] 
**IsHadoop** | Pointer to **NullableBool** |  | [optional] 
**IsK8se** | Pointer to **NullableBool** |  | [optional] 
**Image** | Pointer to [**NullableBcsInstanceV1ApiGetInstanceModelInstanceImageModel**](BcsInstanceV1ApiGetInstanceModelInstanceImageModel.md) |  | [optional] 
**VmState** | Pointer to **NullableString** |  | [optional] 
**TaskState** | Pointer to **NullableString** |  | [optional] 
**PowerState** | Pointer to [**NullablePowerState**](PowerState.md) |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**KeyName** | Pointer to **NullableString** |  | [optional] 
**Hostname** | Pointer to **NullableString** |  | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**AttachedVolumes** | Pointer to [**[]BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel**](BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel.md) |  | [optional] 
**AttachedVolumeCount** | Pointer to **NullableInt64** |  | [optional] 
**SecurityGroups** | Pointer to [**[]BcsInstanceV1ApiGetInstanceModelInstanceSecurityGroupModel**](BcsInstanceV1ApiGetInstanceModelInstanceSecurityGroupModel.md) |  | [optional] 
**SecurityGroupCount** | Pointer to **NullableInt64** |  | [optional] 
**InstanceType** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBcsInstanceV1ApiGetInstanceModelInstanceModel

`func NewBcsInstanceV1ApiGetInstanceModelInstanceModel(id string, ) *BcsInstanceV1ApiGetInstanceModelInstanceModel`

NewBcsInstanceV1ApiGetInstanceModelInstanceModel instantiates a new BcsInstanceV1ApiGetInstanceModelInstanceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsInstanceV1ApiGetInstanceModelInstanceModelWithDefaults

`func NewBcsInstanceV1ApiGetInstanceModelInstanceModelWithDefaults() *BcsInstanceV1ApiGetInstanceModelInstanceModel`

NewBcsInstanceV1ApiGetInstanceModelInstanceModelWithDefaults instantiates a new BcsInstanceV1ApiGetInstanceModelInstanceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetadata

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetFlavor

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetFlavor() BcsInstanceV1ApiGetInstanceModelInstanceFlavorModel`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetFlavorOk() (*BcsInstanceV1ApiGetInstanceModelInstanceFlavorModel, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetFlavor(v BcsInstanceV1ApiGetInstanceModelInstanceFlavorModel)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### SetFlavorNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetFlavorNil(b bool)`

 SetFlavorNil sets the value for Flavor to be an explicit nil

### UnsetFlavor
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetFlavor()`

UnsetFlavor ensures that no value is present for Flavor, not even an explicit nil
### GetAddresses

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetAddresses() []BcsInstanceV1ApiGetInstanceModelInstanceAddressModel`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetAddressesOk() (*[]BcsInstanceV1ApiGetInstanceModelInstanceAddressModel, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetAddresses(v []BcsInstanceV1ApiGetInstanceModelInstanceAddressModel)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### SetAddressesNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetAddressesNil(b bool)`

 SetAddressesNil sets the value for Addresses to be an explicit nil

### UnsetAddresses
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetAddresses()`

UnsetAddresses ensures that no value is present for Addresses, not even an explicit nil
### GetIsHyperThreading

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetIsHyperThreading() bool`

GetIsHyperThreading returns the IsHyperThreading field if non-nil, zero value otherwise.

### GetIsHyperThreadingOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetIsHyperThreadingOk() (*bool, bool)`

GetIsHyperThreadingOk returns a tuple with the IsHyperThreading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyperThreading

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetIsHyperThreading(v bool)`

SetIsHyperThreading sets IsHyperThreading field to given value.

### HasIsHyperThreading

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasIsHyperThreading() bool`

HasIsHyperThreading returns a boolean if a field has been set.

### SetIsHyperThreadingNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetIsHyperThreadingNil(b bool)`

 SetIsHyperThreadingNil sets the value for IsHyperThreading to be an explicit nil

### UnsetIsHyperThreading
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetIsHyperThreading()`

UnsetIsHyperThreading ensures that no value is present for IsHyperThreading, not even an explicit nil
### GetIsHadoop

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetIsHadoop() bool`

GetIsHadoop returns the IsHadoop field if non-nil, zero value otherwise.

### GetIsHadoopOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetIsHadoopOk() (*bool, bool)`

GetIsHadoopOk returns a tuple with the IsHadoop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHadoop

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetIsHadoop(v bool)`

SetIsHadoop sets IsHadoop field to given value.

### HasIsHadoop

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasIsHadoop() bool`

HasIsHadoop returns a boolean if a field has been set.

### SetIsHadoopNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetIsHadoopNil(b bool)`

 SetIsHadoopNil sets the value for IsHadoop to be an explicit nil

### UnsetIsHadoop
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetIsHadoop()`

UnsetIsHadoop ensures that no value is present for IsHadoop, not even an explicit nil
### GetIsK8se

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetIsK8se() bool`

GetIsK8se returns the IsK8se field if non-nil, zero value otherwise.

### GetIsK8seOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetIsK8seOk() (*bool, bool)`

GetIsK8seOk returns a tuple with the IsK8se field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsK8se

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetIsK8se(v bool)`

SetIsK8se sets IsK8se field to given value.

### HasIsK8se

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasIsK8se() bool`

HasIsK8se returns a boolean if a field has been set.

### SetIsK8seNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetIsK8seNil(b bool)`

 SetIsK8seNil sets the value for IsK8se to be an explicit nil

### UnsetIsK8se
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetIsK8se()`

UnsetIsK8se ensures that no value is present for IsK8se, not even an explicit nil
### GetImage

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetImage() BcsInstanceV1ApiGetInstanceModelInstanceImageModel`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetImageOk() (*BcsInstanceV1ApiGetInstanceModelInstanceImageModel, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetImage(v BcsInstanceV1ApiGetInstanceModelInstanceImageModel)`

SetImage sets Image field to given value.

### HasImage

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetVmState

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetVmState() string`

GetVmState returns the VmState field if non-nil, zero value otherwise.

### GetVmStateOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetVmStateOk() (*string, bool)`

GetVmStateOk returns a tuple with the VmState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmState

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetVmState(v string)`

SetVmState sets VmState field to given value.

### HasVmState

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasVmState() bool`

HasVmState returns a boolean if a field has been set.

### SetVmStateNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetVmStateNil(b bool)`

 SetVmStateNil sets the value for VmState to be an explicit nil

### UnsetVmState
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetVmState()`

UnsetVmState ensures that no value is present for VmState, not even an explicit nil
### GetTaskState

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetTaskState() string`

GetTaskState returns the TaskState field if non-nil, zero value otherwise.

### GetTaskStateOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetTaskStateOk() (*string, bool)`

GetTaskStateOk returns a tuple with the TaskState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskState

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetTaskState(v string)`

SetTaskState sets TaskState field to given value.

### HasTaskState

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasTaskState() bool`

HasTaskState returns a boolean if a field has been set.

### SetTaskStateNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetTaskStateNil(b bool)`

 SetTaskStateNil sets the value for TaskState to be an explicit nil

### UnsetTaskState
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetTaskState()`

UnsetTaskState ensures that no value is present for TaskState, not even an explicit nil
### GetPowerState

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetPowerState() PowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetPowerStateOk() (*PowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetPowerState(v PowerState)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### SetPowerStateNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetPowerStateNil(b bool)`

 SetPowerStateNil sets the value for PowerState to be an explicit nil

### UnsetPowerState
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetPowerState()`

UnsetPowerState ensures that no value is present for PowerState, not even an explicit nil
### GetStatus

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUserId

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetProjectId

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetKeyName

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### SetKeyNameNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetHostname

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetAvailabilityZone

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetAttachedVolumes

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetAttachedVolumes() []BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel`

GetAttachedVolumes returns the AttachedVolumes field if non-nil, zero value otherwise.

### GetAttachedVolumesOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetAttachedVolumesOk() (*[]BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel, bool)`

GetAttachedVolumesOk returns a tuple with the AttachedVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedVolumes

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetAttachedVolumes(v []BcsInstanceV1ApiGetInstanceModelInstanceAttachedVolumeModel)`

SetAttachedVolumes sets AttachedVolumes field to given value.

### HasAttachedVolumes

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasAttachedVolumes() bool`

HasAttachedVolumes returns a boolean if a field has been set.

### SetAttachedVolumesNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetAttachedVolumesNil(b bool)`

 SetAttachedVolumesNil sets the value for AttachedVolumes to be an explicit nil

### UnsetAttachedVolumes
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetAttachedVolumes()`

UnsetAttachedVolumes ensures that no value is present for AttachedVolumes, not even an explicit nil
### GetAttachedVolumeCount

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetAttachedVolumeCount() int64`

GetAttachedVolumeCount returns the AttachedVolumeCount field if non-nil, zero value otherwise.

### GetAttachedVolumeCountOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetAttachedVolumeCountOk() (*int64, bool)`

GetAttachedVolumeCountOk returns a tuple with the AttachedVolumeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedVolumeCount

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetAttachedVolumeCount(v int64)`

SetAttachedVolumeCount sets AttachedVolumeCount field to given value.

### HasAttachedVolumeCount

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasAttachedVolumeCount() bool`

HasAttachedVolumeCount returns a boolean if a field has been set.

### SetAttachedVolumeCountNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetAttachedVolumeCountNil(b bool)`

 SetAttachedVolumeCountNil sets the value for AttachedVolumeCount to be an explicit nil

### UnsetAttachedVolumeCount
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetAttachedVolumeCount()`

UnsetAttachedVolumeCount ensures that no value is present for AttachedVolumeCount, not even an explicit nil
### GetSecurityGroups

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetSecurityGroups() []BcsInstanceV1ApiGetInstanceModelInstanceSecurityGroupModel`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetSecurityGroupsOk() (*[]BcsInstanceV1ApiGetInstanceModelInstanceSecurityGroupModel, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetSecurityGroups(v []BcsInstanceV1ApiGetInstanceModelInstanceSecurityGroupModel)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetSecurityGroupCount

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetSecurityGroupCount() int64`

GetSecurityGroupCount returns the SecurityGroupCount field if non-nil, zero value otherwise.

### GetSecurityGroupCountOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetSecurityGroupCountOk() (*int64, bool)`

GetSecurityGroupCountOk returns a tuple with the SecurityGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupCount

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetSecurityGroupCount(v int64)`

SetSecurityGroupCount sets SecurityGroupCount field to given value.

### HasSecurityGroupCount

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasSecurityGroupCount() bool`

HasSecurityGroupCount returns a boolean if a field has been set.

### SetSecurityGroupCountNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetSecurityGroupCountNil(b bool)`

 SetSecurityGroupCountNil sets the value for SecurityGroupCount to be an explicit nil

### UnsetSecurityGroupCount
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetSecurityGroupCount()`

UnsetSecurityGroupCount ensures that no value is present for SecurityGroupCount, not even an explicit nil
### GetInstanceType

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### SetInstanceTypeNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetInstanceTypeNil(b bool)`

 SetInstanceTypeNil sets the value for InstanceType to be an explicit nil

### UnsetInstanceType
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetInstanceType()`

UnsetInstanceType ensures that no value is present for InstanceType, not even an explicit nil
### GetCreatedAt

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BcsInstanceV1ApiGetInstanceModelInstanceModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


