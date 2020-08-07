#  go 1.12
> 单进程锁定  


调用方式 
```golang
plock.Lock()
defer plock.UnLock()
```