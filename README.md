Библиотека для чтения показаний приборов учета энергоресурса через API [АИИС ЭЛДИС](https://eldis24.ru).

Библиотека позволяет:
* получать список приборов учета энергоресурса, связанных с учетной записью пользователя *АИИС ЭЛДИС*;
* читать архивы нормализованных (прошедших достоверизацию, то есть проверку на наличие некорректных и/или отсутствующих показаний и периодов отклонений в нормальной работе прибора) показаний;
* читать архивы "сырых" показаний, т.е. показаний в том виде, в котором они получены непосредственно от прибора учета

Всё это необходимо, если вам нужно интегрировать свою информационную систему с *АИИС ЭЛДИС* и/или использовать свою реализацию алгоритмов достоверизации показаний приборов учета.

### Использование

```go
import "github.com/vitpelekhaty/go-eldis/v2"
```

### Интерфейс Connection

Интерфейс обеспечивает низкоуровневое взаимодействие с API *АИИС ЭЛДИС*.

Пример чтения списка приборов учета в формате JSON:

```go
package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	
	"github.com/vitpelekhaty/go-eldis/v2"
	"github.com/vitpelekhaty/go-eldis/v2/responses"
)

func main() {
	...
	conn, err := eldis.Connect(context.Background(), rawURL, eldis.Credentials{Username: username, Password: password, AccessToken: accessToken})
	
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = conn.Close(context.Background())
	}()

	b, err := conn.ListForDevelopment(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	sb := responses.Extract(responses.SectionListForDevelopment, bytes.NewBuffer(b))

	fmt.Print(string(sb))
}
```

### Пакет responses

Пакет нужен для извлечения и разбора блоков архивов показаний из ответа API *АИИС ЭЛДИС*, чтения отдельных полей показаний и т.п.

Пример ниже демонстрирует чтение часового архива "сырых" показаний по точке учета (тепловому вводу прибора учета):

```go
package main

import (
	"bytes"
	"context"
	"log"
	
	"github.com/vitpelekhaty/go-eldis/v2"
	"github.com/vitpelekhaty/go-eldis/v2/responses"
	"github.com/vitpelekhaty/go-eldis/v2/responses/readings/raw"
)

func main() {
	...
	conn, err := eldis.Connect(context.Background(), rawURL, eldis.Credentials{Username: username, Password: password, AccessToken: accessToken})
	
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = conn.Close(context.Background())
	}()

	b, err := conn.RawReadings(context.Background(), pointID, eldis.HourArchive, from, to)

	if err != nil {
		log.Fatal(err)
	}
	
	// извлекаем из ответа API данные блока "сырых" показаний
	sb := responses.Extract(responses.SectionRaw, bytes.NewBuffer(b))

	// выполняем чтение строк архива "сырых" показаний
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	
	items, err := raw.Parse(ctx, bytes.NewReader(sb))

	if err != nil {
		log.Fatal(err)
	}

	for item := range items {
		// если прочитанный элемент не содержит ошибку разбора
		if !item.IsError() {
			DoSomething(item)
		} else {
			log.Print(item.E)
			cancelFunc()
		}
	}
}
```
