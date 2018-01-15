### Avito Devops Test

#### Description
This repository contains 3 services :
* Factorial TCP Server written in C
* HTTP API consuming Factorial written in Golang
* Front-end SPA app written in Javascript

Inside each folder, you will find the README.md of each service for :
* How to build / compile
* Required dependencies
* How to communicate with other services


#### TODO
As a Devops Engineer, you will need to resolve issues found in each service, more details can be found in the sub-directories README.md, and You have to :

* Automate the whole stack build and packaging
* Use whatever service discovery tool you want in order to register your services
* Automate the provisioning of your stack while making your solution scalable 
* Automate the deployment of your services
* Put in place a Pre and Post deployment process where we can be able to run tests 
* For Database usage create a postgres Master / Slave cluster
* Secure your solution by using SSL 
* Block the access to your solution to the follwing IPs in 3 different layers of network 
  - 81.192.111.243
  - 81.192.111.244
  - 81.192.111.245
  - 81.192.111.246
  - 105.159.249.73
  - 105.159.249.85
  - 105.159.249.86
* Centralize your logs in one place
* Resize both the Root disk and the postgres location of your master
* Allow access to your Database for the following user devops-test only on slave in order to use pgAdmin
  - 81.192.111.243
  - 81.192.111.244
  - 81.192.111.245
  - 81.192.111.246
  - 105.159.249.73
  - 105.159.249.85
  - 105.159.249.86
* Use the configuration management tool in order to install pgtop, and pgbadger
* Generate the postgres report using pgbadger
* Use whatever tool you want to backup your Database
* Execute a tcpdump command in different servers and explains the output
* Document everything that you did step by step

#### Snapshot of the expected Infra
![Stack Infra](https://raw.githubusercontent.com/AvitoMa/devops-test/master/lifecycle.png)

#### Preferred OS / Solutions
* Centos 7
* Nginx
* etcd
* RPM packaging


Good luck!
