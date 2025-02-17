import { useState } from 'react';
import { Typography, Form, Input, Select, Button, Card, Spin } from 'antd';
import { CompassOutlined } from '@ant-design/icons';

const { Title, Paragraph } = Typography;
const { TextArea } = Input;

const divinationTypes = [
  { value: 'general', label: '综合运势' },
  { value: 'career', label: '事业运势' },
  { value: 'love', label: '感情运势' },
  { value: 'wealth', label: '财运预测' },
  { value: 'health', label: '健康预测' }
];

const Divination = () => {
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState<any>(null);

  const onFinish = async (values: any) => {
    setLoading(true);
    try {
      const response = await fetch('/api/divination', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(values),
      });
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      
      const data = await response.json();
      if (data.error) {
        throw new Error(data.error);
      }
      
      setResult({
        ...data,
        hexagram: data.hexagram || '未知卦象',
        interpretation: data.interpretation || '暂无解释',
        result: data.result || '暂无结果',
        ai_analysis: data.ai_analysis || '暂无AI分析'
      });
    } catch (error) {
      console.error('占卜请求失败:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      <Typography>
        <Title level={2}>易经占卜</Title>
        <Paragraph>
          请选择占卜类型并输入您的问题，我们将根据易经原理为您进行解析。
        </Paragraph>
      </Typography>

      <Card style={{ marginTop: '24px' }}>
        <Form layout="vertical" onFinish={onFinish}>
          <Form.Item
            name="type"
            label="占卜类型"
            rules={[{ required: true, message: '请选择占卜类型' }]}
          >
            <Select options={divinationTypes} placeholder="请选择占卜类型" />
          </Form.Item>

          <Form.Item
            name="question"
            label="问题描述"
            rules={[{ required: true, message: '请输入您的问题' }]}
          >
            <TextArea
              rows={4}
              placeholder="请详细描述您的问题，以便我们为您提供更准确的解析"
            />
          </Form.Item>

          <Form.Item>
            <Button
              type="primary"
              htmlType="submit"
              icon={<CompassOutlined />}
              loading={loading}
              block
            >
              开始占卜
            </Button>
          </Form.Item>
        </Form>
      </Card>

      {result && (
        <Card
          title="占卜结果"
          style={{ marginTop: '24px' }}
        >
          <div style={{ textAlign: 'center', marginBottom: '24px' }}>
            <Title level={3}>{result.hexagram}</Title>
          </div>
          
          <div style={{ marginBottom: '24px' }}>
            <Title level={4}>卦象解读</Title>
            <Paragraph>{result.interpretation}</Paragraph>
          </div>

          <div style={{ marginBottom: '24px' }}>
            <Title level={4}>占卜结果</Title>
            <Paragraph>{result.result}</Paragraph>
          </div>

          <div>
            <Title level={4}>AI分析洞见</Title>
            <Paragraph>{result.ai_analysis}</Paragraph>
          </div>
        </Card>
      )}
    </div>
  );
};

export default Divination;