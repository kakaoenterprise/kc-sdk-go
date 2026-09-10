# GetInstanceTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flavor** | [**Flavor**](Flavor.md) | 인스턴스 유형 상세 정보 | 

## Methods

### NewGetInstanceTypeResponse

`func NewGetInstanceTypeResponse(flavor Flavor, ) *GetInstanceTypeResponse`

NewGetInstanceTypeResponse instantiates a new GetInstanceTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceTypeResponseWithDefaults

`func NewGetInstanceTypeResponseWithDefaults() *GetInstanceTypeResponse`

NewGetInstanceTypeResponseWithDefaults instantiates a new GetInstanceTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavor

`func (o *GetInstanceTypeResponse) GetFlavor() Flavor`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *GetInstanceTypeResponse) GetFlavorOk() (*Flavor, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *GetInstanceTypeResponse) SetFlavor(v Flavor)`

SetFlavor sets Flavor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


