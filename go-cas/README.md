# Introduction

A description of the CAS protocol is given here

  https://apereo.github.io/cas/4.2.x/protocol/CAS-Protocol.html

Basically there are 2 layers in this example.
* One layer handling CAS authentication. 

  CAS authentication is handled in this layer, via a Login/Logout mechanism.
  Once the connection is authenticated, the request is "forwarded" to next layer.

* one layer handling GIN, as usual.

  A standard GIN server.

# Building & Execute
```
prompt> go build main.go
prompt> ./main -cas <your_cas_server_url> -iface <ethernet_interface_name>
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (4 handlers)
[GIN-debug] GET    /logout                   --> main.main.func2 (4 handlers)
...
```

