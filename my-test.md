Мои тесты:
```bash
./network.sh cc invoke -c mychannel -ccic '{"Args":["Init"]}'







./network.sh cc query -c mychannel -ccqc '{"Args":["Login","bank"]}'


./network.sh cc query -c mychannel -ccqc '{"Args":["Login","Иванов_Иван_Иванович"]}'


./network.sh cc invoke -c mychannel -ccic '{"Args":["SetCertificate","Иванов_Иван_Иванович","333","5","A"]}'


./network.sh cc invoke -c mychannel -ccic '{"Args":["SetTransport","Иванов_Иван_Иванович","A","500","3"]}'








./network.sh cc query -c mychannel -ccqc '{"Args":["Login","Семенов_Семен_Семенович"]}'


./network.sh cc invoke -c mychannel -ccic '{"Args":["SetCertificate","Семенов_Семен_Семенович","222","8","B"]}'


./network.sh cc invoke -c mychannel -ccic '{"Args":["SetTransport","Семенов_Семен_Семенович","B","200","5"]}'







./network.sh cc invoke -c mychannel -ccic '{"Args":["SetFine","Иванов_Иван_Иванович","222"]}'


./network.sh cc query -c mychannel -ccqc '{"Args":["GetOneFine",""]}'


./network.sh cc query -c mychannel -ccqc '{"Args":["GetAllMyFines","222"]}'


./network.sh cc invoke -c mychannel -ccic '{"Args":["PayFine","Семенов_Семен_Семенович",""]}'

```