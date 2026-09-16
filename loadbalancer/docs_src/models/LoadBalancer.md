# LoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | 리스너 ID | [optional] 
**Name** | Pointer to **NullableString** | 로드 밸런서 이름 | [optional] 
**Description** | Pointer to **NullableString** | 로드 밸런서에 대한 설명 | [optional] 
**Type** | Pointer to [**NullableLoadBalancerType**](LoadBalancerType.md) | 로드 밸런서 유형 | [optional] 
**ListenerIds** | Pointer to **[]string** | 리스너 ID 목록 | [optional] 
**ProjectId** | Pointer to **NullableString** | 프로젝트 ID | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 | [optional] 
**CreatedAt** | Pointer to **NullableTime** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) | 가용 영역 | [optional] 
**AccessLogs** | Pointer to [**NullableAccessLog**](AccessLog.md) | 액세스 로그 상태 | [optional] 
**TargetGroupCount** | Pointer to **NullableInt64** | 연결된 대상 그룹 수 | [optional] 
**ListenerCount** | Pointer to **NullableInt64** | 연결된 리스너 수 | [optional] 
**PrivateVip** | Pointer to **NullableString** | 프라이빗 IP 주소 | [optional] 
**PublicVip** | Pointer to **NullableString** | 퍼블릭 IP 주소 | [optional] 
**SubnetName** | Pointer to **NullableString** | 연결된 서브넷 이름 | [optional] 
**SubnetCidrBlock** | Pointer to **NullableString** | 서브넷의 IPv4 CIDR 블록 | [optional] 
**VpcId** | Pointer to **NullableString** | 연결된 VPC ID | [optional] 
**VpcName** | Pointer to **NullableString** | 연결된 VPC 이름 | [optional] 
**SubnetId** | Pointer to **NullableString** | 연결된 서브넷 ID | [optional] 
**BeyondLoadBalancerId** | Pointer to **NullableString** | 연결된 고가용성 그룹 ID | [optional] 
**BeyondLoadBalancerName** | Pointer to **NullableString** | 연결된 고가용성 그룹 이름 | [optional] 
**BeyondLoadBalancerDnsName** | Pointer to **NullableString** | 연결된 고가용성 그룹 DNS 이름 | [optional] 

## Methods

### NewLoadBalancer

`func NewLoadBalancer() *LoadBalancer`

NewLoadBalancer instantiates a new LoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerWithDefaults

`func NewLoadBalancerWithDefaults() *LoadBalancer`

NewLoadBalancerWithDefaults instantiates a new LoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadBalancer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoadBalancer) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *LoadBalancer) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LoadBalancer) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *LoadBalancer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadBalancer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadBalancer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoadBalancer) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LoadBalancer) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LoadBalancer) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *LoadBalancer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LoadBalancer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LoadBalancer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LoadBalancer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LoadBalancer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LoadBalancer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *LoadBalancer) GetType() LoadBalancerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoadBalancer) GetTypeOk() (*LoadBalancerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoadBalancer) SetType(v LoadBalancerType)`

SetType sets Type field to given value.

### HasType

`func (o *LoadBalancer) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *LoadBalancer) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *LoadBalancer) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetListenerIds

`func (o *LoadBalancer) GetListenerIds() []string`

GetListenerIds returns the ListenerIds field if non-nil, zero value otherwise.

### GetListenerIdsOk

`func (o *LoadBalancer) GetListenerIdsOk() (*[]string, bool)`

GetListenerIdsOk returns a tuple with the ListenerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerIds

`func (o *LoadBalancer) SetListenerIds(v []string)`

SetListenerIds sets ListenerIds field to given value.

### HasListenerIds

`func (o *LoadBalancer) HasListenerIds() bool`

HasListenerIds returns a boolean if a field has been set.

### SetListenerIdsNil

`func (o *LoadBalancer) SetListenerIdsNil(b bool)`

 SetListenerIdsNil sets the value for ListenerIds to be an explicit nil

### UnsetListenerIds
`func (o *LoadBalancer) UnsetListenerIds()`

UnsetListenerIds ensures that no value is present for ListenerIds, not even an explicit nil
### GetProjectId

