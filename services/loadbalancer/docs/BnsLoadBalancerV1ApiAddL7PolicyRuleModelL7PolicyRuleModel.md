# BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 생성된 L7 규칙의 ID | 
**Type** | [**L7RuleType**](L7RuleType.md) | 규칙을 검사할 대상 유형 | 
**CompareType** | [**L7RuleCompareType**](L7RuleCompareType.md) | 규칙 값 비교 방식 | 
**Key** | Pointer to **NullableString** |  | [optional] 
**Value** | **string** | 비교할 대상 값 | 
**IsInverted** | **bool** | 규칙의 비교 결과를 반전할지 여부 | 
**ProvisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
**OperatingStatus** | [**LoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel

`func NewBnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel(id string, type_ L7RuleType, compareType L7RuleCompareType, value string, isInverted bool, provisioningStatus ProvisioningStatus, operatingStatus LoadBalancerOperatingStatus, createdAt time.Time, ) *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel`

NewBnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel instantiates a new BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModelWithDefaults

`func NewBnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModelWithDefaults() *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel`

NewBnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModelWithDefaults instantiates a new BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetType() L7RuleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetTypeOk() (*L7RuleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) SetType(v L7RuleType)`

SetType sets Type field to given value.


### GetCompareType

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetCompareType() L7RuleCompareType`

GetCompareType returns the CompareType field if non-nil, zero value otherwise.

### GetCompareTypeOk

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetCompareTypeOk() (*L7RuleCompareType, bool)`

GetCompareTypeOk returns a tuple with the CompareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareType

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) SetCompareType(v L7RuleCompareType)`

SetCompareType sets CompareType field to given value.


### GetKey

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) SetValue(v string)`

SetValue sets Value field to given value.


### GetIsInverted

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetIsInverted() bool`

GetIsInverted returns the IsInverted field if non-nil, zero value otherwise.

### GetIsInvertedOk

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetIsInvertedOk() (*bool, bool)`

GetIsInvertedOk returns a tuple with the IsInverted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInverted

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) SetIsInverted(v bool)`

SetIsInverted sets IsInverted field to given value.


### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.


### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsLoadBalancerV1ApiAddL7PolicyRuleModelL7PolicyRuleModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


