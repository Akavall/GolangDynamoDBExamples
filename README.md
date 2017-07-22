To play around:

```
go run create_table.go # or python create_table.py
go run put_item.go
go run get_item.go
```

(You might notice that go commands are a little slow,
this is due to loading libraries, this will not 
be a problem in your app)

To delete the created table:

```
aws dynamodb delete-table --table-name Albums
```

And to validate that the table is gone:

```
aws dynamodb list-tables
```

To delete item using go run:

```
go run delete_item.go
```

To validate that the item was deleted run:

```
aws dynamodb get-item --table-name Albums --key '{"year": {"N": "1991"}, "title": {"S": "Generator"}}'
```

We could also delete item from command line:
```
aws dynamodb delete-item --table-name Albums --key '{"year": {"N": "1991"}, "title": {"S": "Generator"}}'
```
