# GetMysqlRestorableTimeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestorableTime** | [**RestorableTime**](RestorableTime.md) | MySQL 인스턴스 그룹의 복원 가능 시간 범위 정보 | 

## Methods

### NewGetMysqlRestorableTimeResponse

`func NewGetMysqlRestorableTimeResponse(restorableTime RestorableTime, ) *GetMysqlRestorableTimeResponse`

NewGetMysqlRestorableTimeResponse instantiates a new GetMysqlRestorableTimeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMysqlRestorableTimeResponseWithDefaults

`func NewGetMysqlRestorableTimeResponseWithDefaults() *GetMysqlRestorableTimeResponse`

NewGetMysqlRestorableTimeResponseWithDefaults instantiates a new GetMysqlRestorableTimeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestorableTime

`func (o *GetMysqlRestorableTimeResponse) GetRestorableTime() RestorableTime`

GetRestorableTime returns the RestorableTime field if non-nil, zero value otherwise.

### GetRestorableTimeOk

`func (o *GetMysqlRestorableTimeResponse) GetRestorableTimeOk() (*RestorableTime, bool)`

GetRestorableTimeOk returns a tuple with the RestorableTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorableTime

`func (o *GetMysqlRestorableTimeResponse) SetRestorableTime(v RestorableTime)`

SetRestorableTime sets RestorableTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


