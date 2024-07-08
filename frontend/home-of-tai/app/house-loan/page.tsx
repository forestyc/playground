import React from 'react';
import Cards from "@/app/house-loan/cards";

// `app/page.tsx` is the UI for the `/` URL
export default function Page() {

    return (
        <main className="flex min-h-screen flex-col items-center justify-between p-24">
            <Cards></Cards>
        </main>
    )
}