# GetKmsServiceStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | [**GetKMSServiceStatusResponseInner**](GetKMSServiceStatusResponseInner.md) | KMS 서비스 활성화 정보 | 

## Methods

### NewGetKmsServiceStatusResponse

`func NewGetKmsServiceStatusResponse(service GetKMSServiceStatusResponseInner, ) *GetKmsServiceStatusResponse`

NewGetKmsServiceStatusResponse instantiates a new GetKmsServiceStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKmsServiceStatusResponseWithDefaults

`func NewGetKmsServiceStatusResponseWithDefaults() *GetKmsServiceStatusResponse`

NewGetKmsServiceStatusResponseWithDefaults instantiates a new GetKmsServiceStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *GetKmsServiceStatusResponse) GetService() GetKMSServiceStatusResponseInner`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetKmsServiceStatusResponse) GetServiceOk() (*GetKMSServiceStatusResponseInner, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetKmsServiceStatusResponse) SetService(v GetKMSServiceStatusResponseInner)`

SetService sets Service field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


