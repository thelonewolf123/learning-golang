import { useCallback, useEffect, useState } from 'react'

function App() {
    const [task, setTask] = useState('')
    const [tasks, setTasks] = useState<string[]>([])

    const fetchTasks = useCallback(() => {
        fetch('http://localhost:8000/tasks')
            .then((r) => r.json())
            .then((x) => {
                setTasks(x)
            })
    }, [])

    const deleteTask = useCallback(
        (idx: number) => {
            fetch(`http://localhost:8000/task/${idx}`, { method: 'DELETE' })
                .then((r) => r.json())
                .then((x) => {
                    console.log(x)
                    fetchTasks()
                })
        },
        [fetchTasks]
    )

    useEffect(() => {
        fetchTasks()
    }, [fetchTasks])

    return (
        <>
            <div className="flex">
                <input
                    type="text"
                    placeholder="task"
                    onChange={(e) => setTask(e.target.value)}
                    value={task}
                    data-id={'task-input'}
                ></input>
                <button
                    type="button"
                    onClick={() => {
                        fetch('http://localhost:8000/task/add', {
                            method: 'POST',
                            body: JSON.stringify({ task }),
                            headers: {
                                'Content-Type': 'application/json'
                            }
                        })
                            .then((r) => r.json())
                            .then((d) => {
                                console.log('Task added! ', d)
                                setTask('')
                                fetchTasks()
                            })
                    }}
                >
                    submit
                </button>
            </div>
            <ul>
                {tasks.map((z, idx) => {
                    return (
                        <li data-id={z}>
                            <div className="flex list-item">
                                <span> {z}</span>
                                <button
                                    type="button"
                                    data-id={'delete-button'}
                                    onClick={() => {
                                        deleteTask(idx)
                                    }}
                                >
                                    delete
                                </button>
                            </div>
                        </li>
                    )
                })}
            </ul>
        </>
    )
}

export default App
