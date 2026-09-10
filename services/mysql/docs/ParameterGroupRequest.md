# ParameterGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ParameterGroupType**](ParameterGroupType.md) | 적용할 파라미터 그룹의 유형 | 
**Id** | **string** | 적용할 파라미터 그룹의 ID &lt;br/&gt;- &#x60;type&#x3D;DEFAULT&#x60;: [List MySQL default parameter groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-default-parameter-groups)에서 확인 &lt;br/&gt;- &#x60;type&#x3D;CUSTOM&#x60;: [List MySQL custom parameter groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-custom-parameter-groups)에서 확인 | 

## Methods

### NewParameterGroupRequest

`func NewParameterGroupRequest(type_ ParameterGroupType, id string, ) *ParameterGroupRequest`

NewParameterGroupRequest instantiates a new ParameterGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterGroupRequestWithDefaults

`func NewParameterGroupRequestWithDefaults() *ParameterGroupRequest`

NewParameterGroupRequestWithDefaults instantiates a new ParameterGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ParameterGroupRequest) GetType() ParameterGroupType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParameterGroupRequest) GetTypeOk() (*ParameterGroupType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParameterGroupRequest) SetType(v ParameterGroupType)`

SetType sets Type field to given value.


### GetId

`func (o *ParameterGroupRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParameterGroupRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParameterGroupRequest) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


