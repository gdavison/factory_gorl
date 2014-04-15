factory_gorl
============

A Go port of Ruby’s [factory_girl](https://github.com/thoughtbot/factory_girl).

Documentation
-------------
Coming soon. In the meantime, take a look at the tests.

Install
-------
```shell
go get github.com/gdavison/factory_gorl
```

Status
------
So far, it can be used to initialize (`Build`) an in-memory object and to persist the
object (`Create`) using [gorp](https://github.com/coopernurse/gorp). Associations have not yet been implemented.

Factory inheritance is implemented.

Copyright © 2014 Graham Davison.
