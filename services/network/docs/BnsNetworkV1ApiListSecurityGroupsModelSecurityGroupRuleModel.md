# BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 보안 그룹 규칙의 ID | 
**Description** | Pointer to **NullableString** |  | [optional] 
**RemoteGroupId** | Pointer to **NullableString** |  | [optional] 
**RemoteGroupName** | Pointer to **NullableString** |  | [optional] 
**Direction** | Pointer to [**NullableSecurityGroupRuleDirection**](SecurityGroupRuleDirection.md) |  | [optional] 
**Protocol** | [**SecurityGroupRuleProtocol**](SecurityGroupRuleProtocol.md) | 프로토콜 | 
**PortRangeMin** | Pointer to **NullableString** |  | [optional] 
**PortRangeMax** | Pointer to **NullableString** |  | [optional] 
**RemoteIpPrefix** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel

`func NewBnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel(id string, protocol SecurityGroupRuleProtocol, ) *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel`

NewBnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel instantiates a new BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModelWithDefaults

`func NewBnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModelWithDefaults() *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel`

NewBnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModelWithDefaults instantiates a new BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRemoteGroupId

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetRemoteGroupId() string`

GetRemoteGroupId returns the RemoteGroupId field if non-nil, zero value otherwise.

### GetRemoteGroupIdOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetRemoteGroupIdOk() (*string, bool)`

GetRemoteGroupIdOk returns a tuple with the RemoteGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGroupId

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetRemoteGroupId(v string)`

SetRemoteGroupId sets RemoteGroupId field to given value.

### HasRemoteGroupId

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) HasRemoteGroupId() bool`

HasRemoteGroupId returns a boolean if a field has been set.

### SetRemoteGroupIdNil

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetRemoteGroupIdNil(b bool)`

 SetRemoteGroupIdNil sets the value for RemoteGroupId to be an explicit nil

### UnsetRemoteGroupId
`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) UnsetRemoteGroupId()`

UnsetRemoteGroupId ensures that no value is present for RemoteGroupId, not even an explicit nil
### GetRemoteGroupName

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetRemoteGroupName() string`

GetRemoteGroupName returns the RemoteGroupName field if non-nil, zero value otherwise.

### GetRemoteGroupNameOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetRemoteGroupNameOk() (*string, bool)`

GetRemoteGroupNameOk returns a tuple with the RemoteGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGroupName

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetRemoteGroupName(v string)`

SetRemoteGroupName sets RemoteGroupName field to given value.

### HasRemoteGroupName

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) HasRemoteGroupName() bool`

HasRemoteGroupName returns a boolean if a field has been set.

### SetRemoteGroupNameNil

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetRemoteGroupNameNil(b bool)`

 SetRemoteGroupNameNil sets the value for RemoteGroupName to be an explicit nil

### UnsetRemoteGroupName
`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) UnsetRemoteGroupName()`

UnsetRemoteGroupName ensures that no value is present for RemoteGroupName, not even an explicit nil
### GetDirection

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetDirection() SecurityGroupRuleDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetDirectionOk() (*SecurityGroupRuleDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetDirection(v SecurityGroupRuleDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetProtocol

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetProtocol() SecurityGroupRuleProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetProtocolOk() (*SecurityGroupRuleProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetProtocol(v SecurityGroupRuleProtocol)`

SetProtocol sets Protocol field to given value.


### GetPortRangeMin

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetPortRangeMin() string`

GetPortRangeMin returns the PortRangeMin field if non-nil, zero value otherwise.

### GetPortRangeMinOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetPortRangeMinOk() (*string, bool)`

GetPortRangeMinOk returns a tuple with the PortRangeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMin

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetPortRangeMin(v string)`

SetPortRangeMin sets PortRangeMin field to given value.

### HasPortRangeMin

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) HasPortRangeMin() bool`

HasPortRangeMin returns a boolean if a field has been set.

### SetPortRangeMinNil

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetPortRangeMinNil(b bool)`

 SetPortRangeMinNil sets the value for PortRangeMin to be an explicit nil

### UnsetPortRangeMin
`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) UnsetPortRangeMin()`

UnsetPortRangeMin ensures that no value is present for PortRangeMin, not even an explicit nil
### GetPortRangeMax

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetPortRangeMax() string`

GetPortRangeMax returns the PortRangeMax field if non-nil, zero value otherwise.

### GetPortRangeMaxOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetPortRangeMaxOk() (*string, bool)`

GetPortRangeMaxOk returns a tuple with the PortRangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeMax

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetPortRangeMax(v string)`

SetPortRangeMax sets PortRangeMax field to given value.

### HasPortRangeMax

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) HasPortRangeMax() bool`

HasPortRangeMax returns a boolean if a field has been set.

### SetPortRangeMaxNil

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetPortRangeMaxNil(b bool)`

 SetPortRangeMaxNil sets the value for PortRangeMax to be an explicit nil

### UnsetPortRangeMax
`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) UnsetPortRangeMax()`

UnsetPortRangeMax ensures that no value is present for PortRangeMax, not even an explicit nil
### GetRemoteIpPrefix

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetRemoteIpPrefix() string`

GetRemoteIpPrefix returns the RemoteIpPrefix field if non-nil, zero value otherwise.

### GetRemoteIpPrefixOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetRemoteIpPrefixOk() (*string, bool)`

GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpPrefix

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetRemoteIpPrefix(v string)`

SetRemoteIpPrefix sets RemoteIpPrefix field to given value.

### HasRemoteIpPrefix

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) HasRemoteIpPrefix() bool`

HasRemoteIpPrefix returns a boolean if a field has been set.

### SetRemoteIpPrefixNil

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetRemoteIpPrefixNil(b bool)`

 SetRemoteIpPrefixNil sets the value for RemoteIpPrefix to be an explicit nil

### UnsetRemoteIpPrefix
`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) UnsetRemoteIpPrefix()`

UnsetRemoteIpPrefix ensures that no value is present for RemoteIpPrefix, not even an explicit nil
### GetCreatedAt

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


