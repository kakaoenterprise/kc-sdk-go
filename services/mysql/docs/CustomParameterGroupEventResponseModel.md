# CustomParameterGroupEventResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 이벤트 이름 | 
**Description** | **string** | 이벤트 상세 설명 | 
**CreatedAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 

## Methods

### NewCustomParameterGroupEventResponseModel

`func NewCustomParameterGroupEventResponseModel(name string, description string, createdAt string, ) *CustomParameterGroupEventResponseModel`

NewCustomParameterGroupEventResponseModel instantiates a new CustomParameterGroupEventResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomParameterGroupEventResponseModelWithDefaults

`func NewCustomParameterGroupEventResponseModelWithDefaults() *CustomParameterGroupEventResponseModel`

NewCustomParameterGroupEventResponseModelWithDefaults instantiates a new CustomParameterGroupEventResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomParameterGroupEventResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomParameterGroupEventResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomParameterGroupEventResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CustomParameterGroupEventResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomParameterGroupEventResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomParameterGroupEventResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedAt

`func (o *CustomParameterGroupEventResponseModel) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomParameterGroupEventResponseModel) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomParameterGroupEventResponseModel) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


