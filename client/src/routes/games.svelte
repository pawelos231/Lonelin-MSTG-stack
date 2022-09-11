<script lang="ts">
	const ExampleGames: Array<string> = [
		'Jungla 1',
		'Stack it!',
		'Arkanoid',
		'space destruction',
		'Diamonds'
	];
	//for now TODO: fetch only all  the names from database and paginate them if the result is greater then 20
</script>

<div class="flex justify-center">
	<div class="mt-20 w-screen flex flex-col items-center">
		<h1 class="text-center">Gry</h1>
		<div class="flex justify-center gap-4 mt-20 w-4/5 flex-wrap">
			{#each ExampleGames as game}
				<div class="flex flex-col basis-1/5 bg-blue-500 rounded overflow-hidden">
					<div class="bg-black text-white text-lg w-full text-center p-2 ">{game}</div>
					<div class="p-5">
						Lorem ipsum dolor sit amet consectetur adipisicing elit. Esse illo veritatis, asperiores
						tempore odit aperiam repudiandae eos, rem a iusto quia maxime! Reiciendis quod atque
						delectus fugiat. Suscipit, repellat doloribus?
					</div>
				</div>
			{/each}
		</div>
	</div>
</div>
