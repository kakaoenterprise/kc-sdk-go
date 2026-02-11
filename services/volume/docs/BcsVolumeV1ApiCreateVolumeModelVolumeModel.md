# BcsVolumeV1ApiCreateVolumeModelVolumeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 생성된 볼륨의 ID | 
**Status** | **string** | 볼륨의 현재 상태 | 
**Size** | **int32** | 볼륨의 크기(GB 단위) | 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**Name** | **string** | 생성 시 사용자가 지정한 볼륨 이름 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**VolumeType** | Pointer to **NullableString** |  | [optional] 
**SnapshotId** | Pointer to **NullableString** |  | [optional] 
**Metadata** | **map[string]string** | 볼륨에 설정된 사용자 정의 메타데이터(Key-Value 쌍) | 
**UserId** | Pointer to **NullableString** |  | [optional] 
**IsBootable** | **bool** | 부팅 가능 여부 | 
**IsEncrypted** | Pointer to **NullableBool** |  | [optional] 
**Attachments** | [**[]Attachment**](Attachment.md) | 볼륨의 인스턴스 연결 상태를 나타내는 Attachment 목록 &lt;br/&gt; - 새로 생성된 볼륨은 일반적으로 빈 배열 | 

## Methods

### NewBcsVolumeV1ApiCreateVolumeModelVolumeModel

`func NewBcsVolumeV1ApiCreateVolumeModelVolumeModel(id string, status string, size int32, name string, metadata map[string]string, isBootable bool, attachments []Attachment, ) *BcsVolumeV1ApiCreateVolumeModelVolumeModel`

NewBcsVolumeV1ApiCreateVolumeModelVolumeModel instantiates a new BcsVolumeV1ApiCreateVolumeModelVolumeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsVolumeV1ApiCreateVolumeModelVolumeModelWithDefaults

`func NewBcsVolumeV1ApiCreateVolumeModelVolumeModelWithDefaults() *BcsVolumeV1ApiCreateVolumeModelVolumeModel`

NewBcsVolumeV1ApiCreateVolumeModelVolumeModelWithDefaults instantiates a new BcsVolumeV1ApiCreateVolumeModelVolumeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSize

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetSize(v int32)`

SetSize sets Size field to given value.


### GetAvailabilityZone

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetCreatedAt

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetName

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVolumeType

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeTypeNil

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
### GetSnapshotId

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetMetadata

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.


### GetUserId

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetIsBootable

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetIsBootable() bool`

GetIsBootable returns the IsBootable field if non-nil, zero value otherwise.

### GetIsBootableOk

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetIsBootableOk() (*bool, bool)`

GetIsBootableOk returns a tuple with the IsBootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBootable

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetIsBootable(v bool)`

SetIsBootable sets IsBootable field to given value.


### GetIsEncrypted

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetIsEncryptedOk() (*bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncrypted

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetIsEncrypted(v bool)`

SetIsEncrypted sets IsEncrypted field to given value.

### HasIsEncrypted

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.

### SetIsEncryptedNil

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetIsEncryptedNil(b bool)`

 SetIsEncryptedNil sets the value for IsEncrypted to be an explicit nil

### UnsetIsEncrypted
`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) UnsetIsEncrypted()`

UnsetIsEncrypted ensures that no value is present for IsEncrypted, not even an explicit nil
### GetAttachments

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *BcsVolumeV1ApiCreateVolumeModelVolumeModel) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


