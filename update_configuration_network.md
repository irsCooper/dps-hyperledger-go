# Настройка конфигурации сети #
Нам понадобится изменить следующие файлы:
```
 1. test-network/compose/compose-bft-test-net.yaml
 2. test-network/compose/compose-test-net.yaml
 3. test-network/configtx/configtx.yaml
 4. test-network/organizations/cryptogen/crypto-config-orderer.yaml
 5. test-network/organizations/cryptogen/crypto-config-org1.yaml
 6. test-network/organizations/cryptogen/crypto-config-org2.yaml
 7. test-network/organizations/ccp-generate.sh
 8. test-network/setOrgEnv.sh
 9. test-network/scripts/deployCC.sh
10. test-network/scripts/deployCCAAS.sh
11. test-network/scripts/envVar.sh
12. test-network/organizations/ccp-generate.sh
13. test-network/organizations/ccp-template.json
14. test-network/organizations/ccp-template.yaml
```
А теперь по порядку, где и какие конкретные строчки нам нужно изменить.

## 1. test-network/compose/compose-bft-test-net.yaml
```yaml
Было:
210       - CORE_PEER_LOCALMSPID=Org1MSP
249       - CORE_PEER_LOCALMSPID=Org2MSP
```

```yaml
Стало:
210      - CORE_PEER_LOCALMSPID=Bank
249      - CORE_PEER_LOCALMSPID=Users
```

## 2. test-network/compose/compose-test-net.yaml
```yaml
Было:
83      - CORE_PEER_LOCALMSPID=Org1MSP
122     - CORE_PEER_LOCALMSPID=Org2MSP
```

```yaml
Стало:
83      - CORE_PEER_LOCALMSPID=Bank
122     - CORE_PEER_LOCALMSPID=Users
```

##  3. test-network/configtx/configtx.yaml
```yaml
Было:
41  - &Org1
    # DefaultOrg defines the organization which is used in the sampleconfig
    # of the fabric.git development environment
    Name: Org1MSP
    # ID to load the MSP definition as
    ID: Org1MSP
    MSPDir: ../organizations/peerOrganizations/org1.example.com/msp
    # Policies defines the set of policies at this level of the config tree
    # For organization policies, their canonical path is usually
    #   /Channel/<Application|Orderer>/<OrgName>/<PolicyName>
    Policies:
      Readers:
        Type: Signature
        Rule: "OR('Org1MSP.admin', 'Org1MSP.peer', 'Org1MSP.client')"
      Writers:
        Type: Signature
        Rule: "OR('Org1MSP.admin', 'Org1MSP.client')"
      Admins:
        Type: Signature
        Rule: "OR('Org1MSP.admin')"
      Endorsement:
        Type: Signature
        Rule: "OR('Org1MSP.peer')"
  - &Org2
    # DefaultOrg defines the organization which is used in the sampleconfig
    # of the fabric.git development environment
    Name: Org2MSP
    # ID to load the MSP definition as
    ID: Org2MSP
    MSPDir: ../organizations/peerOrganizations/org2.example.com/msp
    # Policies defines the set of policies at this level of the config tree
    # For organization policies, their canonical path is usually
    #   /Channel/<Application|Orderer>/<OrgName>/<PolicyName>
    Policies:
      Readers:
        Type: Signature
        Rule: "OR('Org2MSP.admin', 'Org2MSP.peer', 'Org2MSP.client')"
      Writers:
        Type: Signature
        Rule: "OR('Org2MSP.admin', 'Org2MSP.client')"
      Admins:
        Type: Signature
        Rule: "OR('Org2MSP.admin')"
      Endorsement:
        Type: Signature
86        Rule: "OR('Org2MSP.peer')"



274 Application:
        <<: *ApplicationDefaults
        Organizations:
            - *Org1
            - *Org2
279        Capabilities: *ApplicationCapabilities
```

```yaml
Стало:
41  - &Bank
    # DefaultOrg defines the organization which is used in the sampleconfig
    # of the fabric.git development environment
    Name: Bank
    # ID to load the MSP definition as
    ID: Bank
    MSPDir: ../organizations/peerOrganizations/org1.example.com/msp
    # Policies defines the set of policies at this level of the config tree
    # For organization policies, their canonical path is usually
    #   /Channel/<Application|Orderer>/<OrgName>/<PolicyName>
    Policies:
      Readers:
        Type: Signature
        Rule: "OR('Bank.admin', 'Bank.peer', 'Bank.client')"
      Writers:
        Type: Signature
        Rule: "OR('Bank.admin', 'Bank.client')"
      Admins:
        Type: Signature
        Rule: "OR('Bank.admin')"
      Endorsement:
        Type: Signature
        Rule: "OR('Bank.peer')"
  - &Users
    # DefaultOrg defines the organization which is used in the sampleconfig
    # of the fabric.git development environment
    Name: Users
    # ID to load the MSP definition as
    ID: Users
    MSPDir: ../organizations/peerOrganizations/org2.example.com/msp
    # Policies defines the set of policies at this level of the config tree
    # For organization policies, their canonical path is usually
    #   /Channel/<Application|Orderer>/<OrgName>/<PolicyName>
    Policies:
      Readers:
        Type: Signature
        Rule: "OR('Users.admin', 'Users.peer', 'Users.client')"
      Writers:
        Type: Signature
        Rule: "OR('Users.admin', 'Users.client')"
      Admins:
        Type: Signature
        Rule: "OR('Users.admin')"
      Endorsement:
        Type: Signature
86        Rule: "OR('Users.peer')"



274 Application:
        <<: *ApplicationDefaults
        Organizations:
            - *Bank
            - *Users
279        Capabilities: *ApplicationCapabilities
```

