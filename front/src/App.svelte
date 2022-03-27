<script lang="ts">
	import { onMount } from "svelte";
	import { Task, Toast } from "./models";
	import * as bootstrap from "bootstrap";

	let tasks = [];
	let newTitle: string;
	let toastData = new Toast();
	let progress = 0;

	function show() {
		var toastLiveExample = document.getElementById("liveToast");
		var toast = new bootstrap.Toast(toastLiveExample);
		toast.show();
	}

	function getTasksCount() {
		return tasks.length;
	}

	function getTasksDoneCount() {
		return tasks.filter((item) => item.done).length;
	}

	function getProgress() {
		if (getTasksCount() === 0) {
			progress = 0;
			progress = progress;
		}

		if (getTasksCount() > 0) {
			progress = Math.round((100 * getTasksDoneCount()) / getTasksCount());
			progress = progress;
		}
	}

	async function getTasks() {
		const res = await fetch(`http://localhost:10000/tasks`, {
			mode: "cors",
		});
		tasks = await res.json();
		getProgress();
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

		toastData.setTitle("Création");
		toastData.setSubTitle("Maintenant");
		toastData.setBody(`La tâche "${newTask.title}" a été créée.`);
		toastData = toastData;
		show();
		getProgress();
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

		toastData.setTitle("Mise à jour");
		toastData.setSubTitle("Maintenant");
		toastData.setBody(`La tâche "${task.title}" a été mise à jour.`);
		toastData = toastData;
		show();
		getProgress();
	}

	async function deleteTask(id: number, title: string) {
		const res = fetch(`http://localhost:10000/tasks/${id}`, {
			method: "DELETE",
			mode: "cors",
		});
		res.then(() => {
			tasks = tasks.filter((task) => task.id !== id);
			tasks = tasks;
			toastData.setTitle("Suppression");
			toastData.setSubTitle("Maintenant");
			toastData.setBody(`La tâche "${title}" a été supprimée.`);
			toastData = toastData;
			show();
			getProgress();
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
				<i class="fa-solid fa-question-circle" />
			</div>
		</div>
		<div class="row mb-3">
			<div class="col-12">
				<div class="progress">
					{#if progress === 100}
						<div
							class="progress-bar progress-bar-striped progress-bar-animated bg-success"
							role="progressbar"
							style="width: {progress}%;"
							aria-valuenow={progress}
							aria-valuemin="0"
							aria-valuemax="100"
						>
							{progress}%
						</div>
					{:else}
						<div
							class="progress-bar progress-bar-striped progress-bar-animated bg-warning"
							role="progressbar"
							style="width: {progress}%;"
							aria-valuenow={progress}
							aria-valuemin="0"
							aria-valuemax="100"
						>
							{progress}%
						</div>
					{/if}
				</div>
			</div>
		</div>
		<div class="row">
			<div class="col-12">
				<div class="input-group mb-3">
					<input
						bind:value={newTitle}
						type="text"
						class="form-control"
						placeholder="Description de la tâche"
						aria-label="Description de la tâche"
						aria-describedby="new-task"
					/>
					<button
						on:click={adTask}
						class="btn btn-outline-secondary"
						type="button"
						id="new-task">Ajouter une tâche</button
					>
				</div>
			</div>
		</div>
		<div class="row">
			<div class="col-12">
				<ul class="list-group list-group-flush">
					{#each tasks as task}
						<li class="list-group-item">
							<input
								class="form-check-input me-1"
								type="checkbox"
								checked={task.done}
								on:click={tickTask(task)}
							/>

							{#if task.done}
								<del class="text-secondary">{task.title}</del>
							{:else}
								{task.title}
							{/if}

							<button
								class="float-end btn btn-danger m-0"
								on:click={deleteTask(task.id, task.title)}
								>Supprimer</button
							>
						</li>
					{/each}
				</ul>
			</div>
		</div>
	</div>

	<div class="position-fixed bottom-0 end-0 p-3" style="z-index: 11">
		<div
			id="liveToast"
			class="toast"
			role="alert"
			aria-live="assertive"
			aria-atomic="true"
		>
			<div class="toast-header bg-secondary">
				<img src="" class="rounded me-2" alt="" />
				<strong class="me-auto text-light">{toastData.title}</strong>
				<small class="text-light">{toastData.subtitle}</small>
				<button
					type="button"
					class="btn-close"
					data-bs-dismiss="toast"
					aria-label="Close"
				/>
			</div>
			<div class="toast-body">
				{toastData.body}
			</div>
		</div>
	</div>
</main>

<style>
</style>
