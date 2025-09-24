# BcsImageV1ApiListImagesModelImageModel

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
**OsInfo** | Pointer to [**NullableBcsImageV1ApiListImagesModelOsInfoModel**](BcsImageV1ApiListImagesModelOsInfoModel.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBcsImageV1ApiListImagesModelImageModel

`func NewBcsImageV1ApiListImagesModelImageModel(id string, ) *BcsImageV1ApiListImagesModelImageModel`

NewBcsImageV1ApiListImagesModelImageModel instantiates a new BcsImageV1ApiListImagesModelImageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsImageV1ApiListImagesModelImageModelWithDefaults

`func NewBcsImageV1ApiListImagesModelImageModelWithDefaults() *BcsImageV1ApiListImagesModelImageModel`

NewBcsImageV1ApiListImagesModelImageModelWithDefaults instantiates a new BcsImageV1ApiListImagesModelImageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsImageV1ApiListImagesModelImageModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsImageV1ApiListImagesModelImageModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BcsImageV1ApiListImagesModelImageModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BcsImageV1ApiListImagesModelImageModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BcsImageV1ApiListImagesModelImageModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BcsImageV1ApiListImagesModelImageModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BcsImageV1ApiListImagesModelImageModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSize

`func (o *BcsImageV1ApiListImagesModelImageModel) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BcsImageV1ApiListImagesModelImageModel) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *BcsImageV1ApiListImagesModelImageModel) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *BcsImageV1ApiListImagesModelImageModel) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *BcsImageV1ApiListImagesModelImageModel) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetStatus

`func (o *BcsImageV1ApiListImagesModelImageModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BcsImageV1ApiListImagesModelImageModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BcsImageV1ApiListImagesModelImageModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BcsImageV1ApiListImagesModelImageModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BcsImageV1ApiListImagesModelImageModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetOwner

`func (o *BcsImageV1ApiListImagesModelImageModel) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BcsImageV1ApiListImagesModelImageModel) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BcsImageV1ApiListImagesModelImageModel) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *BcsImageV1ApiListImagesModelImageModel) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *BcsImageV1ApiListImagesModelImageModel) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetVisibility

`func (o *BcsImageV1ApiListImagesModelImageModel) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BcsImageV1ApiListImagesModelImageModel) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BcsImageV1ApiListImagesModelImageModel) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *BcsImageV1ApiListImagesModelImageModel) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *BcsImageV1ApiListImagesModelImageModel) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil
### GetDescription

`func (o *BcsImageV1ApiListImagesModelImageModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BcsImageV1ApiListImagesModelImageModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BcsImageV1ApiListImagesModelImageModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BcsImageV1ApiListImagesModelImageModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BcsImageV1ApiListImagesModelImageModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsShared

`func (o *BcsImageV1ApiListImagesModelImageModel) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *BcsImageV1ApiListImagesModelImageModel) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *BcsImageV1ApiListImagesModelImageModel) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### SetIsSharedNil

`func (o *BcsImageV1ApiListImagesModelImageModel) SetIsSharedNil(b bool)`

 SetIsSharedNil sets the value for IsShared to be an explicit nil

### UnsetIsShared
`func (o *BcsImageV1ApiListImagesModelImageModel) UnsetIsShared()`

UnsetIsShared ensures that no value is present for IsShared, not even an explicit nil
### GetDiskFormat

`func (o *BcsImageV1ApiListImagesModelImageModel) GetDiskFormat() string`

GetDiskFormat returns the DiskFormat field if non-nil, zero value otherwise.

### GetDiskFormatOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetDiskFormatOk() (*string, bool)`

GetDiskFormatOk returns a tuple with the DiskFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFormat

`func (o *BcsImageV1ApiListImagesModelImageModel) SetDiskFormat(v string)`

SetDiskFormat sets DiskFormat field to given value.

### HasDiskFormat

`func (o *BcsImageV1ApiListImagesModelImageModel) HasDiskFormat() bool`

HasDiskFormat returns a boolean if a field has been set.

### SetDiskFormatNil

`func (o *BcsImageV1ApiListImagesModelImageModel) SetDiskFormatNil(b bool)`

 SetDiskFormatNil sets the value for DiskFormat to be an explicit nil

