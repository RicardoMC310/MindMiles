# MindMiles 🚀

![brainLogo](./frontend/src/lib/assets/brain.png)

Somos a empresa que cuida da sua educação! 🎓💡

---

## Como nosso site foi feito?

### BackEnd 🖥️

#### Linguagem de escolha
- Usamos a linguagem **_Go_** pela sua eficiência e praticidade ⚡.

#### Onde testar?
- Nosso *BackEnd* já está em produção! Você pode testá-lo separadamente:  
  - 🌐 [https://indiles-ricardomc3107728-byk1p9cq.leapcell.dev](https://indiles-ricardomc3107728-byk1p9cq.leapcell.dev)

#### Entrypoints disponíveis 🔑

1. **getUsers**  
   - Permite visualizar todos os usuários cadastrados no sistema.  
   - **Observações:** Não requer parâmetros nem body.

2. **createUser**  
   - Permite registrar um novo usuário.  
   - **Body esperado:**
     - `name`: string, entre 2 e 255 caracteres  
     - `email`: string, deve ser um e-mail válido  
     - `password`: string, mínimo 8 caracteres  
     - `rules`: tipo de cadastro
       1. **Admin** - Pode realizar todas as ações 🛠️  
       2. **Teacher** (próximas versões) - Pode criar tarefas para seus alunos 📝  
       3. **Student** - Pode realizar suas tarefas e futuramente trocar por cursos de parceiros 🎓

3. **login**  
   - Permite acessar o sistema.  
   - **Body esperado:**
     - `email`: string, deve ser válido  
     - `password`: string, mínimo 8 caracteres alfanuméricos  
   - **Retorno:**
     - `token`: token codificado em JWT para rotas protegidas 🔐  
     - `user`: dados básicos do usuário 👤

#### Como testar na sua máquina? 🛠️

1. **Instalar o Golang**  
   - 🌐 [![Baixe o Go aqui](https://golang.org/lib/godoc/images/go-logo-blue.svg)](https://go.dev/doc/install)

2. **Clonar o repositório**  
   ```bash
   git clone <url-do-repositorio>
3. Rodar o projeto
   - Navegue até a pasta do backend e siga as instruções de build e run do Go.

## FrontEnd

### Linguagens utilizadas

1. **Svelte**  
   - **HTML**: Estrutura principal para o Svelte trabalhar ![app.html](./frontend/src/app.html)
   - **CSS**: Usado para dar vida a o site, dando estilos únicos à ele ![dir css](./frontend/src/lib/styles)
   - **JS**: Usado nessa etapa para comunicação com o BackEnd ![dir js](./frontend/src/lib/scripts)
   - **SVELTE**: Linguagem própria do svelte, que permite escrever _HTML_, _CSS_(nn utilizado, pois foi a parte) e _JS_ ![dir routes](./frontend/src/routes)