`func (o *LoadBalancer) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *LoadBalancer) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *LoadBalancer) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *LoadBalancer) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *LoadBalancer) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *LoadBalancer) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProvisioningStatus

`func (o *LoadBalancer) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *LoadBalancer) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *LoadBalancer) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *LoadBalancer) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *LoadBalancer) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *LoadBalancer) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *LoadBalancer) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *LoadBalancer) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *LoadBalancer) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *LoadBalancer) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *LoadBalancer) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *LoadBalancer) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetCreatedAt

`func (o *LoadBalancer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoadBalancer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoadBalancer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoadBalancer) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *LoadBalancer) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LoadBalancer) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *LoadBalancer) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LoadBalancer) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LoadBalancer) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LoadBalancer) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *LoadBalancer) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *LoadBalancer) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetAvailabilityZone

`func (o *LoadBalancer) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *LoadBalancer) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *LoadBalancer) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *LoadBalancer) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *LoadBalancer) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *LoadBalancer) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetAccessLogs

`func (o *LoadBalancer) GetAccessLogs() AccessLog`

GetAccessLogs returns the AccessLogs field if non-nil, zero value otherwise.

### GetAccessLogsOk

`func (o *LoadBalancer) GetAccessLogsOk() (*AccessLog, bool)`

GetAccessLogsOk returns a tuple with the AccessLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLogs

`func (o *LoadBalancer) SetAccessLogs(v AccessLog)`

SetAccessLogs sets AccessLogs field to given value.

### HasAccessLogs

`func (o *LoadBalancer) HasAccessLogs() bool`

HasAccessLogs returns a boolean if a field has been set.

### SetAccessLogsNil

`func (o *LoadBalancer) SetAccessLogsNil(b bool)`

 SetAccessLogsNil sets the value for AccessLogs to be an explicit nil

### UnsetAccessLogs
`func (o *LoadBalancer) UnsetAccessLogs()`

UnsetAccessLogs ensures that no value is present for AccessLogs, not even an explicit nil
### GetTargetGroupCount

`func (o *LoadBalancer) GetTargetGroupCount() int64`

GetTargetGroupCount returns the TargetGroupCount field if non-nil, zero value otherwise.

### GetTargetGroupCountOk

`func (o *LoadBalancer) GetTargetGroupCountOk() (*int64, bool)`

GetTargetGroupCountOk returns a tuple with the TargetGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupCount

`func (o *LoadBalancer) SetTargetGroupCount(v int64)`

SetTargetGroupCount sets TargetGroupCount field to given value.

### HasTargetGroupCount

`func (o *LoadBalancer) HasTargetGroupCount() bool`

HasTargetGroupCount returns a boolean if a field has been set.

### SetTargetGroupCountNil

`func (o *LoadBalancer) SetTargetGroupCountNil(b bool)`

 SetTargetGroupCountNil sets the value for TargetGroupCount to be an explicit nil

### UnsetTargetGroupCount
`func (o *LoadBalancer) UnsetTargetGroupCount()`

UnsetTargetGroupCount ensures that no value is present for TargetGroupCount, not even an explicit nil
### GetListenerCount

`func (o *LoadBalancer) GetListenerCount() int64`

GetListenerCount returns the ListenerCount field if non-nil, zero value otherwise.

### GetListenerCountOk

`func (o *LoadBalancer) GetListenerCountOk() (*int64, bool)`

GetListenerCountOk returns a tuple with the ListenerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerCount

`func (o *LoadBalancer) SetListenerCount(v int64)`

SetListenerCount sets ListenerCount field to given value.

### HasListenerCount

`func (o *LoadBalancer) HasListenerCount() bool`

HasListenerCount returns a boolean if a field has been set.

### SetListenerCountNil

`func (o *LoadBalancer) SetListenerCountNil(b bool)`

 SetListenerCountNil sets the value for ListenerCount to be an explicit nil

### UnsetListenerCount
`func (o *LoadBalancer) UnsetListenerCount()`

UnsetListenerCount ensures that no value is present for ListenerCount, not even an explicit nil
### GetPrivateVip

`func (o *LoadBalancer) GetPrivateVip() string`

GetPrivateVip returns the PrivateVip field if non-nil, zero value otherwise.

### GetPrivateVipOk

`func (o *LoadBalancer) GetPrivateVipOk() (*string, bool)`

