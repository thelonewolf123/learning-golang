import { test, expect } from '@playwright/test'

test.describe('todo workflow', () => {
    let task: null | string
    test.beforeAll(() => {
        task = task ?? `task-${Math.round(Math.random() * 1000)}`
    })

    test('add new task', async ({ page }) => {
        if (!task) throw new Error('Invalid task!')
        console.log({ task })
        await page.goto('http://localhost:5173/')
        expect(page.getByPlaceholder('task')).toBeTruthy()
        expect(page.getByText('submit')).toBeTruthy()
        await page.getByPlaceholder('task').click()
        await page.getByPlaceholder('task').type(task)
        await page.getByText('submit').click()
        await expect(page.locator(`[data-id=${task}]`)).toBeVisible()
    })

    test('delete task', async ({ page }) => {
        if (!task) throw new Error('Invalid task!')
        console.log({ task })
        await page.goto('http://localhost:5173/')
        await expect(
            page.locator(`[data-id=${task}]`).locator('button')
        ).toBeVisible()
        await page.locator(`[data-id=${task}]`).locator('button').click()
        await expect(page.locator(`[data-id=${task}]`)).toHaveCount(0)
    })
})
