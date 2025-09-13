<script>
	import '$lib/styles/app.css';
	import '$lib/styles/header.css';
	import favicon from '$lib/assets/brain.png';

	let { children } = $props();
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
		</ul>
	</nav>
</header>

{@render children?.()}

<footer id="footer-container">
	<p>&copy; 2025 <strong>MindMiles</strong>. Todos os direitos reservados.</p>
</footer>
