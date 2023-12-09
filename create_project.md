# Инструкция по созданию проекта с использованием hyperledger fabric test-network #

Создадим папку в которой будет находиться наш проект. Откроем терминал и выполним следующие команды:
```
mkdir <имя папки>
cd <имя папки>
```

Отлично, перенесём 3 необходимых каталога из **fabric-samples**:
1.  /bin          - тут лежат бинарники, которые использует сеть
2.  /config       - тут лежат файлы конфигурации, для настройки сеть
3.  /test-network - сама сеть


При необходимости, настроим конфигурацию сети, заменив Org1MSP и Org2MSP на Bank и Users (в моём случае). Для этого [переходим сюда](./update_uonfiguration_network.md)

Для запуска сети и создания канала находясь в корневом каталоге выполним команды:
```bash
cd test-network
./network.sh up createChannel -ca
```

О том, что вы всё сделали правильно сведетельствуют следующие логи (не пугайтесь, их будет больше чем представлено тут, просто инструкция получилась бы слишком длинной):
```bash
Generating certificates using Fabric CA
Creating network "fabric_test" with the default driver
Creating ca_org2    ... done
Creating ca_org1    ... done
Creating ca_orderer ... done
```

```bash
Generating CCP files for Org1 and Org2
Creating volume "compose_orderer.example.com" with default driver
Creating volume "compose_peer0.org1.example.com" with default driver
Creating volume "compose_peer0.org2.example.com" with default driver
Creating peer0.org1.example.com ... done
Creating peer0.org2.example.com ... done
Creating orderer.example.com    ... done
Creating cli                    ... done
```

```bash
Adding orderers
+ . scripts/orderer.sh mychannel
+ '[' 0 -eq 1 ']'
+ res=0
Status: 201
{
	"name": "mychannel",
	"url": "/participation/v1/channels/mychannel",
	"consensusRelation": "consenter",
	"status": "active",
	"height": 1
}
Channel 'mychannel' created
```

```bash
Generating anchor peer update transaction for Org1 on channel mychannel
+ configtxlator proto_encode --input Bankconfig.json --type common.Config --output original_config.pb
+ configtxlator proto_encode --input Bankmodified_config.json --type common.Config --output modified_config.pb
+ configtxlator compute_update --channel_id mychannel --original original_config.pb --updated modified_config.pb --output config_update.pb
+ configtxlator proto_decode --input config_update.pb --type common.ConfigUpdate --output config_update.json
+ jq .
++ cat config_update.json
+ echo '{"payload":{"header":{"channel_header":{"channel_id":"mychannel", "type":2}},"data":{"config_update":{' '"channel_id":' '"mychannel",' '"isolated_data":' '{},' '"read_set":' '{' '"groups":' '{' '"Application":' '{' '"groups":' '{' '"Bank":' '{' '"groups":' '{},' '"mod_policy":' '"",' '"policies":' '{' '"Admins":' '{' '"mod_policy":' '"",' '"policy":' null, '"version":' '"0"' '},' '"Endorsement":' '{' '"mod_policy":' '"",' '"policy":' null, '"version":' '"0"' '},' '"Readers":' '{' '"mod_policy":' '"",' '"policy":' null, '"version":' '"0"' '},' '"Writers":' '{' '"mod_policy":' '"",' '"policy":' null, '"version":' '"0"' '}' '},' '"values":' '{' '"MSP":' '{' '"mod_policy":' '"",' '"value":' null, '"version":' '"0"' '}' '},' '"version":' '"0"' '}' '},' '"mod_policy":' '"",' '"policies":' '{},' '"values":' '{},' '"version":' '"0"' '}' '},' '"mod_policy":' '"",' '"policies":' '{},' '"values":' '{},' '"version":' '"0"' '},' '"write_set":' '{' '"groups":' '{' '"Application":' '{' '"groups":' '{' '"Bank":' '{' '"groups":' '{},' '"mod_policy":' '"Admins",' '"policies":' '{' '"Admins":' '{' '"mod_policy":' '"",' '"policy":' null, '"version":' '"0"' '},' '"Endorsement":' '{' '"mod_policy":' '"",' '"policy":' null, '"version":' '"0"' '},' '"Readers":' '{' '"mod_policy":' '"",' '"policy":' null, '"version":' '"0"' '},' '"Writers":' '{' '"mod_policy":' '"",' '"policy":' null, '"version":' '"0"' '}' '},' '"values":' '{' '"AnchorPeers":' '{' '"mod_policy":' '"Admins",' '"value":' '{' '"anchor_peers":' '[' '{' '"host":' '"peer0.org1.example.com",' '"port":' 7051 '}' ']' '},' '"version":' '"0"' '},' '"MSP":' '{' '"mod_policy":' '"",' '"value":' null, '"version":' '"0"' '}' '},' '"version":' '"1"' '}' '},' '"mod_policy":' '"",' '"policies":' '{},' '"values":' '{},' '"version":' '"0"' '}' '},' '"mod_policy":' '"",' '"policies":' '{},' '"values":' '{},' '"version":' '"0"' '}' '}}}}'
+ configtxlator proto_encode --input config_update_in_envelope.json --type common.Envelope --output Bankanchors.tx
2023-11-07 10:21:25.427 UTC 0001 INFO [channelCmd] InitCmdFactory -> Endorser and orderer connections initialized
2023-11-07 10:21:25.435 UTC 0002 INFO [channelCmd] update -> Successfully submitted channel update
Anchor peer set for org 'Bank' on channel 'mychannel'
```

