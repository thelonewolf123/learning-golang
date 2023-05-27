import { useEffect, useState } from 'react'

export function Tasks() {
    const [tasks, setTasks] = useState<string[]>([])
    useEffect(() => {
        fetch('http://localhost:8000/tasks')
            .then((r) => r.json())
            .then((x) => {
                setTasks(x)
            })
    }, [])
    return (
        <ul>
            {tasks.map((z) => {
                return <li>{z}</li>
            })}
        </ul>
    )
}
