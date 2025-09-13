<script>
	import '$lib/styles/contato.css';
    import { handlerContact } from "$lib/scripts/contato";
</script>

<!-- Fazer links para redes sociais -->

<div id="container-form-contact">
    <h1>Contate a gente!</h1>

	<form id="form-container-contact" onsubmit={async event => handlerContact(event)}>
		<label for="name">Name: </label>
		<input type="text" name="name" placeholder="your name here" required/>

		<label for="email">Email: </label>
		<input type="email" name="email" placeholder="your best email here" required/>

		<label for="content">Content: </label>
		<textarea name="content" rows="10" cols="30" placeholder="content message here" required></textarea>

        <button type="submit">Enviar</button>
	</form>
</div>
