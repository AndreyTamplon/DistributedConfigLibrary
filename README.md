# Клиентская библиотека для работы с [распределенным конфигом](https://github.com/AndreyTamplon/DistributedConfig)

- ### Пример работы
  
  ```go
  package main
  
  import (
    "cfglib"
  )
  
  func main() {
    c := dictributedConfigLibrary.NewConfigServiceWrapper()
    err := c.Connect(8084, "localhost")
    if err != nil {
      return
    }
    _, err = c.CreateConfig("managed-k8s", map[string]string{"k1": "v1"})
    if err != nil {
      return
    }
    _, err = c.UpdateConfig("managed-k8s", map[string]string{"k2": "v2"})
    if err != nil {
      return
    }
  }
  ```
