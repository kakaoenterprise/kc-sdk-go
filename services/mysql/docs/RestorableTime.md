# RestorableTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromTime** | **string** | 복원 가능한 시작 시각 | 
**ToTime** | **string** | 복원 가능한 종료 시각 | 

## Methods

### NewRestorableTime

`func NewRestorableTime(fromTime string, toTime string, ) *RestorableTime`

NewRestorableTime instantiates a new RestorableTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestorableTimeWithDefaults

`func NewRestorableTimeWithDefaults() *RestorableTime`

NewRestorableTimeWithDefaults instantiates a new RestorableTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromTime

`func (o *RestorableTime) GetFromTime() string`

GetFromTime returns the FromTime field if non-nil, zero value otherwise.

### GetFromTimeOk

`func (o *RestorableTime) GetFromTimeOk() (*string, bool)`

GetFromTimeOk returns a tuple with the FromTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTime

`func (o *RestorableTime) SetFromTime(v string)`

SetFromTime sets FromTime field to given value.


### GetToTime

`func (o *RestorableTime) GetToTime() string`

GetToTime returns the ToTime field if non-nil, zero value otherwise.

### GetToTimeOk

`func (o *RestorableTime) GetToTimeOk() (*string, bool)`

GetToTimeOk returns a tuple with the ToTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTime

`func (o *RestorableTime) SetToTime(v string)`

SetToTime sets ToTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


