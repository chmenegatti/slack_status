# Uma simples CLI para alterar status do Slack

## Instalação
```bash
git clone git@github.com:chmenegatti/slack_status.git
cd slack_status
go mod download
go build -o slack_status
```

## Criar token de acesso
Acesse o link [https://api.slack.com/apps](https://api.slack.com/apps) e crie um novo app.

Após criar o app, vá em "OAuth & Permissions" e gere um token de acesso.

Edit o arquivo .bashrc ou .zshrc e adicione a linha abaixo:
```bash
export SLACK_TOKEN="seu_token"
```

## Uso
```bash
./slack_status --help
```

## Exemplo

Definir status do slack
```bash
./slack_status set
```

Limpar status do slack
```bash
./slack_status clear
```

Definir status do slack pro almoço
```bash
./slack_status lunch
```

## Contribuição
Pull requests são bem-vindos. Para mudanças importantes, abra um problema primeiro para discutir o que você gostaria de mudar.

## Licença
[MIT](https://choosealicense.com/licenses/mit/)




