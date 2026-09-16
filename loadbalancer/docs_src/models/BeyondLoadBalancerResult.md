# BeyondLoadBalancerResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 고가용성 로드 밸런서의 ID | 
**Name** | Pointer to **NullableString** | 고가용성 로드 밸런서의 이름 | [optional] 
**Description** | Pointer to **NullableString** | 고가용성 로드 밸런서에 대한 설명 | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 | [optional] 
**ProjectId** | Pointer to **NullableString** | 해당 리소스가 속한 프로젝트의 ID | [optional] 
**DnsName** | Pointer to **NullableString** | 클라이언트가 접근할 수 있는 DNS 도메인 이름 | [optional] 
**VpcId** | Pointer to **NullableString** | 고가용성 로드 밸런서가 연결된 VPC의 ID | [optional] 
**Scheme** | Pointer to [**NullableBeyondLoadBalancerScheme**](BeyondLoadBalancerScheme.md) | 로드 밸런서의 트래픽 처리 방식 | [optional] 
**CreatedAt** | Pointer to **NullableTime** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | [optional] 
**TypeId** | Pointer to **NullableString** | 고가용성 로드 밸런서 유형 식별자 | [optional] 
**Subnets** | Pointer to [**[]VpcSubnet**](VpcSubnet.md) | 연결된 서브넷 목록 | [optional] 

## Methods

### NewBeyondLoadBalancerResult

`func NewBeyondLoadBalancerResult(id string, ) *BeyondLoadBalancerResult`

NewBeyondLoadBalancerResult instantiates a new BeyondLoadBalancerResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeyondLoadBalancerResultWithDefaults

`func NewBeyondLoadBalancerResultWithDefaults() *BeyondLoadBalancerResult`

NewBeyondLoadBalancerResultWithDefaults instantiates a new BeyondLoadBalancerResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BeyondLoadBalancerResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BeyondLoadBalancerResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BeyondLoadBalancerResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BeyondLoadBalancerResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BeyondLoadBalancerResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BeyondLoadBalancerResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BeyondLoadBalancerResult) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BeyondLoadBalancerResult) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BeyondLoadBalancerResult) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BeyondLoadBalancerResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BeyondLoadBalancerResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BeyondLoadBalancerResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BeyondLoadBalancerResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BeyondLoadBalancerResult) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BeyondLoadBalancerResult) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProvisioningStatus

`func (o *BeyondLoadBalancerResult) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BeyondLoadBalancerResult) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BeyondLoadBalancerResult) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BeyondLoadBalancerResult) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BeyondLoadBalancerResult) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BeyondLoadBalancerResult) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *BeyondLoadBalancerResult) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BeyondLoadBalancerResult) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BeyondLoadBalancerResult) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BeyondLoadBalancerResult) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BeyondLoadBalancerResult) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BeyondLoadBalancerResult) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetProjectId

`func (o *BeyondLoadBalancerResult) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BeyondLoadBalancerResult) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BeyondLoadBalancerResult) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BeyondLoadBalancerResult) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BeyondLoadBalancerResult) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BeyondLoadBalancerResult) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetDnsName

`func (o *BeyondLoadBalancerResult) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *BeyondLoadBalancerResult) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *BeyondLoadBalancerResult) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *BeyondLoadBalancerResult) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### SetDnsNameNil

`func (o *BeyondLoadBalancerResult) SetDnsNameNil(b bool)`

 SetDnsNameNil sets the value for DnsName to be an explicit nil

### UnsetDnsName
`func (o *BeyondLoadBalancerResult) UnsetDnsName()`

UnsetDnsName ensures that no value is present for DnsName, not even an explicit nil
### GetVpcId

`func (o *BeyondLoadBalancerResult) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BeyondLoadBalancerResult) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BeyondLoadBalancerResult) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BeyondLoadBalancerResult) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BeyondLoadBalancerResult) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BeyondLoadBalancerResult) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetScheme

`func (o *BeyondLoadBalancerResult) GetScheme() BeyondLoadBalancerScheme`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *BeyondLoadBalancerResult) GetSchemeOk() (*BeyondLoadBalancerScheme, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *BeyondLoadBalancerResult) SetScheme(v BeyondLoadBalancerScheme)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *BeyondLoadBalancerResult) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *BeyondLoadBalancerResult) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *BeyondLoadBalancerResult) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetCreatedAt

`func (o *BeyondLoadBalancerResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BeyondLoadBalancerResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BeyondLoadBalancerResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BeyondLoadBalancerResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BeyondLoadBalancerResult) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BeyondLoadBalancerResult) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BeyondLoadBalancerResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BeyondLoadBalancerResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BeyondLoadBalancerResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BeyondLoadBalancerResult) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BeyondLoadBalancerResult) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BeyondLoadBalancerResult) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetTypeId

`func (o *BeyondLoadBalancerResult) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *BeyondLoadBalancerResult) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *BeyondLoadBalancerResult) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *BeyondLoadBalancerResult) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### SetTypeIdNil

`func (o *BeyondLoadBalancerResult) SetTypeIdNil(b bool)`

 SetTypeIdNil sets the value for TypeId to be an explicit nil

### UnsetTypeId
`func (o *BeyondLoadBalancerResult) UnsetTypeId()`

UnsetTypeId ensures that no value is present for TypeId, not even an explicit nil
### GetSubnets

`func (o *BeyondLoadBalancerResult) GetSubnets() []VpcSubnet`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *BeyondLoadBalancerResult) GetSubnetsOk() (*[]VpcSubnet, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *BeyondLoadBalancerResult) SetSubnets(v []VpcSubnet)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *BeyondLoadBalancerResult) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### SetSubnetsNil

`func (o *BeyondLoadBalancerResult) SetSubnetsNil(b bool)`

 SetSubnetsNil sets the value for Subnets to be an explicit nil

### UnsetSubnets
`func (o *BeyondLoadBalancerResult) UnsetSubnets()`

UnsetSubnets ensures that no value is present for Subnets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


