# RestoreSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**RestoreSourceType**](RestoreSourceType.md) | 생성 소스 유형 | 
**Id** | **string** | 소스로 사용할 백업 또는 리소스 ID &lt;br/&gt;- &#x60;type&#x3D;BACKUP&#x60;: [List MySQL backups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-backups)에서 확인 &lt;br/&gt;- &#x60;type&#x3D;INSTANCE_GROUP&#x60;: [List MySQL instance groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-instance-groups)에서 확인 | 
**Time** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRestoreSourceRequest

`func NewRestoreSourceRequest(type_ RestoreSourceType, id string, ) *RestoreSourceRequest`

NewRestoreSourceRequest instantiates a new RestoreSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreSourceRequestWithDefaults

`func NewRestoreSourceRequestWithDefaults() *RestoreSourceRequest`

NewRestoreSourceRequestWithDefaults instantiates a new RestoreSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RestoreSourceRequest) GetType() RestoreSourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestoreSourceRequest) GetTypeOk() (*RestoreSourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestoreSourceRequest) SetType(v RestoreSourceType)`

SetType sets Type field to given value.


### GetId

`func (o *RestoreSourceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestoreSourceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestoreSourceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetTime

`func (o *RestoreSourceRequest) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *RestoreSourceRequest) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *RestoreSourceRequest) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *RestoreSourceRequest) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTimeNil

`func (o *RestoreSourceRequest) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *RestoreSourceRequest) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


