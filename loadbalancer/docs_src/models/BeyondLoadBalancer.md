# BeyondLoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | 고가용성 그룹 ID | [optional] 
**Name** | Pointer to **NullableString** | 고가용성 그룹 이름 | [optional] 
**Description** | Pointer to **NullableString** | 고가용성 그룹에 대한 설명 | [optional] 
**Provider** | Pointer to **NullableString** | 제공자 정보 | [optional] 
**Scheme** | Pointer to [**NullableBeyondLoadBalancerScheme**](BeyondLoadBalancerScheme.md) | 접근 방식 | [optional] 
**ProjectId** | Pointer to **NullableString** | 프로젝트 ID | [optional] 
**DnsName** | Pointer to **NullableString** | DNS 이름 | [optional] 
**TypeId** | Pointer to **NullableString** | 고가용성 그룹의 유형을 식별하는 ID | [optional] 
**CreatedAt** | Pointer to **NullableTime** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 | [optional] 
**VpcId** | Pointer to **NullableString** | VPC의 고유 ID | [optional] 
**Type** | Pointer to [**NullableLoadBalancerType**](LoadBalancerType.md) | 로드 밸런서 유형 | [optional] 
**VpcName** | Pointer to **NullableString** | VPC 이름 | [optional] 
**VpcCidrBlock** | Pointer to **NullableString** | VPC의 IPv4 CIDR 블록 | [optional] 
**AvailabilityZones** | Pointer to [**[]AvailabilityZone**](AvailabilityZone.md) | 가용 영역 목록 | [optional] 
**LoadBalancers** | [**[]LoadBalancerDetail**](LoadBalancerDetail.md) | 포함된 로드 밸런서 객체 배열 | 

## Methods

### NewBeyondLoadBalancer

`func NewBeyondLoadBalancer(loadBalancers []LoadBalancerDetail, ) *BeyondLoadBalancer`

NewBeyondLoadBalancer instantiates a new BeyondLoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeyondLoadBalancerWithDefaults

`func NewBeyondLoadBalancerWithDefaults() *BeyondLoadBalancer`

NewBeyondLoadBalancerWithDefaults instantiates a new BeyondLoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BeyondLoadBalancer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BeyondLoadBalancer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BeyondLoadBalancer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BeyondLoadBalancer) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BeyondLoadBalancer) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BeyondLoadBalancer) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *BeyondLoadBalancer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BeyondLoadBalancer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BeyondLoadBalancer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BeyondLoadBalancer) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BeyondLoadBalancer) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BeyondLoadBalancer) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BeyondLoadBalancer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BeyondLoadBalancer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BeyondLoadBalancer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BeyondLoadBalancer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BeyondLoadBalancer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BeyondLoadBalancer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProvider

`func (o *BeyondLoadBalancer) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BeyondLoadBalancer) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BeyondLoadBalancer) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *BeyondLoadBalancer) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *BeyondLoadBalancer) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *BeyondLoadBalancer) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetScheme

`func (o *BeyondLoadBalancer) GetScheme() BeyondLoadBalancerScheme`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *BeyondLoadBalancer) GetSchemeOk() (*BeyondLoadBalancerScheme, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *BeyondLoadBalancer) SetScheme(v BeyondLoadBalancerScheme)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *BeyondLoadBalancer) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### SetSchemeNil

`func (o *BeyondLoadBalancer) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *BeyondLoadBalancer) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
### GetProjectId

`func (o *BeyondLoadBalancer) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BeyondLoadBalancer) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BeyondLoadBalancer) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BeyondLoadBalancer) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BeyondLoadBalancer) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BeyondLoadBalancer) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetDnsName

`func (o *BeyondLoadBalancer) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *BeyondLoadBalancer) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *BeyondLoadBalancer) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *BeyondLoadBalancer) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### SetDnsNameNil

`func (o *BeyondLoadBalancer) SetDnsNameNil(b bool)`

 SetDnsNameNil sets the value for DnsName to be an explicit nil

### UnsetDnsName
`func (o *BeyondLoadBalancer) UnsetDnsName()`

UnsetDnsName ensures that no value is present for DnsName, not even an explicit nil
### GetTypeId

