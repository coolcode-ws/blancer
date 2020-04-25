# blancer
go 负载均衡器,支持一下几种算法：
1. 基于随机算法的负载均衡
2. 基于RoundRobin算法的负载均衡
3. 基于带权重RoundRobin算法的负载均衡
4. 基于一致性hash算法的负载均衡
5. 基于洗牌算法的负载均衡
6. 基于优化洗牌算法的负载均衡

# 运行结果
`name:  hash cost time:  3818
192.168.145.0:8080;call times: 2035
192.168.149.1:8080;call times: 2028
192.168.221.2:8080;call times: 1946
192.168.32.3:8080;call times: 2002
192.168.92.4:8080;call times: 1989
name:  random cost time:  542
192.168.145.0:8080;call times: 2031
192.168.149.1:8080;call times: 2031
192.168.221.2:8080;call times: 1975
192.168.32.3:8080;call times: 1999
192.168.92.4:8080;call times: 1964
name:  roundrobin cost time:  231
192.168.145.0:8080;call times: 2000
192.168.149.1:8080;call times: 2000
192.168.221.2:8080;call times: 2000
192.168.32.3:8080;call times: 2000
192.168.92.4:8080;call times: 2000
name:  weight_roundrobin cost time:  4014
192.168.145.0:8080;call times: 1538 ;weight:  4
192.168.149.1:8080;call times: 1923 ;weight:  5
192.168.221.2:8080;call times: 1538 ;weight:  4
192.168.32.3:8080;call times: 1923 ;weight:  5
192.168.92.4:8080;call times: 3078 ;weight:  8
name:  shuffle cost time:  149219
192.168.32.3:8080;call times: 2021
192.168.145.0:8080;call times: 2025
192.168.92.4:8080;call times: 2071
192.168.149.1:8080;call times: 1950
192.168.221.2:8080;call times: 1933
name:  shuffle2 cost time:  145782
192.168.221.2:8080;call times: 2007
192.168.149.1:8080;call times: 2029
192.168.145.0:8080;call times: 1998
192.168.92.4:8080;call times: 1926
192.168.32.3:8080;call times: 2040
`
