# CreateInstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | [**CreatedInstance**](CreatedInstance.md) | 생성된 인스턴스 정보 | 

## Methods

### NewCreateInstanceResponse

`func NewCreateInstanceResponse(instance CreatedInstance, ) *CreateInstanceResponse`

NewCreateInstanceResponse instantiates a new CreateInstanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstanceResponseWithDefaults

`func NewCreateInstanceResponseWithDefaults() *CreateInstanceResponse`

NewCreateInstanceResponseWithDefaults instantiates a new CreateInstanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *CreateInstanceResponse) GetInstance() CreatedInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *CreateInstanceResponse) GetInstanceOk() (*CreatedInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *CreateInstanceResponse) SetInstance(v CreatedInstance)`

SetInstance sets Instance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


