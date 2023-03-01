# GraphObjectWithPaths

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompiledAttributes** | Pointer to **map[string]interface{}** | The graph attributes. | [optional] 
**Id** | **string** | Object ID of this graph object. | 
**Paths** | [**[][]GraphConnection**]([]GraphConnection.md) | A path through the graph between two graph objects. | 
**Type** | [**GraphType**](GraphType.md) |  | 

## Methods

### NewGraphObjectWithPaths

`func NewGraphObjectWithPaths(id string, paths [][]GraphConnection, type_ GraphType, ) *GraphObjectWithPaths`

NewGraphObjectWithPaths instantiates a new GraphObjectWithPaths object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphObjectWithPathsWithDefaults

`func NewGraphObjectWithPathsWithDefaults() *GraphObjectWithPaths`

NewGraphObjectWithPathsWithDefaults instantiates a new GraphObjectWithPaths object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompiledAttributes

`func (o *GraphObjectWithPaths) GetCompiledAttributes() map[string]interface{}`

GetCompiledAttributes returns the CompiledAttributes field if non-nil, zero value otherwise.

### GetCompiledAttributesOk

`func (o *GraphObjectWithPaths) GetCompiledAttributesOk() (*map[string]interface{}, bool)`

GetCompiledAttributesOk returns a tuple with the CompiledAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompiledAttributes

`func (o *GraphObjectWithPaths) SetCompiledAttributes(v map[string]interface{})`

SetCompiledAttributes sets CompiledAttributes field to given value.

### HasCompiledAttributes

`func (o *GraphObjectWithPaths) HasCompiledAttributes() bool`

HasCompiledAttributes returns a boolean if a field has been set.

### GetId

`func (o *GraphObjectWithPaths) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphObjectWithPaths) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphObjectWithPaths) SetId(v string)`

SetId sets Id field to given value.


### GetPaths

`func (o *GraphObjectWithPaths) GetPaths() [][]GraphConnection`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *GraphObjectWithPaths) GetPathsOk() (*[][]GraphConnection, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *GraphObjectWithPaths) SetPaths(v [][]GraphConnection)`

SetPaths sets Paths field to given value.


### GetType

`func (o *GraphObjectWithPaths) GetType() GraphType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphObjectWithPaths) GetTypeOk() (*GraphType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphObjectWithPaths) SetType(v GraphType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


