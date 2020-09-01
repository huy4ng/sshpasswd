# GoPentestTools
Useful golang tools in pentest work especially when internal pentest working

 

```
#######################################
#    SSHPASSWD    |   Author:huy4ng   #
#######################################
Useage Example:

User username and password to connect SSH and execute command

	 ./sshpasswd -ip=127.0.0.1 -port=22 -u=root -p=root -c=id

User username and password to connect SSH and execute command in an interactive method

	 ./sshpasswd -ip=127.0.0.1 -port=22 -u=root -p=root

User username and private key to connect SSH and execute command

	 ./sshpasswd -ip=127.0.0.1 -port=22 -u=root -k=key.pem -c=id

User username and private key to connect SSH and execute command in an interactive method

	 ./sshpasswd -ip=127.0.0.1 -port=22 -u=root -k=key.pem

User username and encryptedprivate key to connect SSH and execute command

	 ./sshpasswd -ip=127.0.0.1 -port=22 -u=root -k=key.pem -p=123456 -c=id

User username and encrypted private key to connect SSH and execute command in an interactive method

	 ./sshpasswd -ip=127.0.0.1 -port=22 -u=root -k=key.pem -p=123456
```

