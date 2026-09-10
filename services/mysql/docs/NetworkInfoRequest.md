# NetworkInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroupIds** | **[]string** | 인스턴스 그룹에 적용할 보안 그룹 ID 목록 &lt;br/&gt;- [List security groups](https://docs.kakaocloud.com/openapi/networking/vpc/list-security-groups)에서 확인 | 
**PrimarySubnetInfo** | [**SubnetInfoRequest**](SubnetInfoRequest.md) | Primary 서브넷 설정 | 
**StandbySubnetInfo** | Pointer to [**[]SubnetInfoRequest**](SubnetInfoRequest.md) |  | [optional] 

## Methods

### NewNetworkInfoRequest

`func NewNetworkInfoRequest(securityGroupIds []string, primarySubnetInfo SubnetInfoRequest, ) *NetworkInfoRequest`

NewNetworkInfoRequest instantiates a new NetworkInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInfoRequestWithDefaults

`func NewNetworkInfoRequestWithDefaults() *NetworkInfoRequest`

NewNetworkInfoRequestWithDefaults instantiates a new NetworkInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroupIds

`func (o *NetworkInfoRequest) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *NetworkInfoRequest) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *NetworkInfoRequest) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.


### GetPrimarySubnetInfo

`func (o *NetworkInfoRequest) GetPrimarySubnetInfo() SubnetInfoRequest`

GetPrimarySubnetInfo returns the PrimarySubnetInfo field if non-nil, zero value otherwise.

### GetPrimarySubnetInfoOk

`func (o *NetworkInfoRequest) GetPrimarySubnetInfoOk() (*SubnetInfoRequest, bool)`

GetPrimarySubnetInfoOk returns a tuple with the PrimarySubnetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySubnetInfo

`func (o *NetworkInfoRequest) SetPrimarySubnetInfo(v SubnetInfoRequest)`

SetPrimarySubnetInfo sets PrimarySubnetInfo field to given value.


### GetStandbySubnetInfo

`func (o *NetworkInfoRequest) GetStandbySubnetInfo() []SubnetInfoRequest`

GetStandbySubnetInfo returns the StandbySubnetInfo field if non-nil, zero value otherwise.

### GetStandbySubnetInfoOk

`func (o *NetworkInfoRequest) GetStandbySubnetInfoOk() (*[]SubnetInfoRequest, bool)`

GetStandbySubnetInfoOk returns a tuple with the StandbySubnetInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbySubnetInfo

`func (o *NetworkInfoRequest) SetStandbySubnetInfo(v []SubnetInfoRequest)`

SetStandbySubnetInfo sets StandbySubnetInfo field to given value.

### HasStandbySubnetInfo

`func (o *NetworkInfoRequest) HasStandbySubnetInfo() bool`

HasStandbySubnetInfo returns a boolean if a field has been set.

### SetStandbySubnetInfoNil

`func (o *NetworkInfoRequest) SetStandbySubnetInfoNil(b bool)`

 SetStandbySubnetInfoNil sets the value for StandbySubnetInfo to be an explicit nil

### UnsetStandbySubnetInfo
`func (o *NetworkInfoRequest) UnsetStandbySubnetInfo()`

UnsetStandbySubnetInfo ensures that no value is present for StandbySubnetInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


