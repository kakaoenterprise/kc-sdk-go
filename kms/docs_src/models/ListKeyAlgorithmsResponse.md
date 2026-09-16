# ListKeyAlgorithmsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithms** | [**[]Algorithm**](Algorithm.md) | 사용할 수 있는 키 알고리즘 목록 | 

## Methods

### NewListKeyAlgorithmsResponse

`func NewListKeyAlgorithmsResponse(algorithms []Algorithm, ) *ListKeyAlgorithmsResponse`

NewListKeyAlgorithmsResponse instantiates a new ListKeyAlgorithmsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKeyAlgorithmsResponseWithDefaults

`func NewListKeyAlgorithmsResponseWithDefaults() *ListKeyAlgorithmsResponse`

NewListKeyAlgorithmsResponseWithDefaults instantiates a new ListKeyAlgorithmsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithms

`func (o *ListKeyAlgorithmsResponse) GetAlgorithms() []Algorithm`

GetAlgorithms returns the Algorithms field if non-nil, zero value otherwise.

### GetAlgorithmsOk

`func (o *ListKeyAlgorithmsResponse) GetAlgorithmsOk() (*[]Algorithm, bool)`

GetAlgorithmsOk returns a tuple with the Algorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithms

`func (o *ListKeyAlgorithmsResponse) SetAlgorithms(v []Algorithm)`

SetAlgorithms sets Algorithms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


