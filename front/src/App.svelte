<script lang="ts">
	import { onMount } from "svelte";
	import { Task } from "./models";
	import {
		getTasks,
		adTask,
		deleteTask,
		updateTask,
	} from "./api";

	let tasks = [];
	let newTitle: string;

	function tickTask(task: Task) {
		task.done = !task.done;
		updateTask(task, tasks);
	}

	onMount(getTasks(tasks));
</script>

<main>
	<ul>
		{#each tasks as task}
			<li>
				{task.title}
				<input type="checkbox" checked={task.done} on:click={tickTask(task)} />
				<button on:click={deleteTask(task.id, tasks)}>Delete</button>
			</li>
		{/each}
	</ul>
	<input bind:value={newTitle}>
	<button on:click={adTask(tasks, newTitle)}>Ajouter</button>
</main>

<style>
</style>