### UnsetDiskFormat
`func (o *BcsImageV1ApiListImagesModelImageModel) UnsetDiskFormat()`

UnsetDiskFormat ensures that no value is present for DiskFormat, not even an explicit nil
### GetContainerFormat

`func (o *BcsImageV1ApiListImagesModelImageModel) GetContainerFormat() string`

GetContainerFormat returns the ContainerFormat field if non-nil, zero value otherwise.

### GetContainerFormatOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetContainerFormatOk() (*string, bool)`

GetContainerFormatOk returns a tuple with the ContainerFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerFormat

`func (o *BcsImageV1ApiListImagesModelImageModel) SetContainerFormat(v string)`

SetContainerFormat sets ContainerFormat field to given value.

### HasContainerFormat

`func (o *BcsImageV1ApiListImagesModelImageModel) HasContainerFormat() bool`

HasContainerFormat returns a boolean if a field has been set.

### SetContainerFormatNil

`func (o *BcsImageV1ApiListImagesModelImageModel) SetContainerFormatNil(b bool)`

 SetContainerFormatNil sets the value for ContainerFormat to be an explicit nil

### UnsetContainerFormat
`func (o *BcsImageV1ApiListImagesModelImageModel) UnsetContainerFormat()`

UnsetContainerFormat ensures that no value is present for ContainerFormat, not even an explicit nil
### GetMinDisk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetMinDisk() int32`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetMinDiskOk() (*int32, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *BcsImageV1ApiListImagesModelImageModel) SetMinDisk(v int32)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *BcsImageV1ApiListImagesModelImageModel) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### SetMinDiskNil

`func (o *BcsImageV1ApiListImagesModelImageModel) SetMinDiskNil(b bool)`

 SetMinDiskNil sets the value for MinDisk to be an explicit nil

### UnsetMinDisk
`func (o *BcsImageV1ApiListImagesModelImageModel) UnsetMinDisk()`

UnsetMinDisk ensures that no value is present for MinDisk, not even an explicit nil
### GetMinRam

`func (o *BcsImageV1ApiListImagesModelImageModel) GetMinRam() int32`

GetMinRam returns the MinRam field if non-nil, zero value otherwise.

### GetMinRamOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetMinRamOk() (*int32, bool)`

GetMinRamOk returns a tuple with the MinRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRam

`func (o *BcsImageV1ApiListImagesModelImageModel) SetMinRam(v int32)`

SetMinRam sets MinRam field to given value.

### HasMinRam

`func (o *BcsImageV1ApiListImagesModelImageModel) HasMinRam() bool`

HasMinRam returns a boolean if a field has been set.

### SetMinRamNil

`func (o *BcsImageV1ApiListImagesModelImageModel) SetMinRamNil(b bool)`

 SetMinRamNil sets the value for MinRam to be an explicit nil

### UnsetMinRam
`func (o *BcsImageV1ApiListImagesModelImageModel) UnsetMinRam()`

UnsetMinRam ensures that no value is present for MinRam, not even an explicit nil
### GetVirtualSize

`func (o *BcsImageV1ApiListImagesModelImageModel) GetVirtualSize() int64`

GetVirtualSize returns the VirtualSize field if non-nil, zero value otherwise.

### GetVirtualSizeOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetVirtualSizeOk() (*int64, bool)`

GetVirtualSizeOk returns a tuple with the VirtualSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSize

`func (o *BcsImageV1ApiListImagesModelImageModel) SetVirtualSize(v int64)`

SetVirtualSize sets VirtualSize field to given value.

### HasVirtualSize

`func (o *BcsImageV1ApiListImagesModelImageModel) HasVirtualSize() bool`

HasVirtualSize returns a boolean if a field has been set.

### SetVirtualSizeNil

`func (o *BcsImageV1ApiListImagesModelImageModel) SetVirtualSizeNil(b bool)`

 SetVirtualSizeNil sets the value for VirtualSize to be an explicit nil

### UnsetVirtualSize
`func (o *BcsImageV1ApiListImagesModelImageModel) UnsetVirtualSize()`

