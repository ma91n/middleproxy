# middleproxy

middleproxy forwards your request another self signed certificated http proxy.

ClientApp(@localhost) -> middleproxy(@localhost) -> sefl signed certificated proxy.

## Usage

```sh
# Windows
set http_proxy=http://{username}:{password}@proxy.example.com:8000
set http_proxy_username={username}
set http_proxy_password={password}

middleproxy.exe
```

middleproxy default port is 9000.

