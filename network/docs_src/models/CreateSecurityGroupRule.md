# CreateSecurityGroupRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | 보안 그룹에 대한 설명 | [optional] 
**Direction** | [**RuleDirection**](RuleDirection.md) | 트래픽 방향 | 
**Protocol** | [**RuleProtocol**](RuleProtocol.md) | 허용할 네트워크 프로토콜 | 
**RemoteIpPrefix** | Pointer to **NullableString** | CIDR 형식의 출발지 또는 목적지 IP 범위 (예: &#x60;0.0.0.0/0&#x60;) - 출발지/목적지를 IP로 설정하고자 할 경우 기입 | [optional] 
**RemoteGroupId** | Pointer to **NullableString** | 허용할 트래픽의 소스 또는 목적지 보안 그룹의 ID - 출발지/목적지를 다른 보안 그룹으로 지정할 경우 기입 - [List security groups](/openapi/networking/vpc/list-security-groups)에서 확인 | [optional] 
**PortRangeMin** | Pointer to **NullableInt32** | 허용할 포트 범위의 시작값 - &#x60;protocol&#x60; 이 &#x60;TCP&#x60;, &#x60;UDP&#x60;인 경우 입력 필요 - &#x60;protocol&#x60;이 &#x60;ALL&#x60;, &#x60;ICMP&#x60;인 경우 입력하지 않으면 &#x60;All&#x60;로 간주됨 | [optional] 
**PortRangeMax** | Pointer to **NullableInt32** | 허용할 포트 범위의 끝값 - &#x60;protocol&#x60;이 &#x60;TCP&#x60;, &#x60;UDP&#x60;인 경우 입력 필요 - &#x60;protocol&#x60;이 &#x60;ALL&#x60;, &#x60;ICMP&#x60;인 경우 입력하지 않으면 &#x60;All&#x60;로 간주됨 | [optional] 

## Methods

### NewCreateSecurityGroupRule

`func NewCreateSecurityGroupRule(direction RuleDirection, protocol RuleProtocol, ) *CreateSecurityGroupRule`

NewCreateSecurityGroupRule instantiates a new CreateSecurityGroupRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecurityGroupRuleWithDefaults

`func NewCreateSecurityGroupRuleWithDefaults() *CreateSecurityGroupRule`

NewCreateSecurityGroupRuleWithDefaults instantiates a new CreateSecurityGroupRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateSecurityGroupRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSecurityGroupRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSecurityGroupRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSecurityGroupRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateSecurityGroupRule) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateSecurityGroupRule) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDirection

`func (o *CreateSecurityGroupRule) GetDirection() RuleDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CreateSecurityGroupRule) GetDirectionOk() (*RuleDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CreateSecurityGroupRule) SetDirection(v RuleDirection)`

SetDirection sets Direction field to given value.


### GetProtocol

`func (o *CreateSecurityGroupRule) GetProtocol() RuleProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateSecurityGroupRule) GetProtocolOk() (*RuleProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateSecurityGroupRule) SetProtocol(v RuleProtocol)`

SetProtocol sets Protocol field to given value.


### GetRemoteIpPrefix

`func (o *CreateSecurityGroupRule) GetRemoteIpPrefix() string`

GetRemoteIpPrefix returns the RemoteIpPrefix field if non-nil, zero value otherwise.

### GetRemoteIpPrefixOk

`func (o *CreateSecurityGroupRule) GetRemoteIpPrefixOk() (*string, bool)`

GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpPrefix

`func (o *CreateSecurityGroupRule) SetRemoteIpPrefix(v string)`

SetRemoteIpPrefix sets RemoteIpPrefix field to given value.

### HasRemoteIpPrefix

`func (o *CreateSecurityGroupRule) HasRemoteIpPrefix() bool`

HasRemoteIpPrefix returns a boolean if a field has been set.

### SetRemoteIpPrefixNil

`func (o *CreateSecurityGroupRule) SetRemoteIpPrefixNil(b bool)`

 SetRemoteIpPrefixNil sets the value for RemoteIpPrefix to be an explicit nil

### UnsetRemoteIpPrefix
`func (o *CreateSecurityGroupRule) UnsetRemoteIpPrefix()`

UnsetRemoteIpPrefix ensures that no value is present for RemoteIpPrefix, not even an explicit nil
### GetRemoteGroupId

`func (o *CreateSecurityGroupRule) GetRemoteGroupId() string`

GetRemoteGroupId returns the RemoteGroupId field if non-nil, zero value otherwise.

### GetRemoteGroupIdOk

`func (o *CreateSecurityGroupRule) GetRemoteGroupIdOk() (*string, bool)`

GetRemoteGroupIdOk returns a tuple with the RemoteGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGroupId

`func (o *CreateSecurityGroupRule) SetRemoteGroupId(v string)`

SetRemoteGroupId sets RemoteGroupId field to given value.

### HasRemoteGroupId

`func (o *CreateSecurityGroupRule) HasRemoteGroupId() bool`

HasRemoteGroupId returns a boolean if a field has been set.

### SetRemoteGroupIdNil

`func (o *CreateSecurityGroupRule) SetRemoteGroupIdNil(b bool)`

 SetRemoteGroupIdNil sets the value for RemoteGroupId to be an explicit nil

### UnsetRemoteGroupId
`func (o *CreateSecurityGroupRule) UnsetRemoteGroupId()`

UnsetRemoteGroupId ensures that no value is present for RemoteGroupId, not even an explicit nil
### GetPortRangeMin

`func (o *CreateSecurityGroupRule) GetPortRangeMin() int32`

GetPortRangeMin returns the PortRangeMin field if non-nil, zero value otherwise.

### GetPortRangeMinOk

`func (o *CreateSecurityGroupRule) GetPortRangeMinOk() (*int32, bool)`

GetPortRangeMinOk returns a tuple with the PortRangeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMin

`func (o *CreateSecurityGroupRule) SetPortRangeMin(v int32)`

SetPortRangeMin sets PortRangeMin field to given value.

### HasPortRangeMin

`func (o *CreateSecurityGroupRule) HasPortRangeMin() bool`

HasPortRangeMin returns a boolean if a field has been set.

### SetPortRangeMinNil

`func (o *CreateSecurityGroupRule) SetPortRangeMinNil(b bool)`

 SetPortRangeMinNil sets the value for PortRangeMin to be an explicit nil

### UnsetPortRangeMin
`func (o *CreateSecurityGroupRule) UnsetPortRangeMin()`

UnsetPortRangeMin ensures that no value is present for PortRangeMin, not even an explicit nil
### GetPortRangeMax

`func (o *CreateSecurityGroupRule) GetPortRangeMax() int32`

GetPortRangeMax returns the PortRangeMax field if non-nil, zero value otherwise.

### GetPortRangeMaxOk

`func (o *CreateSecurityGroupRule) GetPortRangeMaxOk() (*int32, bool)`

GetPortRangeMaxOk returns a tuple with the PortRangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMax

`func (o *CreateSecurityGroupRule) SetPortRangeMax(v int32)`

SetPortRangeMax sets PortRangeMax field to given value.

### HasPortRangeMax

`func (o *CreateSecurityGroupRule) HasPortRangeMax() bool`

HasPortRangeMax returns a boolean if a field has been set.

### SetPortRangeMaxNil

`func (o *CreateSecurityGroupRule) SetPortRangeMaxNil(b bool)`

 SetPortRangeMaxNil sets the value for PortRangeMax to be an explicit nil

### UnsetPortRangeMax
`func (o *CreateSecurityGroupRule) UnsetPortRangeMax()`

UnsetPortRangeMax ensures that no value is present for PortRangeMax, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


