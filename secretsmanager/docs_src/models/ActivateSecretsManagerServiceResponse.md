# ActivateSecretsManagerServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | [**Service**](Service.md) | Secrets Manager 서비스 활성화 정보 | 

## Methods

### NewActivateSecretsManagerServiceResponse

`func NewActivateSecretsManagerServiceResponse(service Service, ) *ActivateSecretsManagerServiceResponse`

NewActivateSecretsManagerServiceResponse instantiates a new ActivateSecretsManagerServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivateSecretsManagerServiceResponseWithDefaults

`func NewActivateSecretsManagerServiceResponseWithDefaults() *ActivateSecretsManagerServiceResponse`

NewActivateSecretsManagerServiceResponseWithDefaults instantiates a new ActivateSecretsManagerServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ActivateSecretsManagerServiceResponse) GetService() Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ActivateSecretsManagerServiceResponse) GetServiceOk() (*Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ActivateSecretsManagerServiceResponse) SetService(v Service)`

SetService sets Service field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


