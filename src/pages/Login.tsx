import React, { useState } from 'react';
import { Form, Input, Button, Card, message, Typography, Space } from 'antd';
import { UserOutlined, LockOutlined, MailOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';

const { Title, Paragraph } = Typography;

interface LoginForm {
  username: string;
  password: string;
}

interface RegisterForm extends LoginForm {
  email: string;
  nickname: string;
  confirmPassword: string;
}

const Login: React.FC = () => {
  const [loading, setLoading] = useState(false);
  const [isRegister, setIsRegister] = useState(false);
  const navigate = useNavigate();

  const onFinish = async (values: LoginForm | RegisterForm) => {
    setLoading(true);
    try {
      const endpoint = isRegister ? '/api/user/register' : '/api/user/login';
      const response = await fetch(endpoint, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(values),
      });

      const data = await response.json();
      
      if (!response.ok) {
        throw new Error(data.error || (isRegister ? '注册失败' : '登录失败'));
      }

      if (isRegister) {
        message.success('注册成功，请登录');
        setIsRegister(false);
      } else {
        // 保存token到localStorage
        localStorage.setItem('token', data.token);
        message.success('登录成功');
        // 刷新页面以重新加载App组件，触发路由重定向
        window.location.reload();
      }
    } catch (error) {
      message.error(error instanceof Error ? error.message : (isRegister ? '注册失败，请重试' : '登录失败，请重试'));
    } finally {
      setLoading(false);
    }
  };

  const toggleForm = () => {
    setIsRegister(!isRegister);
  };

  return (
    <div style={{ maxWidth: 400, margin: '40px auto', padding: '0 16px' }}>
      <Card>
        <Typography style={{ textAlign: 'center', marginBottom: 24 }}>
          <Title level={2}>{isRegister ? '用户注册' : '用户登录'}</Title>
          <Paragraph type="secondary">欢迎使用易经占卜系统</Paragraph>
        </Typography>

        <Form
          name={isRegister ? 'register' : 'login'}
          onFinish={onFinish}
          autoComplete="off"
          layout="vertical"
        >
          <Form.Item
            name="username"
            rules={[{ required: true, message: '请输入用户名' }]}
          >
            <Input
              prefix={<UserOutlined />}
              placeholder="用户名"
              size="large"
            />
          </Form.Item>

          {isRegister && (
            <>
              <Form.Item
                name="nickname"
                rules={[{ required: true, message: '请输入昵称' }]}
              >
                <Input
                  prefix={<UserOutlined />}
                  placeholder="昵称"
                  size="large"
                />
              </Form.Item>

              <Form.Item
                name="email"
                rules={[
                  { required: true, message: '请输入邮箱' },
                  { type: 'email', message: '请输入有效的邮箱地址' }
                ]}
              >
                <Input
                  prefix={<MailOutlined />}
                  placeholder="邮箱"
                  size="large"
                />
              </Form.Item>
            </>
          )}

          <Form.Item
            name="password"
            rules={[{ required: true, message: '请输入密码' }]}
          >
            <Input.Password
              prefix={<LockOutlined />}
              placeholder="密码"
              size="large"
            />
          </Form.Item>

          {isRegister && (
            <Form.Item
              name="confirmPassword"
              dependencies={['password']}
              rules={[
                { required: true, message: '请确认密码' },
                ({ getFieldValue }) => ({
                  validator(_, value) {
                    if (!value || getFieldValue('password') === value) {
                      return Promise.resolve();
                    }
                    return Promise.reject(new Error('两次输入的密码不一致'));
                  },
                }),
              ]}
            >
              <Input.Password
                prefix={<LockOutlined />}
                placeholder="确认密码"
                size="large"
              />
            </Form.Item>
          )}

          <Form.Item>
            <Space direction="vertical" style={{ width: '100%' }}>
              <Button
                type="primary"
                htmlType="submit"
                loading={loading}
                block
                size="large"
              >
                {isRegister ? '注册' : '登录'}
              </Button>
              <Button type="link" onClick={toggleForm} block>
                {isRegister ? '已有账号？去登录' : '没有账号？去注册'}
              </Button>
            </Space>
          </Form.Item>
        </Form>
      </Card>
    </div>
  );
};

export default Login;