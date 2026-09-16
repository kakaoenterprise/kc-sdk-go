# Option

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAutoAcceptSharedAttachments** | Pointer to **NullableBool** | 공유 Attachment 자동 승인 여부 | [optional] 
**IsDefaultRouteTableAssociation** | Pointer to **NullableBool** | 기본 라우팅 테이블 자동 연결 여부 | [optional] 
**AssociationDefaultRouteTableId** | Pointer to **NullableString** | 기본 연결 라우팅 테이블 ID | [optional] 

## Methods

### NewOption

`func NewOption() *Option`

NewOption instantiates a new Option object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionWithDefaults

`func NewOptionWithDefaults() *Option`

NewOptionWithDefaults instantiates a new Option object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAutoAcceptSharedAttachments

`func (o *Option) GetIsAutoAcceptSharedAttachments() bool`

GetIsAutoAcceptSharedAttachments returns the IsAutoAcceptSharedAttachments field if non-nil, zero value otherwise.

### GetIsAutoAcceptSharedAttachmentsOk

`func (o *Option) GetIsAutoAcceptSharedAttachmentsOk() (*bool, bool)`

GetIsAutoAcceptSharedAttachmentsOk returns a tuple with the IsAutoAcceptSharedAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoAcceptSharedAttachments

`func (o *Option) SetIsAutoAcceptSharedAttachments(v bool)`

SetIsAutoAcceptSharedAttachments sets IsAutoAcceptSharedAttachments field to given value.

### HasIsAutoAcceptSharedAttachments

`func (o *Option) HasIsAutoAcceptSharedAttachments() bool`

HasIsAutoAcceptSharedAttachments returns a boolean if a field has been set.

### SetIsAutoAcceptSharedAttachmentsNil

`func (o *Option) SetIsAutoAcceptSharedAttachmentsNil(b bool)`

 SetIsAutoAcceptSharedAttachmentsNil sets the value for IsAutoAcceptSharedAttachments to be an explicit nil

### UnsetIsAutoAcceptSharedAttachments
`func (o *Option) UnsetIsAutoAcceptSharedAttachments()`

UnsetIsAutoAcceptSharedAttachments ensures that no value is present for IsAutoAcceptSharedAttachments, not even an explicit nil
### GetIsDefaultRouteTableAssociation

`func (o *Option) GetIsDefaultRouteTableAssociation() bool`

GetIsDefaultRouteTableAssociation returns the IsDefaultRouteTableAssociation field if non-nil, zero value otherwise.

### GetIsDefaultRouteTableAssociationOk

`func (o *Option) GetIsDefaultRouteTableAssociationOk() (*bool, bool)`

GetIsDefaultRouteTableAssociationOk returns a tuple with the IsDefaultRouteTableAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultRouteTableAssociation

`func (o *Option) SetIsDefaultRouteTableAssociation(v bool)`

SetIsDefaultRouteTableAssociation sets IsDefaultRouteTableAssociation field to given value.

### HasIsDefaultRouteTableAssociation

`func (o *Option) HasIsDefaultRouteTableAssociation() bool`

HasIsDefaultRouteTableAssociation returns a boolean if a field has been set.

### SetIsDefaultRouteTableAssociationNil

`func (o *Option) SetIsDefaultRouteTableAssociationNil(b bool)`

 SetIsDefaultRouteTableAssociationNil sets the value for IsDefaultRouteTableAssociation to be an explicit nil

### UnsetIsDefaultRouteTableAssociation
`func (o *Option) UnsetIsDefaultRouteTableAssociation()`

UnsetIsDefaultRouteTableAssociation ensures that no value is present for IsDefaultRouteTableAssociation, not even an explicit nil
### GetAssociationDefaultRouteTableId

`func (o *Option) GetAssociationDefaultRouteTableId() string`

GetAssociationDefaultRouteTableId returns the AssociationDefaultRouteTableId field if non-nil, zero value otherwise.

### GetAssociationDefaultRouteTableIdOk

`func (o *Option) GetAssociationDefaultRouteTableIdOk() (*string, bool)`

GetAssociationDefaultRouteTableIdOk returns a tuple with the AssociationDefaultRouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationDefaultRouteTableId

`func (o *Option) SetAssociationDefaultRouteTableId(v string)`

SetAssociationDefaultRouteTableId sets AssociationDefaultRouteTableId field to given value.

### HasAssociationDefaultRouteTableId

`func (o *Option) HasAssociationDefaultRouteTableId() bool`

HasAssociationDefaultRouteTableId returns a boolean if a field has been set.

### SetAssociationDefaultRouteTableIdNil

`func (o *Option) SetAssociationDefaultRouteTableIdNil(b bool)`

 SetAssociationDefaultRouteTableIdNil sets the value for AssociationDefaultRouteTableId to be an explicit nil

### UnsetAssociationDefaultRouteTableId
`func (o *Option) UnsetAssociationDefaultRouteTableId()`

UnsetAssociationDefaultRouteTableId ensures that no value is present for AssociationDefaultRouteTableId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


