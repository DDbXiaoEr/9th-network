/**
 * 社团兴趣方向配置文件
 * ---------------------------------------------------------------
 * 本文件配置「兴趣方向」页面的全部内容，修改后保存即热更新，
 * 发布前执行 `npm run build`。
 *
 *   section    - 页面标题区文案
 *   directions - 兴趣方向列表，每项字段：
 *       id / icon / name / tagline / desc / highlights[] / stack[]
 *   ctf        - CTF 战队板块（可整体关闭：enabled: false）
 *       bestRank   历史最好成绩展示
 *       team       战队信息（键值对展示）
 *       members    队员列表
 *       competitions 比赛记录与成绩
 *
 * 说明：战队名称、队员与比赛记录来自社团成员的真实参赛经历（CTFTime）。
 */

export const directionsConfig = {
  section: {
    eyebrow: 'Interest Directions',
    title: '兴趣',
    highlight: '方向',
    subtitle:
      '第九网络组设网络技术与安全技术、编程技术两大兴趣方向，以兴趣聚人、以赛促学，不定期组织训练与参赛。'
  },

  directions: [
    {
      id: 'security',
      icon: 'shield',
      name: '网络技术与安全技术',
      tagline: '攻防兼备，以赛促学',
      desc: '从网络协议到渗透测试、逆向与密码学，系统学习安全技术，并不定期带领同学们一起打 CTF、组队参赛。',
      highlights: [
        '网络协议原理与抓包分析',
        'Web 安全与渗透测试',
        '逆向工程与 Pwn',
        '密码学与杂项（Misc）',
        'CTF 组队训练与实战参赛'
      ],
      stack: ['Wireshark', 'Burp Suite', 'IDA', 'Linux', 'Python', 'CTF']
    },
    {
      id: 'programming',
      icon: 'code',
      name: '编程技术',
      tagline: '从第一行代码到完整项目',
      desc: '面向 C / C++ / Python / 前端与算法，组织刷题训练、项目实战与竞赛组队，把想法变成能跑起来的作品。',
      highlights: [
        '算法与数据结构训练',
        'Web / 小程序项目实战',
        'Git 与开源协作',
        '竞赛组队（ACM / 蓝桥杯）',
        '代码规范与工程化实践'
      ],
      stack: ['C / C++', 'Python', 'Java', 'JavaScript', 'Vue', 'Git']
    }
  ],

  ctf: {
    enabled: true,
    eyebrow: 'CTF Team',
    title: 'CTF 战队',
    subtitle: '以赛代练，在攻防对抗中成长。社团不定期组织 CTF 训练营并带队参赛。',

    bestRank: {
      label: '历史最好成绩',
      value: '92',
      unit: '名',
      title: 'CTFTime 全国第 92 名（待过，没住下）',
      note: '第 92 名体验卡已到期，没续费 · 2022 年度 · 世界第 4463 名'
    },

    team: {
      name: 'the9thnet',
      fullName: '第九网络组 CTF 战队',
      slogan: 'Stay curious, stay secure.',
      url: 'https://ctftime.org/team/166528',
      fields: [
        { label: '战队名称', value: 'the9thnet' },
        { label: '所属社团', value: '第九网络组' },
        { label: '主攻方向', value: 'Web · Misc · Crypto · Reverse · Pwn' },
        { label: '常用平台', value: 'CTFTime' },
        { label: '训练频率', value: '每周一次 + 赛前集训' },
        { label: '加入方式', value: '参加统一迎新报名' }
      ]
    },

    members: [
      { name: '王兆泉', handle: '@V0ker', role: '队长', direction: '社工 · Reverse · Crypto' }
    ],

    competitions: [
      {
        date: '2021',
        name: 'SECCON CTF 2021',
        type: '国际赛',
        result: '第 326 名',
        note: 'CTF 53.0000 · Rating 1.069'
      },
      {
        date: '2021',
        name: 'MetaCTF CyberGames 2021',
        type: '国际赛',
        result: '第 474 名',
        note: 'CTF 3050.0000 · Rating 4.925'
      },
      {
        date: '2021',
        name: 'N1CTF 2021',
        type: '国际赛',
        result: '第 182 名',
        note: 'CTF 68.0000 · Rating 0.663'
      },
      {
        date: '2021',
        name: 'DamCTF 2021',
        type: '国际赛',
        result: '第 326 名',
        note: 'CTF 772.0000 · Rating 2.391'
      },
      {
        date: '2021',
        name: 'ASIS CTF Quals 2021',
        type: '国际赛',
        result: '第 289 名',
        note: 'CTF 51.0000 · Rating 1.760'
      },
      {
        date: '2021',
        name: 'DEADFACE CTF',
        type: '国际赛',
        result: '第 451 名',
        note: 'CTF 725.0000 · Rating 1.687'
      },
      {
        date: '2023',
        name: 'TSG CTF 2023',
        type: '国际赛',
        result: '第 292 名',
        note: 'CTF 100.0000 · Rating 0.642'
      },
      {
        date: '2023',
        name: 'DEADFACE CTF 2023',
        type: '国际赛',
        result: '第 286 名',
        note: 'CTF 965.0000 · Rating 3.253'
      }
    ]
  }
}

export default directionsConfig
