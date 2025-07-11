// Load tasks when page loads
document.addEventListener('DOMContentLoaded', loadTasks);

async function loadTasks() {
    const response = await fetch('/api/tasks');
    const tasks = await response.json();
    const taskList = document.getElementById('task-list');
    taskList.innerHTML = '';
    
    tasks.forEach(task => {
        const li = document.createElement('li');
        li.className = `task-item ${task.done ? 'completed' : ''}`;
        li.innerHTML = `
            <div class="task-text">◻️ ${task.title}</div>
            <div class="task-actions">
                <button onclick="toggleTask(${task.id})">${task.done ? 'Undo' : 'Complete'}</button>
                <button class="delete-btn" onclick="deleteTask(${task.id})">Delete</button>
            </div>
        `;
        taskList.appendChild(li);
    });
}

async function addTask() {
    const input = document.getElementById('task-input');
    const title = input.value.trim();
    
    if (title) {
        const response = await fetch('/api/tasks', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ title: title, done: false })
        });
        
        if (response.ok) {
            input.value = '';
            loadTasks();
        }
    }
}

async function toggleTask(id) {
    const task = await getTaskById(id);
    if (task) {
        const response = await fetch(`/api/tasks/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ done: !task.done })
        });
        
        if (response.ok) {
            loadTasks(); // Перезагружаем список задач
        }
    }
}

async function deleteTask(id) {
    const response = await fetch(`/api/tasks/${id}`, {
        method: 'DELETE'
    });
    
    if (response.ok) {
        loadTasks();
    }
}

async function getTaskById(id) {
    const response = await fetch('/api/tasks');
    const tasks = await response.json();
    return tasks.find(task => task.id === id);
}