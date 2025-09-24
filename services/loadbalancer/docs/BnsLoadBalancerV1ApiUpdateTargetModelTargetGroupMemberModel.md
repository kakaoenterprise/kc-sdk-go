# BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 대상 인스턴스의 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**OperatingStatus** | [**LoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 | 
**ProvisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
**Address** | **string** | 대상 인스턴스의 IP 주소 | 
**ProtocolPort** | **int32** | 대상 인스턴스가 수신하는 포트 번호 | 
**Weight** | **int32** | 트래픽 분산 가중치 | 
**IsBackup** | **bool** | &#x60;true&#x60;일 경우 이 인스턴스는 백업 대상으로 설정되며, 주 멤버가 모두 실패했을 때만 사용 | 
**SubnetId** | **string** | 대상 인스턴스가 속한 서브넷의 ID | 
**ProjectId** | **string** | 해당 멤버 리소스가 속한 프로젝트의 ID | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**MonitorPort** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel

`func NewBnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel(id string, operatingStatus LoadBalancerOperatingStatus, provisioningStatus ProvisioningStatus, address string, protocolPort int32, weight int32, isBackup bool, subnetId string, projectId string, createdAt time.Time, ) *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel`

NewBnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel instantiates a new BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModelWithDefaults

`func NewBnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModelWithDefaults() *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel`

NewBnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModelWithDefaults instantiates a new BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.


### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetAddress

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetProtocolPort

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.


### GetWeight

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetIsBackup

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetIsBackup() bool`

GetIsBackup returns the IsBackup field if non-nil, zero value otherwise.

### GetIsBackupOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetIsBackupOk() (*bool, bool)`

GetIsBackupOk returns a tuple with the IsBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackup

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) SetIsBackup(v bool)`

SetIsBackup sets IsBackup field to given value.


### GetSubnetId

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetProjectId

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetMonitorPort

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetMonitorPort() int32`

GetMonitorPort returns the MonitorPort field if non-nil, zero value otherwise.

### GetMonitorPortOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) GetMonitorPortOk() (*int32, bool)`

GetMonitorPortOk returns a tuple with the MonitorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPort

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) SetMonitorPort(v int32)`

SetMonitorPort sets MonitorPort field to given value.

### HasMonitorPort

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) HasMonitorPort() bool`

HasMonitorPort returns a boolean if a field has been set.

### SetMonitorPortNil

`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) SetMonitorPortNil(b bool)`

 SetMonitorPortNil sets the value for MonitorPort to be an explicit nil

### UnsetMonitorPort
`func (o *BnsLoadBalancerV1ApiUpdateTargetModelTargetGroupMemberModel) UnsetMonitorPort()`

UnsetMonitorPort ensures that no value is present for MonitorPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


