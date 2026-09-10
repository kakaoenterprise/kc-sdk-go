# TgwRouteTableAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Association ID | 
**ResourceId** | **string** | 연결된 리소스의 ID | 
**ResourceType** | [**ResourceType**](ResourceType.md) | 연결된 리소스 유형 | 
**TgwRouteTableId** | **string** | 연결된 TGW 라우팅 테이블 ID | 
**TgwAttachmentId** | **string** | Transit Gateway와 연결된 Attachment의 연결 ID | 
**ProvisioningStatus** | [**TGWAssociationProvisioningStatus**](TGWAssociationProvisioningStatus.md) | Association 프로비저닝 상태 | 
**ProjectId** | **string** | 프로젝트 ID | 

## Methods

### NewTgwRouteTableAssociation

`func NewTgwRouteTableAssociation(id string, resourceId string, resourceType ResourceType, tgwRouteTableId string, tgwAttachmentId string, provisioningStatus TGWAssociationProvisioningStatus, projectId string, ) *TgwRouteTableAssociation`

NewTgwRouteTableAssociation instantiates a new TgwRouteTableAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTgwRouteTableAssociationWithDefaults

`func NewTgwRouteTableAssociationWithDefaults() *TgwRouteTableAssociation`

NewTgwRouteTableAssociationWithDefaults instantiates a new TgwRouteTableAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TgwRouteTableAssociation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TgwRouteTableAssociation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TgwRouteTableAssociation) SetId(v string)`

SetId sets Id field to given value.


### GetResourceId

`func (o *TgwRouteTableAssociation) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *TgwRouteTableAssociation) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *TgwRouteTableAssociation) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *TgwRouteTableAssociation) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *TgwRouteTableAssociation) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *TgwRouteTableAssociation) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.


### GetTgwRouteTableId

`func (o *TgwRouteTableAssociation) GetTgwRouteTableId() string`

GetTgwRouteTableId returns the TgwRouteTableId field if non-nil, zero value otherwise.

### GetTgwRouteTableIdOk

`func (o *TgwRouteTableAssociation) GetTgwRouteTableIdOk() (*string, bool)`

GetTgwRouteTableIdOk returns a tuple with the TgwRouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwRouteTableId

`func (o *TgwRouteTableAssociation) SetTgwRouteTableId(v string)`

SetTgwRouteTableId sets TgwRouteTableId field to given value.


### GetTgwAttachmentId

`func (o *TgwRouteTableAssociation) GetTgwAttachmentId() string`

GetTgwAttachmentId returns the TgwAttachmentId field if non-nil, zero value otherwise.

### GetTgwAttachmentIdOk

`func (o *TgwRouteTableAssociation) GetTgwAttachmentIdOk() (*string, bool)`

GetTgwAttachmentIdOk returns a tuple with the TgwAttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwAttachmentId

`func (o *TgwRouteTableAssociation) SetTgwAttachmentId(v string)`

SetTgwAttachmentId sets TgwAttachmentId field to given value.


### GetProvisioningStatus

`func (o *TgwRouteTableAssociation) GetProvisioningStatus() TGWAssociationProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *TgwRouteTableAssociation) GetProvisioningStatusOk() (*TGWAssociationProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *TgwRouteTableAssociation) SetProvisioningStatus(v TGWAssociationProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetProjectId

`func (o *TgwRouteTableAssociation) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TgwRouteTableAssociation) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TgwRouteTableAssociation) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


