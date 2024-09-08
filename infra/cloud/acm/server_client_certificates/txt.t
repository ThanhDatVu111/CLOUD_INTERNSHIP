
aws acm delete-certificate --certificate-arn arn:aws:acm:us-west-2:431095965741:certificate/4be93f25-8229-4618-8aed-ee734c0eafc7 --output json


aws acm import-certificate --certificate fileb://client1.domain.tld.crt --private-key fileb://client1.domain.tld.key --certificate-chain fileb://ca.crt

 aws rds describe-db-instances --output json | jq "."

 aws acm list-certificates --output json | jq "."    

 