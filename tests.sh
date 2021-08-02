echo "testing BTC transaction signing"

curl --location --request POST 'localhost:8080/api/v1/sign_transaction' \
  --header 'Content-Type: application/json' \
  --data-raw '{
   "gate":"bitcoin",
   "tx":{
      "utxo":[
         {
            "hash":"fff7f7881a8099afa6940d42d1e7f6362bec38171ea3edf433541db4e4ad969f",
            "index":0,
            "sequence":"4294967295",
            "amount":"625000000"
         }
      ],
      "toAddress":"1Bp9U1ogV3A14FMvKbRJms7ctyso4Z4Tcx",
      "ChangeAddress":"1FQc5LdgGHMHEN9nwkjmz6tWkxhPpxBvBU",
      "byteFee":1,
      "Amount":"1000000"
   }
}'

echo "-------------------------------"
echo "testing ETH transaction signing"

curl --location --request POST 'localhost:8080/api/v1/sign_transaction' \
  --header 'Content-Type: application/json' \
  --data-raw '{
   "gate":"ethereum",
   "tx":{
      "chainId": 3,
      "nonce": 1,
      "gasLimit": 21000,
      "gasPrice": 5000000000,
      "toAddress": "0x7788944b6dcd32f8a3042b817cfe7c5588382bd3",
      "value": 133700000000000001
   }
}'
