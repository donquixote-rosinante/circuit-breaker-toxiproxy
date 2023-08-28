1. Install toxiproxy
2. Run toxiproxy -> toxiproxy-server -config ./toxiproxy.json
3. Run app on dummy-endpoint folder
4. Run app on root folder
5. Execute curl -X POST  --data '{"enabled":true}' http://localhost:8474/proxies/dummy_http 
{"name":"dummy_http","listen":"[::]:20000","upstream":"localhost:8080","enabled":true,"Logger":{},"toxics":[]} for enable proxy
6. Execute curl -X POST  --data '{"enabled":false}' http://localhost:8474/proxies/dummy_http
{"name":"dummy_http","listen":"[::]:20000","upstream":"localhost:8080","enabled":false,"Logger":{},"toxics":[]} for disable proxy