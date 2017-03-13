package boxes

import (
	"fmt"
    "os"
    "io/ioutil"
    "encoding/json"
    "time"
)

type data struct {
    List []interface{} `json:"list"`
    Date int64 `json:"date"`
}

type Box struct {
    fullpath string
    write func(interface{}) map[string]interface{}
    read func(map[string]interface{}) interface{}
}

func (b *Box) Add(e ...interface{}) {
    if e == nil || len(e) == 0 {
        return
    }
    list := b.GetAll()
    list = append(list, e...)
    b.save(list)
}

func (b *Box) AddAll(e []interface{}) {
    if e == nil || len(e) == 0 {
        return
    }
    list := b.GetAll()
    list = append(list, e...)
    b.save(list)
}

func (b *Box) GetAll() []interface{} {
    d, e := ioutil.ReadFile(b.fullpath)
    if e != nil {
        panic(e)
    }
    res := data{}
    json.Unmarshal(d, &res)
    for i := range res.List {
        res.List[i] = b.read(res.List[i].(map[string]interface{}))
    }
    return res.List
}

func (b *Box) Query(q func(interface{}) bool) []interface{} {
    d, e := ioutil.ReadFile(b.fullpath)
    if e != nil {
        panic(e)
    }
    res := data{}
    json.Unmarshal(d, &res)
    list := make([]interface{}, 0)
    for i := range res.List {
        v := b.read(res.List[i].(map[string]interface{}))
        if q(v) {
            list = append(list, v)
        }
    }
    return list
}

func (b *Box) Replace(q func(interface{}) bool, e interface{}) {
    list := b.GetAll()
    for i:= range list {
        if q(list[i]) {
            list[i] = e
            b.save(list)
            return;
        }
    }
}

func (b *Box) RemoveFirst(q func(interface{}) bool) {
    list := b.GetAll()
    for i:= range list {
        if q(list[i]) {
            list = append(list[:i], list[i+1:]...)
            b.save(list)
            return;
        }
    }
}

func (b *Box) save(l []interface{}) {
    f, err := os.OpenFile(b.fullpath, os.O_WRONLY|os.O_TRUNC, 0)
    if err != nil {
        panic(err)
    }
    defer func() {
        if err := f.Close(); err != nil {
            panic(err)
        }
    }()
    for i := range l {
        l[i] = b.write(l[i])
    }
    d,_ := json.Marshal(data{l, time.Now().UnixNano() / int64(time.Millisecond)})
    f.Write([]byte(string(d)))
}

func New(write func(interface{}) map[string]interface{}, read func(map[string]interface{}) interface{}, fp string) Box {
    create(fp)
    return Box{fp, write, read}
}

func create(path string) {
    _, e := os.Stat(path)
    if os.IsNotExist(e) {
        fmt.Println(e.Error())
    } else if e != nil {
        panic(e)
    } else {
        return
    }
    fo, err := os.Create(path)
    if err != nil {
        panic(err)
    }
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()
    d,_ := json.Marshal(data{make([]interface{}, 0), time.Now().UnixNano() / int64(time.Millisecond)})
    fo.Write([]byte(string(d)))
}

