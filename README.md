To play around:

```
python make_tables.py
go run put_item.go
go run get_item.go
```

To delete the created table:

```
aws dynamodb delete-table --table-name Albums
```

And to validate that the table is gone:

```
aws dynamodb list-tables
```
