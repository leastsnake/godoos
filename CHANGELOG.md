## 变更记录
- 2024-08-21
内网聊天增加手工添加ip，跨网段通信在ping通的前提下如果发现不了对方可手工添加对方ip
修复思维导图保存的文件每次打开主题又会变成默认主题
新增webdav客户端
新增远程存储
修改选择文件夹会删除文件夹内的文件

- 2024-08-07
1. 新增web端安装
- 2024-08-06
1. 新增docker安装
- 2024-08-04
1. 修复进程管理在windows环境下不停闪窗
- 2024-08-02 
1. 修复桌面更换背景文字颜色问题
2. 修复安装插件没有配置binPath的问题
3. 新增linux环境下获取安装命令的方法

bug梳理：
1. 思维导图保存的文件每次打开主题又会变成默认主题
2. 公网版部署udp转发和文件上传存储等
3. 内网聊天功能，如果是装了虚拟机的电脑，会出现多个虚拟网卡，内网可以看到用户但是IP不对着，发消息收不到的
4. 建议开启webDav
5. web 的甘特图保存不了
6. 场景适用性：系统默认连接gitee远程软件安装包，缺少默认读取本地安装包路径的json配置文件的打包设置选项（比如在安装时自动读取安装U盘当前目录中/或已直接打包进安装包中的mysql/nginx/php/python本地离线安装包、以及企业自用OA WebApp的本地离线安装包，无需联网，直接同步离线安装）
7. 功能完整性：安装的服务型应用（如mysql/nginx/php/python）没有跟随软件环境启动后自动启动的设置选项，依赖于服务型应用的AI Web UI、php /python应用没有创建桌面快捷方式/菜单快捷方式的入口（即：AI Web UI、php /python应用无法安装和自动开箱即用）。
