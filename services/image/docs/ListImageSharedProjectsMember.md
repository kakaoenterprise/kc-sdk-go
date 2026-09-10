# ListImageSharedProjectsMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 이미지를 공유받은 프로젝트 ID | 
**ImageId** | **string** | 이미지의 고유 ID | 
**Status** | Pointer to **NullableString** |  | [optional] 
**IsShared** | **bool** | 공유 여부 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 

## Methods

### NewListImageSharedProjectsMember

`func NewListImageSharedProjectsMember(id string, imageId string, isShared bool, createdAt time.Time, updatedAt time.Time, ) *ListImageSharedProjectsMember`

NewListImageSharedProjectsMember instantiates a new ListImageSharedProjectsMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListImageSharedProjectsMemberWithDefaults

`func NewListImageSharedProjectsMemberWithDefaults() *ListImageSharedProjectsMember`

NewListImageSharedProjectsMemberWithDefaults instantiates a new ListImageSharedProjectsMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListImageSharedProjectsMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListImageSharedProjectsMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListImageSharedProjectsMember) SetId(v string)`

SetId sets Id field to given value.


### GetImageId

`func (o *ListImageSharedProjectsMember) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ListImageSharedProjectsMember) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ListImageSharedProjectsMember) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetStatus

`func (o *ListImageSharedProjectsMember) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListImageSharedProjectsMember) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListImageSharedProjectsMember) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListImageSharedProjectsMember) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ListImageSharedProjectsMember) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ListImageSharedProjectsMember) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetIsShared

`func (o *ListImageSharedProjectsMember) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *ListImageSharedProjectsMember) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *ListImageSharedProjectsMember) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.


### GetCreatedAt

`func (o *ListImageSharedProjectsMember) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListImageSharedProjectsMember) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListImageSharedProjectsMember) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListImageSharedProjectsMember) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListImageSharedProjectsMember) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListImageSharedProjectsMember) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


