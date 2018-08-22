# Back drweb-test

Необходим установленный в системе go (разрабатывал на `go1.9.3 darwin/amd64`) и `glide` (пакетный менеджер)

1. Идем на https://developers.google.com/drive/api/v3/quickstart/go, жмем "Enable The Drive API", выбираем любой проект или на лету создаем новый, получаем credentials.json, который нужно положить в корень этого проекта.
2. `glide i`
3. `go build`
4. Копируем `config.yaml.example` в `config.yaml`. Можно там ничего не менять
5. chmod, если необходимо, и запускаем `./go-server --config-file=config.yaml`
6. Enjoy
