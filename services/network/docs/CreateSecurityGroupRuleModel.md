# CreateSecurityGroupRuleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Direction** | [**SecurityGroupRuleDirection**](SecurityGroupRuleDirection.md) | 트래픽 방향 &lt;br/&gt; - &#x60;ingress&#x60;: 인바운드 (수신)   &lt;br/&gt; - &#x60;egress&#x60;: 아웃바운드 (송신) | 
**Protocol** | [**SecurityGroupRuleProtocol**](SecurityGroupRuleProtocol.md) | 허용할 네트워크 프로토콜 &lt;br/&gt; - &#x60;TCP&#x60;: 전송 제어 프로토콜 &lt;br/&gt; -  &#x60;UDP&#x60;: 사용자 데이터그램 프로토콜 &lt;br/&gt; - &#x60;ICMP&#x60;: 인터넷 제어 메시지 프로토콜 &lt;br/&gt; -  &#x60;IPIP&#x60;: IP-in-IP 터널링 &lt;br/&gt; - &#x60;ALL&#x60;: 모든 프로토콜 | 
**PortRangeMin** | Pointer to **NullableInt32** |  | [optional] 
**PortRangeMax** | Pointer to **NullableInt32** |  | [optional] 
**RemoteIpPrefix** | Pointer to **NullableString** |  | [optional] 
**RemoteGroupId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateSecurityGroupRuleModel

`func NewCreateSecurityGroupRuleModel(direction SecurityGroupRuleDirection, protocol SecurityGroupRuleProtocol, ) *CreateSecurityGroupRuleModel`

NewCreateSecurityGroupRuleModel instantiates a new CreateSecurityGroupRuleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecurityGroupRuleModelWithDefaults

`func NewCreateSecurityGroupRuleModelWithDefaults() *CreateSecurityGroupRuleModel`

NewCreateSecurityGroupRuleModelWithDefaults instantiates a new CreateSecurityGroupRuleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateSecurityGroupRuleModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSecurityGroupRuleModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSecurityGroupRuleModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSecurityGroupRuleModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateSecurityGroupRuleModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateSecurityGroupRuleModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDirection

`func (o *CreateSecurityGroupRuleModel) GetDirection() SecurityGroupRuleDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CreateSecurityGroupRuleModel) GetDirectionOk() (*SecurityGroupRuleDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CreateSecurityGroupRuleModel) SetDirection(v SecurityGroupRuleDirection)`

SetDirection sets Direction field to given value.


### GetProtocol

`func (o *CreateSecurityGroupRuleModel) GetProtocol() SecurityGroupRuleProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateSecurityGroupRuleModel) GetProtocolOk() (*SecurityGroupRuleProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateSecurityGroupRuleModel) SetProtocol(v SecurityGroupRuleProtocol)`

SetProtocol sets Protocol field to given value.


### GetPortRangeMin

`func (o *CreateSecurityGroupRuleModel) GetPortRangeMin() int32`

GetPortRangeMin returns the PortRangeMin field if non-nil, zero value otherwise.

### GetPortRangeMinOk

`func (o *CreateSecurityGroupRuleModel) GetPortRangeMinOk() (*int32, bool)`

GetPortRangeMinOk returns a tuple with the PortRangeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMin

`func (o *CreateSecurityGroupRuleModel) SetPortRangeMin(v int32)`

SetPortRangeMin sets PortRangeMin field to given value.

### HasPortRangeMin

`func (o *CreateSecurityGroupRuleModel) HasPortRangeMin() bool`

HasPortRangeMin returns a boolean if a field has been set.

### SetPortRangeMinNil

`func (o *CreateSecurityGroupRuleModel) SetPortRangeMinNil(b bool)`

 SetPortRangeMinNil sets the value for PortRangeMin to be an explicit nil

### UnsetPortRangeMin
`func (o *CreateSecurityGroupRuleModel) UnsetPortRangeMin()`

UnsetPortRangeMin ensures that no value is present for PortRangeMin, not even an explicit nil
### GetPortRangeMax

`func (o *CreateSecurityGroupRuleModel) GetPortRangeMax() int32`

GetPortRangeMax returns the PortRangeMax field if non-nil, zero value otherwise.

### GetPortRangeMaxOk

`func (o *CreateSecurityGroupRuleModel) GetPortRangeMaxOk() (*int32, bool)`

GetPortRangeMaxOk returns a tuple with the PortRangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMax

`func (o *CreateSecurityGroupRuleModel) SetPortRangeMax(v int32)`

SetPortRangeMax sets PortRangeMax field to given value.

### HasPortRangeMax

`func (o *CreateSecurityGroupRuleModel) HasPortRangeMax() bool`

HasPortRangeMax returns a boolean if a field has been set.

### SetPortRangeMaxNil

`func (o *CreateSecurityGroupRuleModel) SetPortRangeMaxNil(b bool)`

 SetPortRangeMaxNil sets the value for PortRangeMax to be an explicit nil

### UnsetPortRangeMax
`func (o *CreateSecurityGroupRuleModel) UnsetPortRangeMax()`

UnsetPortRangeMax ensures that no value is present for PortRangeMax, not even an explicit nil
### GetRemoteIpPrefix

`func (o *CreateSecurityGroupRuleModel) GetRemoteIpPrefix() string`

GetRemoteIpPrefix returns the RemoteIpPrefix field if non-nil, zero value otherwise.

### GetRemoteIpPrefixOk

`func (o *CreateSecurityGroupRuleModel) GetRemoteIpPrefixOk() (*string, bool)`

GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpPrefix

`func (o *CreateSecurityGroupRuleModel) SetRemoteIpPrefix(v string)`

SetRemoteIpPrefix sets RemoteIpPrefix field to given value.

### HasRemoteIpPrefix

`func (o *CreateSecurityGroupRuleModel) HasRemoteIpPrefix() bool`

HasRemoteIpPrefix returns a boolean if a field has been set.

### SetRemoteIpPrefixNil

`func (o *CreateSecurityGroupRuleModel) SetRemoteIpPrefixNil(b bool)`

 SetRemoteIpPrefixNil sets the value for RemoteIpPrefix to be an explicit nil

### UnsetRemoteIpPrefix
`func (o *CreateSecurityGroupRuleModel) UnsetRemoteIpPrefix()`

UnsetRemoteIpPrefix ensures that no value is present for RemoteIpPrefix, not even an explicit nil
### GetRemoteGroupId

`func (o *CreateSecurityGroupRuleModel) GetRemoteGroupId() string`

GetRemoteGroupId returns the RemoteGroupId field if non-nil, zero value otherwise.

### GetRemoteGroupIdOk

`func (o *CreateSecurityGroupRuleModel) GetRemoteGroupIdOk() (*string, bool)`

GetRemoteGroupIdOk returns a tuple with the RemoteGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGroupId

`func (o *CreateSecurityGroupRuleModel) SetRemoteGroupId(v string)`

SetRemoteGroupId sets RemoteGroupId field to given value.

### HasRemoteGroupId

`func (o *CreateSecurityGroupRuleModel) HasRemoteGroupId() bool`

HasRemoteGroupId returns a boolean if a field has been set.

### SetRemoteGroupIdNil

`func (o *CreateSecurityGroupRuleModel) SetRemoteGroupIdNil(b bool)`

 SetRemoteGroupIdNil sets the value for RemoteGroupId to be an explicit nil

### UnsetRemoteGroupId
`func (o *CreateSecurityGroupRuleModel) UnsetRemoteGroupId()`

UnsetRemoteGroupId ensures that no value is present for RemoteGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


