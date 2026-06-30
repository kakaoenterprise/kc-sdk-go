# GetMySQLCustomParameterGroupEventsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]CustomParameterGroupEventResponseModel**](CustomParameterGroupEventResponseModel.md) | 커스텀 MySQL 파라미터 그룹과 관련된 이벤트 목록 | 

## Methods

### NewGetMySQLCustomParameterGroupEventsResponseModel

`func NewGetMySQLCustomParameterGroupEventsResponseModel(events []CustomParameterGroupEventResponseModel, ) *GetMySQLCustomParameterGroupEventsResponseModel`

NewGetMySQLCustomParameterGroupEventsResponseModel instantiates a new GetMySQLCustomParameterGroupEventsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMySQLCustomParameterGroupEventsResponseModelWithDefaults

`func NewGetMySQLCustomParameterGroupEventsResponseModelWithDefaults() *GetMySQLCustomParameterGroupEventsResponseModel`

NewGetMySQLCustomParameterGroupEventsResponseModelWithDefaults instantiates a new GetMySQLCustomParameterGroupEventsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *GetMySQLCustomParameterGroupEventsResponseModel) GetEvents() []CustomParameterGroupEventResponseModel`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetMySQLCustomParameterGroupEventsResponseModel) GetEventsOk() (*[]CustomParameterGroupEventResponseModel, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetMySQLCustomParameterGroupEventsResponseModel) SetEvents(v []CustomParameterGroupEventResponseModel)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


