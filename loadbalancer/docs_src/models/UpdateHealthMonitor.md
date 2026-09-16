# UpdateHealthMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delay** | Pointer to **NullableInt32** | 헬스 체크 간격 (초) | [optional] 
**ExpectedCodes** | Pointer to **NullableString** | 정상으로 간주할 응답 코드 범위 (type&#x3D;HTTP/HTTPS) | [optional] 
**HttpMethod** | Pointer to [**NullableHealthMonitorMethod**](HealthMonitorMethod.md) | HTTP/HTTPS 헬스 체크 시 사용할 메서드 | [optional] 
**HttpVersion** | Pointer to [**NullableHealthMonitorHttpVersion**](HealthMonitorHttpVersion.md) | HTTP 프로토콜 버전 | [optional] 
**MaxRetries** | Pointer to **NullableInt32** | 실패 허용 횟수 | [optional] 
**MaxRetriesDown** | Pointer to **NullableInt32** | DOWN 상태로 간주하기 위한 연속 실패 횟수 | [optional] 
**Timeout** | Pointer to **NullableInt32** | 응답 대기 시간 (초) | [optional] 
**UrlPath** | Pointer to **NullableString** | 헬스 체크에 사용할 URL 경로 (type&#x3D;HTTP/HTTPS) | [optional] 

## Methods

### NewUpdateHealthMonitor

`func NewUpdateHealthMonitor() *UpdateHealthMonitor`

NewUpdateHealthMonitor instantiates a new UpdateHealthMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHealthMonitorWithDefaults

`func NewUpdateHealthMonitorWithDefaults() *UpdateHealthMonitor`

NewUpdateHealthMonitorWithDefaults instantiates a new UpdateHealthMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelay

`func (o *UpdateHealthMonitor) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *UpdateHealthMonitor) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *UpdateHealthMonitor) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *UpdateHealthMonitor) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### SetDelayNil

`func (o *UpdateHealthMonitor) SetDelayNil(b bool)`

 SetDelayNil sets the value for Delay to be an explicit nil

### UnsetDelay
`func (o *UpdateHealthMonitor) UnsetDelay()`

UnsetDelay ensures that no value is present for Delay, not even an explicit nil
### GetExpectedCodes

`func (o *UpdateHealthMonitor) GetExpectedCodes() string`

GetExpectedCodes returns the ExpectedCodes field if non-nil, zero value otherwise.

### GetExpectedCodesOk

`func (o *UpdateHealthMonitor) GetExpectedCodesOk() (*string, bool)`

GetExpectedCodesOk returns a tuple with the ExpectedCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedCodes

`func (o *UpdateHealthMonitor) SetExpectedCodes(v string)`

SetExpectedCodes sets ExpectedCodes field to given value.

### HasExpectedCodes

`func (o *UpdateHealthMonitor) HasExpectedCodes() bool`

HasExpectedCodes returns a boolean if a field has been set.

### SetExpectedCodesNil

`func (o *UpdateHealthMonitor) SetExpectedCodesNil(b bool)`

 SetExpectedCodesNil sets the value for ExpectedCodes to be an explicit nil

### UnsetExpectedCodes
`func (o *UpdateHealthMonitor) UnsetExpectedCodes()`

UnsetExpectedCodes ensures that no value is present for ExpectedCodes, not even an explicit nil
### GetHttpMethod

`func (o *UpdateHealthMonitor) GetHttpMethod() HealthMonitorMethod`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *UpdateHealthMonitor) GetHttpMethodOk() (*HealthMonitorMethod, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *UpdateHealthMonitor) SetHttpMethod(v HealthMonitorMethod)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *UpdateHealthMonitor) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### SetHttpMethodNil

`func (o *UpdateHealthMonitor) SetHttpMethodNil(b bool)`

 SetHttpMethodNil sets the value for HttpMethod to be an explicit nil

### UnsetHttpMethod
`func (o *UpdateHealthMonitor) UnsetHttpMethod()`

UnsetHttpMethod ensures that no value is present for HttpMethod, not even an explicit nil
### GetHttpVersion

`func (o *UpdateHealthMonitor) GetHttpVersion() HealthMonitorHttpVersion`

GetHttpVersion returns the HttpVersion field if non-nil, zero value otherwise.

### GetHttpVersionOk

`func (o *UpdateHealthMonitor) GetHttpVersionOk() (*HealthMonitorHttpVersion, bool)`

GetHttpVersionOk returns a tuple with the HttpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpVersion

`func (o *UpdateHealthMonitor) SetHttpVersion(v HealthMonitorHttpVersion)`

SetHttpVersion sets HttpVersion field to given value.

### HasHttpVersion

`func (o *UpdateHealthMonitor) HasHttpVersion() bool`

HasHttpVersion returns a boolean if a field has been set.

### SetHttpVersionNil

`func (o *UpdateHealthMonitor) SetHttpVersionNil(b bool)`

 SetHttpVersionNil sets the value for HttpVersion to be an explicit nil

### UnsetHttpVersion
`func (o *UpdateHealthMonitor) UnsetHttpVersion()`

UnsetHttpVersion ensures that no value is present for HttpVersion, not even an explicit nil
### GetMaxRetries

`func (o *UpdateHealthMonitor) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *UpdateHealthMonitor) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *UpdateHealthMonitor) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *UpdateHealthMonitor) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### SetMaxRetriesNil

