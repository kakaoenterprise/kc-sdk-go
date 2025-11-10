# TgwOptionsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAutoAcceptSharedAttachments** | **bool** | 공유 Attachment 자동 승인 여부 | 
**IsDefaultRouteTableAssociation** | **bool** | 기본 라우팅 테이블 자동 연결 여부 | 
**AssociationDefaultRouteTableId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTgwOptionsResponseModel

`func NewTgwOptionsResponseModel(isAutoAcceptSharedAttachments bool, isDefaultRouteTableAssociation bool, ) *TgwOptionsResponseModel`

NewTgwOptionsResponseModel instantiates a new TgwOptionsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTgwOptionsResponseModelWithDefaults

`func NewTgwOptionsResponseModelWithDefaults() *TgwOptionsResponseModel`

NewTgwOptionsResponseModelWithDefaults instantiates a new TgwOptionsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAutoAcceptSharedAttachments

`func (o *TgwOptionsResponseModel) GetIsAutoAcceptSharedAttachments() bool`

GetIsAutoAcceptSharedAttachments returns the IsAutoAcceptSharedAttachments field if non-nil, zero value otherwise.

### GetIsAutoAcceptSharedAttachmentsOk

`func (o *TgwOptionsResponseModel) GetIsAutoAcceptSharedAttachmentsOk() (*bool, bool)`

GetIsAutoAcceptSharedAttachmentsOk returns a tuple with the IsAutoAcceptSharedAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoAcceptSharedAttachments

`func (o *TgwOptionsResponseModel) SetIsAutoAcceptSharedAttachments(v bool)`

SetIsAutoAcceptSharedAttachments sets IsAutoAcceptSharedAttachments field to given value.


### GetIsDefaultRouteTableAssociation

`func (o *TgwOptionsResponseModel) GetIsDefaultRouteTableAssociation() bool`

GetIsDefaultRouteTableAssociation returns the IsDefaultRouteTableAssociation field if non-nil, zero value otherwise.

### GetIsDefaultRouteTableAssociationOk

`func (o *TgwOptionsResponseModel) GetIsDefaultRouteTableAssociationOk() (*bool, bool)`

GetIsDefaultRouteTableAssociationOk returns a tuple with the IsDefaultRouteTableAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultRouteTableAssociation

`func (o *TgwOptionsResponseModel) SetIsDefaultRouteTableAssociation(v bool)`

SetIsDefaultRouteTableAssociation sets IsDefaultRouteTableAssociation field to given value.


### GetAssociationDefaultRouteTableId

`func (o *TgwOptionsResponseModel) GetAssociationDefaultRouteTableId() string`

GetAssociationDefaultRouteTableId returns the AssociationDefaultRouteTableId field if non-nil, zero value otherwise.

### GetAssociationDefaultRouteTableIdOk

`func (o *TgwOptionsResponseModel) GetAssociationDefaultRouteTableIdOk() (*string, bool)`

GetAssociationDefaultRouteTableIdOk returns a tuple with the AssociationDefaultRouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationDefaultRouteTableId

`func (o *TgwOptionsResponseModel) SetAssociationDefaultRouteTableId(v string)`

SetAssociationDefaultRouteTableId sets AssociationDefaultRouteTableId field to given value.

### HasAssociationDefaultRouteTableId

`func (o *TgwOptionsResponseModel) HasAssociationDefaultRouteTableId() bool`

HasAssociationDefaultRouteTableId returns a boolean if a field has been set.

### SetAssociationDefaultRouteTableIdNil

`func (o *TgwOptionsResponseModel) SetAssociationDefaultRouteTableIdNil(b bool)`

 SetAssociationDefaultRouteTableIdNil sets the value for AssociationDefaultRouteTableId to be an explicit nil

### UnsetAssociationDefaultRouteTableId
`func (o *TgwOptionsResponseModel) UnsetAssociationDefaultRouteTableId()`

UnsetAssociationDefaultRouteTableId ensures that no value is present for AssociationDefaultRouteTableId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


