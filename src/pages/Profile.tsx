import { useState } from 'react';
import { Typography, Card, Tabs, List, Tag, Button, Modal, Form, Input } from 'antd';
import { UserOutlined, LockOutlined, HistoryOutlined, StarOutlined } from '@ant-design/icons';

const { Title, Paragraph } = Typography;
const { TabPane } = Tabs;

interface DivinationRecord {
  id: number;
  type: string;
  question: string;
  result: string;
  ai_analysis?: string;
  created_at: string;
}

interface UserInfo {
  username: string;
  email: string;
}

const Profile = () => {
  const [userInfo, setUserInfo] = useState<UserInfo | null>(null);
  const [records, setRecords] = useState<DivinationRecord[]>([]);
  const [favorites, setFavorites] = useState<DivinationRecord[]>([]);
  const [loading, setLoading] = useState(false);
  const [isModalVisible, setIsModalVisible] = useState(false);

  const fetchUserInfo = async () => {
    try {
      // TODO: 调用后端 API 获取用户信息
      const response = await fetch('/api/user/info');
      const data = await response.json();
      setUserInfo(data);
    } catch (error) {
      console.error('获取用户信息失败:', error);
    }
  };

  const fetchRecords = async () => {
    setLoading(true);
    try {
      // TODO: 调用后端 API 获取占卜记录
      const response = await fetch('/api/user/records');
      const data = await response.json();
      setRecords(data);
    } catch (error) {
      console.error('获取占卜记录失败:', error);
    } finally {
      setLoading(false);
    }
  };

  const fetchFavorites = async () => {
    setLoading(true);
    try {
      // TODO: 调用后端 API 获取收藏记录
      const response = await fetch('/api/user/favorites');
      const data = await response.json();
      setFavorites(data);
    } catch (error) {
      console.error('获取收藏记录失败:', error);
    } finally {
      setLoading(false);
    }
  };

  const onUpdatePassword = async (values: any) => {
    try {
      // TODO: 调用后端 API 更新密码
      await fetch('/api/user/password', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(values),
      });
      setIsModalVisible(false);
    } catch (error) {
      console.error('更新密码失败:', error);
    }
  };

  return (
    <div>
      <Typography>
        <Title level={2}>个人中心</Title>
      </Typography>

      <Tabs defaultActiveKey="info">
        <TabPane
          tab={
            <span>
              <UserOutlined />
              个人信息
            </span>
          }
          key="info"
        >
          <Card>
            {userInfo ? (
              <div>
                <p><strong>用户名：</strong>{userInfo.username}</p>
                <p><strong>邮箱：</strong>{userInfo.email}</p>
                <Button type="primary" onClick={() => setIsModalVisible(true)}>
                  修改密码
                </Button>
              </div>
            ) : (
              <Paragraph>加载中...</Paragraph>
            )}
          </Card>
        </TabPane>

        <TabPane
          tab={
            <span>
              <HistoryOutlined />
              占卜历史
            </span>
          }
          key="history"
        >
          <List
            loading={loading}
            itemLayout="horizontal"
            dataSource={records}
            renderItem={(item) => (
              <List.Item
                actions={[
                  <Button type="link" key="view">查看详情</Button>,
                  <Button type="link" key="favorite">收藏</Button>
                ]}
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
                <div>{item.result}</div>
              </List.Item>
            )}
          />
        </TabPane>

        <TabPane
          tab={
            <span>
              <StarOutlined />
              我的收藏
            </span>
          }
          key="favorites"
        >
          <List
            loading={loading}
            itemLayout="horizontal"
            dataSource={favorites}
            renderItem={(item) => (
              <List.Item
                actions={[
                  <Button type="link" key="view">查看详情</Button>,
                  <Button type="link" danger key="unfavorite">取消收藏</Button>
                ]}
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
                <div>{item.result}</div>
              </List.Item>
            )}
          />
        </TabPane>
      </Tabs>

      <Modal
        title="修改密码"
        visible={isModalVisible}
        onCancel={() => setIsModalVisible(false)}
        footer={null}
      >
        <Form onFinish={onUpdatePassword}>
          <Form.Item
            name="oldPassword"
            rules={[{ required: true, message: '请输入原密码' }]}
          >
            <Input.Password prefix={<LockOutlined />} placeholder="原密码" />
          </Form.Item>
          <Form.Item
            name="newPassword"
            rules={[{ required: true, message: '请输入新密码' }]}
          >
            <Input.Password prefix={<LockOutlined />} placeholder="新密码" />
          </Form.Item>
          <Form.Item
            name="confirmPassword"
            rules={[
              { required: true, message: '请确认新密码' },
              ({ getFieldValue }) => ({
                validator(_, value) {
                  if (!value || getFieldValue('newPassword') === value) {
                    return Promise.resolve();
                  }
                  return Promise.reject(new Error('两次输入的密码不一致'));
                },
              }),
            ]}
          >
            <Input.Password prefix={<LockOutlined />} placeholder="确认新密码" />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit" block>
              确认修改
            </Button>
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
};

export default Profile;