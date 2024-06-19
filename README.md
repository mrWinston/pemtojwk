# pemtojwk

Convert public key files in [PEM format](https://en.wikipedia.org/wiki/Privacy-Enhanced_Mail) to [JSON Web Key Sets](https://auth0.com/docs/secure/tokens/json-web-tokens/json-web-key-sets).

This tool uses a function copied from the [kubernetes source code](https://github.com/kubernetes/kubernetes/blob/master/pkg/serviceaccount/jwt.go#L99) to generate the key ids of the jwks exactly as kubernetes would.

## Requirements

- Tested with go 1.22.1, might work with versions lower than that.

## Installation

Build the tool with `go`:
```sh
go build
```

Or, install it to `$GOPATH/bin`:

```sh
go install
```
    
## Usage/Examples

`pemtojwk` accepts only the path to a pem-encoded public key file and prints the jwks to stdout:

```sh
pemtojwk [path/to/pem]
```

For example:
```sh
$ cat <<EOF > pub.pem
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5u+cw+lLRnL85bTZ10Vg
xLKGWk6WtzW2qdtLkdRP2avPxcFWrZdw85YniJWiNdGWoM5ZAXbQNHXySS66pxXl
UDAIBTLNXuiJyslaR8D8xdYHuoPB6mN7lIjQaGFftJo4qK61iEZvjfmR7nnf+dhN
YHhIZqTjcKPvbAMTIW45QK+v06/OubcySGdmHJL/0MkWbYreTkbDqchK0p9UGcOa
C5GUhJ03vpq1bXluQAwEjaKUzXPat7zyGCjHzl3tK8QfS3iFOjb09Kq7ZFDtvSBk
pkQkFqit6wp8pLNbAVSlp3h9KiJv8SB8iKfXHKLCTPThvVSiJKwIxKuaHw9C5+qO
MQIDAQAB
-----END PUBLIC KEY-----
EOF

$ pemtojwk ./pub.pem            
{
    "keys": [
        {
            "use": "sig",
            "kty": "RSA",
            "kid": "pDweZuLqGgPOMoXCqKTCMp2yDmiOF-nsARLfbJZyHN0",
            "alg": "RS256",
            "n": "5u-cw-lLRnL85bTZ10VgxLKGWk6WtzW2qdtLkdRP2avPxcFWrZdw85YniJWiNdGWoM5ZAXbQNHXySS66pxXlUDAIBTLNXuiJyslaR8D8xdYHuoPB6mN7lIjQaGFftJo4qK61iEZvjfmR7nnf-dhNYHhIZqTjcKPvbAMTIW45QK-v06_OubcySGdmHJL_0MkWbYreTkbDqchK0p9UGcOaC5GUhJ03vpq1bXluQAwEjaKUzXPat7zyGCjHzl3tK8QfS3iFOjb09Kq7ZFDtvSBkpkQkFqit6wp8pLNbAVSlp3h9KiJv8SB8iKfXHKLCTPThvVSiJKwIxKuaHw9C5-qOMQ",
            "e": "AQAB"
        }
    ]
}
```

### Reading public key from stdin

`pemtojwk` is only able to read from a given path. So make it read a pem file from stdin, use `/dev/stdin` as input file instead:

```sh
cat ./pub.pem | pemtojwk /dev/stdin
```

## Contributing

Contributions are always welcome!

## License

[Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)