GetPrivateVipOk returns a tuple with the PrivateVip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateVip

`func (o *LoadBalancer) SetPrivateVip(v string)`

SetPrivateVip sets PrivateVip field to given value.

### HasPrivateVip

`func (o *LoadBalancer) HasPrivateVip() bool`

HasPrivateVip returns a boolean if a field has been set.

### SetPrivateVipNil

`func (o *LoadBalancer) SetPrivateVipNil(b bool)`

 SetPrivateVipNil sets the value for PrivateVip to be an explicit nil

### UnsetPrivateVip
`func (o *LoadBalancer) UnsetPrivateVip()`

UnsetPrivateVip ensures that no value is present for PrivateVip, not even an explicit nil
### GetPublicVip

`func (o *LoadBalancer) GetPublicVip() string`

GetPublicVip returns the PublicVip field if non-nil, zero value otherwise.

### GetPublicVipOk

`func (o *LoadBalancer) GetPublicVipOk() (*string, bool)`

GetPublicVipOk returns a tuple with the PublicVip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicVip

`func (o *LoadBalancer) SetPublicVip(v string)`

SetPublicVip sets PublicVip field to given value.

### HasPublicVip

`func (o *LoadBalancer) HasPublicVip() bool`

HasPublicVip returns a boolean if a field has been set.

### SetPublicVipNil

`func (o *LoadBalancer) SetPublicVipNil(b bool)`

 SetPublicVipNil sets the value for PublicVip to be an explicit nil

### UnsetPublicVip
`func (o *LoadBalancer) UnsetPublicVip()`

UnsetPublicVip ensures that no value is present for PublicVip, not even an explicit nil
### GetSubnetName

`func (o *LoadBalancer) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *LoadBalancer) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *LoadBalancer) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.

### HasSubnetName

`func (o *LoadBalancer) HasSubnetName() bool`

HasSubnetName returns a boolean if a field has been set.

### SetSubnetNameNil

`func (o *LoadBalancer) SetSubnetNameNil(b bool)`

 SetSubnetNameNil sets the value for SubnetName to be an explicit nil

### UnsetSubnetName
`func (o *LoadBalancer) UnsetSubnetName()`

UnsetSubnetName ensures that no value is present for SubnetName, not even an explicit nil
### GetSubnetCidrBlock

`func (o *LoadBalancer) GetSubnetCidrBlock() string`

GetSubnetCidrBlock returns the SubnetCidrBlock field if non-nil, zero value otherwise.

### GetSubnetCidrBlockOk

`func (o *LoadBalancer) GetSubnetCidrBlockOk() (*string, bool)`

GetSubnetCidrBlockOk returns a tuple with the SubnetCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetCidrBlock

`func (o *LoadBalancer) SetSubnetCidrBlock(v string)`

SetSubnetCidrBlock sets SubnetCidrBlock field to given value.

### HasSubnetCidrBlock

`func (o *LoadBalancer) HasSubnetCidrBlock() bool`

HasSubnetCidrBlock returns a boolean if a field has been set.

### SetSubnetCidrBlockNil

`func (o *LoadBalancer) SetSubnetCidrBlockNil(b bool)`

 SetSubnetCidrBlockNil sets the value for SubnetCidrBlock to be an explicit nil

### UnsetSubnetCidrBlock
`func (o *LoadBalancer) UnsetSubnetCidrBlock()`

UnsetSubnetCidrBlock ensures that no value is present for SubnetCidrBlock, not even an explicit nil
### GetVpcId

`func (o *LoadBalancer) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *LoadBalancer) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *LoadBalancer) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *LoadBalancer) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *LoadBalancer) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *LoadBalancer) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetVpcName

`func (o *LoadBalancer) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *LoadBalancer) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *LoadBalancer) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *LoadBalancer) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.

### SetVpcNameNil

`func (o *LoadBalancer) SetVpcNameNil(b bool)`

 SetVpcNameNil sets the value for VpcName to be an explicit nil

### UnsetVpcName
`func (o *LoadBalancer) UnsetVpcName()`

UnsetVpcName ensures that no value is present for VpcName, not even an explicit nil
### GetSubnetId

`func (o *LoadBalancer) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *LoadBalancer) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *LoadBalancer) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *LoadBalancer) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *LoadBalancer) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *LoadBalancer) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetBeyondLoadBalancerId

