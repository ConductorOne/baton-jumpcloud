# GraphOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ObjectID of graph object being added or removed as an association. | 
**Op** | **string** | How to modify the graph connection. | 

## Methods

### NewGraphOperation

`func NewGraphOperation(id string, op string, ) *GraphOperation`

NewGraphOperation instantiates a new GraphOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphOperationWithDefaults

`func NewGraphOperationWithDefaults() *GraphOperation`

NewGraphOperationWithDefaults instantiates a new GraphOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GraphOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphOperation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphOperation) SetId(v string)`

SetId sets Id field to given value.


### GetOp

`func (o *GraphOperation) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *GraphOperation) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *GraphOperation) SetOp(v string)`

SetOp sets Op field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


