### Description
This project is a simple TCP server written in C, it accepts a connection, expect a number as an input, and output the factorial of this number as a result.

### Usage
You can open a TCP connection with this server, the expected input is a number, and the response is the factorial of the input number

Example: `print "4" | nc localhost 9999` will ouput `24`

### How to build
`gcc server.c -o server`

### Usage
It accepts a first argument which is the TCP port where to start listening
`./server PORT`

### Dependencies
There is no dependencies of this server, it uses the GNU standard C libraries, but other components depends on this server

### TODO
* Optimize the way we handle the incoming message
* Try to detect any anomaly with the memory management in the code (OUTPUT expected)
* Create an Upstart script for this service, and register it in a service discovery once UP