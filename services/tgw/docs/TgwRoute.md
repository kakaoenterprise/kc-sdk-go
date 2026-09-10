# TgwRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Route ID | 
**DestinationCidrBlock** | **string** | 목적지 CIDR 블록 (예: 10.0.0.0/16) | 
**TgwRouteTableId** | **string** | Route가 속한 Transit Gateway 라우팅 테이블 ID | 
**TgwRouteAttachments** | [**[]TgwRouteAttachment**](TgwRouteAttachment.md) | Route에 연결된 Transit Gateway Attachment 목록 | 
**ProvisioningStatus** | [**TGWRouteProvisioningStatus**](TGWRouteProvisioningStatus.md) | 프로비저닝 상태 | 
**ProjectId** | **string** | Route가 속한 프로젝트 ID | 

## Methods

### NewTgwRoute

`func NewTgwRoute(id string, destinationCidrBlock string, tgwRouteTableId string, tgwRouteAttachments []TgwRouteAttachment, provisioningStatus TGWRouteProvisioningStatus, projectId string, ) *TgwRoute`

NewTgwRoute instantiates a new TgwRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTgwRouteWithDefaults

`func NewTgwRouteWithDefaults() *TgwRoute`

NewTgwRouteWithDefaults instantiates a new TgwRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TgwRoute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TgwRoute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TgwRoute) SetId(v string)`

SetId sets Id field to given value.


### GetDestinationCidrBlock

`func (o *TgwRoute) GetDestinationCidrBlock() string`

GetDestinationCidrBlock returns the DestinationCidrBlock field if non-nil, zero value otherwise.

### GetDestinationCidrBlockOk

`func (o *TgwRoute) GetDestinationCidrBlockOk() (*string, bool)`

GetDestinationCidrBlockOk returns a tuple with the DestinationCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCidrBlock

`func (o *TgwRoute) SetDestinationCidrBlock(v string)`

SetDestinationCidrBlock sets DestinationCidrBlock field to given value.


### GetTgwRouteTableId

`func (o *TgwRoute) GetTgwRouteTableId() string`

GetTgwRouteTableId returns the TgwRouteTableId field if non-nil, zero value otherwise.

### GetTgwRouteTableIdOk

`func (o *TgwRoute) GetTgwRouteTableIdOk() (*string, bool)`

GetTgwRouteTableIdOk returns a tuple with the TgwRouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwRouteTableId

`func (o *TgwRoute) SetTgwRouteTableId(v string)`

SetTgwRouteTableId sets TgwRouteTableId field to given value.


### GetTgwRouteAttachments

`func (o *TgwRoute) GetTgwRouteAttachments() []TgwRouteAttachment`

GetTgwRouteAttachments returns the TgwRouteAttachments field if non-nil, zero value otherwise.

### GetTgwRouteAttachmentsOk

`func (o *TgwRoute) GetTgwRouteAttachmentsOk() (*[]TgwRouteAttachment, bool)`

GetTgwRouteAttachmentsOk returns a tuple with the TgwRouteAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwRouteAttachments

`func (o *TgwRoute) SetTgwRouteAttachments(v []TgwRouteAttachment)`

SetTgwRouteAttachments sets TgwRouteAttachments field to given value.


### GetProvisioningStatus

`func (o *TgwRoute) GetProvisioningStatus() TGWRouteProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *TgwRoute) GetProvisioningStatusOk() (*TGWRouteProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *TgwRoute) SetProvisioningStatus(v TGWRouteProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetProjectId

`func (o *TgwRoute) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TgwRoute) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TgwRoute) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


