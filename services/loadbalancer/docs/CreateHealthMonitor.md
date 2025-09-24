# CreateHealthMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delay** | **int32** | 헬스 체크 간 간격 (초 단위) | 
**MaxRetries** | **int32** | 성공으로 간주되기 위한 연속 성공 횟수 | 
**MaxRetriesDown** | **int32** | 다운으로 간주되기 위한 연속 실패 횟수 | 
**TargetGroupId** | **string** | 헬스 모니터를 연결할 대상 그룹 ID | 
**Timeout** | **int32** | 헬스 체크 응답 대기 시간 (초 단위) | 
**Type** | [**HealthMonitorType**](HealthMonitorType.md) | 헬스 모니터 방식 &lt;br/&gt; - &#x60;HTTP&#x60;: HTTP 방식 &lt;br/&gt; - &#x60;HTTPS&#x60;: HTTPS 방식 &lt;br/&gt; - &#x60;TCP&#x60;: TCP 방식 &lt;br/&gt; - &#x60;PING&#x60;: Ping 방식 | 
**HttpMethod** | Pointer to [**NullableHealthMonitorMethod**](HealthMonitorMethod.md) |  | [optional] 
**HttpVersion** | Pointer to [**NullableHealthMonitorHttpVersion**](HealthMonitorHttpVersion.md) |  | [optional] 
**ExpectedCodes** | Pointer to **NullableString** |  | [optional] 
**UrlPath** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateHealthMonitor

`func NewCreateHealthMonitor(delay int32, maxRetries int32, maxRetriesDown int32, targetGroupId string, timeout int32, type_ HealthMonitorType, ) *CreateHealthMonitor`

NewCreateHealthMonitor instantiates a new CreateHealthMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHealthMonitorWithDefaults

`func NewCreateHealthMonitorWithDefaults() *CreateHealthMonitor`

NewCreateHealthMonitorWithDefaults instantiates a new CreateHealthMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelay

`func (o *CreateHealthMonitor) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *CreateHealthMonitor) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *CreateHealthMonitor) SetDelay(v int32)`

SetDelay sets Delay field to given value.


### GetMaxRetries

`func (o *CreateHealthMonitor) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *CreateHealthMonitor) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *CreateHealthMonitor) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.


### GetMaxRetriesDown

`func (o *CreateHealthMonitor) GetMaxRetriesDown() int32`

GetMaxRetriesDown returns the MaxRetriesDown field if non-nil, zero value otherwise.

### GetMaxRetriesDownOk

`func (o *CreateHealthMonitor) GetMaxRetriesDownOk() (*int32, bool)`

GetMaxRetriesDownOk returns a tuple with the MaxRetriesDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetriesDown

`func (o *CreateHealthMonitor) SetMaxRetriesDown(v int32)`

SetMaxRetriesDown sets MaxRetriesDown field to given value.


### GetTargetGroupId

`func (o *CreateHealthMonitor) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *CreateHealthMonitor) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *CreateHealthMonitor) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.


### GetTimeout

`func (o *CreateHealthMonitor) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CreateHealthMonitor) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CreateHealthMonitor) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetType

`func (o *CreateHealthMonitor) GetType() HealthMonitorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateHealthMonitor) GetTypeOk() (*HealthMonitorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateHealthMonitor) SetType(v HealthMonitorType)`

SetType sets Type field to given value.


### GetHttpMethod

`func (o *CreateHealthMonitor) GetHttpMethod() HealthMonitorMethod`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *CreateHealthMonitor) GetHttpMethodOk() (*HealthMonitorMethod, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *CreateHealthMonitor) SetHttpMethod(v HealthMonitorMethod)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *CreateHealthMonitor) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### SetHttpMethodNil

`func (o *CreateHealthMonitor) SetHttpMethodNil(b bool)`

 SetHttpMethodNil sets the value for HttpMethod to be an explicit nil

### UnsetHttpMethod
`func (o *CreateHealthMonitor) UnsetHttpMethod()`

UnsetHttpMethod ensures that no value is present for HttpMethod, not even an explicit nil
### GetHttpVersion

`func (o *CreateHealthMonitor) GetHttpVersion() HealthMonitorHttpVersion`

GetHttpVersion returns the HttpVersion field if non-nil, zero value otherwise.

### GetHttpVersionOk

`func (o *CreateHealthMonitor) GetHttpVersionOk() (*HealthMonitorHttpVersion, bool)`

GetHttpVersionOk returns a tuple with the HttpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpVersion

`func (o *CreateHealthMonitor) SetHttpVersion(v HealthMonitorHttpVersion)`

SetHttpVersion sets HttpVersion field to given value.

### HasHttpVersion

`func (o *CreateHealthMonitor) HasHttpVersion() bool`

HasHttpVersion returns a boolean if a field has been set.

### SetHttpVersionNil

`func (o *CreateHealthMonitor) SetHttpVersionNil(b bool)`

 SetHttpVersionNil sets the value for HttpVersion to be an explicit nil

### UnsetHttpVersion
`func (o *CreateHealthMonitor) UnsetHttpVersion()`

UnsetHttpVersion ensures that no value is present for HttpVersion, not even an explicit nil
### GetExpectedCodes

`func (o *CreateHealthMonitor) GetExpectedCodes() string`

GetExpectedCodes returns the ExpectedCodes field if non-nil, zero value otherwise.

### GetExpectedCodesOk

`func (o *CreateHealthMonitor) GetExpectedCodesOk() (*string, bool)`

GetExpectedCodesOk returns a tuple with the ExpectedCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedCodes

`func (o *CreateHealthMonitor) SetExpectedCodes(v string)`

SetExpectedCodes sets ExpectedCodes field to given value.

### HasExpectedCodes

`func (o *CreateHealthMonitor) HasExpectedCodes() bool`

HasExpectedCodes returns a boolean if a field has been set.

### SetExpectedCodesNil

`func (o *CreateHealthMonitor) SetExpectedCodesNil(b bool)`

 SetExpectedCodesNil sets the value for ExpectedCodes to be an explicit nil

### UnsetExpectedCodes
`func (o *CreateHealthMonitor) UnsetExpectedCodes()`

UnsetExpectedCodes ensures that no value is present for ExpectedCodes, not even an explicit nil
### GetUrlPath

`func (o *CreateHealthMonitor) GetUrlPath() string`

GetUrlPath returns the UrlPath field if non-nil, zero value otherwise.

### GetUrlPathOk

`func (o *CreateHealthMonitor) GetUrlPathOk() (*string, bool)`

GetUrlPathOk returns a tuple with the UrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPath

`func (o *CreateHealthMonitor) SetUrlPath(v string)`

SetUrlPath sets UrlPath field to given value.

### HasUrlPath

`func (o *CreateHealthMonitor) HasUrlPath() bool`

HasUrlPath returns a boolean if a field has been set.

### SetUrlPathNil

`func (o *CreateHealthMonitor) SetUrlPathNil(b bool)`

 SetUrlPathNil sets the value for UrlPath to be an explicit nil

### UnsetUrlPath
`func (o *CreateHealthMonitor) UnsetUrlPath()`

UnsetUrlPath ensures that no value is present for UrlPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


