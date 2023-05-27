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
                ></input>
                <button
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
                {tasks.map((z) => {
                    return <li>{z}</li>
                })}
            </ul>
        </>
    )
}

export default App
