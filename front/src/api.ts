import { Task } from "./models";


export async function getTasks(tasks) {
    const res = await fetch(`http://localhost:10000/tasks`, {
        mode: "cors",
    });
    tasks = await res.json();
}

export async function adTask(tasks, newTitle) {
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

export async function updateTask(task: Task, tasks) {
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

export async function deleteTask(id: number, tasks) {
    const res = fetch(`http://localhost:10000/tasks/${id}`, {
        method: "DELETE",
        mode: "cors",
    });
    res.then(() => {
        tasks = tasks.filter((task) => task.id !== id);
        tasks = tasks;
    });
}