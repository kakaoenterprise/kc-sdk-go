# BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 보안 그룹 규칙 ID | 
**Direction** | [**SecurityGroupRuleDirection**](SecurityGroupRuleDirection.md) | 트래픽 방향 &lt;br/&gt; - &#x60;ingress&#x60;: 인바운드 (수신)   &lt;br/&gt; - &#x60;egress&#x60;: 아웃바운드 (송신) | 
**Protocol** | [**SecurityGroupRuleProtocol**](SecurityGroupRuleProtocol.md) | 프로토콜 종류 | 
**PortRangeMin** | Pointer to **NullableString** |  | [optional] 
**PortRangeMax** | Pointer to **NullableString** |  | [optional] 
**Description** | **string** | 규칙에 대한 설명 | 
**ProjectId** | **string** | 규칙이 속한 프로젝트 ID | 
**SecurityGroupId** | **string** | 보안 그룹의 고유 ID | 
**RemoteGroupId** | Pointer to **NullableString** |  | [optional] 
**RemoteIpPrefix** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 

## Methods

### NewBnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel

`func NewBnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel(id string, direction SecurityGroupRuleDirection, protocol SecurityGroupRuleProtocol, description string, projectId string, securityGroupId string, createdAt time.Time, updatedAt time.Time, ) *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel`

NewBnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel instantiates a new BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModelWithDefaults

`func NewBnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModelWithDefaults() *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel`

NewBnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModelWithDefaults instantiates a new BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) SetId(v string)`

SetId sets Id field to given value.


### GetDirection

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetDirection() SecurityGroupRuleDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetDirectionOk() (*SecurityGroupRuleDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) SetDirection(v SecurityGroupRuleDirection)`

SetDirection sets Direction field to given value.


### GetProtocol

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetProtocol() SecurityGroupRuleProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetProtocolOk() (*SecurityGroupRuleProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) SetProtocol(v SecurityGroupRuleProtocol)`

SetProtocol sets Protocol field to given value.


### GetPortRangeMin

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetPortRangeMin() string`

GetPortRangeMin returns the PortRangeMin field if non-nil, zero value otherwise.

### GetPortRangeMinOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetPortRangeMinOk() (*string, bool)`

GetPortRangeMinOk returns a tuple with the PortRangeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMin

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) SetPortRangeMin(v string)`

SetPortRangeMin sets PortRangeMin field to given value.

### HasPortRangeMin

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) HasPortRangeMin() bool`

HasPortRangeMin returns a boolean if a field has been set.

### SetPortRangeMinNil

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) SetPortRangeMinNil(b bool)`

 SetPortRangeMinNil sets the value for PortRangeMin to be an explicit nil

### UnsetPortRangeMin
`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) UnsetPortRangeMin()`

UnsetPortRangeMin ensures that no value is present for PortRangeMin, not even an explicit nil
### GetPortRangeMax

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetPortRangeMax() string`

GetPortRangeMax returns the PortRangeMax field if non-nil, zero value otherwise.

### GetPortRangeMaxOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetPortRangeMaxOk() (*string, bool)`

GetPortRangeMaxOk returns a tuple with the PortRangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMax

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) SetPortRangeMax(v string)`

SetPortRangeMax sets PortRangeMax field to given value.

### HasPortRangeMax

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) HasPortRangeMax() bool`

HasPortRangeMax returns a boolean if a field has been set.

### SetPortRangeMaxNil

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) SetPortRangeMaxNil(b bool)`

 SetPortRangeMaxNil sets the value for PortRangeMax to be an explicit nil

### UnsetPortRangeMax
`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) UnsetPortRangeMax()`

UnsetPortRangeMax ensures that no value is present for PortRangeMax, not even an explicit nil
### GetDescription

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProjectId

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSecurityGroupId

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.


### GetRemoteGroupId

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetRemoteGroupId() string`

GetRemoteGroupId returns the RemoteGroupId field if non-nil, zero value otherwise.

### GetRemoteGroupIdOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetRemoteGroupIdOk() (*string, bool)`

GetRemoteGroupIdOk returns a tuple with the RemoteGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGroupId

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) SetRemoteGroupId(v string)`

SetRemoteGroupId sets RemoteGroupId field to given value.

### HasRemoteGroupId

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) HasRemoteGroupId() bool`

HasRemoteGroupId returns a boolean if a field has been set.

### SetRemoteGroupIdNil

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) SetRemoteGroupIdNil(b bool)`

 SetRemoteGroupIdNil sets the value for RemoteGroupId to be an explicit nil

### UnsetRemoteGroupId
`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) UnsetRemoteGroupId()`

UnsetRemoteGroupId ensures that no value is present for RemoteGroupId, not even an explicit nil
### GetRemoteIpPrefix

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetRemoteIpPrefix() string`

GetRemoteIpPrefix returns the RemoteIpPrefix field if non-nil, zero value otherwise.

### GetRemoteIpPrefixOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetRemoteIpPrefixOk() (*string, bool)`

GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpPrefix

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) SetRemoteIpPrefix(v string)`

SetRemoteIpPrefix sets RemoteIpPrefix field to given value.

### HasRemoteIpPrefix

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) HasRemoteIpPrefix() bool`

HasRemoteIpPrefix returns a boolean if a field has been set.

### SetRemoteIpPrefixNil

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) SetRemoteIpPrefixNil(b bool)`

 SetRemoteIpPrefixNil sets the value for RemoteIpPrefix to be an explicit nil

### UnsetRemoteIpPrefix
`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) UnsetRemoteIpPrefix()`

UnsetRemoteIpPrefix ensures that no value is present for RemoteIpPrefix, not even an explicit nil
### GetCreatedAt

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsNetworkV1ApiCreateSecurityGroupRuleModelSecurityGroupRuleModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


