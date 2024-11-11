'use client'

import React from 'react';
// @ts-ignore
import {List, Divider} from 'antd';
import {QueryClientProvider, useQuery} from "@tanstack/react-query";
import {getLoanList, queryClient} from "@/app/loan/api";


type Item = {
    id: number;
    name: string;
    created_at: Date;
    updated_at: Date;
}

function GetList() {
    let rsp: Item[]
    const {data, error, isLoading} = useQuery({queryKey: ['GetList'], queryFn: getLoanList})
    if (isLoading) return <div>Loading...</div>
    if (error) return <div>Error: {error.message}</div>
    rsp = data.object

    return (
        <>
            <Divider orientation="left">Default Size</Divider>
            <List
                itemLayout="horizontal"
                dataSource={rsp}
                bordered
                renderItem={(item) =>
                    <List.Item>
                        <List.Item.Meta
                            title={<a href={"/loan/cards?id=" + item.id}>{item.name}</a>}
                            // description="Ant Design, a design language for background applications, is refined by Ant UED Team"
                            // style={}
                        />
                    </List.Item>}
            />
        </>
    )
}


function LoanList() {
    return (
        <QueryClientProvider client={queryClient}>
            <GetList/>
        </QueryClientProvider>
    )
}

export default LoanList;