## 4. test-network/organizations/cryptogen/crypto-config-orderer.yaml

```yaml
Было:
21     Specs:
        - Hostname: orderer
          SANS:
            - localhost
        - Hostname: orderer2
          SANS:
            - localhost
        - Hostname: orderer3
          SANS:
            - localhost
        - Hostname: orderer4
          SANS:
            - localhost

```

```yaml
Стало:
21     Specs:
        - Hostname: orderer
          SANS:
            - localhost
        # - Hostname: orderer2
        #   SANS:
        #     - localhost
        # - Hostname: orderer3
        #   SANS:
        #     - localhost
        # - Hostname: orderer4
        #   SANS:
        #     - localhost

```

## 5. test-network/organizations/cryptogen/crypto-config-org1.yaml

```yaml
Было:
14   - Name: Org1MSP
```

```yaml
Стало:
14   - Name: Bank
```

##  6. test-network/organizations/cryptogen/crypto-config-org2.yaml
 
```yaml
Было:
13   - Name: Org2MSP
```

```yaml
Стало:
13   - Name: Users
```

## 7. test-network/organizations/ccp-generate.sh

```bash
Было:
29  ORG=1
    P0PORT=7051
    CAPORT=7054
    PEERPEM=organizations/peerOrganizations/org1.example.com/tlsca/tlsca.org1.example.com-cert.pem
    CAPEM=organizations/peerOrganizations/org1.example.com/ca/ca.org1.example.com-cert.pem

    echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/org1.example.com/connection-org1.json
    echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/org1.example.com/connection-org1.yaml

    ORG=2
    P0PORT=9051
    CAPORT=8054
    PEERPEM=organizations/peerOrganizations/org2.example.com/tlsca/tlsca.org2.example.com-cert.pem
    CAPEM=organizations/peerOrganizations/org2.example.com/ca/ca.org2.example.com-cert.pem

    echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/org2.example.com/connection-org2.json
    echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/org2.example.com/connection-org2.yaml

```

```bash
Стало:
29  ORGNAME=Bank
    ORG=1
    P0PORT=7051
    CAPORT=7054
    PEERPEM=organizations/peerOrganizations/org1.example.com/tlsca/tlsca.org1.example.com-cert.pem
    CAPEM=organizations/peerOrganizations/org1.example.com/ca/ca.org1.example.com-cert.pem

    echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/org1.example.com/connection-org1.json
    echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM $ORGNAME)" > organizations/peerOrganizations/org1.example.com/connection-org1.yaml

    ORGNAME=Users
    ORG=2
    P0PORT=9051
    CAPORT=8054
    PEERPEM=organizations/peerOrganizations/org2.example.com/tlsca/tlsca.org2.example.com-cert.pem
    CAPEM=organizations/peerOrganizations/org2.example.com/ca/ca.org2.example.com-cert.pem

    echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/org2.example.com/connection-org2.json
    echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM $ORGNAME)" > organizations/peerOrganizations/org2.example.com/connection-org2.yaml
```

##  8. test-network/setOrgEnv.sh

```bash
Было:
26    CORE_PEER_LOCALMSPID=Org1MSP
33    CORE_PEER_LOCALMSPID=Org2MSP
```

```bash
Стало:
26    CORE_PEER_LOCALMSPID=Bank
33    CORE_PEER_LOCALMSPID=Users
```

##  9. test-network/scripts/deployCC.sh

```bash
Было:
92  checkCommitReadiness 1 "\"Org1MSP\": true" "\"Org2MSP\": false"
    checkCommitReadiness 2 "\"Org1MSP\": true" "\"Org2MSP\": false"

    ## now approve also for org2
    approveForMyOrg 2

    ## check whether the chaincode definition is ready to be committed
    ## expect them both to have approved
    checkCommitReadiness 1 "\"Org1MSP\": true" "\"Org2MSP\": true"
101 checkCommitReadiness 2 "\"Org1MSP\": true" "\"Org2MSP\": true"
```

