# BcsImageV1ApiGetImageModelImageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 이미지의 고유 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableInt64** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsShared** | Pointer to **NullableBool** |  | [optional] 
**DiskFormat** | Pointer to **NullableString** |  | [optional] 
**ContainerFormat** | Pointer to **NullableString** |  | [optional] 
**MinDisk** | Pointer to **NullableInt32** |  | [optional] 
**MinRam** | Pointer to **NullableInt32** |  | [optional] 
**VirtualSize** | Pointer to **NullableInt64** |  | [optional] 
**InstanceType** | Pointer to [**NullableInstanceType**](InstanceType.md) |  | [optional] 
**ImageMemberStatus** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**OsInfo** | Pointer to [**NullableBcsImageV1ApiGetImageModelOsInfoModel**](BcsImageV1ApiGetImageModelOsInfoModel.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBcsImageV1ApiGetImageModelImageModel

`func NewBcsImageV1ApiGetImageModelImageModel(id string, ) *BcsImageV1ApiGetImageModelImageModel`

NewBcsImageV1ApiGetImageModelImageModel instantiates a new BcsImageV1ApiGetImageModelImageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsImageV1ApiGetImageModelImageModelWithDefaults

`func NewBcsImageV1ApiGetImageModelImageModelWithDefaults() *BcsImageV1ApiGetImageModelImageModel`

NewBcsImageV1ApiGetImageModelImageModelWithDefaults instantiates a new BcsImageV1ApiGetImageModelImageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsImageV1ApiGetImageModelImageModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsImageV1ApiGetImageModelImageModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BcsImageV1ApiGetImageModelImageModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BcsImageV1ApiGetImageModelImageModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BcsImageV1ApiGetImageModelImageModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BcsImageV1ApiGetImageModelImageModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BcsImageV1ApiGetImageModelImageModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSize

`func (o *BcsImageV1ApiGetImageModelImageModel) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BcsImageV1ApiGetImageModelImageModel) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *BcsImageV1ApiGetImageModelImageModel) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *BcsImageV1ApiGetImageModelImageModel) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *BcsImageV1ApiGetImageModelImageModel) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetStatus

`func (o *BcsImageV1ApiGetImageModelImageModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BcsImageV1ApiGetImageModelImageModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BcsImageV1ApiGetImageModelImageModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BcsImageV1ApiGetImageModelImageModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BcsImageV1ApiGetImageModelImageModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetOwner

`func (o *BcsImageV1ApiGetImageModelImageModel) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BcsImageV1ApiGetImageModelImageModel) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BcsImageV1ApiGetImageModelImageModel) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BcsImageV1ApiGetImageModelImageModel) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BcsImageV1ApiGetImageModelImageModel) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetVisibility

`func (o *BcsImageV1ApiGetImageModelImageModel) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BcsImageV1ApiGetImageModelImageModel) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BcsImageV1ApiGetImageModelImageModel) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *BcsImageV1ApiGetImageModelImageModel) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *BcsImageV1ApiGetImageModelImageModel) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil
### GetDescription

`func (o *BcsImageV1ApiGetImageModelImageModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BcsImageV1ApiGetImageModelImageModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BcsImageV1ApiGetImageModelImageModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BcsImageV1ApiGetImageModelImageModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BcsImageV1ApiGetImageModelImageModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsShared

`func (o *BcsImageV1ApiGetImageModelImageModel) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *BcsImageV1ApiGetImageModelImageModel) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *BcsImageV1ApiGetImageModelImageModel) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### SetIsSharedNil

`func (o *BcsImageV1ApiGetImageModelImageModel) SetIsSharedNil(b bool)`

 SetIsSharedNil sets the value for IsShared to be an explicit nil

### UnsetIsShared
`func (o *BcsImageV1ApiGetImageModelImageModel) UnsetIsShared()`

UnsetIsShared ensures that no value is present for IsShared, not even an explicit nil
### GetDiskFormat

`func (o *BcsImageV1ApiGetImageModelImageModel) GetDiskFormat() string`

GetDiskFormat returns the DiskFormat field if non-nil, zero value otherwise.

### GetDiskFormatOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetDiskFormatOk() (*string, bool)`

GetDiskFormatOk returns a tuple with the DiskFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFormat

`func (o *BcsImageV1ApiGetImageModelImageModel) SetDiskFormat(v string)`