UnsetVirtualSize ensures that no value is present for VirtualSize, not even an explicit nil
### GetInstanceType

`func (o *BcsImageV1ApiListImagesModelImageModel) GetInstanceType() InstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetInstanceTypeOk() (*InstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *BcsImageV1ApiListImagesModelImageModel) SetInstanceType(v InstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *BcsImageV1ApiListImagesModelImageModel) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### SetInstanceTypeNil

`func (o *BcsImageV1ApiListImagesModelImageModel) SetInstanceTypeNil(b bool)`

 SetInstanceTypeNil sets the value for InstanceType to be an explicit nil

### UnsetInstanceType
`func (o *BcsImageV1ApiListImagesModelImageModel) UnsetInstanceType()`

UnsetInstanceType ensures that no value is present for InstanceType, not even an explicit nil
### GetImageMemberStatus

`func (o *BcsImageV1ApiListImagesModelImageModel) GetImageMemberStatus() string`

GetImageMemberStatus returns the ImageMemberStatus field if non-nil, zero value otherwise.

### GetImageMemberStatusOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetImageMemberStatusOk() (*string, bool)`

GetImageMemberStatusOk returns a tuple with the ImageMemberStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageMemberStatus

`func (o *BcsImageV1ApiListImagesModelImageModel) SetImageMemberStatus(v string)`

SetImageMemberStatus sets ImageMemberStatus field to given value.

### HasImageMemberStatus

`func (o *BcsImageV1ApiListImagesModelImageModel) HasImageMemberStatus() bool`

HasImageMemberStatus returns a boolean if a field has been set.

### SetImageMemberStatusNil

`func (o *BcsImageV1ApiListImagesModelImageModel) SetImageMemberStatusNil(b bool)`

 SetImageMemberStatusNil sets the value for ImageMemberStatus to be an explicit nil

### UnsetImageMemberStatus
`func (o *BcsImageV1ApiListImagesModelImageModel) UnsetImageMemberStatus()`

UnsetImageMemberStatus ensures that no value is present for ImageMemberStatus, not even an explicit nil
### GetProjectId

`func (o *BcsImageV1ApiListImagesModelImageModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BcsImageV1ApiListImagesModelImageModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BcsImageV1ApiListImagesModelImageModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BcsImageV1ApiListImagesModelImageModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BcsImageV1ApiListImagesModelImageModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetOsInfo

`func (o *BcsImageV1ApiListImagesModelImageModel) GetOsInfo() BcsImageV1ApiListImagesModelOsInfoModel`

GetOsInfo returns the OsInfo field if non-nil, zero value otherwise.

### GetOsInfoOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetOsInfoOk() (*BcsImageV1ApiListImagesModelOsInfoModel, bool)`

GetOsInfoOk returns a tuple with the OsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsInfo

`func (o *BcsImageV1ApiListImagesModelImageModel) SetOsInfo(v BcsImageV1ApiListImagesModelOsInfoModel)`

SetOsInfo sets OsInfo field to given value.

### HasOsInfo

`func (o *BcsImageV1ApiListImagesModelImageModel) HasOsInfo() bool`

HasOsInfo returns a boolean if a field has been set.

### SetOsInfoNil

`func (o *BcsImageV1ApiListImagesModelImageModel) SetOsInfoNil(b bool)`

 SetOsInfoNil sets the value for OsInfo to be an explicit nil

### UnsetOsInfo
`func (o *BcsImageV1ApiListImagesModelImageModel) UnsetOsInfo()`

UnsetOsInfo ensures that no value is present for OsInfo, not even an explicit nil
### GetCreatedAt

`func (o *BcsImageV1ApiListImagesModelImageModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsImageV1ApiListImagesModelImageModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BcsImageV1ApiListImagesModelImageModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BcsImageV1ApiListImagesModelImageModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BcsImageV1ApiListImagesModelImageModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BcsImageV1ApiListImagesModelImageModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BcsImageV1ApiListImagesModelImageModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BcsImageV1ApiListImagesModelImageModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BcsImageV1ApiListImagesModelImageModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BcsImageV1ApiListImagesModelImageModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BcsImageV1ApiListImagesModelImageModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


