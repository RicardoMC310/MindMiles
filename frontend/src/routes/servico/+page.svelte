<script>
	import { goto } from '$app/navigation';
	import '$lib/styles/servico.css';
	import brainLogo from '$lib/assets/brain.png';
</script>

<div id="inscriptions-service">
	<button onclick={() => goto('#')}>Cadastrar-se como aluno</button>
</div>

<div id="first-layer-service">
	<p>
		Após o cadastro, você terá acesso imediato aos cursos em que já está matriculado e também poderá
		se inscrever em novos. Além disso, contará com uma aba exclusiva para buscar cursos. Não deixe
		de conferir as atividades propostas por professores cadastrados: ao concluí-las, você receberá <strong
			>MindCoins</strong
		>, que podem ser utilizados em nossa loja. E lembre-se:
		<strong>Aprenda. Cresça. Alcance!</strong>
	</p>
	<p>
		Conecte-se também com outros alunos e professores, compartilhe experiências e faça parte de uma
		comunidade que valoriza o conhecimento. Juntos, vocês poderão trocar ideias, colaborar em
		projetos e ampliar suas oportunidades.
	</p>
</div>

<div id="second-layer-service">
	<img src={brainLogo} alt="imagem de cérebro treinando" />
	<p>
		Não perca o foco: o maior desafio é superar a si mesmo. Com dedicação e esforço, você alcançará
		seus objetivos. Nada vem fácil, mas isso nunca deve ser motivo para desistirmos dos nossos
		sonhos.
	</p>
</div>