SetDiskFormat sets DiskFormat field to given value.

### HasDiskFormat

`func (o *BcsImageV1ApiGetImageModelImageModel) HasDiskFormat() bool`

HasDiskFormat returns a boolean if a field has been set.

### SetDiskFormatNil

`func (o *BcsImageV1ApiGetImageModelImageModel) SetDiskFormatNil(b bool)`

 SetDiskFormatNil sets the value for DiskFormat to be an explicit nil

### UnsetDiskFormat
`func (o *BcsImageV1ApiGetImageModelImageModel) UnsetDiskFormat()`

UnsetDiskFormat ensures that no value is present for DiskFormat, not even an explicit nil
### GetContainerFormat

`func (o *BcsImageV1ApiGetImageModelImageModel) GetContainerFormat() string`

GetContainerFormat returns the ContainerFormat field if non-nil, zero value otherwise.

### GetContainerFormatOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetContainerFormatOk() (*string, bool)`

GetContainerFormatOk returns a tuple with the ContainerFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerFormat

`func (o *BcsImageV1ApiGetImageModelImageModel) SetContainerFormat(v string)`

SetContainerFormat sets ContainerFormat field to given value.

### HasContainerFormat

`func (o *BcsImageV1ApiGetImageModelImageModel) HasContainerFormat() bool`

HasContainerFormat returns a boolean if a field has been set.

### SetContainerFormatNil

`func (o *BcsImageV1ApiGetImageModelImageModel) SetContainerFormatNil(b bool)`

 SetContainerFormatNil sets the value for ContainerFormat to be an explicit nil

### UnsetContainerFormat
`func (o *BcsImageV1ApiGetImageModelImageModel) UnsetContainerFormat()`

UnsetContainerFormat ensures that no value is present for ContainerFormat, not even an explicit nil
### GetMinDisk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetMinDisk() int32`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetMinDiskOk() (*int32, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *BcsImageV1ApiGetImageModelImageModel) SetMinDisk(v int32)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *BcsImageV1ApiGetImageModelImageModel) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### SetMinDiskNil

`func (o *BcsImageV1ApiGetImageModelImageModel) SetMinDiskNil(b bool)`

 SetMinDiskNil sets the value for MinDisk to be an explicit nil

### UnsetMinDisk
`func (o *BcsImageV1ApiGetImageModelImageModel) UnsetMinDisk()`

UnsetMinDisk ensures that no value is present for MinDisk, not even an explicit nil
### GetMinRam

`func (o *BcsImageV1ApiGetImageModelImageModel) GetMinRam() int32`

GetMinRam returns the MinRam field if non-nil, zero value otherwise.

### GetMinRamOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetMinRamOk() (*int32, bool)`

GetMinRamOk returns a tuple with the MinRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRam

`func (o *BcsImageV1ApiGetImageModelImageModel) SetMinRam(v int32)`

SetMinRam sets MinRam field to given value.

### HasMinRam

`func (o *BcsImageV1ApiGetImageModelImageModel) HasMinRam() bool`

HasMinRam returns a boolean if a field has been set.

### SetMinRamNil

`func (o *BcsImageV1ApiGetImageModelImageModel) SetMinRamNil(b bool)`

 SetMinRamNil sets the value for MinRam to be an explicit nil

### UnsetMinRam
`func (o *BcsImageV1ApiGetImageModelImageModel) UnsetMinRam()`

UnsetMinRam ensures that no value is present for MinRam, not even an explicit nil
### GetVirtualSize

`func (o *BcsImageV1ApiGetImageModelImageModel) GetVirtualSize() int64`

GetVirtualSize returns the VirtualSize field if non-nil, zero value otherwise.

### GetVirtualSizeOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetVirtualSizeOk() (*int64, bool)`

GetVirtualSizeOk returns a tuple with the VirtualSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSize

`func (o *BcsImageV1ApiGetImageModelImageModel) SetVirtualSize(v int64)`

SetVirtualSize sets VirtualSize field to given value.

### HasVirtualSize

`func (o *BcsImageV1ApiGetImageModelImageModel) HasVirtualSize() bool`

HasVirtualSize returns a boolean if a field has been set.

### SetVirtualSizeNil

`func (o *BcsImageV1ApiGetImageModelImageModel) SetVirtualSizeNil(b bool)`

 SetVirtualSizeNil sets the value for VirtualSize to be an explicit nil

