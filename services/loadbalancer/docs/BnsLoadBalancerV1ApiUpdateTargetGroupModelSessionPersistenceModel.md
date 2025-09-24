# BnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | 세션 유형 &lt;br/&gt; - 예시: &#x60;APP_COOKIE&#x60;, &#x60;HTTP_COOKIE&#x60;, &#x60;SOURCE_IP&#x60; | 
**CookieName** | **string** | 쿠키 이름 (세션 유형이 &#x60;APP_COOKIE&#x60;, &#x60;HTTP_COOKIE&#x60;인 경우) | 
**PersistenceTimeout** | **int32** | 세션 유지 시간 (초) | 
**PersistenceGranularity** | **string** | 세션 유지 기준이 되는 IP 서브넷 마스크 (세션 유형이 &#x60;SOURCE_IP&#x60;인 경우) | 

## Methods

### NewBnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel

`func NewBnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel(type_ string, cookieName string, persistenceTimeout int32, persistenceGranularity string, ) *BnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel`

NewBnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel instantiates a new BnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModelWithDefaults

`func NewBnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModelWithDefaults() *BnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel`

NewBnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModelWithDefaults instantiates a new BnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel) SetType(v string)`

SetType sets Type field to given value.


### GetCookieName

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel) GetCookieName() string`

GetCookieName returns the CookieName field if non-nil, zero value otherwise.

### GetCookieNameOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel) GetCookieNameOk() (*string, bool)`

GetCookieNameOk returns a tuple with the CookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieName

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel) SetCookieName(v string)`

SetCookieName sets CookieName field to given value.


### GetPersistenceTimeout

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel) GetPersistenceTimeout() int32`

GetPersistenceTimeout returns the PersistenceTimeout field if non-nil, zero value otherwise.

### GetPersistenceTimeoutOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel) GetPersistenceTimeoutOk() (*int32, bool)`

GetPersistenceTimeoutOk returns a tuple with the PersistenceTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceTimeout

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel) SetPersistenceTimeout(v int32)`

SetPersistenceTimeout sets PersistenceTimeout field to given value.


### GetPersistenceGranularity

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel) GetPersistenceGranularity() string`

GetPersistenceGranularity returns the PersistenceGranularity field if non-nil, zero value otherwise.

### GetPersistenceGranularityOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel) GetPersistenceGranularityOk() (*string, bool)`

GetPersistenceGranularityOk returns a tuple with the PersistenceGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceGranularity

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelSessionPersistenceModel) SetPersistenceGranularity(v string)`

SetPersistenceGranularity sets PersistenceGranularity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


