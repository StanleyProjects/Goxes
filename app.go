package main

import (
	"fmt"
    "os"
    "runtime"
    "github.com/StanleyProjects/Goxes/boxes"
)

type TransactionModel interface {
    GetId() int
    GetDate() int64
    GetCount() int
}

type Transaction struct {
    lala_id int
    lala_date int64
    lala_count int
}
func (t *Transaction) GetId() int {
    return t.lala_id
}
func (t *Transaction) GetDate() int64 {
    return t.lala_date
}
func (t *Transaction) GetCount() int {
    return t.lala_count
}

func userHomeDir() string {
    env := "HOME"
    if runtime.GOOS == "windows" {
        env = "USERPROFILE"
    } else if runtime.GOOS == "plan9" {
        env = "home"
    }
    return os.Getenv(env)
}

func main() {

    box := boxes.New(func(i interface{}) map[string]interface{} {
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
        }, userHomeDir() + "/stan/goxes/testbox")
    //fmt.Println(box.GetAll())
    box.Add(&Transaction{1, 100, 123},
            &Transaction{2, 58, 124},
            &Transaction{3, -45, 125},
            &Transaction{4, 23, 126},
            &Transaction{5, -78, 127})
    //box.AddAll([]interface{}{&Transaction{76, 12, 90}, &Transaction{-45, 23, 76}})
    /*
    for _, v := range box.GetAll() {
        transaction := v.(*Transaction)
        fmt.Println(i, ")", "count:", transaction.GetCount())
    }
    */

    for i, v := range box.Query(func(e interface{}) bool {
        transaction := e.(*Transaction)
        return transaction.GetCount() < 75
    }) {
        transaction := v.(*Transaction)
        fmt.Println(i, ")", "count:", transaction.GetCount())
    }

    box.Replace(func(e interface{}) bool {
        transaction := e.(*Transaction)
        return transaction.GetId() == 3
    }, &Transaction{3, 23, 12})
    box.RemoveFirst(func(e interface{}) bool {
        transaction := e.(*Transaction)
        return transaction.GetId() == 4
    })

    fmt.Println(" HOME: ", userHomeDir())
    fmt.Println(" box ", box)
    fmt.Println("\tend")
}
