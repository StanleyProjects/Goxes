# Goxes
Object-Relational Mapping experiment for golang
(Redesign [Boxes](https://github.com/StanleyProjects/Boxes))

<img src="media/icon.png" width="128" height="128" />

# Box

##### Example 1.1. Create *Box*

```go
transactionsBox := boxes.New(func(i interface{}) map[string]interface{} {
        transaction := i.(TransactionModel)
        m := make(map[string]interface{})
        m["id"] = transaction.GetId()
        m["date"] = transaction.GetDate()
        m["count"] = transaction.GetCount()
        return m
    }, func (m map[string]interface{}) interface{} {
        return &Transaction{
            int(m["id"].(float64)),
            int64(m["date"].(float64)),
            int(m["count"].(float64)),
        }
    }, "/home/user/stan/goxes/transactionsbox")
```

##### Example 1.2. Put objects to *Box*

```go
transactionsBox.Add(&Transaction{1, 100, 123},
					&Transaction{2, 58, 124},
					&Transaction{3, -45, 125},
					&Transaction{4, 23, 126},
					&Transaction{5, -78, 127})
```

##### Example 1.3. Get list objects from *Box*

```go
for i, v := range transactionsBox.GetAll() {
    transaction := v.(*Transaction)
    fmt.Println(i, ")", "count:", transaction.GetCount())
}
```

##### Example 1.4. Get list objects from *Box* with query

```go
for i, v := range transactionsBox.Query(func(e interface{}) bool {
    transaction := e.(*Transaction)
    return transaction.GetCount() < 75
}) {
    transaction := v.(*Transaction)
    fmt.Println(i, ")", "count:", transaction.GetCount())
}
```

##### Example 1.5. Replace object in *Box*

```java
transactionsBox.Replace(func(e interface{}) bool {
    transaction := e.(*Transaction)
    return transaction.GetId() == 3
}, &Transaction{3, 23, 12})
```

##### Example 1.6. Remove object from *Box*

```java
transactionsBox.RemoveFirst(func(e interface{}) bool {
    transaction := e.(*Transaction)
    return transaction.GetId() == 4
})
```
