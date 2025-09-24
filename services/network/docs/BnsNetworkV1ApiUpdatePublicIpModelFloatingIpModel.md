# BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 퍼블릭 IP의 고유 ID | 
**Status** | [**PublicIpStatus**](PublicIpStatus.md) | 퍼블릭 IP 상태 | 
**Description** | **string** | 퍼블릭 IP에 대한 설명 | 
**ProjectId** | **string** | 퍼블릭 IP가 속한 프로젝트 ID | 
**PublicIp** | **string** | 할당된 퍼블릭 IPv4 주소 | 
**PrivateIp** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 

## Methods

### NewBnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel

`func NewBnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel(id string, status PublicIpStatus, description string, projectId string, publicIp string, createdAt time.Time, updatedAt time.Time, ) *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel`

NewBnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel instantiates a new BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsNetworkV1ApiUpdatePublicIpModelFloatingIpModelWithDefaults

`func NewBnsNetworkV1ApiUpdatePublicIpModelFloatingIpModelWithDefaults() *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel`

NewBnsNetworkV1ApiUpdatePublicIpModelFloatingIpModelWithDefaults instantiates a new BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) GetStatus() PublicIpStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) GetStatusOk() (*PublicIpStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) SetStatus(v PublicIpStatus)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProjectId

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetPublicIp

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.


### GetPrivateIp

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### SetPrivateIpNil

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) SetPrivateIpNil(b bool)`

 SetPrivateIpNil sets the value for PrivateIp to be an explicit nil

### UnsetPrivateIp
`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) UnsetPrivateIp()`

UnsetPrivateIp ensures that no value is present for PrivateIp, not even an explicit nil
### GetCreatedAt

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsNetworkV1ApiUpdatePublicIpModelFloatingIpModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


