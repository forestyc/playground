import React from 'react';
import { Card, Space } from 'antd';

const App: React.FC = () => (
    <Space direction="vertical" size={16}>
        <Card title="Period 1" extra={<a href="#">More</a>} style={{ width: 300 }}>
            <p>Card content</p>
            <p>Card content</p>
            {/*<p>Card content</p>*/}
        </Card>
        <Card title="Period 2" extra={<a href="#">More</a>} style={{ width: 300 }}>
            <p>Card content</p>
            <p>Card content</p>
            {/*<p>Card content</p>*/}
        </Card>
    </Space>
);

export default App;