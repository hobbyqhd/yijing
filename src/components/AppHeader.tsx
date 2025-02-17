import { Layout, Menu, Avatar, Dropdown } from 'antd';
import { UserOutlined } from '@ant-design/icons';
import { useNavigate } from 'react-router-dom';

const { Header } = Layout;

const AppHeader = () => {
  const navigate = useNavigate();

  const userMenu = [
    {
      key: 'profile',
      label: '个人中心',
      onClick: () => navigate('/profile')
    },
    {
      key: 'logout',
      label: '退出登录',
      onClick: () => {
        // 先导航到登录页面
        navigate('/login');
        // 清除所有用户相关的本地存储数据
        localStorage.clear();
        // 刷新页面以确保所有状态被重置
        window.location.reload();
      }
    }
  ];

  return (
    <Header style={{ background: '#fff', padding: '0 24px', display: 'flex', justifyContent: 'flex-end', alignItems: 'center' }}>
      <Dropdown menu={{ items: userMenu }} placement="bottomRight">
        <Avatar icon={<UserOutlined />} style={{ cursor: 'pointer' }} />
      </Dropdown>
    </Header>
  );
};

export default AppHeader;