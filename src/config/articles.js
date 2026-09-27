/**
 * 学习资源默认配置（兜底）
 * ---------------------------------------------------------------
 * 运行时优先加载 /config/articles.json；失败时回退本文件。
 * 正文 Markdown 从 /articles/*.md 按 file 字段加载。
 *
 *   groups    - 资源分组（标题、图标、主题色、链接标题列表）
 *   articles  - 文章元数据，每项字段：
 *       id / category / title / author / summary / cover / tags[] / file
 */

export const articlesConfig = {
  groups: [
    {
      title: '系统教学',
      icon: '/images/1.png',
      accent: 'cyan',
      links: [
        'Windows 系统安装教程',
        'Linux CentOS 系统安装教程',
        'U 盘启动盘制作',
        '系统激活与驱动安装'
      ]
    },
    {
      title: '网络安全',
      icon: '/images/2.png',
      accent: 'violet',
      links: ['不安全的 WiFi', '社会工程学', '密码安全指南', '常见木马识别']
    },
    {
      title: '建大热点',
      icon: '/images/3.png',
      accent: 'blue',
      links: ['思科杯虚拟网络架构大赛', '建信歌手大赛开赛', '校园科技文化节', '创新创业大赛']
    },
    {
      title: '社团动态',
      icon: '/images/4.png',
      accent: 'amber',
      links: ['网络技术交流会', '新生见面会', '技术讲座预告', '社团纳新进行时']
    }
  ],

  articles: [
    {
      id: 'windows-install',
      category: '系统教学',
      title: 'Windows 系统安装教程',
      author: '董威',
      summary: '通过 PE 工具箱制作启动 U 盘，完成磁盘分区、系统安装、驱动安装与激活的完整流程。',
      cover: '/images/t91.jpg',
      tags: ['U盘启动', 'Windows', '教程'],
      file: '/articles/windows-install.md'
    },
    {
      id: 'unsafe-wifi',
      category: '网络安全',
      title: '不安全的 WiFi',
      author: '贺钧威',
      summary: '从 WPS PIN 码到抓包跑字典，看清公共 WiFi 背后的风险，提高网络安全意识。',
      cover: '/images/t92.jpg',
      tags: ['WiFi', '安全', '意识'],
      file: '/articles/unsafe-wifi.md'
    },
    {
      id: 'social-engineering',
      category: '网络安全',
      title: '社会工程学简介',
      author: '贺钧威',
      summary: '人类思维也可以被看作运行中的软件，社会工程学正是利用人性漏洞的一种攻击方式。',
      cover: '/images/t93.jpg',
      tags: ['社会工程学', '诈骗', '防护'],
      file: '/articles/social-engineering.md'
    }
  ]
}

export default articlesConfig
