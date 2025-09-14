# MindMiles 🚀

![brainLogo](./frontend/src/lib/assets/brain.png)

Somos a empresa que cuida da sua educação! 🎓💡

-------

*Etec*: Dr. Geraldo Jośe Rodrigues Alckmin

*Integrantes*:
- Ricardo Matos Costa 1º DS/AMS
- Mariana Peres Silva 1º DS/AMS
- Nicolas Luis da Silva 1º DS/AMS

*Qual a nossa proposta:*
- A nossa proposta é ajudar os estudantes de todo o país a conseguir se adaptar melhor com a sua forma de aprendizado.
- Com auxilio da nossa aplicação, o estudante terá mais oportunidade de entender e compreender aquela disciplina.
- Pois terá acesso a várias atividades. além de recompensas. E é claro, podendo tirar dúvidas com outros alunos e professores.

-------

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
   - Navegue até a pasta do backend e rode `go run main.go`

-------

## FrontEnd

### Linguagens utilizadas

1. **Svelte**  
   - **HTML**: Estrutura principal para o Svelte trabalhar ![app.html](./frontend/src/app.html)
   - **CSS**: Usado para dar vida a o site, dando estilos únicos à ele ![dir css](./frontend/src/lib/styles)
   - **JS**: Usado nessa etapa para comunicação com o BackEnd ![dir js](./frontend/src/lib/scripts)
   - **SVELTE**: Linguagem própria do svelte, que permite escrever _HTML_, _CSS_(nn utilizado, pois foi a parte) e _JS_ ![dir routes](./frontend/src/routes)

### Aonde eu testo?
  1. Também já em producão, e está disponível em:
     - link: [https://mindmiles-ldqr1pyy8-ricardomc310s-projects.vercel.app](https://mindmiles-ldqr1pyy8-ricardomc310s-projects.vercel.app)

### E como executo na minha máquina?
  1. Tenha o *NodeJS* instalado:
     - Link de download: [Download do NodeJS](https://nodejs.org/pt/download)
  2. Com o repositório já clonado:
     - entrar na pasta frontend e executar os seguintes comandos:
       1. `npm install`
       2. `npm run dev -- --open`
          - Fará com que abra no navegador padrão.
          - Se não abrir só entrar no navegador e digitar `http://localhost:5173`
  
