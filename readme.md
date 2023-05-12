# kproxy: 基于TCP的IPv6 HTTP代理服务器 / TCP-based IPv6 HTTP Proxy Server

kproxy是一个简单的基于TCP的HTTP代理服务器，使用Golang编写，支持IPv6地址和自定义端口。 / kproxy is a simple TCP-based HTTP proxy server written in Golang, supporting IPv6 addresses and custom ports.

## 功能 / Features

- 支持HTTP和HTTPS请求 / Support for HTTP and HTTPS requests
- 支持IPv6出口 / Support for IPv6 exit
- 支持自定义端口 / Support for custom ports

## 用法 / Usage

1. 安装Golang环境 / Install the Golang environment

2. 编译程序 / Compile the program

    ```bash
    go build -o kproxy main.go
    ```

3. 运行代理服务器 / Run the proxy server

    ```bash
    ./kproxy -b [绑定IP:端口] -i [IPv6网段]
    ```

   示例 / Example:

    ```bash
    ./kproxy -b :8080 -i 2001:db8::/64
    ```

## 参数说明 / Parameter Description

- `-b`: 用于绑定代理服务器的IP地址和端口。默认值为":8080"。 / Used to bind the proxy server's IP address and port. The default value is ":8080".
- `-i`: 用于指定IPv6网段，以便在该网段内生成随机IPv6地址。 / Used to specify the IPv6 subnet for generating random IPv6 addresses within that subnet.

## 注意 / Note

为了确保最佳性能，您可能需要根据您的具体需求和应用场景对代码进行调整和优化。 / To ensure optimal performance, you may need to adjust and optimize the code according to your specific needs and application scenarios.

方案一：使用 HE.net 的隧道服务获取 IPv6

Hurricane Electric（HE.net）提供了一种免费的隧道服务，用于获取 IPv6 地址。以下是详细的操作步骤：

注册 HE.net 账号

访问 HE.net 隧道服务页面，点击 "Create Account" 按钮创建一个新账号。

创建隧道

登录账号后，点击 "Create Regular Tunnel"，输入您的 IPv4 地址，选择离您最近的隧道服务器，然后点击 "Create Tunnel"。

配置隧道

在创建的隧道的详细信息页面，按照提供的示例配置您的操作系统。这些示例包括了 Windows、macOS、Linux 等操作系统的配置方法。

验证配置

配置完成后，使用 ping 命令测试 IPv6 连接是否正常工作。例如，可以尝试 ping ipv6.google.com。

方案二：在 Vultr 上使用 ndppd 进行转发获取 IPv6
Vultr 提供了一种在其 VPS 上通过 ndppd 转发 IPv6 地址的方法。以下是详细的操作步骤：

在 Vultr 购买 VPS

访问 Vultr 官网，购买一台支持 IPv6 的 VPS。

安装 ndppd

以 root 用户身份登录 VPS，然后执行以下命令安装 ndppd：


```bash
sudo apt-get update
sudo apt-get install ndppd
```
配置 ndppd

创建一个名为 ndppd.conf 的配置文件，并编辑内容如下：

```bash
route-ttl 30000
proxy eth0 {
  router yes
  timeout 500
  ttl 30000
  rule 2001:19f0:xxxx:xxxx::/64 {
    static
  }
}

```
请将 2001:19f0:xxxx:xxxx::/64 替换为您的 VPS 分配的 IPv6 地址段。

启动 ndppd

将 ndppd.conf 文件上传至 VPS 的 /etc/ 目录下，然后执行以下命令启动 ndppd 服务：

```bash
sudo systemctl enable ndppd
sudo systemctl start ndppd
```
验证配置

在本地设备上配置静态 IPv6 地址，然后使用 ping 命令测试连接是否正常工作。例如，尝试 ping ipv6.google.com。


## 许可

Alconna 采用 [MIT](LICENSE) 许可协议

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FArcletProject%2FAlconna.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2FArcletProject%2FAlconna?ref=badge_large)

## 鸣谢

[JetBrains](https://www.jetbrains.com/): 为本项目提供 [GoLand](https://www.jetbrains.com/goland/) 等 IDE 的授权<br>
[<img src="https://cdn.jsdelivr.net/gh/Kyomotoi/CDN@master/noting/jetbrains-variant-3.png" width="200"/>](https://www.jetbrains.com/)

