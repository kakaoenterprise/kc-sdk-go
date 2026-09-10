# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActivated** | **bool** | Secrets Manager 서비스 활성화 여부 | 

## Methods

### NewService

`func NewService(isActivated bool, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActivated

`func (o *Service) GetIsActivated() bool`

GetIsActivated returns the IsActivated field if non-nil, zero value otherwise.

### GetIsActivatedOk

`func (o *Service) GetIsActivatedOk() (*bool, bool)`

GetIsActivatedOk returns a tuple with the IsActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActivated

`func (o *Service) SetIsActivated(v bool)`

SetIsActivated sets IsActivated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


