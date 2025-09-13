<script>
	import '$lib/styles/app.css';
	import '$lib/styles/header.css';
	import favicon from '$lib/assets/brain.png';
	import { user, token } from '$lib/store/store';
	import { goto } from '$app/navigation';

	let { children } = $props();

	const logout = () => {
		user.set(null);
		token.set(null);
		goto("/");
	};
</script>

<svelte:head>
	<title>MindMiles</title>
	<link rel="icon" href={favicon} type="image/png" />
</svelte:head>

<header id="header-container">
	<div class="logo">
		<h1>MindMiles</h1>
	</div>

	<nav aria-label="Main Navigation">
		<ul>
			<li>
				<a href="/" aria-current="page">Home</a>
			</li>
			<li>
				<a href="/sobreNos">Sobre Nós</a>
			</li>
			<li>
				<a href="/servico">Serviço</a>
			</li>
			<li>
				<a href="/contato">Contato</a>
			</li>
			{#if $user && $token}
				<li>
					<a href="/dashboard">Dashboard</a>
				</li>
				<li>
					<a href="/#" onclick={logout}>Logout</a>
				</li>
			{:else}
				<li>
					<a href="/login">Login</a>
				</li>
				<li>
					<a href="/cadastro">Cadastrar-se</a>
				</li>
			{/if}
		</ul>
	</nav>
</header>

{@render children?.()}

<footer id="footer-container">
	<p>&copy; 2025 <strong>MindMiles</strong>. Todos os direitos reservados.</p>
</footer>
