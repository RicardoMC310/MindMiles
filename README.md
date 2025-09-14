# MindMiles

![brainLogo](./frontend/src/lib/assets/brain.png)

Somos a empresa que cuida da sua educação!

## Como nosso site foi feito?

### BackEnd

#### Linguagem de escolha
- Usamos a linguagem _Go_ pela sua eficiência e praticidade.

#### Onde testar?
- Meu *BackEnd* já está em produção! Você pode testá-lo separadamente a partir da URL:  
  - [https://indiles-ricardomc3107728-byk1p9cq.leapcell.dev](https://indiles-ricardomc3107728-byk1p9cq.leapcell.dev)

#### Entrypoints disponíveis
1. **getUsers**  
   - Permite visualizar todos os usuários cadastrados no sistema.  
   - **Observações:** Não requer parâmetros nem body.

2. **createUser**  
   - Permite registrar um novo usuário.  
   - **Body esperado:**
     - `name`: string, entre 2 e 255 caracteres  
     - `email`: string, deve ser um e-mail válido  
     - `password`: string, no mínimo 8 caracteres  
     - `rules`: tipo de cadastro
       1. **Admin** - Pode realizar todas as ações  
       2. **Teacher** (próximas versões) - Pode criar tarefas para seus alunos  
       3. **Student** - Pode realizar suas tarefas e futuramente trocar por cursos de parceiros

3. **login**  
   - Permite acessar o sistema.  
   - **Body esperado:**
     - `email`: string, deve ser válido  
     - `password`: string, mínimo 8 caracteres alfanuméricos  
   - **Retorno:**
     - `token`: token codificado em JWT para rotas protegidas  
     - `user`: dados básicos do usuário

#### Como testo na minha máquina?
  1. Instalar o `Golang` na sua máquina.
     - link: [![site para download do Golang](https://go.dev/doc/install)](https://go.dev/doc/install)
  2. Dar comando pra clonar o repositório:
    ```bash
       git clone <url do repositório>
    ```
