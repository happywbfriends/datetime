
# Datetime utils


### Простой парсинг времени

```
    /*
        Понимает следующие форматы:
        1. 2006-01-02               => 2006-01-02T00:00:00 UTC
        2. 2006-01-02T15:04         => 2006-01-02T15:04:00 UTC
        3. 2006-01-02T15:04:05      => 2006-01-02T15:04:05 UTC
        4. 2006-01-02T15:04:05TZ    => 2006-01-02T15:04:05 TZ
        5. 2006-01-02T15:04:05.999  => 2006-01-02T15:04:05.999 UTC
        6. 2006-01-02T15:04:05.999TZ=> 2006-01-02T15:04:05.999 TZ
        
        Всегда возвращает часовой пояс UTC
    */
    dtm, err := ParseTime("...") 
    
    // Сериализация
    fmt.Print(SerializeTime(dtm))
```

### Парсинг времени из JSON

```
    type Foo struct {
        Bar SerializedTime `json:"t,omitempty"`
    }
    
    foo := Foo{}
    _ = json.Unmarshal(..., &foo)
    
    fmt.Print(foo.Bar.Time)
    
    data, _ := json.Marshal(&foo)

    fmt.Print(data)
```
