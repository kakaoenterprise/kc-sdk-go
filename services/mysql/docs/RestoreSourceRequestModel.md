# RestoreSourceRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**RestoreSourceType**](RestoreSourceType.md) |  | 
**Id** | **string** | 소스로 사용할 백업 또는 리소스 ID | 
**Time** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRestoreSourceRequestModel

`func NewRestoreSourceRequestModel(type_ RestoreSourceType, id string, ) *RestoreSourceRequestModel`

NewRestoreSourceRequestModel instantiates a new RestoreSourceRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreSourceRequestModelWithDefaults

`func NewRestoreSourceRequestModelWithDefaults() *RestoreSourceRequestModel`

NewRestoreSourceRequestModelWithDefaults instantiates a new RestoreSourceRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RestoreSourceRequestModel) GetType() RestoreSourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestoreSourceRequestModel) GetTypeOk() (*RestoreSourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestoreSourceRequestModel) SetType(v RestoreSourceType)`

SetType sets Type field to given value.


### GetId

`func (o *RestoreSourceRequestModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestoreSourceRequestModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestoreSourceRequestModel) SetId(v string)`

SetId sets Id field to given value.


### GetTime

`func (o *RestoreSourceRequestModel) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *RestoreSourceRequestModel) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *RestoreSourceRequestModel) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *RestoreSourceRequestModel) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTimeNil

`func (o *RestoreSourceRequestModel) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *RestoreSourceRequestModel) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


