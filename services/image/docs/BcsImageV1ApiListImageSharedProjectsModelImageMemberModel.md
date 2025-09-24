# BcsImageV1ApiListImageSharedProjectsModelImageMemberModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 이미지를 공유받은 프로젝트 ID | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**ImageId** | **string** | 이미지의 고유 ID | 
**Status** | Pointer to **NullableString** |  | [optional] 
**IsShared** | **bool** | 공유 여부 | 

## Methods

### NewBcsImageV1ApiListImageSharedProjectsModelImageMemberModel

`func NewBcsImageV1ApiListImageSharedProjectsModelImageMemberModel(id string, createdAt time.Time, updatedAt time.Time, imageId string, isShared bool, ) *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel`

NewBcsImageV1ApiListImageSharedProjectsModelImageMemberModel instantiates a new BcsImageV1ApiListImageSharedProjectsModelImageMemberModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsImageV1ApiListImageSharedProjectsModelImageMemberModelWithDefaults

`func NewBcsImageV1ApiListImageSharedProjectsModelImageMemberModelWithDefaults() *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel`

NewBcsImageV1ApiListImageSharedProjectsModelImageMemberModelWithDefaults instantiates a new BcsImageV1ApiListImageSharedProjectsModelImageMemberModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetImageId

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetStatus

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetIsShared

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *BcsImageV1ApiListImageSharedProjectsModelImageMemberModel) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


