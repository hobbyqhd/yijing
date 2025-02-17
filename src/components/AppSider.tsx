import { Layout, Menu } from 'antd';
import { HomeOutlined, CompassOutlined, LineChartOutlined, UserOutlined } from '@ant-design/icons';
import { useNavigate, useLocation } from 'react-router-dom';

const { Sider } = Layout;

const AppSider = () => {
  const navigate = useNavigate();
  const location = useLocation();

  const menuItems = [
    {
      key: '/',
      icon: <HomeOutlined />,
      label: '首页'
    },
    {
      key: '/divination',
      icon: <CompassOutlined />,
      label: '占卜'
    },
    {
      key: '/fortune',
      icon: <LineChartOutlined />,
      label: '运势分析'
    },
    {
      key: '/profile',
      icon: <UserOutlined />,
      label: '个人中心'
    }
  ];

  return (
    <Sider theme="light" style={{ boxShadow: '2px 0 8px 0 rgba(29,35,41,.05)' }}>
      <div style={{ height: '64px', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
        <h1 style={{ margin: 0, fontSize: '20px' }}>易经占卜</h1>
      </div>
      <Menu
        mode="inline"
        selectedKeys={[location.pathname]}
        items={menuItems}
        onClick={({ key }) => navigate(key)}
      />
    </Sider>
  );
};

export default AppSider;