### UnsetVirtualSize
`func (o *BcsImageV1ApiGetImageModelImageModel) UnsetVirtualSize()`

UnsetVirtualSize ensures that no value is present for VirtualSize, not even an explicit nil
### GetInstanceType

`func (o *BcsImageV1ApiGetImageModelImageModel) GetInstanceType() InstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetInstanceTypeOk() (*InstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *BcsImageV1ApiGetImageModelImageModel) SetInstanceType(v InstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *BcsImageV1ApiGetImageModelImageModel) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### SetInstanceTypeNil

`func (o *BcsImageV1ApiGetImageModelImageModel) SetInstanceTypeNil(b bool)`

 SetInstanceTypeNil sets the value for InstanceType to be an explicit nil

### UnsetInstanceType
`func (o *BcsImageV1ApiGetImageModelImageModel) UnsetInstanceType()`

UnsetInstanceType ensures that no value is present for InstanceType, not even an explicit nil
### GetImageMemberStatus

`func (o *BcsImageV1ApiGetImageModelImageModel) GetImageMemberStatus() string`

GetImageMemberStatus returns the ImageMemberStatus field if non-nil, zero value otherwise.

### GetImageMemberStatusOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetImageMemberStatusOk() (*string, bool)`

GetImageMemberStatusOk returns a tuple with the ImageMemberStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageMemberStatus

`func (o *BcsImageV1ApiGetImageModelImageModel) SetImageMemberStatus(v string)`

SetImageMemberStatus sets ImageMemberStatus field to given value.

### HasImageMemberStatus

`func (o *BcsImageV1ApiGetImageModelImageModel) HasImageMemberStatus() bool`

HasImageMemberStatus returns a boolean if a field has been set.

### SetImageMemberStatusNil

`func (o *BcsImageV1ApiGetImageModelImageModel) SetImageMemberStatusNil(b bool)`

 SetImageMemberStatusNil sets the value for ImageMemberStatus to be an explicit nil

### UnsetImageMemberStatus
`func (o *BcsImageV1ApiGetImageModelImageModel) UnsetImageMemberStatus()`

UnsetImageMemberStatus ensures that no value is present for ImageMemberStatus, not even an explicit nil
### GetProjectId

`func (o *BcsImageV1ApiGetImageModelImageModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BcsImageV1ApiGetImageModelImageModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BcsImageV1ApiGetImageModelImageModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BcsImageV1ApiGetImageModelImageModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BcsImageV1ApiGetImageModelImageModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetOsInfo

`func (o *BcsImageV1ApiGetImageModelImageModel) GetOsInfo() BcsImageV1ApiGetImageModelOsInfoModel`

GetOsInfo returns the OsInfo field if non-nil, zero value otherwise.

### GetOsInfoOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetOsInfoOk() (*BcsImageV1ApiGetImageModelOsInfoModel, bool)`

GetOsInfoOk returns a tuple with the OsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsInfo

`func (o *BcsImageV1ApiGetImageModelImageModel) SetOsInfo(v BcsImageV1ApiGetImageModelOsInfoModel)`

SetOsInfo sets OsInfo field to given value.

### HasOsInfo

`func (o *BcsImageV1ApiGetImageModelImageModel) HasOsInfo() bool`

HasOsInfo returns a boolean if a field has been set.

### SetOsInfoNil

`func (o *BcsImageV1ApiGetImageModelImageModel) SetOsInfoNil(b bool)`

 SetOsInfoNil sets the value for OsInfo to be an explicit nil

### UnsetOsInfo
`func (o *BcsImageV1ApiGetImageModelImageModel) UnsetOsInfo()`

UnsetOsInfo ensures that no value is present for OsInfo, not even an explicit nil
### GetCreatedAt

`func (o *BcsImageV1ApiGetImageModelImageModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsImageV1ApiGetImageModelImageModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BcsImageV1ApiGetImageModelImageModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BcsImageV1ApiGetImageModelImageModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BcsImageV1ApiGetImageModelImageModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BcsImageV1ApiGetImageModelImageModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BcsImageV1ApiGetImageModelImageModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BcsImageV1ApiGetImageModelImageModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BcsImageV1ApiGetImageModelImageModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BcsImageV1ApiGetImageModelImageModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BcsImageV1ApiGetImageModelImageModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