`func (o *LoadBalancer) GetBeyondLoadBalancerId() string`

GetBeyondLoadBalancerId returns the BeyondLoadBalancerId field if non-nil, zero value otherwise.

### GetBeyondLoadBalancerIdOk

`func (o *LoadBalancer) GetBeyondLoadBalancerIdOk() (*string, bool)`

GetBeyondLoadBalancerIdOk returns a tuple with the BeyondLoadBalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeyondLoadBalancerId

`func (o *LoadBalancer) SetBeyondLoadBalancerId(v string)`

SetBeyondLoadBalancerId sets BeyondLoadBalancerId field to given value.

### HasBeyondLoadBalancerId

`func (o *LoadBalancer) HasBeyondLoadBalancerId() bool`

HasBeyondLoadBalancerId returns a boolean if a field has been set.

### SetBeyondLoadBalancerIdNil

`func (o *LoadBalancer) SetBeyondLoadBalancerIdNil(b bool)`

 SetBeyondLoadBalancerIdNil sets the value for BeyondLoadBalancerId to be an explicit nil

### UnsetBeyondLoadBalancerId
`func (o *LoadBalancer) UnsetBeyondLoadBalancerId()`

UnsetBeyondLoadBalancerId ensures that no value is present for BeyondLoadBalancerId, not even an explicit nil
### GetBeyondLoadBalancerName

`func (o *LoadBalancer) GetBeyondLoadBalancerName() string`

GetBeyondLoadBalancerName returns the BeyondLoadBalancerName field if non-nil, zero value otherwise.

### GetBeyondLoadBalancerNameOk

`func (o *LoadBalancer) GetBeyondLoadBalancerNameOk() (*string, bool)`

GetBeyondLoadBalancerNameOk returns a tuple with the BeyondLoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeyondLoadBalancerName

`func (o *LoadBalancer) SetBeyondLoadBalancerName(v string)`

SetBeyondLoadBalancerName sets BeyondLoadBalancerName field to given value.

### HasBeyondLoadBalancerName

`func (o *LoadBalancer) HasBeyondLoadBalancerName() bool`

HasBeyondLoadBalancerName returns a boolean if a field has been set.

### SetBeyondLoadBalancerNameNil

`func (o *LoadBalancer) SetBeyondLoadBalancerNameNil(b bool)`

 SetBeyondLoadBalancerNameNil sets the value for BeyondLoadBalancerName to be an explicit nil

### UnsetBeyondLoadBalancerName
`func (o *LoadBalancer) UnsetBeyondLoadBalancerName()`

UnsetBeyondLoadBalancerName ensures that no value is present for BeyondLoadBalancerName, not even an explicit nil
### GetBeyondLoadBalancerDnsName

`func (o *LoadBalancer) GetBeyondLoadBalancerDnsName() string`

GetBeyondLoadBalancerDnsName returns the BeyondLoadBalancerDnsName field if non-nil, zero value otherwise.

### GetBeyondLoadBalancerDnsNameOk

`func (o *LoadBalancer) GetBeyondLoadBalancerDnsNameOk() (*string, bool)`

GetBeyondLoadBalancerDnsNameOk returns a tuple with the BeyondLoadBalancerDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeyondLoadBalancerDnsName

`func (o *LoadBalancer) SetBeyondLoadBalancerDnsName(v string)`

SetBeyondLoadBalancerDnsName sets BeyondLoadBalancerDnsName field to given value.

### HasBeyondLoadBalancerDnsName

`func (o *LoadBalancer) HasBeyondLoadBalancerDnsName() bool`

HasBeyondLoadBalancerDnsName returns a boolean if a field has been set.

### SetBeyondLoadBalancerDnsNameNil

`func (o *LoadBalancer) SetBeyondLoadBalancerDnsNameNil(b bool)`

 SetBeyondLoadBalancerDnsNameNil sets the value for BeyondLoadBalancerDnsName to be an explicit nil

### UnsetBeyondLoadBalancerDnsName
`func (o *LoadBalancer) UnsetBeyondLoadBalancerDnsName()`

UnsetBeyondLoadBalancerDnsName ensures that no value is present for BeyondLoadBalancerDnsName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


