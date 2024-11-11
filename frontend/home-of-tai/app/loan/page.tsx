import React from 'react';
import LoanList from "@/app/loan/list";


// `app/page.tsx` is the UI for the `/` URL
export default function Page() {

    return (
        // <main className="flex min-h-screen flex-col items-center justify-between p-24">
            <LoanList ></LoanList>
        // </main>
    )
}