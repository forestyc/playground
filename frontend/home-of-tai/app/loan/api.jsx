import {QueryClient} from "@tanstack/react-query";

export const queryClient = new QueryClient()
export async function getLoanList() {
    const response = await fetch('/api/home/loan/list')
    if (!response.ok) {
        throw new Error(response.statusText)
    }
    return await response.json()
}

export async function getLoanInfo(id) {
    console.log(id)
    const response = await fetch('/api/home/loan/?id=' + id)
    if (!response.ok) {
        throw new Error(response.statusText)
    }
    return await response.json()
}