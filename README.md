## Router OS v7 Service Using Go and Fiber

This is a simple example of a service that can be used to manage a Router OS device using the [RouterOS API](https://wiki.mikrotik.com/wiki/Manual:API). It is written in Go and uses the [Fiber](https://gofiber.io/) web framework.

### Usage
#### Build and Run
```bash
task dev
```

### Test POST | Adding PPPoE User Using Curl
```bash
curl -sS -X POST --location "http://localhost:3000/api/v1/ppp/secret/add" \
     -H "Content-Type: application/json" \
     -d '{
            "router_ip_addr": "192.168.88.1",
            "router_username": "admin",
            "router_password": "password",
            "username_pppoe": "user_pppoe",
            "password_pppoe": "password_pppoe
            "profile_pppoe": "FO_10M",
            "remote-address_pppoe": "10.0.0.35"
        }' | jq
```

### Test DELETE | Removing PPPoE User Using Curl
```bash
curl -sS -X DELETE --location "http://localhost:3000/api/v1/ppp/secret/delete" \
     -H "Content-Type: application/json" \
     -d '{
            "router_ip_addr": "192.168.88.1",
            "router_username": "admin",
            "router_password": "password",
            "remote-address_pppoe": "10.0.0.35"
        }' | jq
```