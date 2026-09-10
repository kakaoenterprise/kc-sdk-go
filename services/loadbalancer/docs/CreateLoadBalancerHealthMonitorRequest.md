# CreateLoadBalancerHealthMonitorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**HealthMonitorType**](HealthMonitorType.md) | 헬스 체크 방식 | 
**Delay** | **int32** | 헬스 체크 간격 (초) | 
**MaxRetries** | **int32** | 실패 허용 횟수 | 
**Timeout** | **int32** | 응답 대기 시간 (초) | 
**HttpMethod** | Pointer to [**NullableHealthMonitorMethod**](HealthMonitorMethod.md) |  | [optional] 
**HttpVersion** | Pointer to [**NullableHealthMonitorHttpVersion**](HealthMonitorHttpVersion.md) |  | [optional] 
**ExpectedCodes** | Pointer to **NullableString** |  | [optional] 
**UrlPath** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateLoadBalancerHealthMonitorRequest

`func NewCreateLoadBalancerHealthMonitorRequest(type_ HealthMonitorType, delay int32, maxRetries int32, timeout int32, ) *CreateLoadBalancerHealthMonitorRequest`

NewCreateLoadBalancerHealthMonitorRequest instantiates a new CreateLoadBalancerHealthMonitorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerHealthMonitorRequestWithDefaults

`func NewCreateLoadBalancerHealthMonitorRequestWithDefaults() *CreateLoadBalancerHealthMonitorRequest`

NewCreateLoadBalancerHealthMonitorRequestWithDefaults instantiates a new CreateLoadBalancerHealthMonitorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateLoadBalancerHealthMonitorRequest) GetType() HealthMonitorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateLoadBalancerHealthMonitorRequest) GetTypeOk() (*HealthMonitorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateLoadBalancerHealthMonitorRequest) SetType(v HealthMonitorType)`

SetType sets Type field to given value.


### GetDelay

`func (o *CreateLoadBalancerHealthMonitorRequest) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *CreateLoadBalancerHealthMonitorRequest) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *CreateLoadBalancerHealthMonitorRequest) SetDelay(v int32)`

SetDelay sets Delay field to given value.


### GetMaxRetries

`func (o *CreateLoadBalancerHealthMonitorRequest) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *CreateLoadBalancerHealthMonitorRequest) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *CreateLoadBalancerHealthMonitorRequest) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.


### GetTimeout

`func (o *CreateLoadBalancerHealthMonitorRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CreateLoadBalancerHealthMonitorRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CreateLoadBalancerHealthMonitorRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetHttpMethod

`func (o *CreateLoadBalancerHealthMonitorRequest) GetHttpMethod() HealthMonitorMethod`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *CreateLoadBalancerHealthMonitorRequest) GetHttpMethodOk() (*HealthMonitorMethod, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *CreateLoadBalancerHealthMonitorRequest) SetHttpMethod(v HealthMonitorMethod)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *CreateLoadBalancerHealthMonitorRequest) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### SetHttpMethodNil

`func (o *CreateLoadBalancerHealthMonitorRequest) SetHttpMethodNil(b bool)`

 SetHttpMethodNil sets the value for HttpMethod to be an explicit nil

### UnsetHttpMethod
`func (o *CreateLoadBalancerHealthMonitorRequest) UnsetHttpMethod()`

UnsetHttpMethod ensures that no value is present for HttpMethod, not even an explicit nil
### GetHttpVersion

`func (o *CreateLoadBalancerHealthMonitorRequest) GetHttpVersion() HealthMonitorHttpVersion`

GetHttpVersion returns the HttpVersion field if non-nil, zero value otherwise.

### GetHttpVersionOk

`func (o *CreateLoadBalancerHealthMonitorRequest) GetHttpVersionOk() (*HealthMonitorHttpVersion, bool)`

GetHttpVersionOk returns a tuple with the HttpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpVersion

`func (o *CreateLoadBalancerHealthMonitorRequest) SetHttpVersion(v HealthMonitorHttpVersion)`

SetHttpVersion sets HttpVersion field to given value.

### HasHttpVersion

`func (o *CreateLoadBalancerHealthMonitorRequest) HasHttpVersion() bool`

HasHttpVersion returns a boolean if a field has been set.

### SetHttpVersionNil

`func (o *CreateLoadBalancerHealthMonitorRequest) SetHttpVersionNil(b bool)`

 SetHttpVersionNil sets the value for HttpVersion to be an explicit nil

### UnsetHttpVersion
`func (o *CreateLoadBalancerHealthMonitorRequest) UnsetHttpVersion()`

UnsetHttpVersion ensures that no value is present for HttpVersion, not even an explicit nil
### GetExpectedCodes

`func (o *CreateLoadBalancerHealthMonitorRequest) GetExpectedCodes() string`

GetExpectedCodes returns the ExpectedCodes field if non-nil, zero value otherwise.

### GetExpectedCodesOk

`func (o *CreateLoadBalancerHealthMonitorRequest) GetExpectedCodesOk() (*string, bool)`

GetExpectedCodesOk returns a tuple with the ExpectedCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedCodes

`func (o *CreateLoadBalancerHealthMonitorRequest) SetExpectedCodes(v string)`

SetExpectedCodes sets ExpectedCodes field to given value.

### HasExpectedCodes

`func (o *CreateLoadBalancerHealthMonitorRequest) HasExpectedCodes() bool`

HasExpectedCodes returns a boolean if a field has been set.

### SetExpectedCodesNil

`func (o *CreateLoadBalancerHealthMonitorRequest) SetExpectedCodesNil(b bool)`

 SetExpectedCodesNil sets the value for ExpectedCodes to be an explicit nil

### UnsetExpectedCodes
`func (o *CreateLoadBalancerHealthMonitorRequest) UnsetExpectedCodes()`

UnsetExpectedCodes ensures that no value is present for ExpectedCodes, not even an explicit nil
### GetUrlPath

`func (o *CreateLoadBalancerHealthMonitorRequest) GetUrlPath() string`

GetUrlPath returns the UrlPath field if non-nil, zero value otherwise.

### GetUrlPathOk

`func (o *CreateLoadBalancerHealthMonitorRequest) GetUrlPathOk() (*string, bool)`

GetUrlPathOk returns a tuple with the UrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPath

`func (o *CreateLoadBalancerHealthMonitorRequest) SetUrlPath(v string)`

SetUrlPath sets UrlPath field to given value.

### HasUrlPath

`func (o *CreateLoadBalancerHealthMonitorRequest) HasUrlPath() bool`

HasUrlPath returns a boolean if a field has been set.

### SetUrlPathNil

`func (o *CreateLoadBalancerHealthMonitorRequest) SetUrlPathNil(b bool)`

 SetUrlPathNil sets the value for UrlPath to be an explicit nil

### UnsetUrlPath
`func (o *CreateLoadBalancerHealthMonitorRequest) UnsetUrlPath()`

UnsetUrlPath ensures that no value is present for UrlPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


