# Cors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedOrigins** | **[]string** |  | 

## Methods

### NewCors

`func NewCors(allowedOrigins []string, ) *Cors`

NewCors instantiates a new Cors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorsWithDefaults

`func NewCorsWithDefaults() *Cors`

NewCorsWithDefaults instantiates a new Cors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedOrigins

`func (o *Cors) GetAllowedOrigins() []string`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *Cors) GetAllowedOriginsOk() (*[]string, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *Cors) SetAllowedOrigins(v []string)`

SetAllowedOrigins sets AllowedOrigins field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


