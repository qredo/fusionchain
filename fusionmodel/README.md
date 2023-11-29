# Fusion - MPC model service

This is a mini server that mocks the response from `fusiond` when 

## Example usage

```
$ go run .
```

```
$ curl 'localhost:9090/fusionchain/treasury/signature_request_by_id?id=312'
{"id":"3155","creator":"qredo1d652c9nngq5cneak2whyaqa4g9ehr8psyl0t7j","key_id":"0000000000000000000000000000000000000000000000000000000000000001","data_for_signing":"tSR4wa1srbASeiRWjzEKKC1PgSuPBuzuWosOEdj3NB0=","status":"SIGN_REQUEST_STATUS_PENDING"}

```

## Run with docker

TODO

