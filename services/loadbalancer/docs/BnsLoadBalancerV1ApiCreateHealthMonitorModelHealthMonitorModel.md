# BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 헬스 모니터 ID | 
**Name** | **string** | 헬스 모니터 이름 | 
**Type** | [**HealthMonitorType**](HealthMonitorType.md) | 헬스 모니터 방식 | 
**Delay** | **int32** | 헬스 체크 간격 (초) | 
**Timeout** | **int32** | 응답 대기 시간 (초) | 
**MaxRetries** | **int32** | 성공으로 간주될 연속 성공 횟수 | 
**MaxRetriesDown** | **int32** | 실패로 간주될 연속 실패 횟수 | 
**HttpMethod** | Pointer to [**NullableHealthMonitorMethod**](HealthMonitorMethod.md) |  | [optional] 
**UrlPath** | Pointer to **NullableString** |  | [optional] 
**ExpectedCodes** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | **string** | 리소스가 속한 프로젝트 ID | 
**TargetGroups** | [**[]BnsLoadBalancerV1ApiCreateHealthMonitorModelTargetGroupModel**](BnsLoadBalancerV1ApiCreateHealthMonitorModelTargetGroupModel.md) | 연결된 대상 그룹 목록 | 
**ProvisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
**OperatingStatus** | [**LoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**HttpVersion** | Pointer to [**NullableHealthMonitorHttpVersion**](HealthMonitorHttpVersion.md) |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel

`func NewBnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel(id string, name string, type_ HealthMonitorType, delay int32, timeout int32, maxRetries int32, maxRetriesDown int32, projectId string, targetGroups []BnsLoadBalancerV1ApiCreateHealthMonitorModelTargetGroupModel, provisioningStatus ProvisioningStatus, operatingStatus LoadBalancerOperatingStatus, createdAt time.Time, ) *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel`

NewBnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel instantiates a new BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModelWithDefaults

`func NewBnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModelWithDefaults() *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel`

NewBnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModelWithDefaults instantiates a new BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetType() HealthMonitorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetTypeOk() (*HealthMonitorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetType(v HealthMonitorType)`

SetType sets Type field to given value.


### GetDelay

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetDelay(v int32)`

SetDelay sets Delay field to given value.


### GetTimeout

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetMaxRetries

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.


### GetMaxRetriesDown

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetMaxRetriesDown() int32`

GetMaxRetriesDown returns the MaxRetriesDown field if non-nil, zero value otherwise.

### GetMaxRetriesDownOk

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetMaxRetriesDownOk() (*int32, bool)`

GetMaxRetriesDownOk returns a tuple with the MaxRetriesDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetriesDown

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetMaxRetriesDown(v int32)`

SetMaxRetriesDown sets MaxRetriesDown field to given value.


### GetHttpMethod

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetHttpMethod() HealthMonitorMethod`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetHttpMethodOk() (*HealthMonitorMethod, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetHttpMethod(v HealthMonitorMethod)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### SetHttpMethodNil

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetHttpMethodNil(b bool)`

 SetHttpMethodNil sets the value for HttpMethod to be an explicit nil

### UnsetHttpMethod
`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) UnsetHttpMethod()`

UnsetHttpMethod ensures that no value is present for HttpMethod, not even an explicit nil
### GetUrlPath

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetUrlPath() string`

GetUrlPath returns the UrlPath field if non-nil, zero value otherwise.

### GetUrlPathOk

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetUrlPathOk() (*string, bool)`

GetUrlPathOk returns a tuple with the UrlPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPath

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetUrlPath(v string)`

SetUrlPath sets UrlPath field to given value.

### HasUrlPath

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) HasUrlPath() bool`

HasUrlPath returns a boolean if a field has been set.

### SetUrlPathNil

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetUrlPathNil(b bool)`

 SetUrlPathNil sets the value for UrlPath to be an explicit nil

### UnsetUrlPath
`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) UnsetUrlPath()`

UnsetUrlPath ensures that no value is present for UrlPath, not even an explicit nil
### GetExpectedCodes

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetExpectedCodes() string`

GetExpectedCodes returns the ExpectedCodes field if non-nil, zero value otherwise.

### GetExpectedCodesOk

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetExpectedCodesOk() (*string, bool)`

GetExpectedCodesOk returns a tuple with the ExpectedCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedCodes

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetExpectedCodes(v string)`

SetExpectedCodes sets ExpectedCodes field to given value.

### HasExpectedCodes

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) HasExpectedCodes() bool`

HasExpectedCodes returns a boolean if a field has been set.

### SetExpectedCodesNil

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetExpectedCodesNil(b bool)`

 SetExpectedCodesNil sets the value for ExpectedCodes to be an explicit nil

### UnsetExpectedCodes
`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) UnsetExpectedCodes()`

UnsetExpectedCodes ensures that no value is present for ExpectedCodes, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetTargetGroups

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetTargetGroups() []BnsLoadBalancerV1ApiCreateHealthMonitorModelTargetGroupModel`

GetTargetGroups returns the TargetGroups field if non-nil, zero value otherwise.

### GetTargetGroupsOk

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetTargetGroupsOk() (*[]BnsLoadBalancerV1ApiCreateHealthMonitorModelTargetGroupModel, bool)`

GetTargetGroupsOk returns a tuple with the TargetGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroups

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetTargetGroups(v []BnsLoadBalancerV1ApiCreateHealthMonitorModelTargetGroupModel)`

SetTargetGroups sets TargetGroups field to given value.


### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.


### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetHttpVersion

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetHttpVersion() HealthMonitorHttpVersion`

GetHttpVersion returns the HttpVersion field if non-nil, zero value otherwise.

### GetHttpVersionOk

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) GetHttpVersionOk() (*HealthMonitorHttpVersion, bool)`

GetHttpVersionOk returns a tuple with the HttpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpVersion

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetHttpVersion(v HealthMonitorHttpVersion)`

SetHttpVersion sets HttpVersion field to given value.

### HasHttpVersion

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) HasHttpVersion() bool`

HasHttpVersion returns a boolean if a field has been set.

### SetHttpVersionNil

`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) SetHttpVersionNil(b bool)`

 SetHttpVersionNil sets the value for HttpVersion to be an explicit nil

### UnsetHttpVersion
`func (o *BnsLoadBalancerV1ApiCreateHealthMonitorModelHealthMonitorModel) UnsetHttpVersion()`

UnsetHttpVersion ensures that no value is present for HttpVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


