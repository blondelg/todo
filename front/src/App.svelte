<script lang="ts">
	import { onMount } from "svelte";
	import { Task } from "./models";

	let tasks = [];
	let newTitle: string;

	export async function getTasks() {
    const res = await fetch(`http://localhost:10000/tasks`, {
        mode: "cors",
    });
    tasks = await res.json();

    return tasks;
}

export async function adTask() {
    console.log('DELETINBG');
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

    return tasks;
}

export async function updateTask(task: Task) {
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

    return tasks;
}

export async function deleteTask(id: number) {
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
	<div class="container">
		<div class="row">
			<div class="col-12 text-center">
				<h1>Ma TODO Liste</h1>
			</div>
		</div>
		<div class="row">
			<div class="col-12 text-center">
				<ul class="list-group">
					{#each tasks as task}
						<li class="list-group-item">
							{task.title}
							<input type="checkbox" checked={task.done} on:click={tickTask(task)} />
							<button on:click={deleteTask(task.id)}>Delete</button>
						</li>
					{/each}
				</ul>
			</div>
		</div>

		<input bind:value={newTitle}>
		<button on:click={adTask}>Ajouter</button>
	</div>
</main>
<style>
</style>
