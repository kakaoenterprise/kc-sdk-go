# TargetGroupHealthMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delay** | Pointer to **NullableInt32** | 헬스 체크 간격 (초) | [optional] 
**Id** | Pointer to **NullableString** | 헬스 모니터 ID | [optional] 
**Timeout** | Pointer to **NullableInt32** | 응답 대기 시간 (초) | [optional] 
**Type** | Pointer to [**NullableHealthMonitorType**](HealthMonitorType.md) | 헬스 체크 방식 (예: HTTP, TCP 등) | [optional] 
**ExpectedCodes** | Pointer to **NullableString** | 정상으로 간주할 응답 코드 범위 (type&#x3D;HTTP/HTTPS) | [optional] 
**FallThreshold** | Pointer to **NullableInt32** | 연속 실패 허용 횟수 | [optional] 
**HttpMethod** | Pointer to [**NullableHealthMonitorMethod**](HealthMonitorMethod.md) | HTTP 요청 방식 (예: GET, HEAD) | [optional] 
**HttpVersion** | Pointer to **NullableFloat32** | HTTP 버전 (예: 1.0, 1.1) | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 | [optional] 
**ProjectId** | Pointer to **NullableString** | 프로젝트 ID | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | [optional] 
**RiseThreshold** | Pointer to **NullableInt32** | 연속 성공 기준 횟수 | [optional] 
**UrlPath** | Pointer to **NullableString** | 요청 경로 | [optional] 

## Methods

### NewTargetGroupHealthMonitor

`func NewTargetGroupHealthMonitor() *TargetGroupHealthMonitor`

NewTargetGroupHealthMonitor instantiates a new TargetGroupHealthMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetGroupHealthMonitorWithDefaults

`func NewTargetGroupHealthMonitorWithDefaults() *TargetGroupHealthMonitor`

NewTargetGroupHealthMonitorWithDefaults instantiates a new TargetGroupHealthMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelay

`func (o *TargetGroupHealthMonitor) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *TargetGroupHealthMonitor) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *TargetGroupHealthMonitor) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *TargetGroupHealthMonitor) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### SetDelayNil

`func (o *TargetGroupHealthMonitor) SetDelayNil(b bool)`

 SetDelayNil sets the value for Delay to be an explicit nil

### UnsetDelay
`func (o *TargetGroupHealthMonitor) UnsetDelay()`

UnsetDelay ensures that no value is present for Delay, not even an explicit nil
### GetId

`func (o *TargetGroupHealthMonitor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TargetGroupHealthMonitor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TargetGroupHealthMonitor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TargetGroupHealthMonitor) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TargetGroupHealthMonitor) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TargetGroupHealthMonitor) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimeout

`func (o *TargetGroupHealthMonitor) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TargetGroupHealthMonitor) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TargetGroupHealthMonitor) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TargetGroupHealthMonitor) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *TargetGroupHealthMonitor) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *TargetGroupHealthMonitor) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *TargetGroupHealthMonitor) GetType() HealthMonitorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TargetGroupHealthMonitor) GetTypeOk() (*HealthMonitorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TargetGroupHealthMonitor) SetType(v HealthMonitorType)`

SetType sets Type field to given value.

### HasType

