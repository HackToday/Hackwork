Simple Test of Authrization Of KeyStone Intergrated with Docker
==============================================================

As docker provide authrization plugin ways, so it is possible to use
keystone to verify related access role under certain docker operations.
It is useful for multi-tenant environments to isolation resouces and
operation.

> **Note**: This is simple test and play ways, not strict for any user
> environment test and develop.

## How to Run it?

Start docker with following way:

	docker daemon -D --authorization-plugin=test_authz_plugin  -H tcp://127.0.0.1:4243

Start Plugin this way:
	
	./authz_plugin <absolute path dockerule.json>

## TODO List

Track docker commuity for ACL full support and related multi-tenants support

## Bugs

Not know now :)


