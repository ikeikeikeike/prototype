# injector

Experimental project

### Installation

**Install Packages**

```elixir
$ cd $GOPATH/src/github.com/ikeikeikeike/prototype/injector
$ dep ensure
```

### Database migration

**Install migration tool**

```elixir
$ go get -u github.com/mattes/migrate
$ go get -u github.com/mattes/migrate/cli
```

##### Usage

https://github.com/mattes/migrate/tree/master/cli

**Create migration file**

```elixir
$ cli create -dir ./migration -ext sql user_additions
```
