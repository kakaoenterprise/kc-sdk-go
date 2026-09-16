# CreateNetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 네트워크 인터페이스의 이름 | 
**Description** | Pointer to **NullableString** | 네트워크 인터페이스에 대한 설명 | [optional] 
**PrivateIp** | Pointer to **NullableString** | 수동으로 지정할 프라이빗 IP 주소, 서브넷 내 유효한 IP - 지정한 &#x60;private_ip&#x60;가 이미 사용 중이거나 서브넷 CIDR 범위 밖인 경우 요청 실패 | [optional] 
**SubnetId** | **string** | 네트워크 인터페이스를 연결할 서브넷의 ID - [List subnets](/openapi/networking/vpc/list-subnets)에서 확인 | 
**SecurityGroups** | **[]string** | 연결할 보안 그룹의 ID 목록 - 최소 1개 이상 지정 권장 - ※ 생략 시 기본 보안 그룹이 자동 적용되지 않음 - [List security groups](/openapi/networking/vpc/list-security-groups)에서 확인 | 

## Methods

### NewCreateNetworkInterface

`func NewCreateNetworkInterface(name string, subnetId string, securityGroups []string, ) *CreateNetworkInterface`

NewCreateNetworkInterface instantiates a new CreateNetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkInterfaceWithDefaults

`func NewCreateNetworkInterfaceWithDefaults() *CreateNetworkInterface`

NewCreateNetworkInterfaceWithDefaults instantiates a new CreateNetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkInterface) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateNetworkInterface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkInterface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkInterface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNetworkInterface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateNetworkInterface) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateNetworkInterface) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPrivateIp

`func (o *CreateNetworkInterface) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *CreateNetworkInterface) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *CreateNetworkInterface) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *CreateNetworkInterface) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### SetPrivateIpNil

`func (o *CreateNetworkInterface) SetPrivateIpNil(b bool)`

 SetPrivateIpNil sets the value for PrivateIp to be an explicit nil

### UnsetPrivateIp
`func (o *CreateNetworkInterface) UnsetPrivateIp()`

UnsetPrivateIp ensures that no value is present for PrivateIp, not even an explicit nil
### GetSubnetId

`func (o *CreateNetworkInterface) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreateNetworkInterface) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreateNetworkInterface) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetSecurityGroups

`func (o *CreateNetworkInterface) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CreateNetworkInterface) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CreateNetworkInterface) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


