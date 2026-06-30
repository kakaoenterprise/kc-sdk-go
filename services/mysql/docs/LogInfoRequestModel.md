# LogInfoRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogType** | [**LogType**](LogType.md) | 내보낼 로그 유형 | 
**StartDate** | Pointer to **NullableString** |  | [optional] 
**EndDate** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLogInfoRequestModel

`func NewLogInfoRequestModel(logType LogType, ) *LogInfoRequestModel`

NewLogInfoRequestModel instantiates a new LogInfoRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogInfoRequestModelWithDefaults

`func NewLogInfoRequestModelWithDefaults() *LogInfoRequestModel`

NewLogInfoRequestModelWithDefaults instantiates a new LogInfoRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogType

`func (o *LogInfoRequestModel) GetLogType() LogType`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *LogInfoRequestModel) GetLogTypeOk() (*LogType, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *LogInfoRequestModel) SetLogType(v LogType)`

SetLogType sets LogType field to given value.


### GetStartDate

`func (o *LogInfoRequestModel) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LogInfoRequestModel) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *LogInfoRequestModel) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *LogInfoRequestModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *LogInfoRequestModel) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *LogInfoRequestModel) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *LogInfoRequestModel) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *LogInfoRequestModel) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *LogInfoRequestModel) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *LogInfoRequestModel) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *LogInfoRequestModel) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *LogInfoRequestModel) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


