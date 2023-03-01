# Search

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to **map[string]interface{}** |  | [optional] 
**SearchFilter** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSearch

`func NewSearch() *Search`

NewSearch instantiates a new Search object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchWithDefaults

`func NewSearchWithDefaults() *Search`

NewSearchWithDefaults instantiates a new Search object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *Search) GetFields() string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Search) GetFieldsOk() (*string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Search) SetFields(v string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Search) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFilter

`func (o *Search) GetFilter() map[string]interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Search) GetFilterOk() (*map[string]interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Search) SetFilter(v map[string]interface{})`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *Search) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSearchFilter

`func (o *Search) GetSearchFilter() map[string]interface{}`

GetSearchFilter returns the SearchFilter field if non-nil, zero value otherwise.

### GetSearchFilterOk

`func (o *Search) GetSearchFilterOk() (*map[string]interface{}, bool)`

GetSearchFilterOk returns a tuple with the SearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilter

`func (o *Search) SetSearchFilter(v map[string]interface{})`

SetSearchFilter sets SearchFilter field to given value.

### HasSearchFilter

`func (o *Search) HasSearchFilter() bool`

HasSearchFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


