# Back drweb-test

1. Идем на https://developers.google.com/drive/api/v3/quickstart/go, жмем "Enable The Drive API", выбираем любой проект или на лету создаем новый, получаем credentials.json, который нужно положить в корень этого проекта.

2. `go build`

3. Копируем config.yaml.example в config.yaml. Можно там ничего не менять

4. chmod, если необходимо, и запускаем `./go-server --config-file=config.yaml`

5. Enjoy