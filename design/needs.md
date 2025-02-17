# 占卜功能网站系统设计方案

## 1. 系统概述

本系统是一个现代化的占卜功能网站，集成多种占卜方式，并结合 AI 技术提供智能化的占卜服务。系统采用微服务架构，使用 Go 语言开发后端服务，MySQL 作为主要数据存储，支持多平台访问。

## 2. 系统架构

### 2.1 整体架构

- **前端层**：
  - Web 前端（React.js + TypeScript）
  - 移动端（React Native）
  - 小程序（统一接口适配）

- **后端层**：
  - API 网关服务（Go）
  - 用户服务（Go）
  - 占卜核心服务（Go）
  - AI 服务适配器（Go）
  - 运势分析服务（Go）

- **数据层**：
  - MySQL：核心业务数据
  - Redis：缓存层
  - Elasticsearch：搜索服务

- **基础设施**：
  - Kubernetes：容器编排
  - Docker：容器化部署
  - Nginx：负载均衡
  - RabbitMQ：消息队列

### 2.2 技术栈选型

- **后端**：
  - 主要语言：Go 1.20+
  - Web 框架：Gin
  - ORM：GORM
  - 缓存：Redis
  - 消息队列：RabbitMQ

- **数据库**：
  - 主数据库：MySQL 8.0+
  - 分库分表：ShardingSphere

- **前端**：
  - Web：React.js + TypeScript
  - 移动端：React Native
  - UI 框架：Ant Design

## 3. 核心功能模块

### 3.1 占卜功能模块

1. **多种占卜方式**：
   - 星座占卜
   - 塔罗牌占卜
   - 易经卦象
   - 八字分析

2. **AI 智能解读**：
   - 接入 OpenAI API
   - 自定义提示词模板
   - 结果智能分析

3. **运势分析**：
   - 周运预测
   - 月运分析
   - 年度运势

### 3.2 用户系统

1. **用户管理**：
   - 账号注册/登录
   - 身份认证
   - 用户画像

2. **个人中心**：
   - 历史记录
   - 收藏管理
   - 个性化设置

## 4. 数据库设计

### 4.1 核心表结构

```sql
-- 用户表
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_email` (`email`)
);

-- 占卜记录表
CREATE TABLE `divination_records` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `type` varchar(20) NOT NULL,
  `question` text NOT NULL,
  `result` text NOT NULL,
  `ai_analysis` text,
  `created_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
);

-- 运势分析表
CREATE TABLE `fortune_analysis` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `period_type` varchar(10) NOT NULL,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `content` text NOT NULL,
  `created_at` timestamp NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_period` (`user_id`, `period_type`, `start_date`)
);
```

## 5. API 接口设计

### 5.1 RESTful API

1. **用户相关**：
```
POST /api/v1/users/register     # 用户注册
POST /api/v1/users/login        # 用户登录
GET  /api/v1/users/profile      # 获取用户信息
```

2. **占卜相关**：
```
POST /api/v1/divination/start   # 开始占卜
GET  /api/v1/divination/{id}    # 获取占卜结果
GET  /api/v1/divination/history # 获取历史记录
```

3. **运势分析**：
```
GET  /api/v1/fortune/weekly     # 获取周运
GET  /api/v1/fortune/monthly    # 获取月运
POST /api/v1/fortune/custom     # 自定义时间段运势
```

## 6. AI 集成方案

### 6.1 OpenAI API 集成

1. **提示词模板**：
   - 为不同占卜类型设计专属提示词
   - 结合用户输入动态生成完整提示
   - 支持多轮对话优化结果

2. **结果处理**：
   - 响应解析和格式化
   - 结果缓存策略
   - 失败重试机制

### 6.2 AI 服务优化

1. **性能优化**：
   - API 调用并发控制
   - 结果缓存机制
   - 异步处理长时间运算

2. **成本控制**：
   - Token 使用优化
   - 缓存策略优化
   - 调用频率控制

## 7. 多平台适配

### 7.1 响应式设计

- 采用移动优先的设计理念
- 使用响应式布局适配不同设备
- 统一的用户体验设计

### 7.2 多端适配策略

1. **Web 端**：
   - PWA 支持
   - 响应式布局
   - 浏览器兼容性处理

2. **移动端**：
   - React Native 跨平台方案
   - 原生功能集成
   - 性能优化

3. **小程序**：
   - 统一后端接口
   - 平台特性适配
   - 分包加载优化

## 8. 系统安全

### 8.1 安全措施

1. **用户认证**：
   - JWT 认证
   - OAuth2.0 集成
   - Session 管理

2. **数据安全**：
   - 数据加密存储
   - 传输加密（HTTPS）
   - 敏感信息脱敏

3. **接口安全**：
   - 接口限流
   - CORS 配置
   - 参数验证

## 9. 监控和运维

### 9.1 监控体系

1. **系统监控**：
   - 服务器监控
   - 容器监控
   - 数据库监控

2. **业务监控**：
   - 用户行为分析
   - 接口调用统计
   - 错误日志收集

### 9.2 运维支持

1. **部署策略**：
   - 蓝绿部署
   - 灰度发布
   - 回滚机制

2. **日志管理**：
   - ELK 日志收集
   - 日志分析
   - 告警机制

## 10. 扩展性考虑

### 10.1 系统扩展

1. **服务扩展**：
   - 微服务架构支持
   - 服务自动扩缩容
   - 新功能快速集成

2. **数据扩展**：
   - 分库分表方案
   - 读写分离
   - 数据归档

### 10.2 业务扩展

1. **新占卜方式接入**：
   - 插件化架构
   - 标准化接口
   - 配置化管理

2. **多语言支持**：
   - i18n 框架集成
   - 多语言内容管理
   - 动态语言切换