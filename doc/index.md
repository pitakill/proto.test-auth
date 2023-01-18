# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [oyster/entities/pb/payment/domain/v1/currency.proto](#oyster_entities_pb_payment_domain_v1_currency-proto)
    - [CurrencyRate](#oyster-entities-pb-payment-domain-v1-CurrencyRate)
  
    - [Currency](#oyster-entities-pb-payment-domain-v1-Currency)
  
- [oyster/entities/pb/payment/domain/v1/payment.proto](#oyster_entities_pb_payment_domain_v1_payment-proto)
    - [Payment](#oyster-entities-pb-payment-domain-v1-Payment)
  
- [oyster/entities/pb/v1/commons.proto](#oyster_entities_pb_v1_commons-proto)
    - [Coordinates](#oyster-entities-pb-v1-Coordinates)
    - [Debug](#oyster-entities-pb-v1-Debug)
    - [Permission](#oyster-entities-pb-v1-Permission)
  
- [oyster/entities/pb/v1/auth.proto](#oyster_entities_pb_v1_auth-proto)
    - [Authorization](#oyster-entities-pb-v1-Authorization)
  
- [oyster/entities/pb/payment/v1/payment_detail_request.proto](#oyster_entities_pb_payment_v1_payment_detail_request-proto)
    - [GetPaymentTransactionRequest](#oyster-entities-pb-payment-v1-GetPaymentTransactionRequest)
  
- [oyster/entities/pb/v1/errors.proto](#oyster_entities_pb_v1_errors-proto)
    - [ErrorMessage](#oyster-entities-pb-v1-ErrorMessage)
  
    - [ErrorCode](#oyster-entities-pb-v1-ErrorCode)
  
- [oyster/entities/pb/payment/v1/payment_detail_response.proto](#oyster_entities_pb_payment_v1_payment_detail_response-proto)
    - [GetPaymentTransactionResponse](#oyster-entities-pb-payment-v1-GetPaymentTransactionResponse)
  
- [oyster/entities/pb/payment/v1/process_payment_request.proto](#oyster_entities_pb_payment_v1_process_payment_request-proto)
    - [MakePaymentRequest](#oyster-entities-pb-payment-v1-MakePaymentRequest)
  
- [oyster/entities/pb/payment/v1/process_payment_response.proto](#oyster_entities_pb_payment_v1_process_payment_response-proto)
    - [MakePaymentResponse](#oyster-entities-pb-payment-v1-MakePaymentResponse)
  
- [oyster/rpc/payment/v1/payment_info_rpc.proto](#oyster_rpc_payment_v1_payment_info_rpc-proto)
    - [PaymentInfoService](#oyster-rpc-payment-v1-PaymentInfoService)
  
- [oyster/rpc/payment/v1/process_payment_rpc.proto](#oyster_rpc_payment_v1_process_payment_rpc-proto)
    - [ProcessPaymentService](#oyster-rpc-payment-v1-ProcessPaymentService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="oyster_entities_pb_payment_domain_v1_currency-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## oyster/entities/pb/payment/domain/v1/currency.proto



<a name="oyster-entities-pb-payment-domain-v1-CurrencyRate"></a>

### CurrencyRate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| currency | [Currency](#oyster-entities-pb-payment-domain-v1-Currency) |  |  |
| rate | [double](#double) |  |  |





 


<a name="oyster-entities-pb-payment-domain-v1-Currency"></a>

### Currency


| Name | Number | Description |
| ---- | ------ | ----------- |
| CURRENCY_UNSPECIFIED | 0 |  |
| CURRENCY_MXN | 1 |  |
| CURRENCY_USD | 2 |  |
| CURRENCY_CAD | 3 |  |
| CURRENCY_BRL | 4 |  |
| CURRENCY_COP | 5 |  |


 

 

 



<a name="oyster_entities_pb_payment_domain_v1_payment-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## oyster/entities/pb/payment/domain/v1/payment.proto



<a name="oyster-entities-pb-payment-domain-v1-Payment"></a>

### Payment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_transaction_id | [string](#string) |  |  |
| business_id | [string](#string) |  |  |
| gross_amount | [string](#string) |  |  |
| requester_ref_id | [string](#string) |  |  |
| currency | [Currency](#oyster-entities-pb-payment-domain-v1-Currency) |  |  |
| payment_method | [string](#string) |  |  |
| status | [string](#string) |  |  |
| fee | [string](#string) |  |  |
| net_amount | [string](#string) |  |  |
| settled_at | [string](#string) |  |  |
| created_at | [string](#string) |  |  |
| updated_at | [string](#string) |  |  |





 

 

 

 



<a name="oyster_entities_pb_v1_commons-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## oyster/entities/pb/v1/commons.proto
Metadata related messages about permission and debug

This file have the Permission and Debug messages as base for all the gRPC communications


<a name="oyster-entities-pb-v1-Coordinates"></a>

### Coordinates
Coordinates message to hold information about the client location


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lat | [string](#string) |  | latitude |
| lon | [string](#string) |  | longitude |






<a name="oyster-entities-pb-v1-Debug"></a>

### Debug
Debug message to hold information about the request/response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| debug_id | [string](#string) |  | debug id to search in logs |
| user_id | [string](#string) |  | user id (DEPRECATED) |
| client | [string](#string) |  | client (DEPRECATED) |
| locale | [string](#string) |  | locale (DEPRECATED) |
| trace_id | [string](#string) |  | trace id to search in the system |
| app_version | [string](#string) |  | application version (DEPRECATED) |
| app_requested_api_version | [string](#string) |  | application requested api version (DEPRECATED) |
| device_id | [string](#string) |  | device id client |
| device_type | [string](#string) |  | device type (DEPRECATED) |
| ip | [string](#string) |  | ip client (DEPRECATED) |
| coordinates | [Coordinates](#oyster-entities-pb-v1-Coordinates) |  | coordinates (DEPRECATED) |






<a name="oyster-entities-pb-v1-Permission"></a>

### Permission
Permission message to hold information from the client
(DEPRECATED)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| permission_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| create_at | [int64](#int64) |  |  |
| updated_at | [int64](#int64) |  |  |





 

 

 

 



<a name="oyster_entities_pb_v1_auth-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## oyster/entities/pb/v1/auth.proto
Authorization related messages.

This file has the Authorization message as base for all the gRPC communications


<a name="oyster-entities-pb-v1-Authorization"></a>

### Authorization
Authorization base message to hold the information from the client


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  | token of type JWT |
| permission | [Permission](#oyster-entities-pb-v1-Permission) |  |  |
| ip | [string](#string) |  | ip client |
| user_agent | [string](#string) |  | user agent client |
| device_id | [string](#string) |  | device id client |
| debug | [Debug](#oyster-entities-pb-v1-Debug) |  | debug information with: trace_id app_version etc. |





 

 

 

 



<a name="oyster_entities_pb_payment_v1_payment_detail_request-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## oyster/entities/pb/payment/v1/payment_detail_request.proto



<a name="oyster-entities-pb-payment-v1-GetPaymentTransactionRequest"></a>

### GetPaymentTransactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| authorization | [oyster.entities.pb.v1.Authorization](#oyster-entities-pb-v1-Authorization) |  |  |
| debug | [oyster.entities.pb.v1.Debug](#oyster-entities-pb-v1-Debug) |  |  |
| payment_id | [string](#string) |  |  |





 

 

 

 



<a name="oyster_entities_pb_v1_errors-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## oyster/entities/pb/v1/errors.proto
Error related messages.

This file has the error messages for the communications between services


<a name="oyster-entities-pb-v1-ErrorMessage"></a>

### ErrorMessage
ErrorMessage message to hold the detailed information about the error


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [ErrorCode](#oyster-entities-pb-v1-ErrorCode) |  | code reference |
| messages | [string](#string) | repeated | messages is a repeated detailed information about the error |





 


<a name="oyster-entities-pb-v1-ErrorCode"></a>

### ErrorCode
ErrorCode message to hold a coded information about the specific error

| Name | Number | Description |
| ---- | ------ | ----------- |
| ERROR_CODE_UNSPECIFIED | 0 | Just default value for enum in proto. Should not be used as error code. |
| ERROR_CODE_NPE | 4 | When any required input is null. |
| ERROR_CODE_ILLEGAL | 5 | Mostly used with validation check. |
| ERROR_CODE_DUPLICATE | 6 | For duplicate records. |
| ERROR_CODE_FAILED | 7 | When any action fails. Like while saving record on db |
| ERROR_CODE_NOTFOUND | 8 | When records are not found. |
| ERROR_CODE_INVALID | 9 | For invalid inputs. |
| ERROR_CODE_EXCEPTION | 10 | For any runtime exception |
| ERROR_CODE_TIMEOUT | 11 | For timeouts on calling any third party services |
| ERROR_CODE_AUTH_ERROR | 12 | If authorization failed |
| ERROR_CODE_PERMISSION_ERROR | 13 | If user don&#39;t have permission |
| ERROR_CODE_FORBIDDEN | 14 | If user don&#39;t have permission |


 

 

 



<a name="oyster_entities_pb_payment_v1_payment_detail_response-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## oyster/entities/pb/payment/v1/payment_detail_response.proto



<a name="oyster-entities-pb-payment-v1-GetPaymentTransactionResponse"></a>

### GetPaymentTransactionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| debug | [oyster.entities.pb.v1.Debug](#oyster-entities-pb-v1-Debug) |  |  |
| success | [bool](#bool) |  |  |
| error | [bool](#bool) |  |  |
| error_message | [oyster.entities.pb.v1.ErrorMessage](#oyster-entities-pb-v1-ErrorMessage) |  |  |
| payment | [oyster.entities.pb.payment.domain.v1.Payment](#oyster-entities-pb-payment-domain-v1-Payment) |  |  |





 

 

 

 



<a name="oyster_entities_pb_payment_v1_process_payment_request-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## oyster/entities/pb/payment/v1/process_payment_request.proto



<a name="oyster-entities-pb-payment-v1-MakePaymentRequest"></a>

### MakePaymentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| authorization | [oyster.entities.pb.v1.Authorization](#oyster-entities-pb-v1-Authorization) |  |  |
| debug | [oyster.entities.pb.v1.Debug](#oyster-entities-pb-v1-Debug) |  |  |
| business_id | [string](#string) |  |  |
| requester_ref_id | [string](#string) |  |  |
| gross_amount | [int64](#int64) |  |  |
| description | [string](#string) |  |  |
| currency | [oyster.entities.pb.payment.domain.v1.Currency](#oyster-entities-pb-payment-domain-v1-Currency) |  |  |
| payment_method | [string](#string) |  |  |
| payeer_name | [string](#string) |  |  |
| payeer_email | [string](#string) |  |  |





 

 

 

 



<a name="oyster_entities_pb_payment_v1_process_payment_response-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## oyster/entities/pb/payment/v1/process_payment_response.proto



<a name="oyster-entities-pb-payment-v1-MakePaymentResponse"></a>

### MakePaymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| debug | [oyster.entities.pb.v1.Debug](#oyster-entities-pb-v1-Debug) |  |  |
| error | [bool](#bool) |  |  |
| success | [bool](#bool) |  |  |
| error_message | [oyster.entities.pb.v1.ErrorMessage](#oyster-entities-pb-v1-ErrorMessage) |  |  |
| timestamp | [int64](#int64) |  |  |
| business_id | [string](#string) |  |  |
| requester_ref_id | [string](#string) |  |  |
| gross_amount | [int64](#int64) |  |  |
| currency | [oyster.entities.pb.payment.domain.v1.Currency](#oyster-entities-pb-payment-domain-v1-Currency) |  |  |
| status | [string](#string) |  |  |
| payment_id | [string](#string) |  |  |





 

 

 

 



<a name="oyster_rpc_payment_v1_payment_info_rpc-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## oyster/rpc/payment/v1/payment_info_rpc.proto


 

 

 


<a name="oyster-rpc-payment-v1-PaymentInfoService"></a>

### PaymentInfoService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetPaymentTransaction | [.oyster.entities.pb.payment.v1.GetPaymentTransactionRequest](#oyster-entities-pb-payment-v1-GetPaymentTransactionRequest) | [.oyster.entities.pb.payment.v1.GetPaymentTransactionResponse](#oyster-entities-pb-payment-v1-GetPaymentTransactionResponse) |  |

 



<a name="oyster_rpc_payment_v1_process_payment_rpc-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## oyster/rpc/payment/v1/process_payment_rpc.proto


 

 

 


<a name="oyster-rpc-payment-v1-ProcessPaymentService"></a>

### ProcessPaymentService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| MakePayment | [.oyster.entities.pb.payment.v1.MakePaymentRequest](#oyster-entities-pb-payment-v1-MakePaymentRequest) | [.oyster.entities.pb.payment.v1.MakePaymentResponse](#oyster-entities-pb-payment-v1-MakePaymentResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

