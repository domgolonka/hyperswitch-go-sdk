# ArrayOfRoutingRulesInner

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethodsIncl** | [**[]PaymentMethod**](PaymentMethod.md) | The List of payment methods to include for this routing rule | [optional] [default to null]
**PaymentMethodsExcl** | [**[]PaymentMethod**](PaymentMethod.md) | The List of payment methods to exclude for this routing rule. If there is conflict between include and exclude lists, include list overrides the exclude list. | [optional] [default to null]
**PaymentMethodTypesIncl** | [**[]PaymentMethodType**](PaymentMethodType.md) | The List of payment method types to include for this routing rule | [optional] [default to null]
**PaymentMethodTypesExcl** | [**[]PaymentMethodType**](PaymentMethodType.md) | The List of payment method types to exclude for this routing rule. If there is conflict between include and exclude lists, include list overrides the exclude list. | [optional] [default to null]
**PaymentMethodIssuersIncl** | **[]string** | The List of payment method issuers to include for this routing rule | [optional] [default to null]
**PaymentMethodIssuersExcl** | **[]string** | The List of payment method issuers to exclude for this routing rule. If there is conflict between include and exclude lists, include list overrides the exclude list. | [optional] [default to null]
**CountriesIncl** | **[]string** | The List of countries to include for this routing rule | [optional] [default to null]
**CountriesExcl** | **[]string** | The List of countries to exclude for this routing rule. If there is conflict between include and exclude lists, include list overrides the exclude list. | [optional] [default to null]
**CurrenciesIncl** | **[]string** | The List of currencies to include for this routing rule | [optional] [default to null]
**CurrenciesExcl** | **[]string** | The List of currencies to exclude for this routing rule. If there is conflict between include and exclude lists, include list overrides the exclude list. | [optional] [default to null]
**MetadataFiltersKeys** | **[]string** | List of Metadata Filters to apply for the Routing Rule. The filters are presented as 2 arrays of keys and value. This property contains all the keys. | [optional] [default to null]
**MetadataFiltersValues** | **[]string** | List of Metadata Filters to apply for the Routing Rule. The filters are presented as 2 arrays of keys and value. This property contains all the keys. | [optional] [default to null]
**ConnectorsPeckingOrder** | [**[]PaymentConnector**](PaymentConnector.md) | The pecking order of payment connectors (or processors) to be used for routing. The first connector in the array will be attempted for routing. If it fails, the second connector will be used till the list is exhausted. | [optional] [default to null]
**ConnectorsTrafficWeightageKey** | [**[]PaymentConnector**](PaymentConnector.md) | An Array of Connectors (as Keys) with the associated percentage of traffic to be routed through the given connector (Expressed as an array of values) | [optional] [default to null]
**ConnectorsTrafficWeightageValue** | **[]int32** | An Array of Weightage (expressed in percentage) that needs to be associated with the respective connectors (Expressed as an array of keys) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

