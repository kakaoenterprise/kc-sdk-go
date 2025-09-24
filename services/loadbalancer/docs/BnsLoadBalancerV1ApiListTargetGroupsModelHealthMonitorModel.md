# BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel

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

### NewBnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel

`func NewBnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel(id string, ) *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel`

NewBnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel instantiates a new BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModelWithDefaults

`func NewBnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModelWithDefaults() *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel`

NewBnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModelWithDefaults instantiates a new BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelay

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### SetDelayNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetDelayNil(b bool)`

 SetDelayNil sets the value for Delay to be an explicit nil

### UnsetDelay
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) UnsetDelay()`

UnsetDelay ensures that no value is present for Delay, not even an explicit nil
### GetId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetId(v string)`

SetId sets Id field to given value.


### GetTimeout

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetExpectedCodes

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetExpectedCodes() string`

GetExpectedCodes returns the ExpectedCodes field if non-nil, zero value otherwise.

### GetExpectedCodesOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetExpectedCodesOk() (*string, bool)`

GetExpectedCodesOk returns a tuple with the ExpectedCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedCodes

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetExpectedCodes(v string)`

SetExpectedCodes sets ExpectedCodes field to given value.

### HasExpectedCodes

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) HasExpectedCodes() bool`

HasExpectedCodes returns a boolean if a field has been set.

### SetExpectedCodesNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetExpectedCodesNil(b bool)`

 SetExpectedCodesNil sets the value for ExpectedCodes to be an explicit nil

### UnsetExpectedCodes
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) UnsetExpectedCodes()`

UnsetExpectedCodes ensures that no value is present for ExpectedCodes, not even an explicit nil
### GetFallThreshold

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetFallThreshold() int32`

GetFallThreshold returns the FallThreshold field if non-nil, zero value otherwise.

### GetFallThresholdOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetFallThresholdOk() (*int32, bool)`

GetFallThresholdOk returns a tuple with the FallThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallThreshold

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetFallThreshold(v int32)`

SetFallThreshold sets FallThreshold field to given value.

### HasFallThreshold

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) HasFallThreshold() bool`

HasFallThreshold returns a boolean if a field has been set.

### SetFallThresholdNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetFallThresholdNil(b bool)`

 SetFallThresholdNil sets the value for FallThreshold to be an explicit nil

### UnsetFallThreshold
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) UnsetFallThreshold()`

UnsetFallThreshold ensures that no value is present for FallThreshold, not even an explicit nil
### GetHttpMethod

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetHttpMethod(v string)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### SetHttpMethodNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetHttpMethodNil(b bool)`

 SetHttpMethodNil sets the value for HttpMethod to be an explicit nil

### UnsetHttpMethod
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) UnsetHttpMethod()`

UnsetHttpMethod ensures that no value is present for HttpMethod, not even an explicit nil
### GetHttpVersion

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetHttpVersion() float32`

GetHttpVersion returns the HttpVersion field if non-nil, zero value otherwise.

### GetHttpVersionOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetHttpVersionOk() (*float32, bool)`

GetHttpVersionOk returns a tuple with the HttpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpVersion

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetHttpVersion(v float32)`

SetHttpVersion sets HttpVersion field to given value.

### HasHttpVersion

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) HasHttpVersion() bool`

HasHttpVersion returns a boolean if a field has been set.

### SetHttpVersionNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetHttpVersionNil(b bool)`

 SetHttpVersionNil sets the value for HttpVersion to be an explicit nil

### UnsetHttpVersion
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) UnsetHttpVersion()`

UnsetHttpVersion ensures that no value is present for HttpVersion, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetRiseThreshold

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetRiseThreshold() int32`

GetRiseThreshold returns the RiseThreshold field if non-nil, zero value otherwise.

### GetRiseThresholdOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetRiseThresholdOk() (*int32, bool)`

GetRiseThresholdOk returns a tuple with the RiseThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiseThreshold

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetRiseThreshold(v int32)`

SetRiseThreshold sets RiseThreshold field to given value.

### HasRiseThreshold

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) HasRiseThreshold() bool`

HasRiseThreshold returns a boolean if a field has been set.

### SetRiseThresholdNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetRiseThresholdNil(b bool)`

 SetRiseThresholdNil sets the value for RiseThreshold to be an explicit nil

### UnsetRiseThreshold
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) UnsetRiseThreshold()`

UnsetRiseThreshold ensures that no value is present for RiseThreshold, not even an explicit nil
### GetUrlPath

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetUrlPath() string`

GetUrlPath returns the UrlPath field if non-nil, zero value otherwise.

### GetUrlPathOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) GetUrlPathOk() (*string, bool)`

GetUrlPathOk returns a tuple with the UrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPath

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetUrlPath(v string)`

SetUrlPath sets UrlPath field to given value.

### HasUrlPath

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) HasUrlPath() bool`

HasUrlPath returns a boolean if a field has been set.

### SetUrlPathNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) SetUrlPathNil(b bool)`

 SetUrlPathNil sets the value for UrlPath to be an explicit nil

### UnsetUrlPath
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel) UnsetUrlPath()`

UnsetUrlPath ensures that no value is present for UrlPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


