# Turn server for SoChat
Made on Go using library pion/turn

**It's not done yet**

### How to use / Как использовать
- Download application from releases page / Скачать приложение с 
- On Windows:
    -  ```.\soturn.exe --public-ip (public ip) --realm (realm ip) --jwt (secret string)```
- On Linux or Mac(NOT TESTED):
    -  ```./soturn --public-ip (public ip) --realm (realm ip) --jwt (secret string)```

### Build from Source / Компиляция из исходников
- Install Go
- Clone sources with git clone
- Install requirements:
    - ```go get github.com/pion/turn/v5```
    - ```go install github.com/pion/turn/v5```
    - ```go get github.com/pion/logging```
    - ```go install github.com/pion/logging```
- Run in terminal
    - ```go build .```
- And now you can run it with instructions from above

[**Made for SoChat Server / Сделано для SoChat Server**](https://github.com/yomirein/SoChatServer)

## License - Лицензия