`func (o *UpdateHealthMonitor) SetMaxRetriesNil(b bool)`

 SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil

### UnsetMaxRetries
`func (o *UpdateHealthMonitor) UnsetMaxRetries()`

UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
### GetMaxRetriesDown

`func (o *UpdateHealthMonitor) GetMaxRetriesDown() int32`

GetMaxRetriesDown returns the MaxRetriesDown field if non-nil, zero value otherwise.

### GetMaxRetriesDownOk

`func (o *UpdateHealthMonitor) GetMaxRetriesDownOk() (*int32, bool)`

GetMaxRetriesDownOk returns a tuple with the MaxRetriesDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetriesDown

`func (o *UpdateHealthMonitor) SetMaxRetriesDown(v int32)`

SetMaxRetriesDown sets MaxRetriesDown field to given value.

### HasMaxRetriesDown

`func (o *UpdateHealthMonitor) HasMaxRetriesDown() bool`

HasMaxRetriesDown returns a boolean if a field has been set.

### SetMaxRetriesDownNil

`func (o *UpdateHealthMonitor) SetMaxRetriesDownNil(b bool)`

 SetMaxRetriesDownNil sets the value for MaxRetriesDown to be an explicit nil

### UnsetMaxRetriesDown
`func (o *UpdateHealthMonitor) UnsetMaxRetriesDown()`

UnsetMaxRetriesDown ensures that no value is present for MaxRetriesDown, not even an explicit nil
### GetTimeout

`func (o *UpdateHealthMonitor) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *UpdateHealthMonitor) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *UpdateHealthMonitor) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *UpdateHealthMonitor) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *UpdateHealthMonitor) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *UpdateHealthMonitor) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetUrlPath

`func (o *UpdateHealthMonitor) GetUrlPath() string`

GetUrlPath returns the UrlPath field if non-nil, zero value otherwise.

### GetUrlPathOk

`func (o *UpdateHealthMonitor) GetUrlPathOk() (*string, bool)`

GetUrlPathOk returns a tuple with the UrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPath

`func (o *UpdateHealthMonitor) SetUrlPath(v string)`

SetUrlPath sets UrlPath field to given value.

### HasUrlPath

`func (o *UpdateHealthMonitor) HasUrlPath() bool`

HasUrlPath returns a boolean if a field has been set.

### SetUrlPathNil

`func (o *UpdateHealthMonitor) SetUrlPathNil(b bool)`

 SetUrlPathNil sets the value for UrlPath to be an explicit nil

### UnsetUrlPath
`func (o *UpdateHealthMonitor) UnsetUrlPath()`

UnsetUrlPath ensures that no value is present for UrlPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


