# BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 생성된 퍼블릭 IP의 고유 ID | 
**Status** | **string** | 현재 퍼블릭 IP의 상태 | 
**Description** | **string** | 퍼블릭 IP에 대한 설명 | 
**ProjectId** | **string** | 퍼블릭 IP가 속한 프로젝트 ID | 
**PublicIp** | **string** | 외부에서 접근 가능한 퍼블릭 IP 주소 (IPv4 형식) | 
**PrivateIp** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 

## Methods

### NewBnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel

`func NewBnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel(id string, status string, description string, projectId string, publicIp string, createdAt time.Time, updatedAt time.Time, ) *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel`

NewBnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel instantiates a new BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModelWithDefaults

`func NewBnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModelWithDefaults() *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel`

NewBnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModelWithDefaults instantiates a new BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProjectId

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetPublicIp

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.


### GetPrivateIp

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### SetPrivateIpNil

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) SetPrivateIpNil(b bool)`

 SetPrivateIpNil sets the value for PrivateIp to be an explicit nil

### UnsetPrivateIp
`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) UnsetPrivateIp()`

UnsetPrivateIp ensures that no value is present for PrivateIp, not even an explicit nil
### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiAssociateNewPublicIpModelPublicIpModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


