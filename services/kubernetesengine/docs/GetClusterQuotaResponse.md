# GetClusterQuotaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quota** | [**ClusterQuota**](ClusterQuota.md) | 클러스터 할당량 정보 | 

## Methods

### NewGetClusterQuotaResponse

`func NewGetClusterQuotaResponse(quota ClusterQuota, ) *GetClusterQuotaResponse`

NewGetClusterQuotaResponse instantiates a new GetClusterQuotaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterQuotaResponseWithDefaults

`func NewGetClusterQuotaResponseWithDefaults() *GetClusterQuotaResponse`

NewGetClusterQuotaResponseWithDefaults instantiates a new GetClusterQuotaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuota

`func (o *GetClusterQuotaResponse) GetQuota() ClusterQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *GetClusterQuotaResponse) GetQuotaOk() (*ClusterQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *GetClusterQuotaResponse) SetQuota(v ClusterQuota)`

SetQuota sets Quota field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