```bash
Стало:
92  checkCommitReadiness 1 "\"Bank\": true" "\"Users\": false"
    checkCommitReadiness 2 "\"Bank\": true" "\"Users\": false"

    ## now approve also for org2
    approveForMyOrg 2

    ## check whether the chaincode definition is ready to be committed
    ## expect them both to have approved
    checkCommitReadiness 1 "\"Bank\": true" "\"Users\": true"
101 checkCommitReadiness 2 "\"Bank\": true" "\"Users\": true"
```

## 10. test-network/scripts/deployCCAAS.sh

```bash
Было:
188     checkCommitReadiness 1 "\"Org1MSP\": true" "\"Org2MSP\": false"
        checkCommitReadiness 2 "\"Org1MSP\": true" "\"Org2MSP\": false"

        ## now approve also for org2
        approveForMyOrg 2

        ## check whether the chaincode definition is ready to be committed
        ## expect them both to have approved
        checkCommitReadiness 1 "\"Org1MSP\": true" "\"Org2MSP\": true"
        checkCommitReadiness 2 "\"Org1MSP\": true" "\"Org2MSP\": true"

```

```bash
Стало:
188     checkCommitReadiness 1 "\"Bank\": true" "\"Users\": false"
        checkCommitReadiness 2 "\"Bank\": true" "\"Users\": false"

        ## now approve also for org2
        approveForMyOrg 2

        ## check whether the chaincode definition is ready to be committed
        ## expect them both to have approved
        checkCommitReadiness 1 "\"Bank\": true" "\"Users\": true"
        checkCommitReadiness 2 "\"Bank\": true" "\"Users\": true"
```

## 11. test-network/scripts/envVar.sh

```bash
Было:
29     export CORE_PEER_LOCALMSPID="Org1MSP"
34     export CORE_PEER_LOCALMSPID="Org2MSP"
```

```bash
Стало:
29     export CORE_PEER_LOCALMSPID="Bank"
34     export CORE_PEER_LOCALMSPID="Users"
```

## 12. test-network/organizations/ccp-generate.sh

```bash
Было:
29     export CORE_PEER_LOCALMSPID="Org1MSP"
34     export CORE_PEER_LOCALMSPID="Org2MSP"
```


знаком  ################### отмечены изменения
```bash
Стало:
function json_ccp {
    local PP=$(one_line_pem $4)
    local CP=$(one_line_pem $5)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${CAPORT}/$3/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        -e "s/\${ORGNAME}/$6/" \  ###################
        organizations/ccp-template.json
}

function yaml_ccp {
    local PP=$(one_line_pem $4)
    local CP=$(one_line_pem $5)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${CAPORT}/$3/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        -e "s/\${ORGNAME}/$6/" \  ###################
        organizations/ccp-template.yaml | sed -e $'s/\\\\n/\\\n          /g'
}

ORGNAME=Bank  ###################
ORG=1
P0PORT=7051
CAPORT=7054
PEERPEM=organizations/peerOrganizations/org1.example.com/tlsca/tlsca.org1.example.com-cert.pem
CAPEM=organizations/peerOrganizations/org1.example.com/ca/ca.org1.example.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM $ORGNAME)" > organizations/peerOrganizations/org1.example.com/connection-org1.json  ###################
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM $ORGNAME)" > organizations/peerOrganizations/org1.example.com/connection-org1.yaml  ###################

ORGNAME=Users  ###################
ORG=2
P0PORT=9051
CAPORT=8054
PEERPEM=organizations/peerOrganizations/org2.example.com/tlsca/tlsca.org2.example.com-cert.pem
CAPEM=organizations/peerOrganizations/org2.example.com/ca/ca.org2.example.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM $ORGNAME)" > organizations/peerOrganizations/org2.example.com/connection-org2.json  ###################
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM $ORGNAME)" > organizations/peerOrganizations/org2.example.com/connection-org2.yaml  ###################

```



## 13. test-network/organizations/ccp-template.json

```bash
Было:
29     export CORE_PEER_LOCALMSPID="Org1MSP"
34     export CORE_PEER_LOCALMSPID="Org2MSP"
```

```json
"name": "test-network-org${ORG}",
    "version": "1.0.0",
    "client": {
        "organization": "${ORGNAME}",  ###################
        "connection": {
            "timeout": {
                "peer": {
                    "endorser": "300"
                }
            }
        }
    },
    "organizations": {
        "${ORGNAME}": {  ###################
            "mspid": "${ORGNAME}",  ###################
```



## 14. test-network/organizations/ccp-template.yaml

```bash
Было:
29     export CORE_PEER_LOCALMSPID="Org1MSP"
34     export CORE_PEER_LOCALMSPID="Org2MSP"
```

```yaml
Стало:
name: test-network-org${ORG}
version: 1.0.0
client:
  organization: ${ORGNAME}  ###################
  connection:
    timeout:
      peer:
        endorser: '300'
organizations:
  ${ORGNAME}:  ###################
    mspid: ${ORGNAME}  ###################
```



Отлично, настройка завершена! Вы __умничка__! Можем вернуться назад и продолжить работу над проектом [тык](./CREATE_PROJECT.md).