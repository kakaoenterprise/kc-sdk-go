# BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delay** | Pointer to **NullableInt32** |  | [optional] 
**Id** | **string** | 헬스 모니터 ID | 
**Timeout** | Pointer to **NullableInt32** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**ExpectedCodes** | Pointer to **NullableString** |  | [optional] 
**FallThreshold** | Pointer to **NullableInt32** |  | [optional] 
**HttpMethod** | Pointer to **NullableString** |  | [optional] 
**HttpVersion** | Pointer to **NullableFloat32** |  | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**RiseThreshold** | Pointer to **NullableInt32** |  | [optional] 
**UrlPath** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel

`func NewBnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel(id string, ) *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel`

NewBnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel instantiates a new BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModelWithDefaults

`func NewBnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModelWithDefaults() *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel`

NewBnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModelWithDefaults instantiates a new BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelay

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### SetDelayNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetDelayNil(b bool)`

 SetDelayNil sets the value for Delay to be an explicit nil

### UnsetDelay
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) UnsetDelay()`

UnsetDelay ensures that no value is present for Delay, not even an explicit nil
### GetId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetId(v string)`

SetId sets Id field to given value.


### GetTimeout

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetExpectedCodes

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetExpectedCodes() string`

GetExpectedCodes returns the ExpectedCodes field if non-nil, zero value otherwise.

### GetExpectedCodesOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetExpectedCodesOk() (*string, bool)`

GetExpectedCodesOk returns a tuple with the ExpectedCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedCodes

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetExpectedCodes(v string)`

SetExpectedCodes sets ExpectedCodes field to given value.

### HasExpectedCodes

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) HasExpectedCodes() bool`

HasExpectedCodes returns a boolean if a field has been set.

### SetExpectedCodesNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetExpectedCodesNil(b bool)`

 SetExpectedCodesNil sets the value for ExpectedCodes to be an explicit nil

### UnsetExpectedCodes
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) UnsetExpectedCodes()`

UnsetExpectedCodes ensures that no value is present for ExpectedCodes, not even an explicit nil
### GetFallThreshold

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetFallThreshold() int32`

GetFallThreshold returns the FallThreshold field if non-nil, zero value otherwise.

### GetFallThresholdOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetFallThresholdOk() (*int32, bool)`

GetFallThresholdOk returns a tuple with the FallThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallThreshold

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetFallThreshold(v int32)`

SetFallThreshold sets FallThreshold field to given value.

### HasFallThreshold

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) HasFallThreshold() bool`

HasFallThreshold returns a boolean if a field has been set.

### SetFallThresholdNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetFallThresholdNil(b bool)`

 SetFallThresholdNil sets the value for FallThreshold to be an explicit nil

### UnsetFallThreshold
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) UnsetFallThreshold()`

UnsetFallThreshold ensures that no value is present for FallThreshold, not even an explicit nil
### GetHttpMethod

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetHttpMethod(v string)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### SetHttpMethodNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetHttpMethodNil(b bool)`

 SetHttpMethodNil sets the value for HttpMethod to be an explicit nil

### UnsetHttpMethod
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) UnsetHttpMethod()`

UnsetHttpMethod ensures that no value is present for HttpMethod, not even an explicit nil
### GetHttpVersion

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetHttpVersion() float32`

GetHttpVersion returns the HttpVersion field if non-nil, zero value otherwise.

### GetHttpVersionOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetHttpVersionOk() (*float32, bool)`

GetHttpVersionOk returns a tuple with the HttpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpVersion

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetHttpVersion(v float32)`

SetHttpVersion sets HttpVersion field to given value.

### HasHttpVersion

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) HasHttpVersion() bool`

HasHttpVersion returns a boolean if a field has been set.

### SetHttpVersionNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetHttpVersionNil(b bool)`

 SetHttpVersionNil sets the value for HttpVersion to be an explicit nil

### UnsetHttpVersion
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) UnsetHttpVersion()`

UnsetHttpVersion ensures that no value is present for HttpVersion, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetRiseThreshold

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetRiseThreshold() int32`

GetRiseThreshold returns the RiseThreshold field if non-nil, zero value otherwise.

### GetRiseThresholdOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetRiseThresholdOk() (*int32, bool)`

GetRiseThresholdOk returns a tuple with the RiseThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiseThreshold

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetRiseThreshold(v int32)`

SetRiseThreshold sets RiseThreshold field to given value.

### HasRiseThreshold

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) HasRiseThreshold() bool`

HasRiseThreshold returns a boolean if a field has been set.

### SetRiseThresholdNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetRiseThresholdNil(b bool)`

 SetRiseThresholdNil sets the value for RiseThreshold to be an explicit nil

### UnsetRiseThreshold
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) UnsetRiseThreshold()`

UnsetRiseThreshold ensures that no value is present for RiseThreshold, not even an explicit nil
### GetUrlPath

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetUrlPath() string`

GetUrlPath returns the UrlPath field if non-nil, zero value otherwise.

### GetUrlPathOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) GetUrlPathOk() (*string, bool)`

GetUrlPathOk returns a tuple with the UrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPath

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetUrlPath(v string)`

SetUrlPath sets UrlPath field to given value.

### HasUrlPath

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) HasUrlPath() bool`

HasUrlPath returns a boolean if a field has been set.

### SetUrlPathNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) SetUrlPathNil(b bool)`

 SetUrlPathNil sets the value for UrlPath to be an explicit nil

### UnsetUrlPath
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel) UnsetUrlPath()`

UnsetUrlPath ensures that no value is present for UrlPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