```bash
Generating anchor peer update transaction for Org2 on channel mychannel
+ configtxlator proto_encode --input Usersconfig.json --type common.Config --output original_config.pb
+ configtxlator proto_encode --input Usersmodified_config.json --type common.Config --output modified_config.pb
+ configtxlator compute_update --channel_id mychannel --original original_config.pb --updated modified_config.pb --output config_update.pb
+ configtxlator proto_decode --input config_update.pb --type common.ConfigUpdate --output config_update.json
+ jq .
++ cat config_update.json
+ echo '{"payload":{"header":{"channel_header":{"channel_id":"mychannel", "type":2}},"data":{"config_update":{' '"channel_id":' '"mychannel",' '"isolated_data":' '{},' '"read_set":' '{' '"groups":' '{' '"Application":' '{' '"groups":' '{' '"Users":' '{' '"groups":' '{},' '"mod_policy":' '"",' '"policies":' '{' '"Admins":' '{' '"mod_policy":' '"",' '"policy":' null, '"version":' '"0"' '},' '"Endorsement":' '{' '"mod_policy":' '"",' '"policy":' null, '"version":' '"0"' '},' '"Readers":' '{' '"mod_policy":' '"",' '"policy":' null, '"version":' '"0"' '},' '"Writers":' '{' '"mod_policy":' '"",' '"policy":' null, '"version":' '"0"' '}' '},' '"values":' '{' '"MSP":' '{' '"mod_policy":' '"",' '"value":' null, '"version":' '"0"' '}' '},' '"version":' '"0"' '}' '},' '"mod_policy":' '"",' '"policies":' '{},' '"values":' '{},' '"version":' '"0"' '}' '},' '"mod_policy":' '"",' '"policies":' '{},' '"values":' '{},' '"version":' '"0"' '},' '"write_set":' '{' '"groups":' '{' '"Application":' '{' '"groups":' '{' '"Users":' '{' '"groups":' '{},' '"mod_policy":' '"Admins",' '"policies":' '{' '"Admins":' '{' '"mod_policy":' '"",' '"policy":' null, '"version":' '"0"' '},' '"Endorsement":' '{' '"mod_policy":' '"",' '"policy":' null, '"version":' '"0"' '},' '"Readers":' '{' '"mod_policy":' '"",' '"policy":' null, '"version":' '"0"' '},' '"Writers":' '{' '"mod_policy":' '"",' '"policy":' null, '"version":' '"0"' '}' '},' '"values":' '{' '"AnchorPeers":' '{' '"mod_policy":' '"Admins",' '"value":' '{' '"anchor_peers":' '[' '{' '"host":' '"peer0.org2.example.com",' '"port":' 9051 '}' ']' '},' '"version":' '"0"' '},' '"MSP":' '{' '"mod_policy":' '"",' '"value":' null, '"version":' '"0"' '}' '},' '"version":' '"1"' '}' '},' '"mod_policy":' '"",' '"policies":' '{},' '"values":' '{},' '"version":' '"0"' '}' '},' '"mod_policy":' '"",' '"policies":' '{},' '"values":' '{},' '"version":' '"0"' '}' '}}}}'
+ configtxlator proto_encode --input config_update_in_envelope.json --type common.Envelope --output Usersanchors.tx
2023-11-07 10:21:25.807 UTC 0001 INFO [channelCmd] InitCmdFactory -> Endorser and orderer connections initialized
2023-11-07 10:21:25.817 UTC 0002 INFO [channelCmd] update -> Successfully submitted channel update
Anchor peer set for org 'Users' on channel 'mychannel'
Channel 'mychannel' joined
```


Если у вас не возникло ошибок, поздравляю, вы __умничка__! Если ошибки всё же возникли, проверьте, правильно ли вы изменили конфигурацию, и установлен ли у вас docker-compose. Попробуйте загуглить ошибки и найти решение самостоятельно. Да-да, возможно даже на китайских сайтах. __У вас обязательно получится!!__

Для остановки и очистки сети, при необходимости, используем команду:
```bash
./network.sh down
```
По правильному, очистку сети нужно делать перед каждым запуском нашего тестового решения.

---

Отлично, сеть у нас запущена, теперь можем приступить к написанию нашего контракта. Для этого в корневом каталоге нашего проекта откроем терминал и создадим папку, где в будущем будет лежать наш контракт:
```bash
mkdir chaincode-go
cd chaincode-go
```
В появившейся папке создадим main.go, который в будущем будет запускать наш контракт. А так же сгенерируем go.mod. Для этого в терминале напишем:

```bash
go mod init < желательно ссылка на ваш git.hub репозиторий с проектом, но можно и локальное имя, например app >
go mod tidy -go=1.20
```

Зачем бы добавили флаг -go=1.20? Вы можете использовать любую версию go, главное, чтобы она не выглядела так: 1.20.4, иначе вы не сможете задеплоить ваш чейнкод в сети, она будет ругаться

Далее, создадим папку, в которой будет наш контракт:
```bash
mkdir chaincode
```
А в только что созданной папке создаём chaincode.go, где мы и будем писать основную логику проекта. Расписывать подробно этот момент я не буду, при желании можете разобраться в моём коде.



Запуск сети и создание канала
```bash
 ./network.sh up createChannel -ca
```



Деплой нашего контракта в сети:
```bash
./network.sh  deployCC -ccn basic -ccp ../chaincode-go/ -ccv 1 -ccl go -c  mychannel
```
