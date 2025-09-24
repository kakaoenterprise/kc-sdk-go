# BcsImageV1ApiAddImageShareModelImageMemberModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | **string** | 이미지의 고유 ID | 
**Status** | **string** | 이미지 공유 상태 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**Id** | **string** | 이미지를 공유받은 프로젝트의 ID | 

## Methods

### NewBcsImageV1ApiAddImageShareModelImageMemberModel

`func NewBcsImageV1ApiAddImageShareModelImageMemberModel(imageId string, status string, createdAt time.Time, updatedAt time.Time, id string, ) *BcsImageV1ApiAddImageShareModelImageMemberModel`

NewBcsImageV1ApiAddImageShareModelImageMemberModel instantiates a new BcsImageV1ApiAddImageShareModelImageMemberModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsImageV1ApiAddImageShareModelImageMemberModelWithDefaults

`func NewBcsImageV1ApiAddImageShareModelImageMemberModelWithDefaults() *BcsImageV1ApiAddImageShareModelImageMemberModel`

NewBcsImageV1ApiAddImageShareModelImageMemberModelWithDefaults instantiates a new BcsImageV1ApiAddImageShareModelImageMemberModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *BcsImageV1ApiAddImageShareModelImageMemberModel) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *BcsImageV1ApiAddImageShareModelImageMemberModel) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *BcsImageV1ApiAddImageShareModelImageMemberModel) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetStatus

`func (o *BcsImageV1ApiAddImageShareModelImageMemberModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BcsImageV1ApiAddImageShareModelImageMemberModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BcsImageV1ApiAddImageShareModelImageMemberModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *BcsImageV1ApiAddImageShareModelImageMemberModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsImageV1ApiAddImageShareModelImageMemberModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsImageV1ApiAddImageShareModelImageMemberModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BcsImageV1ApiAddImageShareModelImageMemberModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BcsImageV1ApiAddImageShareModelImageMemberModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BcsImageV1ApiAddImageShareModelImageMemberModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetId

`func (o *BcsImageV1ApiAddImageShareModelImageMemberModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsImageV1ApiAddImageShareModelImageMemberModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsImageV1ApiAddImageShareModelImageMemberModel) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


