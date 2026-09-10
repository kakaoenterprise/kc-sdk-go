# SessionPersistence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**LoadBalancerSessionPersistenceType**](LoadBalancerSessionPersistenceType.md) | 세션 유형 &lt;br/&gt; - 예시: &#x60;APP_COOKIE&#x60;, &#x60;HTTP_COOKIE&#x60;, &#x60;SOURCE_IP&#x60; | 
**CookieName** | Pointer to **NullableString** |  | [optional] 
**PersistenceTimeout** | **int32** | 세션 유지 시간 (초) | 
**PersistenceGranularity** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSessionPersistence

`func NewSessionPersistence(type_ LoadBalancerSessionPersistenceType, persistenceTimeout int32, ) *SessionPersistence`

NewSessionPersistence instantiates a new SessionPersistence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionPersistenceWithDefaults

`func NewSessionPersistenceWithDefaults() *SessionPersistence`

NewSessionPersistenceWithDefaults instantiates a new SessionPersistence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SessionPersistence) GetType() LoadBalancerSessionPersistenceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SessionPersistence) GetTypeOk() (*LoadBalancerSessionPersistenceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SessionPersistence) SetType(v LoadBalancerSessionPersistenceType)`

SetType sets Type field to given value.


### GetCookieName

`func (o *SessionPersistence) GetCookieName() string`

GetCookieName returns the CookieName field if non-nil, zero value otherwise.

### GetCookieNameOk

`func (o *SessionPersistence) GetCookieNameOk() (*string, bool)`

GetCookieNameOk returns a tuple with the CookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieName

`func (o *SessionPersistence) SetCookieName(v string)`

SetCookieName sets CookieName field to given value.

### HasCookieName

`func (o *SessionPersistence) HasCookieName() bool`

HasCookieName returns a boolean if a field has been set.

### SetCookieNameNil

`func (o *SessionPersistence) SetCookieNameNil(b bool)`

 SetCookieNameNil sets the value for CookieName to be an explicit nil

### UnsetCookieName
`func (o *SessionPersistence) UnsetCookieName()`

UnsetCookieName ensures that no value is present for CookieName, not even an explicit nil
### GetPersistenceTimeout

`func (o *SessionPersistence) GetPersistenceTimeout() int32`

GetPersistenceTimeout returns the PersistenceTimeout field if non-nil, zero value otherwise.

### GetPersistenceTimeoutOk

`func (o *SessionPersistence) GetPersistenceTimeoutOk() (*int32, bool)`

GetPersistenceTimeoutOk returns a tuple with the PersistenceTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceTimeout

`func (o *SessionPersistence) SetPersistenceTimeout(v int32)`

SetPersistenceTimeout sets PersistenceTimeout field to given value.


### GetPersistenceGranularity

`func (o *SessionPersistence) GetPersistenceGranularity() string`

GetPersistenceGranularity returns the PersistenceGranularity field if non-nil, zero value otherwise.

### GetPersistenceGranularityOk

`func (o *SessionPersistence) GetPersistenceGranularityOk() (*string, bool)`

GetPersistenceGranularityOk returns a tuple with the PersistenceGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceGranularity

`func (o *SessionPersistence) SetPersistenceGranularity(v string)`

SetPersistenceGranularity sets PersistenceGranularity field to given value.

### HasPersistenceGranularity

`func (o *SessionPersistence) HasPersistenceGranularity() bool`

HasPersistenceGranularity returns a boolean if a field has been set.

### SetPersistenceGranularityNil

`func (o *SessionPersistence) SetPersistenceGranularityNil(b bool)`

 SetPersistenceGranularityNil sets the value for PersistenceGranularity to be an explicit nil

### UnsetPersistenceGranularity
`func (o *SessionPersistence) UnsetPersistenceGranularity()`

UnsetPersistenceGranularity ensures that no value is present for PersistenceGranularity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


