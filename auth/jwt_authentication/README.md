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
Since we are using ECDSA as the signing method, we need to provide both a private key and a public key. The private key is used to sign the token, while the public key is used to verify it. If you use other signing methods, such as HS256, you only need to provide a shared secret string.

Typically, a JWT is used as a means of authentication. For example, when a client logs into an application, it receives a JWT. The client can then include this token in subsequent requests to prove its identity, allowing the server to recognize and authenticate the client. Additionally, because the JWT payload can include custom data, it can also be used to transmit extra information to the server—such as the client’s role, whether the token is associated with a user, and account-related properties.

## How to use this code
### Requirements
- You need Go version >=1.24.4
- Install make

### Usage
Run the code with the following command:
```bash
make start
```

### Output
After running the code, you should see the signed token output along with the parsed and logged claims. Feel free to experiment with the code—try adding custom attributes to your claims and explore other signing methods, such as HS256.
