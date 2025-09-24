# BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 퍼블릭 IP의 고유 ID | 
**Description** | **string** | 퍼블릭 IP에 대한 설명 | 
**Status** | **string** | 퍼블릭 IP의 현재 상태 | 
**ProjectId** | **string** | 퍼블릭 IP가 속한 프로젝트의 ID | 
**PublicIp** | **string** | 실제로 부여된 퍼블릭 IPv4 주소 | 
**PrivateIp** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 

## Methods

### NewBcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel

`func NewBcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel(id string, description string, status string, projectId string, publicIp string, createdAt time.Time, updatedAt time.Time, ) *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel`

NewBcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel instantiates a new BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModelWithDefaults

`func NewBcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModelWithDefaults() *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel`

NewBcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModelWithDefaults instantiates a new BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatus

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProjectId

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetPublicIp

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.


### GetPrivateIp

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### SetPrivateIpNil

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) SetPrivateIpNil(b bool)`

 SetPrivateIpNil sets the value for PrivateIp to be an explicit nil

### UnsetPrivateIp
`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) UnsetPrivateIp()`

UnsetPrivateIp ensures that no value is present for PrivateIp, not even an explicit nil
### GetCreatedAt

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BcsInstanceV1ApiAssociateNewPublicIpModelPublicIpModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


