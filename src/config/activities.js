/**
 * 活动通知与历史配置文件
 * ---------------------------------------------------------------
 * 数据加载优先级：
 *   1. 若配置了 api（非空字符串），则优先 GET 该地址获取 JSON 列表；
 *   2. api 为空或请求失败时，回退使用下方 activities 本地数据。
 *
 * 期望 API 返回以下任意一种结构（数组元素字段见下）：
 *   [ {...}, {...} ]
 *   { "activities": [ ... ] }
 *   { "data": [ ... ] }
 *
 * activities 每个字段：
 *   id       唯一标识，可选，缺省自动生成
 *   title    活动主题
 *   date     时间，如 '2026-09-20 19:00'，用于展示与排序
 *   location 地点
 *   status   'upcoming' 即将开始 / 'past' 往期；缺省时按 date 自动判断
 *   type     活动类型，可选，如「分享会」「训练营」
 *   summary  活动简介，可选
 */

export const activitiesConfig = {
  // 示例：api: 'https://api.9thnet.com/activities'
  api: '',

  section: {
    eyebrow: 'Activities',
    title: '活动通知',
    highlight: '与回顾',
    subtitle: '社团训练、分享会与纳新的时间线，最新动态抢先看。',
    upcomingLabel: '即将开始',
    pastLabel: '往期回顾',
    emptyText: '暂无活动安排，敬请期待。'
  },

  activities: [
    {
      id: 'recruit-2026-fall',
      title: '2026 秋季纳新宣讲会',
      date: '2026-09-20 19:00',
      location: '草堂校区 · 社团活动中心',
      status: 'upcoming',
      type: '纳新',
      summary: '介绍社团方向、服务与训练计划，现场答疑，欢迎新同学到场了解。'
    },
    {
      id: 'ctf-camp-2026-1',
      title: 'CTF 新手训练营（第一期）',
      date: '2026-10-11 14:00',
      location: '逸夫楼 · 机房',
      status: 'upcoming',
      type: '训练营',
      summary: '从环境搭建到 Web / Misc 入门，带你完成第一道 CTF 题目。'
    },
    {
      id: 'network-talk-2026',
      title: '网络与安全技术分享会',
      date: '2026-11-08 15:00',
      location: '图书馆 · 报告厅',
      status: 'upcoming',
      type: '分享会',
      summary: '网络协议、抓包分析与常见安全风险，一次讲清基础与实战。'
    },
    {
      id: 'summary-2026-spring',
      title: '期末技术总结会',
      date: '2026-06-15 19:00',
      location: '草堂校区 · 社团活动中心',
      status: 'past',
      type: '例会',
      summary: '回顾本学期训练与参赛情况，表彰积极成员并规划暑期安排。'
    },
    {
      id: 'network-exchange-2026',
      title: '第十届网络技术交流会',
      date: '2026-04-20 19:00',
      location: '逸夫楼 · 报告厅',
      status: 'past',
      type: '交流会',
      summary: '社团成员分享系统、网络、开发方向的学习心得与实践项目。'
    },
    {
      id: 'tech-festival-2025',
      title: '校园科技文化节参展',
      date: '2025-11-10 10:00',
      location: '草堂校区 · 中心广场',
      status: 'past',
      type: '展览',
      summary: '现场展示社团项目并提供电脑义诊、系统安装等公益服务。'
    },
    {
      id: 'meetup-2025',
      title: '新生见面会',
      date: '2025-09-18 19:30',
      location: '草堂校区 · 社团活动中心',
      status: 'past',
      type: '见面会',
      summary: '新老成员初次见面，介绍社团文化与各兴趣方向。'
    }
  ]
}

export default activitiesConfig
