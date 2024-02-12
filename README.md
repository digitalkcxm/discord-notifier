# How to Install

Para poder utilizar o pacote local, precisa configurar as credenciais para que o GO possa ter acesso ao repositório.

1. Precisamos adicionar a pasta e todos os seus repositórios na env do go.
```
go env -w GOPRIVATE="gitlab.digitalk.com.br/digitalk/pacotes-customizados/*"
```

2. Em seguida precisamos criar um arquivo chamado `~/.netrc`. Esse arquivo deve conter as credenciais de acesso aos repositórios de pacotes customizados. O conteúdo desse arquivo deve seguir o seguinte padrão:
```bash
machine gitlab.digitalk.com.br login <username> password <gitlab_access_token>
```

3. Precisamos definir permissão para esse arquivo criado com o seguinte comando: 
```
chmod 600 ~/.netrc
```

4. Com esses passos feitos, podemos finalmente adicionar o pacote ao projeto.

```
go get gitlab.digitalk.com.br/digitalk/pacotes-customizados/discord-notifier
```
