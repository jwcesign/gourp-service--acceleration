# gourp-service-acceleration

## Background
Knative recently lacks a way to handle group-service(i.e. A and B are different knative services, which form a complex function. The requesting path is: user-client requests ksvc A, then ksvc A requests ksvc B). As a result, cloud providers may encounter difficulties with very long response times(ksvc A cold start time + ksvc B cold start time). To alleviate this issue, this feature track proposes implementing a group-service-acceleration functionality that can scale up group-service at same time when they need to call each other. This will reduce the response time of requesting group-service. Also, the longer the requesting path, the more response time is reduced.

## Example Description
There is 3 ksvc, a, b and c. The request order is: ksvc a requests ksvc b, ksvc b requests ksvc c.

## How to deploy
Run:
```shell
git clone https://github.com/jwcesign/group-service-acceleration.git
cd group-service-acceleration
ko apply -Rf config
```

## Test results
Request with cold start: 29.143s
```shell
root@cesign [12:55:22 AM] [+47.0°C] [~/git/group-service-acceleration] [main *]
-> # time curl -H "Host:group-service-a.default.example.com" 172.19.0.2:31775
group-service-a -> group-service-b -> group-service-c  0.01s user 0.00s system 0% cpu 29.143 total
```

Request with single cold start: 6.881s
```shell
root@cesign [01:04:23 AM] [+48.0°C] [~/git/group-service-acceleration] [main *]
-> # time curl -H "Host:group-service-c.default.example.com" 172.19.0.2:31775
group-service-c  0.00s user 0.01s system 0% cpu 6.881 total
```

Request with hot pods:  0.009s
```shell
root@cesign [01:01:46 AM] [+45.0°C] [~/git/group-service-acceleration] [main *]
-> # time curl -H "Host:group-service-a.default.example.com" 172.19.0.2:31775
group-service-a -> group-service-b -> group-service-c 0.00s user 0.00s system 67% cpu 0.009 total
```
