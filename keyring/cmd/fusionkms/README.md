# Fusion KMS

A lightweight keyring client for Fusion. This application runs as a server-side key management service for ECDSA and EDDSA key pairs.


## Run

```
go run .
```

## Configuration

```
cd ../../../blockchain
./init.sh
```

Then

```
cd ~$HOME/go/src/github.com/qredo/fusionchain/keyring/cmd/fusionkms
go run .
```

to start the MPC relayer service.

## APIs

### 1) /status (GET)

The `/status` call requests information about the liveness of the fusionkms and will always repond "OK" if the service is up. 

```go
StatusCode: 200
JSON:
{
    "service":"fusionkms",
    "version":"0.1.0",
    "message":"OK"
}
```

### 2) /healthcheck (GET)

The `/healthcheck` call requests information about the current health of the fusionkms and its connections. On receiving this request the fusionkms pings its local `fusiond` client.

```go
StatusCode: 200
JSON:
{
    "service":"fusionkms",
    "version":"0.1.0",
    "message":"OK",
    "failures": []
}
```

If one or more of the checks fail then the response will contain an array of failure messages

```go
StatusCode: 503
JSON: 
{
    "service":"fusionkms",
    "version":"0.1.0",
    "message" "",
    "failures": ["'key':<failure error message>"]
} 
```

### 3) /pubkeys (GET)

The `/pubkeys` call requests a list of workspace keys that have been saved to the  application's local database. Note that this call is password protected

```go
StatusCode: 200
JSON:
{
    "service":"fusionkms",
    "version":"0.1.0",
    "message":"OK",
    "pubkeys": []
}
```

Example

```
$ curl -H "password: my_password" localhost:8080/pubkeys
> TODO
```

### 4) /mnemonic (GET)

The `/mnemonic` call requests a list of workspace keys that have been saved to the  application's local database.

```go
StatusCode: 200
JSON:
{
    "service":"fusionkms",
    "version":"0.1.0",
    "message":"OK",
    "mnemonic": <mnemonic_seed_phrase>
    "password_protected": true
}
```

Example

```
$ curl -H "password: my_password" localhost:8080/mnemonic
> TODO
```

### 5) /keyring (GET)

The `/keyring` call requests a list of workspace keys that have been saved to the  application's local database.

```go
StatusCode: 200
JSON:
{
    "service":"fusionkms",
    "version":"0.1.0",
    "message":"OK",
    "mnemonic": <mnemonic_seed_phrase>
    "password_protected": true
}
```

Example

```
$ curl -H "password: <my_password>" localhost:8080/keyring
> TODO
```
