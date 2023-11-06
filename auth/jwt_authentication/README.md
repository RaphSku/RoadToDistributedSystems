# Understanding Authentication with JSON Web Token (JWT)

## General
JSON Web Token (JWT) was established with the RFC 7519 in May 2015. What is a JSON Web Token?

A JWT constitutes of 3 parts in the form of: xxx.yyy.zzz where the first part is the header, the second part is representing the payload and the third part is the signature. Every part is encoded in base64url and represents a JSON object. This may sound confusing, so let's do an example.

Imagine, we have the following signed token: `eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTYyNzkwMTEsIm5hbWUiOiJSYXBoYWVsIn0.NfxMLXE772UjxEFIPyHsU-Ov1ICOvjrYkZJyA-zPU0imVyudnkvcXu066A8vEEhIVd4f-_k-LGFrHDdwOYgcyw`

Paste this into the debugger of https://jwt.io/. The debugger will decode the token for you and the header contains the following information:
```json
{
    "alg": "ES256",
    "typ": "JWT"
}
```
`ES256` was the algorithm used to sign the token and the type is JWT. 

The payload is the data that we wanted to transport with this token:
```json
{
    "exp": 1696279011,
    "name": "Raphael"
}
```
namely `exp` which is the expiration date of the token and `name` which is the username associated with that token.

The last part is the signature which has the following form:
```json
ECDSASHA256(
    base64UrlEncode(header) + "." + base64UrlEncode(payload), public key in SPKI or JWK string format, private key in PKCS or JWK string format
)
```
Since we are using ECDSA as a signing method, we have to provide a private key and public key. The private key is used for signing and the public key for verification. If you use other signing methods like HS256, you will only need to provide a secret string. 

Typically, a JWT serves as a means of authentication. For instance, when a client logs into an application, it obtains a JWT. Subsequently, the client can include this token in its requests to confirm its identity. This enables the server to identify the client. Furthermore, due to the ability to include custom data in the payload, we can use JWTs to convey additional information to the server. For instance, this information may include details about the client's role, whether it is linked to a user, and the properties of the account.

## How to use this code
### Requirements
- You need Go version >=1.20.x
- Install make

### Usage
First of all, to fetch the module dependencies run 
```bash
make install_dependencies
```

Then to run the code, just run
```bash
make start
```

### Output
After running the code, you should see an output where your signed token is display and also the claims are parsed and logged. I invite you to play around with the code, create additional attributes in your claims and also try other signing methods, e.g. HS256.