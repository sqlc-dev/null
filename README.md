## Tested

Starting with [drivers](https://zchee.github.io/golang-wiki/SQLDrivers/) that
pass [go-sql-test](https://github.com/bradfitz/go-sql-test).

* https://github.com/mattn/go-sqlite3
* https://github.com/jackc/pgx
* https://github.com/lib/pq
* https://github.com/ziutek/mymysql
* https://github.com/go-sql-driver/mysql/

## Design doc

pgx

```go
type Status byte

const (
	Undefined Status = iota
	Null
	Present
)

type ACLItem struct {
	String string
	Status Status
}
```

stdlib

```go
type NullString struct {
    String string
    Valid  bool // Valid is true if String is not NULL
}
```


* Have another null type for pgtype?

