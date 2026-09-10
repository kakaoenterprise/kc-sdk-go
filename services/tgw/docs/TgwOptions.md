# TgwOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAutoAcceptSharedAttachments** | **bool** | 공유 Attachment 자동 승인 여부 | 
**IsDefaultRouteTableAssociation** | **bool** | 기본 라우팅 테이블 자동 연결 여부 | 
**AssociationDefaultRouteTableId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTgwOptions

`func NewTgwOptions(isAutoAcceptSharedAttachments bool, isDefaultRouteTableAssociation bool, ) *TgwOptions`

NewTgwOptions instantiates a new TgwOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTgwOptionsWithDefaults

`func NewTgwOptionsWithDefaults() *TgwOptions`

NewTgwOptionsWithDefaults instantiates a new TgwOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAutoAcceptSharedAttachments

`func (o *TgwOptions) GetIsAutoAcceptSharedAttachments() bool`

GetIsAutoAcceptSharedAttachments returns the IsAutoAcceptSharedAttachments field if non-nil, zero value otherwise.

### GetIsAutoAcceptSharedAttachmentsOk

`func (o *TgwOptions) GetIsAutoAcceptSharedAttachmentsOk() (*bool, bool)`

GetIsAutoAcceptSharedAttachmentsOk returns a tuple with the IsAutoAcceptSharedAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoAcceptSharedAttachments

`func (o *TgwOptions) SetIsAutoAcceptSharedAttachments(v bool)`

SetIsAutoAcceptSharedAttachments sets IsAutoAcceptSharedAttachments field to given value.


### GetIsDefaultRouteTableAssociation

`func (o *TgwOptions) GetIsDefaultRouteTableAssociation() bool`

GetIsDefaultRouteTableAssociation returns the IsDefaultRouteTableAssociation field if non-nil, zero value otherwise.

### GetIsDefaultRouteTableAssociationOk

`func (o *TgwOptions) GetIsDefaultRouteTableAssociationOk() (*bool, bool)`

GetIsDefaultRouteTableAssociationOk returns a tuple with the IsDefaultRouteTableAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultRouteTableAssociation

`func (o *TgwOptions) SetIsDefaultRouteTableAssociation(v bool)`

SetIsDefaultRouteTableAssociation sets IsDefaultRouteTableAssociation field to given value.


### GetAssociationDefaultRouteTableId

`func (o *TgwOptions) GetAssociationDefaultRouteTableId() string`

GetAssociationDefaultRouteTableId returns the AssociationDefaultRouteTableId field if non-nil, zero value otherwise.

### GetAssociationDefaultRouteTableIdOk

`func (o *TgwOptions) GetAssociationDefaultRouteTableIdOk() (*string, bool)`

GetAssociationDefaultRouteTableIdOk returns a tuple with the AssociationDefaultRouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationDefaultRouteTableId

`func (o *TgwOptions) SetAssociationDefaultRouteTableId(v string)`

SetAssociationDefaultRouteTableId sets AssociationDefaultRouteTableId field to given value.

### HasAssociationDefaultRouteTableId

`func (o *TgwOptions) HasAssociationDefaultRouteTableId() bool`

HasAssociationDefaultRouteTableId returns a boolean if a field has been set.

### SetAssociationDefaultRouteTableIdNil

`func (o *TgwOptions) SetAssociationDefaultRouteTableIdNil(b bool)`

 SetAssociationDefaultRouteTableIdNil sets the value for AssociationDefaultRouteTableId to be an explicit nil

### UnsetAssociationDefaultRouteTableId
`func (o *TgwOptions) UnsetAssociationDefaultRouteTableId()`

UnsetAssociationDefaultRouteTableId ensures that no value is present for AssociationDefaultRouteTableId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


