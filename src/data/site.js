export const brand = {
  name: '第九网络组',
  enName: 'THE 9TH NETWORK',
  fullEnName: 'THE 9TH NETWORK TEAM',
  slogan: '资源共享 · 共学习 · 共提高 · 共进步',
  school: '西安建筑科技大学 · 子午社联第九网络组',
  logo: '/images/t9.gif'
}

export const navLinks = [
  { id: 'home', label: '首页', kind: 'section' },
  { id: 'about', label: '社团', kind: 'section' },
  { id: 'services', label: '服务', kind: 'section' },
  { id: 'directions', label: '方向', kind: 'route', to: '/directions' },
  { id: 'resources', label: '资源', kind: 'section' },
  { id: 'join', label: '加入我们', kind: 'section' }
]

export const footerLinks = [
  {
    title: '社团',
    items: [
      { label: '社团简介', to: 'about' },
      { label: '兴趣方向', to: '/directions' },
      { label: '社团成员', to: 'about' },
      { label: '加入我们', to: 'join' }
    ]
  },
  {
    title: '服务',
    items: [
      { label: '系统安装', to: 'services' },
      { label: '网络诊断', to: 'services' },
      { label: '网站开发', to: 'services' },
      { label: '技术培训', to: 'services' }
    ]
  },
  {
    title: '资源',
    items: [
      { label: '系统教学', to: 'resources' },
      { label: '网络安全', to: 'resources' },
      { label: '建大热点', to: 'resources' },
      { label: '社团动态', to: 'resources' }
    ]
  }
]

export const heroSlides = [
  {
    image: '/images/1.jpg',
    tag: '西安建筑科技大学',
    title: '第九网络组',
    subtitle: '资源共享 · 共学习 · 共提高 · 共进步',
    desc: '一群热爱计算机科学、刻苦专研的大学生聚集地。'
  },
  {
    image: '/images/2.jpg',
    tag: '启明于今天',
    title: '相信这是新的开始',
    subtitle: '让想法付诸行动',
    desc: '现在我们是计算机的新主人，未来我们就是 IT 的精英。'
  },
  {
    image: '/images/3.jpg',
    tag: '技术与热爱',
    title: '与时代同频',
    subtitle: '我要拼搏 · 我要进取 · 我要与时俱进',
    desc: '在系统、网络、开发与安全中，找到属于你的方向。'
  }
]

export const stats = [
  { key: 'directions', value: '6', suffix: '大', label: '技术方向' },
  { key: 'services', value: '12', suffix: '项', label: '对外服务' },
  { key: 'members', value: '60', suffix: '+', label: '社团成员' },
  { key: 'served', value: '800', suffix: '+', label: '累计服务人次' }
]

export const aboutTabs = [
  {
    key: 'intro',
    label: '社团简介',
    images: ['/images/t91.jpg'],
    paragraphs: [
      '我们是一群热爱计算机科学，有着刻苦专研精神的大学生。我们社团的宗旨是「资源共享，共学习，共提高，共进步」。我们的口号是：现在我们是计算机的新主人，未来我们就是 IT 的精英！',
      '社团旨在同学中间普及计算机知识、提高计算机技术，提供各种计算机技术辅导，立志于为电脑爱好者提供便利的学习机会和条件，使同学们的计算机水平得到提高。',
      '社团以增强广大同学的科技意识、激发计算机兴趣、全面提高电脑知识水平和应用能力为目标，重点发展同学在基本操作、网络、图形处理、硬件维护和编程方面的技能，并发掘和发展电脑人才，积极回报服务社会。'
    ]
  },
  {
    key: 'members',
    label: '社团成员',
    images: ['/images/t94.jpg', '/images/t95.jpg'],
    groups: [
      { role: '社长', names: ['张浩'] },
      { role: '副社长', names: [] },
      { role: '顾问团', names: ['童凯', '贺均威'] },
      { role: '社团骨干', names: ['李欣', '郭通', '张景涛', '张婵'] },
      { role: '社团成员', names: ['王兆泉', '杨军', '翟亮', '薛亚鹏', '叶闻晶', '拜晓萌', '等等多位同学'] }
    ]
  },
  {
    key: 'development',
    label: '社团发展',
    images: ['/images/t92.jpg'],
    items: [
      '组织会员学习计算机知识，交流并不断提高社员的理论和技术水平。',
      '不定期组织会员进行各项计算机知识竞赛或其他相关活动。',
      '适时举办最新技术讲座及咨询活动。',
      '建立本社团的网站、论坛、QQ 群、微信平台，以便进行资料信息提供及知识交流。',
      '进行各种电脑实践，帮助同学解决电脑日常使用中的疑难问题、网页制作等。',
      '组织社员共同学习软件、开发软件，加强分工协作的精神。',
      '沟通社员与学校和社会相关部门的联系，扩大社员的社交范围和社会影响。'
    ]
  }
]

export const joinInfo = {
  eyebrow: 'Join Us',
  title: '加入',
  highlight: '我们',
  subtitle: '第九网络组每年在统一的迎新活动中招新，无需在线填写信息。',
  notice: {
    title: '统一迎新 · 社团处现场报名',
    desc: '每年开学季，社团会参加学校统一的迎新纳新活动。届时到第九网络组社团摊位现场报名即可，没有线上表单，也不用提前填写个人信息。'
  },
  steps: [
    { title: '关注迎新时间', desc: '留意学校统一迎新活动通知与社团招新安排。' },
    { title: '来到社团处', desc: '在迎新活动现场找到第九网络组的社团摊位。' },
    { title: '现场报名', desc: '登记基本信息，选择自己感兴趣的方向。' },
    { title: '参加见面会', desc: '参加迎新见面会与技术交流，正式开启社团生活。' }
  ],
  note: '如有疑问，可在迎新现场向社团工作人员咨询。',
  cta: { label: '先看看兴趣方向', to: '/directions' }
}

export const friendLinks = [
  { label: '西安建筑科技大学', url: 'https://www.xauat.edu.cn/' },
  { label: '中国教育和科研计算机网', url: 'https://www.cernet.edu.cn/' }
]

export const icp = '陕ICP备XXXXXXXX号'
