<script lang="ts">
	import { onMount } from "svelte";

	class Task {
		id: number;
		title: string;
		done: boolean;

		deserialize(input: object) {
			Object.assign(this, input);
			return this;
		}
	}

	let tasks = [];
	let newTitle: string;

	async function getTasks() {
		const res = await fetch(`http://localhost:10000/tasks`, {
			mode: "cors",
		});
		tasks = await res.json();
	}

	async function adTask() {
		let newTask = new Task();
		newTask.title = newTitle;
		newTask.done = false;
		const res = await fetch(`http://localhost:10000/tasks/create`, {
			method: "POST",
			mode: "cors",
			body: JSON.stringify(newTask),
		});
		newTask = new Task().deserialize(await res.json());
		tasks = [...tasks, newTask];
		newTitle = "";
	}

	async function updateTask(task: Task) {
		const res = await fetch(`http://localhost:10000/tasks/update`, {
			method: "PUT",
			mode: "cors",
			body: JSON.stringify(task),
		});
		task = new Task().deserialize(await res.json());
		const index = tasks.findIndex((element, index) => {
			if (element.id === task.id) {
				return true;
			}
		});
		tasks[index] = task;
		tasks = tasks;
	}

	async function deleteTask(id: number) {
		const res = fetch(`http://localhost:10000/tasks/${id}`, {
			method: "DELETE",
			mode: "cors",
		});
		res.then(() => {
			tasks = tasks.filter((task) => task.id !== id);
			tasks = tasks;
		});
	}

	function tickTask(task: Task) {
		task.done = !task.done;
		updateTask(task);
	}

	onMount(getTasks);
</script>

<main>
	<ul>
		{#each tasks as task}
			<li>
				{task.title}
				<input type="checkbox" checked={task.done} on:click={tickTask(task)} />
				<button on:click={deleteTask(task.id)}>Delete</button>
			</li>
		{/each}
	</ul>
	<input bind:value={newTitle} />
	<button on:click={adTask}>Ajouter</button>
</main>

<style>
</style>
