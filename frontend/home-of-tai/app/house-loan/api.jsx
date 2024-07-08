export async function getLoanInfo() {
    const response = await fetch('/api/house    loan/repayment?date=2053-08-15')
    if (!response.ok) {

        throw new Error(response.statusText)
    }
    return await response.json()
}
