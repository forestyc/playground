'use client'

import React from 'react';
import {Card, Space} from 'antd';

import {
    useQuery,
    QueryClientProvider,
} from '@tanstack/react-query'
import {getLoanInfo, queryClient} from "@/app/loan/detail/api";


function GetCards() {
    const searchParams = new URLSearchParams(window.location.search);
    const id = searchParams.get("id")
    const {data, error, isLoading} = useQuery({queryKey: ['GetCards'], queryFn: () => getLoanInfo(id)})
    if (isLoading) return <div>Loading...</div>
    if (error) return <div>Error: {error.message}</div>
    if (data.code === -1) return <div>Error: {data.message}</div>
    return (
        <Space direction="vertical" size={16}>
            <Card title={data.object.loan_info.name} extra={<a href="#">More</a>} style={{width: 300}}>
                <p>Total: {data.object.Total}</p>
                <p>ProvidentFund: {data.object.ProvidentFund}</p>
                <p>Business: {data.object.Business}</p>
            </Card>
        </Space>
    )
}


function Cards() {
    return (
        <QueryClientProvider client={queryClient}>
            <GetCards/>
        </QueryClientProvider>
        // <Space direction="vertical" size={16}>
        //     <Card title="Period 1" extra={<a href="#">More</a>} style={{width: 300}}>
        //         {/*<p>Total: {data.Total}</p>*/}
        //         {/*<p>ProvidentFund: {data.ProvidentFund}</p>*/}
        //         {/*<p>Business: {data.Business}</p>*/}
        //     </Card>
        // </Space>
    )
}


export default Cards;