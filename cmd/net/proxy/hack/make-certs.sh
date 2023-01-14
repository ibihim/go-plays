#!/usr/bin/env bash

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes \
    -keyout ca.key -out ca.cert \
    -subj "/C=DE/ST=Brandenburg/L=Falkensee/CN=*.signer.ostrowski.example.com/emailAddress=signer.ostrowski@example.com"

echo "CA's self-signed certificate"
openssl x509 -in ca.cert -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes \
    -keyout server.key -out server.csr \
    -subj "/C=DE/ST=Brandenburg/L=Falkensee/CN=*.cert.ostrowski.example.com/emailAddress=cert.ostrowski@example.com"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req \
    -in server.csr -days 60 \
    -CA ca.cert -CAkey ca.key -CAcreateserial -out server.cert -extfile server-ext.cnf

echo "Server's signed certificate"
openssl x509 -in server.cert -noout -text

echo "Verify Certs"
openssl verify -CAfile ca.cert server.cert

