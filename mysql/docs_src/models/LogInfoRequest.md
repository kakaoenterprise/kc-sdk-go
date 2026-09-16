# LogInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogType** | [**LogType**](LogType.md) | 내보낼 로그 유형 | 
**StartDate** | Pointer to **NullableString** | 로그 추출 시작 시각 (&#x60;yyyy-mm-dd&#x60;, UTC 기준) - &#x60;BIN_LOG&#x60;는 날짜 지정 불가 - 최근 7일 이내의 날짜만 지정 가능 | [optional] 
**EndDate** | Pointer to **NullableString** | 로그 추출 종료 시각 (&#x60;yyyy-mm-dd&#x60;, UTC 기준) - &#x60;BIN_LOG&#x60;는 날짜 지정 불가 - 최근 7일 이내의 날짜만 지정 가능 | [optional] 

## Methods

### NewLogInfoRequest

`func NewLogInfoRequest(logType LogType, ) *LogInfoRequest`

NewLogInfoRequest instantiates a new LogInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogInfoRequestWithDefaults

`func NewLogInfoRequestWithDefaults() *LogInfoRequest`

NewLogInfoRequestWithDefaults instantiates a new LogInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogType

`func (o *LogInfoRequest) GetLogType() LogType`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *LogInfoRequest) GetLogTypeOk() (*LogType, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *LogInfoRequest) SetLogType(v LogType)`

SetLogType sets LogType field to given value.


### GetStartDate

`func (o *LogInfoRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LogInfoRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *LogInfoRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *LogInfoRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *LogInfoRequest) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *LogInfoRequest) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *LogInfoRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *LogInfoRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *LogInfoRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *LogInfoRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *LogInfoRequest) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *LogInfoRequest) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