`func (o *TargetGroupHealthMonitor) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *TargetGroupHealthMonitor) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TargetGroupHealthMonitor) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetExpectedCodes

`func (o *TargetGroupHealthMonitor) GetExpectedCodes() string`

GetExpectedCodes returns the ExpectedCodes field if non-nil, zero value otherwise.

### GetExpectedCodesOk

`func (o *TargetGroupHealthMonitor) GetExpectedCodesOk() (*string, bool)`

GetExpectedCodesOk returns a tuple with the ExpectedCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedCodes

`func (o *TargetGroupHealthMonitor) SetExpectedCodes(v string)`

SetExpectedCodes sets ExpectedCodes field to given value.

### HasExpectedCodes

`func (o *TargetGroupHealthMonitor) HasExpectedCodes() bool`

HasExpectedCodes returns a boolean if a field has been set.

### SetExpectedCodesNil

`func (o *TargetGroupHealthMonitor) SetExpectedCodesNil(b bool)`

 SetExpectedCodesNil sets the value for ExpectedCodes to be an explicit nil

### UnsetExpectedCodes
`func (o *TargetGroupHealthMonitor) UnsetExpectedCodes()`

UnsetExpectedCodes ensures that no value is present for ExpectedCodes, not even an explicit nil
### GetFallThreshold

`func (o *TargetGroupHealthMonitor) GetFallThreshold() int32`

GetFallThreshold returns the FallThreshold field if non-nil, zero value otherwise.

### GetFallThresholdOk

`func (o *TargetGroupHealthMonitor) GetFallThresholdOk() (*int32, bool)`

GetFallThresholdOk returns a tuple with the FallThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallThreshold

`func (o *TargetGroupHealthMonitor) SetFallThreshold(v int32)`

SetFallThreshold sets FallThreshold field to given value.

### HasFallThreshold

`func (o *TargetGroupHealthMonitor) HasFallThreshold() bool`

HasFallThreshold returns a boolean if a field has been set.

### SetFallThresholdNil

`func (o *TargetGroupHealthMonitor) SetFallThresholdNil(b bool)`

 SetFallThresholdNil sets the value for FallThreshold to be an explicit nil

### UnsetFallThreshold
`func (o *TargetGroupHealthMonitor) UnsetFallThreshold()`

UnsetFallThreshold ensures that no value is present for FallThreshold, not even an explicit nil
### GetHttpMethod

`func (o *TargetGroupHealthMonitor) GetHttpMethod() HealthMonitorMethod`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *TargetGroupHealthMonitor) GetHttpMethodOk() (*HealthMonitorMethod, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *TargetGroupHealthMonitor) SetHttpMethod(v HealthMonitorMethod)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *TargetGroupHealthMonitor) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### SetHttpMethodNil

`func (o *TargetGroupHealthMonitor) SetHttpMethodNil(b bool)`

 SetHttpMethodNil sets the value for HttpMethod to be an explicit nil

### UnsetHttpMethod
`func (o *TargetGroupHealthMonitor) UnsetHttpMethod()`

UnsetHttpMethod ensures that no value is present for HttpMethod, not even an explicit nil
### GetHttpVersion

`func (o *TargetGroupHealthMonitor) GetHttpVersion() float32`

GetHttpVersion returns the HttpVersion field if non-nil, zero value otherwise.

### GetHttpVersionOk

`func (o *TargetGroupHealthMonitor) GetHttpVersionOk() (*float32, bool)`

GetHttpVersionOk returns a tuple with the HttpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpVersion

`func (o *TargetGroupHealthMonitor) SetHttpVersion(v float32)`

SetHttpVersion sets HttpVersion field to given value.

### HasHttpVersion

`func (o *TargetGroupHealthMonitor) HasHttpVersion() bool`

HasHttpVersion returns a boolean if a field has been set.

### SetHttpVersionNil

`func (o *TargetGroupHealthMonitor) SetHttpVersionNil(b bool)`

 SetHttpVersionNil sets the value for HttpVersion to be an explicit nil

### UnsetHttpVersion
`func (o *TargetGroupHealthMonitor) UnsetHttpVersion()`

UnsetHttpVersion ensures that no value is present for HttpVersion, not even an explicit nil
### GetOperatingStatus

`func (o *TargetGroupHealthMonitor) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *TargetGroupHealthMonitor) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *TargetGroupHealthMonitor) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *TargetGroupHealthMonitor) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *TargetGroupHealthMonitor) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *TargetGroupHealthMonitor) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetProjectId

`func (o *TargetGroupHealthMonitor) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TargetGroupHealthMonitor) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TargetGroupHealthMonitor) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *TargetGroupHealthMonitor) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *TargetGroupHealthMonitor) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *TargetGroupHealthMonitor) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProvisioningStatus

`func (o *TargetGroupHealthMonitor) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *TargetGroupHealthMonitor) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *TargetGroupHealthMonitor) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *TargetGroupHealthMonitor) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *TargetGroupHealthMonitor) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *TargetGroupHealthMonitor) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetRiseThreshold

`func (o *TargetGroupHealthMonitor) GetRiseThreshold() int32`

GetRiseThreshold returns the RiseThreshold field if non-nil, zero value otherwise.

### GetRiseThresholdOk

`func (o *TargetGroupHealthMonitor) GetRiseThresholdOk() (*int32, bool)`

GetRiseThresholdOk returns a tuple with the RiseThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiseThreshold

`func (o *TargetGroupHealthMonitor) SetRiseThreshold(v int32)`

SetRiseThreshold sets RiseThreshold field to given value.

### HasRiseThreshold

`func (o *TargetGroupHealthMonitor) HasRiseThreshold() bool`

HasRiseThreshold returns a boolean if a field has been set.

### SetRiseThresholdNil

`func (o *TargetGroupHealthMonitor) SetRiseThresholdNil(b bool)`

 SetRiseThresholdNil sets the value for RiseThreshold to be an explicit nil

### UnsetRiseThreshold
`func (o *TargetGroupHealthMonitor) UnsetRiseThreshold()`

UnsetRiseThreshold ensures that no value is present for RiseThreshold, not even an explicit nil
### GetUrlPath

`func (o *TargetGroupHealthMonitor) GetUrlPath() string`

GetUrlPath returns the UrlPath field if non-nil, zero value otherwise.

### GetUrlPathOk

`func (o *TargetGroupHealthMonitor) GetUrlPathOk() (*string, bool)`

GetUrlPathOk returns a tuple with the UrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPath

`func (o *TargetGroupHealthMonitor) SetUrlPath(v string)`

SetUrlPath sets UrlPath field to given value.

### HasUrlPath

`func (o *TargetGroupHealthMonitor) HasUrlPath() bool`

HasUrlPath returns a boolean if a field has been set.

### SetUrlPathNil

`func (o *TargetGroupHealthMonitor) SetUrlPathNil(b bool)`

 SetUrlPathNil sets the value for UrlPath to be an explicit nil

### UnsetUrlPath
`func (o *TargetGroupHealthMonitor) UnsetUrlPath()`

UnsetUrlPath ensures that no value is present for UrlPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


