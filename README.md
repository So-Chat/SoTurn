# Turn server for SoChat
Made on Go using library pion/turn

**It's not done yet**

### How to use / Как использовать
- Download application from releases page / Скачать приложение со страницы релизов
- On Windows:
    -  ```.\soturn.exe --public-ip (public ip) --realm (realm ip) --jwt (secret string)```
- On Linux or Mac(NOT TESTED):
    -  ```./soturn --public-ip (public ip) --realm (realm ip) --jwt (secret string)```

### Build from Source / Компиляция из исходников
- Install Go 1.26.3 or higher / Установить Go 1.26.3 или выше
- Clone sources with git clone / Клонировать исходиники с git clone
- Install requirements / Установить зависимости: 
    - ```go get github.com/pion/turn/v5```
    - ```go install github.com/pion/turn/v5```
    - ```go get github.com/pion/logging```
    - ```go install github.com/pion/logging```
- Run in terminal / Написать в терминале
    - ```go build .```
- And now you can run it with instructions from above / И теперь вы можете запустить его с инструкциями выше

[**Made for SoChat Server / Сделано для SoChat Server**](https://github.com/yomirein/SoChatServer)

## License - Лицензия
- **This project is licensed under the [MIT License](https://github.com/So-Chat/SoTurn/blob/master/LICENSE)** 
- **Этот проект находится распространяется под [лицензией MIT](https://github.com/So-Chat/SoTurn/blob/master/LICENSE)**
