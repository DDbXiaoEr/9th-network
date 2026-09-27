CentOS 适合作为服务器与实验环境。本文按「准备镜像 → 制作启动盘 → 分区安装 → 首次配置」走完一遍基础安装。

## 准备材料

- 一张 8GB 以上的 U 盘
- CentOS 官方 ISO（建议 7 / Stream 对应版本）
- 目标电脑支持从 USB 启动

## 制作启动盘

Windows 下可用 Rufus：选择 U 盘、选择 ISO、分区方案选 GPT（UEFI）或 MBR（传统 BIOS），开始写入。macOS / Linux 可用 `dd`：

```bash
sudo dd if=CentOS.iso of=/dev/sdX bs=4M status=progress
```

写入前务必确认 `sdX` 是 U 盘设备，写错会清空硬盘。

## 安装要点

1. 开机进入 BIOS / 启动菜单，选 U 盘启动。
2. 选择 Install CentOS，语言选中文或 English。
3. 安装目标磁盘：服务器建议单独划分 `/`、`/home`、`swap`。
4. 网络与主机名：校园网环境可先跳过，装完再配。
5. 设置 root 密码，并至少创建一个普通用户。

## 装完后建议

- 先更新：`sudo yum update` 或 `sudo dnf update`
- 关闭不需要的服务，只开放必要端口
- 用 SSH 公钥登录，尽量少用密码直接登 root

安装过程卡住时，优先检查启动模式（UEFI / Legacy）是否与启动盘分区方案一致。
