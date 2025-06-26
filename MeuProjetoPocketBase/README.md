# Meu Projeto PocketBase Customizado

Este projeto usa PocketBase com customizações locais.

## 🚀 Como executar

### Opção 1: Usando JavaScript (Mais simples)

1. Baixe o PocketBase executável: https://github.com/pocketbase/pocketbase/releases
2. Coloque na raiz do projeto
3. Execute com os hooks JavaScript:
   ```bash
   ./pocketbase serve --hooksDir=./pb_hooks
   ```

### Opção 2: Compilando com Go (Mais poderoso)

1. Instale Go 1.23+: https://go.dev/download/
2. Instale as dependências:
   ```bash
   go mod tidy
   ```
3. Compile o projeto:
   ```bash
   go build -o meu-pocketbase.exe
   ```
4. Execute:
   ```bash
   ./meu-pocketbase serve
   ```

## 📁 Estrutura do Projeto

```
MeuProjetoPocketBase/
├── pb_hooks/           # Customizações JavaScript (versionado)
│   └── main.pb.js      # Hooks e rotas customizadas JS
├── pb_migrations/      # Migrações do banco (versionado)
├── pb_public/          # Arquivos estáticos (versionado)
├── pb_data/            # Dados e uploads (NÃO versionado)
├── main.go             # Customizações em Go
├── go.mod              # Dependências Go
├── .gitignore          # Arquivos ignorados
└── README.md           # Este arquivo
```

## 🛠️ Customizações Disponíveis

### JavaScript (pb_hooks/main.pb.js)
- Hooks para eventos do sistema
- Rotas customizadas
- Validações adicionais
- Integrações com APIs externas

### Go (main.go)
- Controle total sobre o PocketBase
- Middleware customizado
- Extensões complexas
- Performance otimizada

## 🔗 URLs

- Admin Dashboard: http://127.0.0.1:8090/_/
- API REST: http://127.0.0.1:8090/api/
- Rota customizada: http://127.0.0.1:8090/api/custom/hello

## 🔒 Versionamento

### ✅ Versionar:
- `pb_hooks/` - Scripts JavaScript
- `pb_migrations/` - Migrações do banco
- `pb_public/` - Frontend/assets
- `main.go` - Código Go customizado
- `go.mod` - Dependências

### ❌ NÃO versionar:
- `pb_data/` - Dados do banco
- Executáveis (`.exe`)
- `vendor/` - Dependências Go
- `.env` - Variáveis de ambiente

## 📝 Desenvolvimento

1. Faça suas customizações em `pb_hooks/` ou `main.go`
2. Teste localmente
3. Commite apenas os arquivos de código
4. Outros desenvolvedores devem baixar/compilar o executável 