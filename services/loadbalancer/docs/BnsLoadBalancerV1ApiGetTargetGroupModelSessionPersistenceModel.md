# BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | 세션 유형 &lt;br/&gt; - 예시: &#x60;APP_COOKIE&#x60;, &#x60;HTTP_COOKIE&#x60;, &#x60;SOURCE_IP&#x60; | 
**CookieName** | Pointer to **NullableString** |  | [optional] 
**PersistenceTimeout** | **int32** | 세션 유지 시간 (초) | 
**PersistenceGranularity** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel

`func NewBnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel(type_ string, persistenceTimeout int32, ) *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel`

NewBnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel instantiates a new BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModelWithDefaults

`func NewBnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModelWithDefaults() *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel`

NewBnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModelWithDefaults instantiates a new BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel) SetType(v string)`

SetType sets Type field to given value.


### GetCookieName

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel) GetCookieName() string`

GetCookieName returns the CookieName field if non-nil, zero value otherwise.

### GetCookieNameOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel) GetCookieNameOk() (*string, bool)`

GetCookieNameOk returns a tuple with the CookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieName

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel) SetCookieName(v string)`

SetCookieName sets CookieName field to given value.

### HasCookieName

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel) HasCookieName() bool`

HasCookieName returns a boolean if a field has been set.

### SetCookieNameNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel) SetCookieNameNil(b bool)`

 SetCookieNameNil sets the value for CookieName to be an explicit nil

### UnsetCookieName
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel) UnsetCookieName()`

UnsetCookieName ensures that no value is present for CookieName, not even an explicit nil
### GetPersistenceTimeout

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel) GetPersistenceTimeout() int32`

GetPersistenceTimeout returns the PersistenceTimeout field if non-nil, zero value otherwise.

### GetPersistenceTimeoutOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel) GetPersistenceTimeoutOk() (*int32, bool)`

GetPersistenceTimeoutOk returns a tuple with the PersistenceTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceTimeout

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel) SetPersistenceTimeout(v int32)`

SetPersistenceTimeout sets PersistenceTimeout field to given value.


### GetPersistenceGranularity

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel) GetPersistenceGranularity() string`

GetPersistenceGranularity returns the PersistenceGranularity field if non-nil, zero value otherwise.

### GetPersistenceGranularityOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel) GetPersistenceGranularityOk() (*string, bool)`

GetPersistenceGranularityOk returns a tuple with the PersistenceGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceGranularity

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel) SetPersistenceGranularity(v string)`

SetPersistenceGranularity sets PersistenceGranularity field to given value.

### HasPersistenceGranularity

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel) HasPersistenceGranularity() bool`

HasPersistenceGranularity returns a boolean if a field has been set.

### SetPersistenceGranularityNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel) SetPersistenceGranularityNil(b bool)`

 SetPersistenceGranularityNil sets the value for PersistenceGranularity to be an explicit nil

### UnsetPersistenceGranularity
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelSessionPersistenceModel) UnsetPersistenceGranularity()`

UnsetPersistenceGranularity ensures that no value is present for PersistenceGranularity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


