### Requisitos - Contexto do Cliente

A startup **QuickCheck**, que está desenvolvendo um sistema de avaliação de serviços, precisa de um **sistema de autenticação simples** para permitir que usuários se cadastrem e acessem sua plataforma de forma segura.

Atualmente, o time não possui nenhuma camada de login protegida e deseja começar pela base: um **serviço backend minimalista** que permita criar usuários e gerar tokens JWT para autenticação. Esse serviço deve ser facilmente integrável ao futuro frontend da empresa.

### Requisitos Funcionais

1. **Cadastro de Usuário**
   - Deve receber nome, email e senha.
   - Não pode permitir o cadastro de dois usuários com o mesmo email.
   - A senha deve ser armazenada de forma segura (hash futuro, mas pode começar simples).
   - Retornar mensagem de erro clara se o JSON enviado for inválido.
2. **Login de Usuário**
   - Deve aceitar email e senha.
   - Se válidos, retornar um **JWT** com 24h de validade.
   - Se inválidos, retornar erro 401 com mensagem amigável.
3. **Validações**
   - O JSON de entrada deve ser validado (todos os campos obrigatórios).
   - As mensagens de erro devem seguir um padrão único (estrutura padronizada para mensagens e códigos de erro).

### Requisitos Não Funcionais

- A API deve seguir o padrão REST.
- As rotas e a lógica devem ser organizadas em camadas:
  - `model` (estrutura dos dados)
  - `controller` (lógica dos endpoints)
  - `usecase` (lógica de negócio)
  - `repository` (acesso ao banco de dados)
  - `utils` (JWT, mensagens de erro)
  - `router` (registro de rotas)
- O projeto deve carregar configurações sensíveis de um arquivo `.env` (como conexão com banco e chave JWT).
- O banco de dados é um **PostgreSQL já rodando via Docker**, disponível para conexão.

### Exemplos de Mensagens de Erro

```json
json
{
  "message": "Credenciais inválidas",
  "code": 401
}

```

```json
json
{
  "message": "Usuário já existe",
  "code": 409
}

```

```json
json
{
  "message": "Dados inválidos",
  "code": 400
}

```

### Critérios de aceite

- Cadastro com JSON incompleto → Erro 400
- Cadastro com email já existente → Erro 409
- Login com email errado → Erro 401
- Login com sucesso → Token JWT no corpo da resposta
- Cadastro e login feitos separadamente com sucesso
