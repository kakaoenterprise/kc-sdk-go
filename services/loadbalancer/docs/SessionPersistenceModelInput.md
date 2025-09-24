# SessionPersistenceModelInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | 세션 유형 &lt;br/&gt; - 예시: &#x60;APP_COOKIE&#x60;, &#x60;HTTP_COOKIE&#x60;, &#x60;SOURCE_IP&#x60; | 
**CookieName** | **string** | 쿠키 이름 (세션 유형이 &#x60;APP_COOKIE&#x60;, &#x60;HTTP_COOKIE&#x60;인 경우) | 
**PersistenceTimeout** | **int32** | 세션 유지 시간 (초) | 
**PersistenceGranularity** | **string** | 세션 유지 기준이 되는 IP 서브넷 마스크 (세션 유형이 &#x60;SOURCE_IP&#x60;인 경우) | 

## Methods

### NewSessionPersistenceModelInput

`func NewSessionPersistenceModelInput(type_ string, cookieName string, persistenceTimeout int32, persistenceGranularity string, ) *SessionPersistenceModelInput`

NewSessionPersistenceModelInput instantiates a new SessionPersistenceModelInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionPersistenceModelInputWithDefaults

`func NewSessionPersistenceModelInputWithDefaults() *SessionPersistenceModelInput`

NewSessionPersistenceModelInputWithDefaults instantiates a new SessionPersistenceModelInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SessionPersistenceModelInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SessionPersistenceModelInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SessionPersistenceModelInput) SetType(v string)`

SetType sets Type field to given value.


### GetCookieName

`func (o *SessionPersistenceModelInput) GetCookieName() string`

GetCookieName returns the CookieName field if non-nil, zero value otherwise.

### GetCookieNameOk

`func (o *SessionPersistenceModelInput) GetCookieNameOk() (*string, bool)`

GetCookieNameOk returns a tuple with the CookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieName

`func (o *SessionPersistenceModelInput) SetCookieName(v string)`

SetCookieName sets CookieName field to given value.


### GetPersistenceTimeout

`func (o *SessionPersistenceModelInput) GetPersistenceTimeout() int32`

GetPersistenceTimeout returns the PersistenceTimeout field if non-nil, zero value otherwise.

### GetPersistenceTimeoutOk

`func (o *SessionPersistenceModelInput) GetPersistenceTimeoutOk() (*int32, bool)`

GetPersistenceTimeoutOk returns a tuple with the PersistenceTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceTimeout

`func (o *SessionPersistenceModelInput) SetPersistenceTimeout(v int32)`

SetPersistenceTimeout sets PersistenceTimeout field to given value.


### GetPersistenceGranularity

`func (o *SessionPersistenceModelInput) GetPersistenceGranularity() string`

GetPersistenceGranularity returns the PersistenceGranularity field if non-nil, zero value otherwise.

### GetPersistenceGranularityOk

`func (o *SessionPersistenceModelInput) GetPersistenceGranularityOk() (*string, bool)`

GetPersistenceGranularityOk returns a tuple with the PersistenceGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceGranularity

`func (o *SessionPersistenceModelInput) SetPersistenceGranularity(v string)`

SetPersistenceGranularity sets PersistenceGranularity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


