# EditHealthMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delay** | Pointer to **NullableInt32** |  | [optional] 
**ExpectedCodes** | Pointer to **string** | 정상으로 간주할 응답 코드 범위 (type&#x3D;HTTP/HTTPS) | [optional] 
**HttpMethod** | Pointer to [**HealthMonitorMethod**](HealthMonitorMethod.md) | HTTP/HTTPS 헬스 체크 시 사용할 메서드 &lt;br/&gt; - &#x60;CONNECT&#x60;: 서버와 터널 연결 시도 &lt;br/&gt; - &#x60;GET&#x60;: 리소스 조회 요청 (기본적으로 가장 많이 사용) &lt;br/&gt; - &#x60;POST&#x60;: 데이터 전송 요청 &lt;br/&gt; - &#x60;DELETE&#x60;: 리소스 삭제 요청 &lt;br/&gt; - &#x60;PATCH&#x60;: 리소스 일부 업데이트 요청 &lt;br/&gt; - &#x60;PUT&#x60;: 리소스 전체 업데이트 요청 &lt;br/&gt; - &#x60;HEAD&#x60;: 헤더만 조회 (본문 없음, 응답 속도 확인에 유용) &lt;br/&gt; - &#x60;OPTIONS&#x60;: 서버가 지원하는 메서드 확인 &lt;br/&gt; - &#x60;TRACE&#x60;: 요청/응답 루프백 테스트 | [optional] 
**HttpVersion** | Pointer to [**HealthMonitorHttpVersion**](HealthMonitorHttpVersion.md) | HTTP 프로토콜 버전 &lt;br/&gt; - &#x60;1.0&#x60;: HTTP/1.0 프로토콜&lt;br/&gt;- &#x60;1.1&#x60;: HTTP/1.1 프로토콜(기본적으로 가장 널리 사용되는 버전) | [optional] 
**MaxRetries** | Pointer to **NullableInt32** |  | [optional] 
**MaxRetriesDown** | Pointer to **NullableInt32** |  | [optional] 
**Timeout** | Pointer to **NullableInt32** |  | [optional] 
**UrlPath** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEditHealthMonitor

`func NewEditHealthMonitor() *EditHealthMonitor`

NewEditHealthMonitor instantiates a new EditHealthMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditHealthMonitorWithDefaults

`func NewEditHealthMonitorWithDefaults() *EditHealthMonitor`

NewEditHealthMonitorWithDefaults instantiates a new EditHealthMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelay

`func (o *EditHealthMonitor) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *EditHealthMonitor) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *EditHealthMonitor) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *EditHealthMonitor) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### SetDelayNil

`func (o *EditHealthMonitor) SetDelayNil(b bool)`

 SetDelayNil sets the value for Delay to be an explicit nil

### UnsetDelay
`func (o *EditHealthMonitor) UnsetDelay()`

UnsetDelay ensures that no value is present for Delay, not even an explicit nil
### GetExpectedCodes

`func (o *EditHealthMonitor) GetExpectedCodes() string`

GetExpectedCodes returns the ExpectedCodes field if non-nil, zero value otherwise.

### GetExpectedCodesOk

`func (o *EditHealthMonitor) GetExpectedCodesOk() (*string, bool)`

GetExpectedCodesOk returns a tuple with the ExpectedCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedCodes

`func (o *EditHealthMonitor) SetExpectedCodes(v string)`

SetExpectedCodes sets ExpectedCodes field to given value.

### HasExpectedCodes

`func (o *EditHealthMonitor) HasExpectedCodes() bool`

HasExpectedCodes returns a boolean if a field has been set.

### GetHttpMethod

`func (o *EditHealthMonitor) GetHttpMethod() HealthMonitorMethod`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *EditHealthMonitor) GetHttpMethodOk() (*HealthMonitorMethod, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *EditHealthMonitor) SetHttpMethod(v HealthMonitorMethod)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *EditHealthMonitor) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetHttpVersion

`func (o *EditHealthMonitor) GetHttpVersion() HealthMonitorHttpVersion`

GetHttpVersion returns the HttpVersion field if non-nil, zero value otherwise.

### GetHttpVersionOk

`func (o *EditHealthMonitor) GetHttpVersionOk() (*HealthMonitorHttpVersion, bool)`

GetHttpVersionOk returns a tuple with the HttpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpVersion

`func (o *EditHealthMonitor) SetHttpVersion(v HealthMonitorHttpVersion)`

SetHttpVersion sets HttpVersion field to given value.

### HasHttpVersion

`func (o *EditHealthMonitor) HasHttpVersion() bool`

HasHttpVersion returns a boolean if a field has been set.

### GetMaxRetries

`func (o *EditHealthMonitor) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *EditHealthMonitor) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *EditHealthMonitor) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *EditHealthMonitor) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### SetMaxRetriesNil

`func (o *EditHealthMonitor) SetMaxRetriesNil(b bool)`

 SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil

### UnsetMaxRetries
`func (o *EditHealthMonitor) UnsetMaxRetries()`

UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
### GetMaxRetriesDown

`func (o *EditHealthMonitor) GetMaxRetriesDown() int32`

GetMaxRetriesDown returns the MaxRetriesDown field if non-nil, zero value otherwise.

### GetMaxRetriesDownOk

`func (o *EditHealthMonitor) GetMaxRetriesDownOk() (*int32, bool)`

GetMaxRetriesDownOk returns a tuple with the MaxRetriesDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetriesDown

`func (o *EditHealthMonitor) SetMaxRetriesDown(v int32)`

SetMaxRetriesDown sets MaxRetriesDown field to given value.

### HasMaxRetriesDown

`func (o *EditHealthMonitor) HasMaxRetriesDown() bool`

HasMaxRetriesDown returns a boolean if a field has been set.

### SetMaxRetriesDownNil

`func (o *EditHealthMonitor) SetMaxRetriesDownNil(b bool)`

 SetMaxRetriesDownNil sets the value for MaxRetriesDown to be an explicit nil

### UnsetMaxRetriesDown
`func (o *EditHealthMonitor) UnsetMaxRetriesDown()`

UnsetMaxRetriesDown ensures that no value is present for MaxRetriesDown, not even an explicit nil
### GetTimeout

`func (o *EditHealthMonitor) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *EditHealthMonitor) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *EditHealthMonitor) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *EditHealthMonitor) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *EditHealthMonitor) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *EditHealthMonitor) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetUrlPath

`func (o *EditHealthMonitor) GetUrlPath() string`

GetUrlPath returns the UrlPath field if non-nil, zero value otherwise.

### GetUrlPathOk

`func (o *EditHealthMonitor) GetUrlPathOk() (*string, bool)`

GetUrlPathOk returns a tuple with the UrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPath

`func (o *EditHealthMonitor) SetUrlPath(v string)`

SetUrlPath sets UrlPath field to given value.

### HasUrlPath

`func (o *EditHealthMonitor) HasUrlPath() bool`

HasUrlPath returns a boolean if a field has been set.

### SetUrlPathNil

`func (o *EditHealthMonitor) SetUrlPathNil(b bool)`

 SetUrlPathNil sets the value for UrlPath to be an explicit nil

### UnsetUrlPath
`func (o *EditHealthMonitor) UnsetUrlPath()`

UnsetUrlPath ensures that no value is present for UrlPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


