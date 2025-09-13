<script>
    import "$lib/styles/cadastro.css";
    import { handlerCadastro } from "$lib/scripts/cadastro";
	import { goto } from "$app/navigation";

    const middlewareHandlerCadastro = async (event) => {
        event.preventDefault();
        const res = await (await handlerCadastro(event)).json();

        alert(res?.error??res?.message);
    }

</script>

<h1 id="title-cadastro">Cadastre-Se</h1>

<div id="container-form-cadastro">
    <form onsubmit={async (event) => middlewareHandlerCadastro(event)}>
        <label for="name">Name:</label>
        <input type="text" name="name" placeholder="your name here" min="2" max="255" required>

        <label for="email">Email:</label>
        <input type="email" name="email" placeholder="your best email here" max="255" required>

        <label for="password">Password:</label>
        <input type="password" name="password" placeholder="your password here" required>

        <a href="/cadastro/#" >Já pussui uma conta?</a>

        <button type="submit">Cadastrar-Se</button>
    </form>
</div>