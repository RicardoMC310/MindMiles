<script>
	import '$lib/styles/dashboard.css';
	import { user, token } from '$lib/store/store';
	import { goto } from '$app/navigation';
	import { get } from 'svelte/store';

	$: if (typeof window !== "undefined" && (!$user || !$token)) {
		goto("/login");
	}
</script>

<h1 id="welcome-dashboard">Olá {$user?.name}, seja muito Bem-Vindo(a)!</h1>

<div id="leassons-dashboard">
    <h1>Você não tem nenhuma atividade ainda!</h1><br>
    <p>Volte mais tarde para checar atualizações</p>
</div>
