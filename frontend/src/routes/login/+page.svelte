<script>
	import '$lib/styles/login.css';
	import { handlerLogin } from '$lib/scripts/login';
	import { setContext } from 'svelte';
	import { writable } from 'svelte/store';
	import { user, token } from '$lib/store/store';
	import { goto } from '$app/navigation';

	let clickButtonLogin = false;

	const middlewareHandlerLogin = async (event) => {
		clickButtonLogin = true;
		event.preventDefault();
		const res = await handlerLogin(event);

		user.set(res.user ?? null);
		token.set(res.token ?? null);

		goto('/');
	};
</script>

<h1 id="title-login">Logar-Se</h1>

<div id="container-form-login">
	<form method="post" onsubmit={async (event) => middlewareHandlerLogin(event)}>
		<label for="email">Email:</label>
		<input type="email" name="email" min="2" placeholder="your email here" required />

		<label for="password">Password:</label>
		<input type="password" name="password" min="8" required />

		<a href="/cadastro">Não tem conta?</a>

		{#if !clickButtonLogin}
			<button type="submit">Logar</button>
		{/if}
	</form>
</div>
