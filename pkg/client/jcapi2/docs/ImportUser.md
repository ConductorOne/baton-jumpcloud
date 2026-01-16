# ImportUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**[]ImportUserAddress**](ImportUserAddress.md) |  | [optional] 
**AlternateEmail** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**CostCenter** | Pointer to **string** |  | [optional] 
**Department** | Pointer to **string** |  | [optional] 
**Displayname** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EmployeeIdentifier** | Pointer to **string** |  | [optional] 
**EmployeeType** | Pointer to **string** |  | [optional] 
**Firstname** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**JobTitle** | Pointer to **string** |  | [optional] 
**Lastname** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Manager** | Pointer to **string** |  | [optional] 
**Middlename** | Pointer to **string** |  | [optional] 
**PhoneNumbers** | Pointer to [**[]ImportUserPhoneNumber**](ImportUserPhoneNumber.md) |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewImportUser

`func NewImportUser() *ImportUser`

NewImportUser instantiates a new ImportUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportUserWithDefaults

`func NewImportUserWithDefaults() *ImportUser`

NewImportUserWithDefaults instantiates a new ImportUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ImportUser) GetAddresses() []ImportUserAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ImportUser) GetAddressesOk() (*[]ImportUserAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ImportUser) SetAddresses(v []ImportUserAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *ImportUser) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAlternateEmail

`func (o *ImportUser) GetAlternateEmail() string`

GetAlternateEmail returns the AlternateEmail field if non-nil, zero value otherwise.

### GetAlternateEmailOk

`func (o *ImportUser) GetAlternateEmailOk() (*string, bool)`

GetAlternateEmailOk returns a tuple with the AlternateEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateEmail

`func (o *ImportUser) SetAlternateEmail(v string)`

SetAlternateEmail sets AlternateEmail field to given value.

### HasAlternateEmail

`func (o *ImportUser) HasAlternateEmail() bool`

HasAlternateEmail returns a boolean if a field has been set.

### GetCompany

`func (o *ImportUser) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ImportUser) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ImportUser) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ImportUser) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCostCenter

`func (o *ImportUser) GetCostCenter() string`

GetCostCenter returns the CostCenter field if non-nil, zero value otherwise.

### GetCostCenterOk

`func (o *ImportUser) GetCostCenterOk() (*string, bool)`

GetCostCenterOk returns a tuple with the CostCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCenter

`func (o *ImportUser) SetCostCenter(v string)`

SetCostCenter sets CostCenter field to given value.

### HasCostCenter

`func (o *ImportUser) HasCostCenter() bool`

HasCostCenter returns a boolean if a field has been set.

### GetDepartment

`func (o *ImportUser) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *ImportUser) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *ImportUser) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *ImportUser) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetDisplayname

`func (o *ImportUser) GetDisplayname() string`

GetDisplayname returns the Displayname field if non-nil, zero value otherwise.

### GetDisplaynameOk

`func (o *ImportUser) GetDisplaynameOk() (*string, bool)`

GetDisplaynameOk returns a tuple with the Displayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayname

`func (o *ImportUser) SetDisplayname(v string)`

SetDisplayname sets Displayname field to given value.

### HasDisplayname

`func (o *ImportUser) HasDisplayname() bool`

HasDisplayname returns a boolean if a field has been set.

### GetEmail

`func (o *ImportUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ImportUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ImportUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ImportUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmployeeIdentifier

`func (o *ImportUser) GetEmployeeIdentifier() string`

GetEmployeeIdentifier returns the EmployeeIdentifier field if non-nil, zero value otherwise.

### GetEmployeeIdentifierOk

`func (o *ImportUser) GetEmployeeIdentifierOk() (*string, bool)`

GetEmployeeIdentifierOk returns a tuple with the EmployeeIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeIdentifier

`func (o *ImportUser) SetEmployeeIdentifier(v string)`

SetEmployeeIdentifier sets EmployeeIdentifier field to given value.

### HasEmployeeIdentifier

`func (o *ImportUser) HasEmployeeIdentifier() bool`

HasEmployeeIdentifier returns a boolean if a field has been set.

### GetEmployeeType

`func (o *ImportUser) GetEmployeeType() string`

GetEmployeeType returns the EmployeeType field if non-nil, zero value otherwise.

### GetEmployeeTypeOk

`func (o *ImportUser) GetEmployeeTypeOk() (*string, bool)`

GetEmployeeTypeOk returns a tuple with the EmployeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeType

`func (o *ImportUser) SetEmployeeType(v string)`

SetEmployeeType sets EmployeeType field to given value.

### HasEmployeeType

`func (o *ImportUser) HasEmployeeType() bool`

HasEmployeeType returns a boolean if a field has been set.

### GetFirstname

`func (o *ImportUser) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *ImportUser) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *ImportUser) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *ImportUser) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetId

`func (o *ImportUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImportUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobTitle

`func (o *ImportUser) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *ImportUser) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *ImportUser) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *ImportUser) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetLastname

`func (o *ImportUser) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *ImportUser) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *ImportUser) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *ImportUser) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetLocation

`func (o *ImportUser) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ImportUser) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ImportUser) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ImportUser) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetManager

`func (o *ImportUser) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *ImportUser) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *ImportUser) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *ImportUser) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetMiddlename

`func (o *ImportUser) GetMiddlename() string`

GetMiddlename returns the Middlename field if non-nil, zero value otherwise.

### GetMiddlenameOk

`func (o *ImportUser) GetMiddlenameOk() (*string, bool)`

GetMiddlenameOk returns a tuple with the Middlename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddlename

`func (o *ImportUser) SetMiddlename(v string)`

SetMiddlename sets Middlename field to given value.

### HasMiddlename

`func (o *ImportUser) HasMiddlename() bool`

HasMiddlename returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *ImportUser) GetPhoneNumbers() []ImportUserPhoneNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *ImportUser) GetPhoneNumbersOk() (*[]ImportUserPhoneNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *ImportUser) SetPhoneNumbers(v []ImportUserPhoneNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *ImportUser) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetUsername

`func (o *ImportUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ImportUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ImportUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ImportUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


