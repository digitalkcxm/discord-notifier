# DiscordNotifier Documentation

Para poder utilizar o pacote local, precisa configurar as credenciais para que o GO possa ter acesso ao repositório.

Precisamos adicionar a pasta e todos os seus repositórios na env do go.

`go env -w GOPRIVATE="[gitlab.digitalk.com.br/digitalk/pacotes-customizados/*](http://gitlab.digitalk.com.br/digitalk/pacotes-customizados/*)`

Em seguida precisamos criar um arquivo chamado `~/.netrc` 
Esse arquivo contém as credenciais de acesso aos repositórios de pacotes customizados.

O conteúdo desse arquivo é o seguinte:

```bash
machine gitlab.digitalk.com.br login <username> password <gitlab_access_token>
```

Precisamos definir permissão para esse arquivo criado com o seguinte comando: `chmod 600 ~/.netrc`

Com esses passos feitos, podemos finalmente adicionar o pacote ao projeto.
`go get [gitlab.digitalk.com.br/digitalk/pacotes-customizados/discord-notifier](http://gitlab.digitalk.com.br/digitalk/pacotes-customizados/discord-notifier)`
