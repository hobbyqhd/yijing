import { Typography, Card, Row, Col, List, Tag } from 'antd';
import { CompassOutlined, LineChartOutlined, UserOutlined, RightOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';

const { Title, Paragraph } = Typography;

const Home = () => {
  const navigate = useNavigate();

  const features = [
    {
      icon: <CompassOutlined style={{ fontSize: '32px' }} />,
      title: '易经占卜',
      description: '基于易经原理的智能占卜系统，为您提供准确的预测和指导。',
      path: '/divination'
    },
    {
      icon: <LineChartOutlined style={{ fontSize: '32px' }} />,
      title: '运势分析',
      description: '全面的运势分析，包括事业、感情、财运等多个维度的预测。',
      path: '/fortune'
    },
    {
      icon: <UserOutlined style={{ fontSize: '32px' }} />,
      title: '个性化服务',
      description: '根据您的个人信息和历史记录，提供更准确的占卜和分析服务。',
      path: '/profile'
    }
  ];

  const recentDivinations = [
    {
      type: '易经占卜',
      question: '近期事业发展如何？',
      created_at: '2024-01-20'
    },
    {
      type: '塔罗牌',
      question: '感情运势分析',
      created_at: '2024-01-19'
    },
    {
      type: '八字分析',
      question: '财运预测',
      created_at: '2024-01-18'
    }
  ];

  return (
    <div>
      <Typography>
        <Title level={2}>欢迎使用易经占卜系统</Title>
        <Paragraph>
          易经占卜系统是一个基于古老易经智慧，结合现代技术的智能预测平台。
          我们致力于为您提供准确、专业的占卜服务和运势分析。
        </Paragraph>
      </Typography>

      <Row gutter={[24, 24]} style={{ marginTop: '32px' }}>
        {features.map((feature, index) => (
          <Col key={index} xs={24} sm={12} md={8}>
            <Card 
              hoverable 
              onClick={() => navigate(feature.path)}
              style={{ cursor: 'pointer' }}
            >
              <div style={{ textAlign: 'center', marginBottom: '16px' }}>
                {feature.icon}
              </div>
              <Title level={4} style={{ textAlign: 'center' }}>{feature.title}</Title>
              <Paragraph style={{ textAlign: 'center' }}>{feature.description}</Paragraph>
            </Card>
          </Col>
        ))}
      </Row>

      <div style={{ marginTop: '48px' }}>
        <Title level={3}>最新占卜动态</Title>
        <Card>
          <List
            itemLayout="horizontal"
            dataSource={recentDivinations}
            renderItem={(item) => (
              <List.Item
                actions={[<RightOutlined />]}
                style={{ cursor: 'pointer' }}
                onClick={() => navigate('/divination')}
              >
                <List.Item.Meta
                  title={item.question}
                  description={
                    <div>
                      <Tag color="blue">{item.type}</Tag>
                      <span style={{ marginLeft: '8px' }}>{item.created_at}</span>
                    </div>
                  }
                />
              </List.Item>
            )}
          />
        </Card>
      </div>
    </div>
  );
};

export default Home;