# ResizeInstanceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 변경할 인스턴스 유형 ID &lt;br/&gt; - [List instance types](https://docs.kakaocloud.com/openapi/bcs/list-instance-types) API를 통해 확인 가능 &lt;br/&gt; - [인스턴스 유형별 사양](https://docs.kakaocloud.com/service/bcs/bcs-instance/bcs-type) 참고 | 

## Methods

### NewResizeInstanceModel

`func NewResizeInstanceModel(id string, ) *ResizeInstanceModel`

NewResizeInstanceModel instantiates a new ResizeInstanceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResizeInstanceModelWithDefaults

`func NewResizeInstanceModelWithDefaults() *ResizeInstanceModel`

NewResizeInstanceModelWithDefaults instantiates a new ResizeInstanceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResizeInstanceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResizeInstanceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResizeInstanceModel) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


