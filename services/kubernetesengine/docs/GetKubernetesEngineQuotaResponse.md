# GetKubernetesEngineQuotaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quota** | [**Quota**](Quota.md) | 프로젝트 전체 Kubernetes Engine 할당량 정보 | 

## Methods

### NewGetKubernetesEngineQuotaResponse

`func NewGetKubernetesEngineQuotaResponse(quota Quota, ) *GetKubernetesEngineQuotaResponse`

NewGetKubernetesEngineQuotaResponse instantiates a new GetKubernetesEngineQuotaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKubernetesEngineQuotaResponseWithDefaults

`func NewGetKubernetesEngineQuotaResponseWithDefaults() *GetKubernetesEngineQuotaResponse`

NewGetKubernetesEngineQuotaResponseWithDefaults instantiates a new GetKubernetesEngineQuotaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuota

`func (o *GetKubernetesEngineQuotaResponse) GetQuota() Quota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *GetKubernetesEngineQuotaResponse) GetQuotaOk() (*Quota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *GetKubernetesEngineQuotaResponse) SetQuota(v Quota)`

SetQuota sets Quota field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


