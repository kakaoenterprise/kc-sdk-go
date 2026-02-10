# BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Attachment의 ID | 
**TgwId** | **string** | Transit Gateway ID | 
**VpcId** | **string** | 연결된 VPC ID | 
**ProvisioningStatus** | [**TGWProvisioningStatus**](TGWProvisioningStatus.md) | 프로비저닝 상태 | 
**TgwProjectId** | **string** | VPC 소유자 ID | 
**VpcName** | **string** | VPC 이름 | 
**CidrBlock** | **string** | VPC CIDR 블록 | 
**ProjectId** | **string** | 프로젝트 ID | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 

## Methods

### NewBnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel

`func NewBnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel(id string, tgwId string, vpcId string, provisioningStatus TGWProvisioningStatus, tgwProjectId string, vpcName string, cidrBlock string, projectId string, createdAt time.Time, updatedAt time.Time, ) *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel`

NewBnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel instantiates a new BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModelWithDefaults

`func NewBnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModelWithDefaults() *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel`

NewBnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModelWithDefaults instantiates a new BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetTgwId

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetTgwId() string`

GetTgwId returns the TgwId field if non-nil, zero value otherwise.

### GetTgwIdOk

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetTgwIdOk() (*string, bool)`

GetTgwIdOk returns a tuple with the TgwId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwId

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) SetTgwId(v string)`

SetTgwId sets TgwId field to given value.


### GetVpcId

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetProvisioningStatus

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetProvisioningStatus() TGWProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetProvisioningStatusOk() (*TGWProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) SetProvisioningStatus(v TGWProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetTgwProjectId

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetTgwProjectId() string`

GetTgwProjectId returns the TgwProjectId field if non-nil, zero value otherwise.

### GetTgwProjectIdOk

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetTgwProjectIdOk() (*string, bool)`

GetTgwProjectIdOk returns a tuple with the TgwProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwProjectId

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) SetTgwProjectId(v string)`

SetTgwProjectId sets TgwProjectId field to given value.


### GetVpcName

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.


### GetCidrBlock

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetProjectId

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCreatedAt

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsTgwV1ApiApproveTgwAttachmentModelTgwAttachmentResponseModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


