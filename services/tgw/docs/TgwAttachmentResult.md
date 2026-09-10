# TgwAttachmentResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Attachment ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**TgwId** | **string** | Transit Gateway ID | 
**VpcId** | **string** | 연결된 VPC ID | 
**ProvisioningStatus** | [**TGWAttachmentProvisioningStatus**](TGWAttachmentProvisioningStatus.md) | 프로비저닝 상태 | 
**ProjectId** | **string** | Attachment가 속한 프로젝트 ID | 
**VpcName** | **string** | 연결된 VPC 이름 | 
**CidrBlock** | **string** | 연결된 VPC의 CIDR 블록 (예: 10.0.0.0/16) | 
**TgwProjectId** | **string** | VPC 소유자 프로젝트 ID | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 

## Methods

### NewTgwAttachmentResult

`func NewTgwAttachmentResult(id string, tgwId string, vpcId string, provisioningStatus TGWAttachmentProvisioningStatus, projectId string, vpcName string, cidrBlock string, tgwProjectId string, createdAt time.Time, updatedAt time.Time, ) *TgwAttachmentResult`

NewTgwAttachmentResult instantiates a new TgwAttachmentResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTgwAttachmentResultWithDefaults

`func NewTgwAttachmentResultWithDefaults() *TgwAttachmentResult`

NewTgwAttachmentResultWithDefaults instantiates a new TgwAttachmentResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TgwAttachmentResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TgwAttachmentResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TgwAttachmentResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TgwAttachmentResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TgwAttachmentResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TgwAttachmentResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TgwAttachmentResult) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TgwAttachmentResult) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TgwAttachmentResult) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTgwId

`func (o *TgwAttachmentResult) GetTgwId() string`

GetTgwId returns the TgwId field if non-nil, zero value otherwise.

### GetTgwIdOk

`func (o *TgwAttachmentResult) GetTgwIdOk() (*string, bool)`

GetTgwIdOk returns a tuple with the TgwId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwId

`func (o *TgwAttachmentResult) SetTgwId(v string)`

SetTgwId sets TgwId field to given value.


### GetVpcId

`func (o *TgwAttachmentResult) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *TgwAttachmentResult) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *TgwAttachmentResult) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetProvisioningStatus

`func (o *TgwAttachmentResult) GetProvisioningStatus() TGWAttachmentProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *TgwAttachmentResult) GetProvisioningStatusOk() (*TGWAttachmentProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *TgwAttachmentResult) SetProvisioningStatus(v TGWAttachmentProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetProjectId

`func (o *TgwAttachmentResult) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TgwAttachmentResult) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TgwAttachmentResult) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVpcName

`func (o *TgwAttachmentResult) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *TgwAttachmentResult) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *TgwAttachmentResult) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.


### GetCidrBlock

`func (o *TgwAttachmentResult) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *TgwAttachmentResult) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *TgwAttachmentResult) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetTgwProjectId

`func (o *TgwAttachmentResult) GetTgwProjectId() string`

GetTgwProjectId returns the TgwProjectId field if non-nil, zero value otherwise.

### GetTgwProjectIdOk

`func (o *TgwAttachmentResult) GetTgwProjectIdOk() (*string, bool)`

GetTgwProjectIdOk returns a tuple with the TgwProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwProjectId

`func (o *TgwAttachmentResult) SetTgwProjectId(v string)`

SetTgwProjectId sets TgwProjectId field to given value.


### GetCreatedAt

`func (o *TgwAttachmentResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TgwAttachmentResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TgwAttachmentResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TgwAttachmentResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TgwAttachmentResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TgwAttachmentResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


