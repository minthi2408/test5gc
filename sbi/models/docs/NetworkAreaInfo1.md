# NetworkAreaInfo1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ecgis** | Pointer to [**[]Ecgi1**](Ecgi1.md) | Contains a list of E-UTRA cell identities. | [optional] 
**Ncgis** | Pointer to [**[]Ncgi1**](Ncgi1.md) | Contains a list of NR cell identities. | [optional] 
**GRanNodeIds** | Pointer to [**[]GlobalRanNodeId1**](GlobalRanNodeId1.md) | Contains a list of NG RAN nodes. | [optional] 
**Tais** | Pointer to [**[]Tai1**](Tai1.md) | Contains a list of tracking area identities. | [optional] 

## Methods

### NewNetworkAreaInfo1

`func NewNetworkAreaInfo1() *NetworkAreaInfo1`

NewNetworkAreaInfo1 instantiates a new NetworkAreaInfo1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkAreaInfo1WithDefaults

`func NewNetworkAreaInfo1WithDefaults() *NetworkAreaInfo1`

NewNetworkAreaInfo1WithDefaults instantiates a new NetworkAreaInfo1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEcgis

`func (o *NetworkAreaInfo1) GetEcgis() []Ecgi1`

GetEcgis returns the Ecgis field if non-nil, zero value otherwise.

### GetEcgisOk

`func (o *NetworkAreaInfo1) GetEcgisOk() (*[]Ecgi1, bool)`

GetEcgisOk returns a tuple with the Ecgis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcgis

`func (o *NetworkAreaInfo1) SetEcgis(v []Ecgi1)`

SetEcgis sets Ecgis field to given value.

### HasEcgis

`func (o *NetworkAreaInfo1) HasEcgis() bool`

HasEcgis returns a boolean if a field has been set.

### GetNcgis

`func (o *NetworkAreaInfo1) GetNcgis() []Ncgi1`

GetNcgis returns the Ncgis field if non-nil, zero value otherwise.

### GetNcgisOk

`func (o *NetworkAreaInfo1) GetNcgisOk() (*[]Ncgi1, bool)`

GetNcgisOk returns a tuple with the Ncgis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcgis

`func (o *NetworkAreaInfo1) SetNcgis(v []Ncgi1)`

SetNcgis sets Ncgis field to given value.

### HasNcgis

`func (o *NetworkAreaInfo1) HasNcgis() bool`

HasNcgis returns a boolean if a field has been set.

### GetGRanNodeIds

`func (o *NetworkAreaInfo1) GetGRanNodeIds() []GlobalRanNodeId1`

GetGRanNodeIds returns the GRanNodeIds field if non-nil, zero value otherwise.

### GetGRanNodeIdsOk

`func (o *NetworkAreaInfo1) GetGRanNodeIdsOk() (*[]GlobalRanNodeId1, bool)`

GetGRanNodeIdsOk returns a tuple with the GRanNodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGRanNodeIds

`func (o *NetworkAreaInfo1) SetGRanNodeIds(v []GlobalRanNodeId1)`

SetGRanNodeIds sets GRanNodeIds field to given value.

### HasGRanNodeIds

`func (o *NetworkAreaInfo1) HasGRanNodeIds() bool`

HasGRanNodeIds returns a boolean if a field has been set.

### GetTais

`func (o *NetworkAreaInfo1) GetTais() []Tai1`

GetTais returns the Tais field if non-nil, zero value otherwise.

### GetTaisOk

`func (o *NetworkAreaInfo1) GetTaisOk() (*[]Tai1, bool)`

GetTaisOk returns a tuple with the Tais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTais

`func (o *NetworkAreaInfo1) SetTais(v []Tai1)`

SetTais sets Tais field to given value.

### HasTais

`func (o *NetworkAreaInfo1) HasTais() bool`

HasTais returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


