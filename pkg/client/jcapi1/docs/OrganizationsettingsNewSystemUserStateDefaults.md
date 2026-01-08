# OrganizationsettingsNewSystemUserStateDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationImport** | Pointer to **string** | The default user state for a user created using the [Bulk Users Create](https://docs.jumpcloud.com/api/2.0/index.html#operation/bulk_usersCreate) endpoint. See endpoint documentation for more details. | [optional] 
**CsvImport** | Pointer to **string** | The default user state for a user created using the [Bulk Users Create](https://docs.jumpcloud.com/api/2.0/index.html#operation/bulk_usersCreate) endpoint. See endpoint documentation for more details. | [optional] 
**ManualEntry** | Pointer to **string** | The default state for a user that is created using the [Create a system user](https://docs.jumpcloud.com/api/1.0/index.html#operation/systemusers_post) endpoint. See endpoint documentation for more details. | [optional] 

## Methods

### NewOrganizationsettingsNewSystemUserStateDefaults

`func NewOrganizationsettingsNewSystemUserStateDefaults() *OrganizationsettingsNewSystemUserStateDefaults`

NewOrganizationsettingsNewSystemUserStateDefaults instantiates a new OrganizationsettingsNewSystemUserStateDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsettingsNewSystemUserStateDefaultsWithDefaults

`func NewOrganizationsettingsNewSystemUserStateDefaultsWithDefaults() *OrganizationsettingsNewSystemUserStateDefaults`

NewOrganizationsettingsNewSystemUserStateDefaultsWithDefaults instantiates a new OrganizationsettingsNewSystemUserStateDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationImport

`func (o *OrganizationsettingsNewSystemUserStateDefaults) GetApplicationImport() string`

GetApplicationImport returns the ApplicationImport field if non-nil, zero value otherwise.

### GetApplicationImportOk

`func (o *OrganizationsettingsNewSystemUserStateDefaults) GetApplicationImportOk() (*string, bool)`

GetApplicationImportOk returns a tuple with the ApplicationImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationImport

`func (o *OrganizationsettingsNewSystemUserStateDefaults) SetApplicationImport(v string)`

SetApplicationImport sets ApplicationImport field to given value.

### HasApplicationImport

`func (o *OrganizationsettingsNewSystemUserStateDefaults) HasApplicationImport() bool`

HasApplicationImport returns a boolean if a field has been set.

### GetCsvImport

`func (o *OrganizationsettingsNewSystemUserStateDefaults) GetCsvImport() string`

GetCsvImport returns the CsvImport field if non-nil, zero value otherwise.

### GetCsvImportOk

`func (o *OrganizationsettingsNewSystemUserStateDefaults) GetCsvImportOk() (*string, bool)`

GetCsvImportOk returns a tuple with the CsvImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvImport

`func (o *OrganizationsettingsNewSystemUserStateDefaults) SetCsvImport(v string)`

SetCsvImport sets CsvImport field to given value.

### HasCsvImport

`func (o *OrganizationsettingsNewSystemUserStateDefaults) HasCsvImport() bool`

HasCsvImport returns a boolean if a field has been set.

### GetManualEntry

`func (o *OrganizationsettingsNewSystemUserStateDefaults) GetManualEntry() string`

GetManualEntry returns the ManualEntry field if non-nil, zero value otherwise.

### GetManualEntryOk

`func (o *OrganizationsettingsNewSystemUserStateDefaults) GetManualEntryOk() (*string, bool)`

GetManualEntryOk returns a tuple with the ManualEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualEntry

`func (o *OrganizationsettingsNewSystemUserStateDefaults) SetManualEntry(v string)`

SetManualEntry sets ManualEntry field to given value.

### HasManualEntry

`func (o *OrganizationsettingsNewSystemUserStateDefaults) HasManualEntry() bool`

HasManualEntry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


