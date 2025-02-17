import { useState, useEffect } from 'react';
import { Typography, Card, DatePicker, Button, Table, Tag, Row, Col, Statistic } from 'antd';
import { Line } from '@ant-design/plots';
import type { Dayjs } from 'dayjs';
import { LineChartOutlined } from '@ant-design/icons';

const { Title, Paragraph } = Typography;
const { RangePicker } = DatePicker;

interface FortuneRecord {
  id: number;
  periodType: string;
  startDate: string;
  endDate: string;
  content: string;
  createdAt: string;
}

const Fortune = () => {
  const [loading, setLoading] = useState(false);
  const [records, setRecords] = useState<FortuneRecord[]>([]);
  const [trendData, setTrendData] = useState<any[]>([]);
  const [statistics, setStatistics] = useState<any>({
    overall: 0,
    career: 0,
    love: 0,
    wealth: 0,
    health: 0
  });

  useEffect(() => {
    fetchFortuneRecords();
  }, []);

  const config = {
    data: trendData,
    xField: 'date',
    yField: 'value',
    seriesField: 'type',
    smooth: true,
    animation: {
      appear: {
        animation: 'path-in',
        duration: 1000,
      },
    },
  };

  const onDateRangeChange = async (dates: [Dayjs, Dayjs] | null) => {
    if (!dates) return;
    setLoading(true);
    try {
      const response = await fetch('/api/fortune/analyze', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          startDate: dates[0].format('YYYY-MM-DD'),
          endDate: dates[1].format('YYYY-MM-DD')
        }),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      if (data.error) {
        throw new Error(data.error);
      }

      setTrendData(data.trends || []);
      setStatistics(data.statistics || {
        overall: 0,
        career: 0,
        love: 0,
        wealth: 0,
        health: 0
      });
      await fetchFortuneRecords();
    } catch (error) {
      console.error('运势分析请求失败:', error);
    } finally {
      setLoading(false);
    }
  };

  const columns = [
    {
      title: '周期类型',
      dataIndex: 'periodType',
      key: 'periodType',
      render: (type: string) => (
        <Tag color={type === '日' ? 'blue' : type === '周' ? 'green' : 'purple'}>
          {type}运势
        </Tag>
      ),
    },
    {
      title: '起始日期',
      dataIndex: 'startDate',
      key: 'startDate',
    },
    {
      title: '结束日期',
      dataIndex: 'endDate',
      key: 'endDate',
    },
    {
      title: '分析内容',
      dataIndex: 'content',
      key: 'content',
      ellipsis: true,
    },
    {
      title: '创建时间',
      dataIndex: 'createdAt',
      key: 'createdAt',
    },
  ];

  const fetchFortuneRecords = async () => {
    setLoading(true);
    try {
      const response = await fetch('/api/fortune/records');
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      
      const data = await response.json();
      if (data.error) {
        throw new Error(data.error);
      }
      
      setRecords(data.map((record: any) => ({
        ...record,
        key: record.id
      })));
    } catch (error) {
      console.error('获取运势记录失败:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      <Typography>
        <Title level={2}>运势分析</Title>
        <Paragraph>
          选择日期范围，查看您的运势分析结果。我们提供日、周、月、年等多个维度的运势预测。
        </Paragraph>
      </Typography>

      <Row gutter={[16, 16]}>
        <Col span={24}>
          <Card title="运势趋势" style={{ marginBottom: '24px' }}>
            {trendData.length > 0 ? (
              <Line {...config} />
            ) : (
              <div style={{ textAlign: 'center', padding: '20px' }}>暂无趋势数据</div>
            )}
          </Card>
        </Col>

        <Col span={24}>
          <Card title="运势指数" style={{ marginBottom: '24px' }}>
            <Row gutter={[16, 16]}>
              <Col span={4}>
                <Statistic title="综合运势" value={statistics.overall} suffix="/100" />
              </Col>
              <Col span={4}>
                <Statistic title="事业运" value={statistics.career} suffix="/100" />
              </Col>
              <Col span={4}>
                <Statistic title="感情运" value={statistics.love} suffix="/100" />
              </Col>
              <Col span={4}>
                <Statistic title="财运" value={statistics.wealth} suffix="/100" />
              </Col>
              <Col span={4}>
                <Statistic title="健康运" value={statistics.health} suffix="/100" />
              </Col>
            </Row>
          </Card>
        </Col>

        <Col span={24}>
          <Card style={{ marginTop: '24px' }}>
            <div style={{ marginBottom: '24px' }}>
              <RangePicker onChange={onDateRangeChange} style={{ marginRight: '16px' }} />
              <Button
                type="primary"
                icon={<LineChartOutlined />}
                onClick={fetchFortuneRecords}
                loading={loading}
              >
                分析运势
              </Button>
            </div>

            <Table
              columns={columns}
              dataSource={records}
              rowKey="id"
              loading={loading}
              pagination={{
                defaultPageSize: 10,
                showSizeChanger: true,
                showTotal: (total) => `共 ${total} 条记录`,
              }}
            />
          </Card>
        </Col>
      </Row>
    </div>
  );
};

export default Fortune;