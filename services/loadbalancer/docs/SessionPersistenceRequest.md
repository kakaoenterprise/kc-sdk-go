# SessionPersistenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**LoadBalancerSessionPersistenceType**](LoadBalancerSessionPersistenceType.md) | 세션 유형 &lt;br/&gt; - 예시: &#x60;APP_COOKIE&#x60;, &#x60;HTTP_COOKIE&#x60;, &#x60;SOURCE_IP&#x60; | 
**CookieName** | Pointer to **NullableString** |  | [optional] 
**PersistenceTimeout** | **int32** | 세션 유지 시간 (초) | 
**PersistenceGranularity** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSessionPersistenceRequest

`func NewSessionPersistenceRequest(type_ LoadBalancerSessionPersistenceType, persistenceTimeout int32, ) *SessionPersistenceRequest`

NewSessionPersistenceRequest instantiates a new SessionPersistenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionPersistenceRequestWithDefaults

`func NewSessionPersistenceRequestWithDefaults() *SessionPersistenceRequest`

NewSessionPersistenceRequestWithDefaults instantiates a new SessionPersistenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SessionPersistenceRequest) GetType() LoadBalancerSessionPersistenceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SessionPersistenceRequest) GetTypeOk() (*LoadBalancerSessionPersistenceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SessionPersistenceRequest) SetType(v LoadBalancerSessionPersistenceType)`

SetType sets Type field to given value.


### GetCookieName

`func (o *SessionPersistenceRequest) GetCookieName() string`

GetCookieName returns the CookieName field if non-nil, zero value otherwise.

### GetCookieNameOk

`func (o *SessionPersistenceRequest) GetCookieNameOk() (*string, bool)`

GetCookieNameOk returns a tuple with the CookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieName

`func (o *SessionPersistenceRequest) SetCookieName(v string)`

SetCookieName sets CookieName field to given value.

### HasCookieName

`func (o *SessionPersistenceRequest) HasCookieName() bool`

HasCookieName returns a boolean if a field has been set.

### SetCookieNameNil

`func (o *SessionPersistenceRequest) SetCookieNameNil(b bool)`

 SetCookieNameNil sets the value for CookieName to be an explicit nil

### UnsetCookieName
`func (o *SessionPersistenceRequest) UnsetCookieName()`

UnsetCookieName ensures that no value is present for CookieName, not even an explicit nil
### GetPersistenceTimeout

`func (o *SessionPersistenceRequest) GetPersistenceTimeout() int32`

GetPersistenceTimeout returns the PersistenceTimeout field if non-nil, zero value otherwise.

### GetPersistenceTimeoutOk

`func (o *SessionPersistenceRequest) GetPersistenceTimeoutOk() (*int32, bool)`

GetPersistenceTimeoutOk returns a tuple with the PersistenceTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceTimeout

`func (o *SessionPersistenceRequest) SetPersistenceTimeout(v int32)`

SetPersistenceTimeout sets PersistenceTimeout field to given value.


### GetPersistenceGranularity

`func (o *SessionPersistenceRequest) GetPersistenceGranularity() string`

GetPersistenceGranularity returns the PersistenceGranularity field if non-nil, zero value otherwise.

### GetPersistenceGranularityOk

`func (o *SessionPersistenceRequest) GetPersistenceGranularityOk() (*string, bool)`

GetPersistenceGranularityOk returns a tuple with the PersistenceGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceGranularity

`func (o *SessionPersistenceRequest) SetPersistenceGranularity(v string)`

SetPersistenceGranularity sets PersistenceGranularity field to given value.

### HasPersistenceGranularity

`func (o *SessionPersistenceRequest) HasPersistenceGranularity() bool`

HasPersistenceGranularity returns a boolean if a field has been set.

### SetPersistenceGranularityNil

`func (o *SessionPersistenceRequest) SetPersistenceGranularityNil(b bool)`

 SetPersistenceGranularityNil sets the value for PersistenceGranularity to be an explicit nil

### UnsetPersistenceGranularity
`func (o *SessionPersistenceRequest) UnsetPersistenceGranularity()`

UnsetPersistenceGranularity ensures that no value is present for PersistenceGranularity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


