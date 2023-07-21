
# Datetime utils

Утилиты для парсинга и сериализации времени по стандарту WB Seller API.

Стандарт подразумевает, что время может быть передано в одном из следующих форматов:

```
1. 2006-01-02               => 2006-01-02T00:00:00 UTC
2. 2006-01-02T15:04         => 2006-01-02T15:04:00 UTC
3. 2006-01-02T15:04:05      => 2006-01-02T15:04:05 UTC
4. 2006-01-02T15:04:05TZ    => 2006-01-02T15:04:05 TZ
5. 2006-01-02T15:04:05.999  => 2006-01-02T15:04:05.999 UTC
6. 2006-01-02T15:04:05.999TZ=> 2006-01-02T15:04:05.999 TZ
```

Парсер ВСЕГДА возвращает время в часовом поясе UTC.


# Примеры

## Ручной парсинг и сериализация времени 

```
    import github.com/happywbfriends/datetime/datetime

    // Parse
    dtm, err := datetime.ParseTime("2023-01-01T13:00")
    
    // Serialize
    fmt.Print(datetime.SerializeTime(time.Now(), true))
     
```

## Парсинг и сериализация в JSON поле

Чтобы автоматически распарсить время в поле структуры при использовании `json.Unmarshal`, поле следует объявить типа `SerializedTime`

```
    import github.com/happywbfriends/datetime/datetime

    type Foo struct {
        Bar datetime.ParsedTime `json:"t"`
    }
    
    // Unmarshal
    foo := Foo{}

    _ = json.Unmarshal([]byte{`{"t": "2023-01-01"}`}, &foo)
    
    fmt.Print(foo.Bar.Time) // Prints "2023-01-01 00:00:00 +0000 UTC"
    
    // Marshal
    foo.Bar.Time = time.Date(2030, 1, 2, 0, 0, 0, 0, time.UTC)
    
    data, _ := json.Marshal(&foo)

    fmt.Print(data) // Prints `{"t":"2030-01-02T00:00:00Z"}`
```
