# Igw

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 인터넷 게이트웨이의 고유 ID | 
**Name** | **string** | 인터넷 게이트웨이 이름 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Vpc** | Pointer to [**NullableInternetGatewayVpc**](InternetGatewayVpc.md) |  | [optional] 
**Attachment** | Pointer to [**NullableAttachment**](Attachment.md) |  | [optional] 
**NatIp** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**DomainId** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**AttachmentStatus** | **string** | IGW Attachment의 연결 상태 | 

## Methods

### NewIgw

`func NewIgw(id string, name string, attachmentStatus string, ) *Igw`

NewIgw instantiates a new Igw object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIgwWithDefaults

`func NewIgwWithDefaults() *Igw`

NewIgwWithDefaults instantiates a new Igw object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Igw) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Igw) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Igw) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Igw) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Igw) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Igw) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Igw) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Igw) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Igw) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Igw) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Igw) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Igw) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVpc

`func (o *Igw) GetVpc() InternetGatewayVpc`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *Igw) GetVpcOk() (*InternetGatewayVpc, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *Igw) SetVpc(v InternetGatewayVpc)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *Igw) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### SetVpcNil

`func (o *Igw) SetVpcNil(b bool)`

 SetVpcNil sets the value for Vpc to be an explicit nil

### UnsetVpc
`func (o *Igw) UnsetVpc()`

UnsetVpc ensures that no value is present for Vpc, not even an explicit nil
### GetAttachment

`func (o *Igw) GetAttachment() Attachment`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *Igw) GetAttachmentOk() (*Attachment, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *Igw) SetAttachment(v Attachment)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *Igw) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### SetAttachmentNil

`func (o *Igw) SetAttachmentNil(b bool)`

 SetAttachmentNil sets the value for Attachment to be an explicit nil

### UnsetAttachment
`func (o *Igw) UnsetAttachment()`

UnsetAttachment ensures that no value is present for Attachment, not even an explicit nil
### GetNatIp

`func (o *Igw) GetNatIp() string`

GetNatIp returns the NatIp field if non-nil, zero value otherwise.

### GetNatIpOk

`func (o *Igw) GetNatIpOk() (*string, bool)`

GetNatIpOk returns a tuple with the NatIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatIp

`func (o *Igw) SetNatIp(v string)`

SetNatIp sets NatIp field to given value.

### HasNatIp

`func (o *Igw) HasNatIp() bool`

HasNatIp returns a boolean if a field has been set.

### SetNatIpNil

`func (o *Igw) SetNatIpNil(b bool)`

 SetNatIpNil sets the value for NatIp to be an explicit nil

### UnsetNatIp
`func (o *Igw) UnsetNatIp()`

UnsetNatIp ensures that no value is present for NatIp, not even an explicit nil
### GetProjectId

`func (o *Igw) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Igw) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Igw) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Igw) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *Igw) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *Igw) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectName

`func (o *Igw) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *Igw) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *Igw) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *Igw) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *Igw) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *Igw) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetDomainId

`func (o *Igw) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *Igw) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *Igw) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *Igw) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### SetDomainIdNil

`func (o *Igw) SetDomainIdNil(b bool)`

 SetDomainIdNil sets the value for DomainId to be an explicit nil

### UnsetDomainId
`func (o *Igw) UnsetDomainId()`

UnsetDomainId ensures that no value is present for DomainId, not even an explicit nil
### GetProvisioningStatus

`func (o *Igw) GetProvisioningStatus() string`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *Igw) GetProvisioningStatusOk() (*string, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *Igw) SetProvisioningStatus(v string)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *Igw) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *Igw) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *Igw) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetCreatedAt

`func (o *Igw) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Igw) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Igw) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Igw) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Igw) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Igw) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Igw) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Igw) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Igw) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Igw) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Igw) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Igw) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetAttachmentStatus

`func (o *Igw) GetAttachmentStatus() string`

GetAttachmentStatus returns the AttachmentStatus field if non-nil, zero value otherwise.

### GetAttachmentStatusOk

`func (o *Igw) GetAttachmentStatusOk() (*string, bool)`

GetAttachmentStatusOk returns a tuple with the AttachmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentStatus

`func (o *Igw) SetAttachmentStatus(v string)`

SetAttachmentStatus sets AttachmentStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


