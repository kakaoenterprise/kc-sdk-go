# BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Attachment ID | 
**TgwId** | **string** | Transit Gateway ID | 
**VpcId** | **string** | 연결된 VPC ID | 
**ProvisioningStatus** | [**TGWProvisioningStatus**](TGWProvisioningStatus.md) | 프로비저닝 상태 | 
**TgwProjectId** | **string** | VPC 소유자 프로젝트 ID | 
**VpcName** | **string** | 연결된 VPC 이름 | 
**CidrBlock** | **string** | 연결된 VPC의 CIDR 블록 (예: 10.0.0.0/16) | 
**ProjectId** | **string** | Attachment가 속한 프로젝트 ID | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 

## Methods

### NewBnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel

`func NewBnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel(id string, tgwId string, vpcId string, provisioningStatus TGWProvisioningStatus, tgwProjectId string, vpcName string, cidrBlock string, projectId string, createdAt time.Time, updatedAt time.Time, ) *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel`

NewBnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel instantiates a new BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModelWithDefaults

`func NewBnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModelWithDefaults() *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel`

NewBnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModelWithDefaults instantiates a new BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetTgwId

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetTgwId() string`

GetTgwId returns the TgwId field if non-nil, zero value otherwise.

### GetTgwIdOk

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetTgwIdOk() (*string, bool)`

GetTgwIdOk returns a tuple with the TgwId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwId

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) SetTgwId(v string)`

SetTgwId sets TgwId field to given value.


### GetVpcId

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetProvisioningStatus

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetProvisioningStatus() TGWProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetProvisioningStatusOk() (*TGWProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) SetProvisioningStatus(v TGWProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetTgwProjectId

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetTgwProjectId() string`

GetTgwProjectId returns the TgwProjectId field if non-nil, zero value otherwise.

### GetTgwProjectIdOk

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetTgwProjectIdOk() (*string, bool)`

GetTgwProjectIdOk returns a tuple with the TgwProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwProjectId

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) SetTgwProjectId(v string)`

SetTgwProjectId sets TgwProjectId field to given value.


### GetVpcName

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.


### GetCidrBlock

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetProjectId

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCreatedAt

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsTgwV1ApiUpdateTgwAttachmentModelTgwAttachmentResponseModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


