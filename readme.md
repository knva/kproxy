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
