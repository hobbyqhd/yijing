import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { Layout } from 'antd';
import AppHeader from './components/AppHeader';
import AppSider from './components/AppSider';
import Home from './pages/Home';
import Divination from './pages/Divination';
import Fortune from './pages/Fortune';
import Profile from './pages/Profile';
import Login from './pages/Login';
import './App.css';

const { Content } = Layout;

function App() {
  const isAuthenticated = !!localStorage.getItem('token');

  return (
    <Router>
      {isAuthenticated ? (
        <Layout style={{ minHeight: '100vh' }}>
          <AppSider />
          <Layout>
            <AppHeader />
            <Content style={{ margin: '24px 16px', padding: 24, background: '#fff' }}>
              <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/divination" element={<Divination />} />
                <Route path="/fortune" element={<Fortune />} />
                <Route path="/profile" element={<Profile />} />
                <Route path="/login" element={<Navigate to="/" replace />} />
                <Route path="/register" element={<Navigate to="/" replace />} />
                <Route path="*" element={<Navigate to="/" replace />} />
              </Routes>
            </Content>
          </Layout>
        </Layout>
      ) : (
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="*" element={<Navigate to="/login" replace />} />
        </Routes>
      )}
    </Router>
  );
}

export default App;