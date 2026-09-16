# CustomParameterGroupEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 이벤트 이름 | 
**Description** | **string** | 이벤트 상세 설명 | 
**CreatedAt** | **string** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 

## Methods

### NewCustomParameterGroupEvent

`func NewCustomParameterGroupEvent(name string, description string, createdAt string, ) *CustomParameterGroupEvent`

NewCustomParameterGroupEvent instantiates a new CustomParameterGroupEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomParameterGroupEventWithDefaults

`func NewCustomParameterGroupEventWithDefaults() *CustomParameterGroupEvent`

NewCustomParameterGroupEventWithDefaults instantiates a new CustomParameterGroupEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomParameterGroupEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomParameterGroupEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomParameterGroupEvent) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CustomParameterGroupEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomParameterGroupEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomParameterGroupEvent) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedAt

`func (o *CustomParameterGroupEvent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomParameterGroupEvent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomParameterGroupEvent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


