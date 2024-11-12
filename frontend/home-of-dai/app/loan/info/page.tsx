'use client'

import React from 'react';
import {Card, Space} from 'antd';

import {QueryClientProvider, useQuery,} from '@tanstack/react-query'
import {getLoanInfo, queryClient} from "@/app/loan/info/api";

type Repayment = {
    amount: number
    repayment_date: Date
    name: string
}

let repaymentList: Repayment[]

function GetCards() {
    const searchParams = new URLSearchParams(window.location.search);
    const id = searchParams.get("id")
    const {data, error, isLoading} = useQuery({queryKey: ['GetCards'], queryFn: () => getLoanInfo(id)})
    if (isLoading) return <div>Loading...</div>
    if (error) return <div>Error: {error.message}</div>
    if (data.code === -1) return <div>Error: {data.message}</div>
    repaymentList = data.object.repayment_list
    let total = 0

    return (
        <Space direction="vertical" size={16}>
            <Card title={data.object.loan_info.name} extra={<a href="#">More</a>} style={{width: 300}}>
                    {
                        repaymentList.map(item => {
                            total += item.amount
                            return <><p>{item.name}: {item.amount}</p></>
                        })
                    }
                    <p>总计: {total}</p>
            </Card>
        </Space>
    )
}


function LoanInfo() {
    return (
        <QueryClientProvider client={queryClient}>
            <GetCards/>
        </QueryClientProvider>
    )
}


export default LoanInfo;