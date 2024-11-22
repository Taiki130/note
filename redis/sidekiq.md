### redis-cliでRedisに入る。

keyをすべて表示するには以下

```bash
> keys *
```

### queueの数を確認する

まず、sidekiq:queuesのtypeを確認する

```bash
> type "sidekiq:queues"
set
```

setなため、smembersで確認

```bash
> smembers "queues"
<queue名>
<queue名>
<queue名>
```

それぞれのqueueの数を確認する

```bash
> llen<queue名>"
(integer) 0
```

### **現在スケジューリングされている job を確認する**

```bash
> type "schedule"
zset
```

最初のjobは、

```bash
> ZRANGE "schedule" 0 0
```

すべてのjobは、

```bash
> ZRANGEBYSCORE "schedule" -inf +inf
```