`func (o *BeyondLoadBalancer) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *BeyondLoadBalancer) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *BeyondLoadBalancer) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *BeyondLoadBalancer) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### SetTypeIdNil

`func (o *BeyondLoadBalancer) SetTypeIdNil(b bool)`

 SetTypeIdNil sets the value for TypeId to be an explicit nil

### UnsetTypeId
`func (o *BeyondLoadBalancer) UnsetTypeId()`

UnsetTypeId ensures that no value is present for TypeId, not even an explicit nil
### GetCreatedAt

`func (o *BeyondLoadBalancer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BeyondLoadBalancer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BeyondLoadBalancer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BeyondLoadBalancer) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BeyondLoadBalancer) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BeyondLoadBalancer) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BeyondLoadBalancer) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BeyondLoadBalancer) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BeyondLoadBalancer) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BeyondLoadBalancer) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BeyondLoadBalancer) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BeyondLoadBalancer) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetProvisioningStatus

`func (o *BeyondLoadBalancer) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BeyondLoadBalancer) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BeyondLoadBalancer) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BeyondLoadBalancer) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BeyondLoadBalancer) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BeyondLoadBalancer) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *BeyondLoadBalancer) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BeyondLoadBalancer) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BeyondLoadBalancer) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BeyondLoadBalancer) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BeyondLoadBalancer) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BeyondLoadBalancer) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetVpcId

`func (o *BeyondLoadBalancer) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BeyondLoadBalancer) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BeyondLoadBalancer) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BeyondLoadBalancer) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BeyondLoadBalancer) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BeyondLoadBalancer) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetType

`func (o *BeyondLoadBalancer) GetType() LoadBalancerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BeyondLoadBalancer) GetTypeOk() (*LoadBalancerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BeyondLoadBalancer) SetType(v LoadBalancerType)`

SetType sets Type field to given value.

### HasType

`func (o *BeyondLoadBalancer) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BeyondLoadBalancer) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BeyondLoadBalancer) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetVpcName

`func (o *BeyondLoadBalancer) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *BeyondLoadBalancer) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *BeyondLoadBalancer) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *BeyondLoadBalancer) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.

### SetVpcNameNil

`func (o *BeyondLoadBalancer) SetVpcNameNil(b bool)`

 SetVpcNameNil sets the value for VpcName to be an explicit nil

### UnsetVpcName
`func (o *BeyondLoadBalancer) UnsetVpcName()`

UnsetVpcName ensures that no value is present for VpcName, not even an explicit nil
### GetVpcCidrBlock

`func (o *BeyondLoadBalancer) GetVpcCidrBlock() string`

GetVpcCidrBlock returns the VpcCidrBlock field if non-nil, zero value otherwise.

### GetVpcCidrBlockOk

`func (o *BeyondLoadBalancer) GetVpcCidrBlockOk() (*string, bool)`

GetVpcCidrBlockOk returns a tuple with the VpcCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcCidrBlock

`func (o *BeyondLoadBalancer) SetVpcCidrBlock(v string)`

SetVpcCidrBlock sets VpcCidrBlock field to given value.

### HasVpcCidrBlock

`func (o *BeyondLoadBalancer) HasVpcCidrBlock() bool`

HasVpcCidrBlock returns a boolean if a field has been set.

### SetVpcCidrBlockNil

`func (o *BeyondLoadBalancer) SetVpcCidrBlockNil(b bool)`

 SetVpcCidrBlockNil sets the value for VpcCidrBlock to be an explicit nil

### UnsetVpcCidrBlock
`func (o *BeyondLoadBalancer) UnsetVpcCidrBlock()`

UnsetVpcCidrBlock ensures that no value is present for VpcCidrBlock, not even an explicit nil
### GetAvailabilityZones

`func (o *BeyondLoadBalancer) GetAvailabilityZones() []AvailabilityZone`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *BeyondLoadBalancer) GetAvailabilityZonesOk() (*[]AvailabilityZone, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *BeyondLoadBalancer) SetAvailabilityZones(v []AvailabilityZone)`

SetAvailabilityZones sets AvailabilityZones field to given value.

### HasAvailabilityZones

`func (o *BeyondLoadBalancer) HasAvailabilityZones() bool`

HasAvailabilityZones returns a boolean if a field has been set.

### SetAvailabilityZonesNil

`func (o *BeyondLoadBalancer) SetAvailabilityZonesNil(b bool)`

 SetAvailabilityZonesNil sets the value for AvailabilityZones to be an explicit nil

### UnsetAvailabilityZones
`func (o *BeyondLoadBalancer) UnsetAvailabilityZones()`

UnsetAvailabilityZones ensures that no value is present for AvailabilityZones, not even an explicit nil
### GetLoadBalancers

`func (o *BeyondLoadBalancer) GetLoadBalancers() []LoadBalancerDetail`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *BeyondLoadBalancer) GetLoadBalancersOk() (*[]LoadBalancerDetail, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *BeyondLoadBalancer) SetLoadBalancers(v []LoadBalancerDetail)`

SetLoadBalancers sets LoadBalancers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


