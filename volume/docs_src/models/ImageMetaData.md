# ImageMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerFormat** | Pointer to **NullableString** | 이미지의 컨테이너 형식 | [optional] 
**DiskFormat** | Pointer to **NullableString** | 이미지의 디스크 형식 | [optional] 
**ImageId** | Pointer to **NullableString** | 이미지의 고유 ID | [optional] 
**ImageName** | Pointer to **NullableString** | 이미지 이름 | [optional] 
**MinDisk** | Pointer to **NullableString** | 이미지 실행에 필요한 최소 디스크 크기 (GB 단위) | [optional] 
**OsType** | Pointer to **NullableString** | 운영체제 유형 | [optional] 
**MinRam** | Pointer to **NullableString** | 이미지 실행에 필요한 최소 메모리 크기 (MB 단위) | [optional] 
**Size** | Pointer to **NullableString** | 이미지 파일 크기 (Byte 단위) | [optional] 

## Methods

### NewImageMetaData

`func NewImageMetaData() *ImageMetaData`

NewImageMetaData instantiates a new ImageMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageMetaDataWithDefaults

`func NewImageMetaDataWithDefaults() *ImageMetaData`

NewImageMetaDataWithDefaults instantiates a new ImageMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerFormat

`func (o *ImageMetaData) GetContainerFormat() string`

GetContainerFormat returns the ContainerFormat field if non-nil, zero value otherwise.

### GetContainerFormatOk

`func (o *ImageMetaData) GetContainerFormatOk() (*string, bool)`

GetContainerFormatOk returns a tuple with the ContainerFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerFormat

`func (o *ImageMetaData) SetContainerFormat(v string)`

SetContainerFormat sets ContainerFormat field to given value.

### HasContainerFormat

`func (o *ImageMetaData) HasContainerFormat() bool`

HasContainerFormat returns a boolean if a field has been set.

### SetContainerFormatNil

`func (o *ImageMetaData) SetContainerFormatNil(b bool)`

 SetContainerFormatNil sets the value for ContainerFormat to be an explicit nil

### UnsetContainerFormat
`func (o *ImageMetaData) UnsetContainerFormat()`

UnsetContainerFormat ensures that no value is present for ContainerFormat, not even an explicit nil
### GetDiskFormat

`func (o *ImageMetaData) GetDiskFormat() string`

GetDiskFormat returns the DiskFormat field if non-nil, zero value otherwise.

### GetDiskFormatOk

`func (o *ImageMetaData) GetDiskFormatOk() (*string, bool)`

GetDiskFormatOk returns a tuple with the DiskFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFormat

`func (o *ImageMetaData) SetDiskFormat(v string)`

SetDiskFormat sets DiskFormat field to given value.

### HasDiskFormat

`func (o *ImageMetaData) HasDiskFormat() bool`

HasDiskFormat returns a boolean if a field has been set.

### SetDiskFormatNil

`func (o *ImageMetaData) SetDiskFormatNil(b bool)`

 SetDiskFormatNil sets the value for DiskFormat to be an explicit nil

### UnsetDiskFormat
`func (o *ImageMetaData) UnsetDiskFormat()`

UnsetDiskFormat ensures that no value is present for DiskFormat, not even an explicit nil
### GetImageId

`func (o *ImageMetaData) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ImageMetaData) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ImageMetaData) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *ImageMetaData) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageIdNil

`func (o *ImageMetaData) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *ImageMetaData) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
### GetImageName

`func (o *ImageMetaData) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ImageMetaData) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ImageMetaData) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ImageMetaData) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### SetImageNameNil

`func (o *ImageMetaData) SetImageNameNil(b bool)`

 SetImageNameNil sets the value for ImageName to be an explicit nil

### UnsetImageName
`func (o *ImageMetaData) UnsetImageName()`

UnsetImageName ensures that no value is present for ImageName, not even an explicit nil
### GetMinDisk

`func (o *ImageMetaData) GetMinDisk() string`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *ImageMetaData) GetMinDiskOk() (*string, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *ImageMetaData) SetMinDisk(v string)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *ImageMetaData) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### SetMinDiskNil

`func (o *ImageMetaData) SetMinDiskNil(b bool)`

 SetMinDiskNil sets the value for MinDisk to be an explicit nil

### UnsetMinDisk
`func (o *ImageMetaData) UnsetMinDisk()`

UnsetMinDisk ensures that no value is present for MinDisk, not even an explicit nil
### GetOsType

`func (o *ImageMetaData) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ImageMetaData) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ImageMetaData) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ImageMetaData) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### SetOsTypeNil

`func (o *ImageMetaData) SetOsTypeNil(b bool)`

 SetOsTypeNil sets the value for OsType to be an explicit nil

### UnsetOsType
`func (o *ImageMetaData) UnsetOsType()`

UnsetOsType ensures that no value is present for OsType, not even an explicit nil
### GetMinRam

`func (o *ImageMetaData) GetMinRam() string`

GetMinRam returns the MinRam field if non-nil, zero value otherwise.

### GetMinRamOk

`func (o *ImageMetaData) GetMinRamOk() (*string, bool)`

GetMinRamOk returns a tuple with the MinRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRam

`func (o *ImageMetaData) SetMinRam(v string)`

SetMinRam sets MinRam field to given value.

### HasMinRam

`func (o *ImageMetaData) HasMinRam() bool`

HasMinRam returns a boolean if a field has been set.

### SetMinRamNil

`func (o *ImageMetaData) SetMinRamNil(b bool)`

 SetMinRamNil sets the value for MinRam to be an explicit nil

### UnsetMinRam
`func (o *ImageMetaData) UnsetMinRam()`

UnsetMinRam ensures that no value is present for MinRam, not even an explicit nil
### GetSize

`func (o *ImageMetaData) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ImageMetaData) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ImageMetaData) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ImageMetaData) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *ImageMetaData) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *ImageMetaData) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


