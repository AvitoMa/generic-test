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

* Automate the whole stack build
* Register services in a service discovery and use it
* Do Packaging of each service after build
* Provisioning of required resources for each service
* Deployment of the whole stack to the provisioned resources
* Put in place a Pre and Post deployment process where we can be able to run tests
* For future development, the stack will need a Postgres service with a Master / Slave nodes

#### Preferred OS / Solutions
* Centos 7
* Nginx
* etcd
* RPM packaging

Many folks completed this in a short time, you can do it as well :)
Good luck!