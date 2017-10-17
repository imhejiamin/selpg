# selpg:CLI 命令行实用程序开发基础
------
## 概述
CLI（Command Line Interface）实用程序是Linux下应用开发的基础。正确的编写命令行程序让应用与操作系统融为一体，通过shell或script使得应用获得最大的灵活性与开发效率。Linux提供了cat、ls、copy等命令与操作系统交互；go语言提供一组实用程序完成从编码、编译、库管理、产品发布全过程支持；容器服务如docker、k8s提供了大量实用程序支撑云服务的开发、部署、监控、访问等管理任务；git、npm等都是大家比较熟悉的工具。尽管操作系统与应用系统服务可视化、图形化，但在开发领域，CLI在编程、调试、运维、管理中提供了图形化程序不可替代的灵活性与效率。

## 参考
https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html

## 要求实现功能
$ selpg -s1 -e1 input_file

该命令将把“input_file”的第 1 页写至标准输出（也就是屏幕），因为这里没有重定向或管道。

$ selpg -s1 -e1 < input_file

该命令与示例 1 所做的工作相同，但在本例中，selpg 读取标准输入，而标准输入已被 shell／内核重定向为来自“input_file”而不是显式命名的文件名参数。输入的第 1 页被写至屏幕。

$ other_command | selpg -s10 -e20

“other_command”的标准输出被 shell／内核重定向至 selpg 的标准输入。将第 10 页到第 20 页写至 selpg 的标准输出（屏幕）。

$ selpg -s10 -e20 input_file >output_file

selpg 将第 10 页到第 20 页写至标准输出；标准输出被 shell／内核重定向至“output_file”。

$ selpg -s10 -e20 input_file 2>error_file

selpg 将第 10 页到第 20 页写至标准输出（屏幕）；所有的错误消息被 shell／内核重定向至“error_file”。请注意：在“2”和“>”之间不能有空格；这是 shell 语法的一部分（请参阅“man bash”或“man sh”）。

$ selpg -s10 -e20 input_file >output_file 2>error_file

selpg 将第 10 页到第 20 页写至标准输出，标准输出被重定向至“output_file”；selpg 写至标准错误的所有内容都被重定向至“error_file”。当“input_file”很大时可使用这种调用；您不会想坐在那里等着 selpg 完成工作，并且您希望对输出和错误都进行保存。

$ selpg -s10 -e20 input_file >output_file 2>/dev/null

selpg 将第 10 页到第 20 页写至标准输出，标准输出被重定向至“output_file”；selpg 写至标准错误的所有内容都被重定向至 /dev/null（空设备），这意味着错误消息被丢弃了。设备文件 /dev/null 废弃所有写至它的输出，当从该设备文件读取时，会立即返回 EOF。

$ selpg -s10 -e20 input_file >/dev/null

selpg 将第 10 页到第 20 页写至标准输出，标准输出被丢弃；错误消息在屏幕出现。这可作为测试 selpg 的用途，此时您也许只想（对一些测试情况）检查错误消息，而不想看到正常输出。

$ selpg -s10 -e20 input_file | other_command

selpg 的标准输出透明地被 shell／内核重定向，成为“other_command”的标准输入，第 10 页到第 20 页被写至该标准输入。“other_command”的示例可以是 lp，它使输出在系统缺省打印机上打印。“other_command”的示例也可以 wc，它会显示选定范围的页中包含的行数、字数和字符数。“other_command”可以是任何其它能从其标准输入读取的命令。错误消息仍在屏幕显示。

$ selpg -s10 -e20 input_file 2>error_file | other_command

与上面的示例 9 相似，只有一点不同：错误消息被写至“error_file”。

$ selpg -s10 -e20 -l66 input_file
该命令将页长设置为 66 行，这样 selpg 就可以把输入当作被定界为该长度的页那样处理。第 10 页到第 20 页被写至 selpg 的标准输出（屏幕）。

$ selpg -s10 -e20 -f input_file

假定页由换页符定界。第 10 页到第 20 页被写至 selpg 的标准输出（屏幕）。

$ selpg -s10 -e20 -dlp1 input_file

第 10 页到第 20 页由管道输送至命令“lp -dlp1”，该命令将使输出在打印机 lp1 上打印。

最后一个示例将演示 Linux shell 的另一特性：

$ selpg -s10 -e20 input_file > output_file 2>error_file &

该命令利用了 Linux 的一个强大特性，即：在“后台”运行进程的能力。在这个例子中发生的情况是：“进程标识”（pid）如 1234 将被显示，然后 shell 提示符几乎立刻会出现，使得您能向 shell 输入更多命令。同时，selpg 进程在后台运行，并且标准输出和标准错误都被重定向至文件。这样做的好处是您可以在 selpg 运行时继续做其它工作。

## 实验结果与测试
### 实验前
测试前项目下包含的文件内容，test.txt为测试使用的输入文件，selpg.go为go源代码，selpg.go为编译出来的go的可执行程序。

![-1](https://github.com/imhejiamin/selpg/blob/master/my_images/-1.png)

下面是c源代码提供给用户的的usage（）函数

![-2](https://github.com/imhejiamin/selpg/blob/master/my_images/-2.png)

这是go语言下提供给用户的usage（）函数参考

![-3](https://github.com/imhejiamin/selpg/blob/master/my_images/-3.png)

作为input file 的 test.txt文件，一共100行，每行一个数字，一共是1-100.

![-4](https://github.com/imhejiamin/selpg/blob/master/my_images/-4.png)

### 1、selpg -s 1 -e 1 test.txt
因为是没有传l参数，就默认l页长度为72行，所以第一页打印到第一页结束，共一页，即72行。

![1](https://github.com/imhejiamin/selpg/blob/master/my_images/1.png)

### 2、selpg -s 1 -e 1 < test.txt
< 符号selpg 读取标准输入，而标准输入已被 shell／内核重定向为来自“input_file”而不是显式命名的文件名参数。输入的第 1 页被写至屏幕。与1输出相同。

![2](https://github.com/imhejiamin/selpg/blob/master/my_images/2.png)

### 3、type test.txt | selpg -s 1 -e 1
也是跟1输出相同，“type test.txt”的标准输出被 shell／内核重定向至 selpg 的标准输入。将第 10 页到第 20 页写至 selpg 的标准输出。

![3](https://github.com/imhejiamin/selpg/blob/master/my_images/3.png)

### 4、selpg -s 10 -e 20 -l 1 test.txt > out.txt
> 符号将输出重定向到out.txt文件。

![4](https://github.com/imhejiamin/selpg/blob/master/my_images/4.png)

会发现project下面多了一个新文件，out.txt。

![4-2](https://github.com/imhejiamin/selpg/blob/master/my_images/4-2.png)

打开out.txt，就是输出的内容。

![4-3](https://github.com/imhejiamin/selpg/blob/master/my_images/4-3.png)

### 5、selpg -s 10 -e 20 test.txt 2>error.txt
这个命令实现将前面的命令执行的错误结果重定向输出到error.txt文件中。

![5](https://github.com/imhejiamin/selpg/blob/master/my_images/5.png)

会发现project下面多了一个新文件，error.txt。

![5-2](https://github.com/imhejiamin/selpg/blob/master/my_images/5-2.png)

打开文件，出现报错，因为输入文件太短，该命令需要读取11页的数据（每一页初始默认72行）。

![5-3](https://github.com/imhejiamin/selpg/blob/master/my_images/5-3.png)

### 6、selpg -s 10 -e 20 test.txt >out.txt 2>error.txt
两次重定向输出，屏幕无输出，out.txt没有内容，error.txt有存储错误信息。

![6](https://github.com/imhejiamin/selpg/blob/master/my_images/66.png)

![6-2](https://github.com/imhejiamin/selpg/blob/master/my_images/6-2.png)

### 7、selpg -s 10 -e 20 test.txt >out.txt 2>null
7的输出结果是与6一样的。

![7](https://github.com/imhejiamin/selpg/blob/master/my_images/7.png)

### 8、selpg -s 10 -e 20 rest.txt >null
屏幕标准输出错误信息，没有Output信息。

![8](https://github.com/imhejiamin/selpg/blob/master/my_images/8.png)

### 9、selpg -s 10 -e 20 test.txt | type test.txt
一直把test.txt的文件显示完毕，最后还输出报错信息，报错信息来自于前一个命令。

![9](https://github.com/imhejiamin/selpg/blob/master/my_images/9.png)

![9-2](https://github.com/imhejiamin/selpg/blob/master/my_images/9-2.png)

### 10、selpg -s 10 -e 20 test.txt 2>error.txt | type test.txt

![10](https://github.com/imhejiamin/selpg/blob/master/my_images/10.png)

![10-2](https://github.com/imhejiamin/selpg/blob/master/my_images/10-2.png)

与9类似，但是，错误信息在error.txt里面了。

![10-3](https://github.com/imhejiamin/selpg/blob/master/my_images/10-3.png)

### 11、selpg -s 5 -e 10 -l 1 test.txt
把页长度定义为1行，输出第5页到第10页。

![11](https://github.com/imhejiamin/selpg/blob/master/my_images/11.png)

### 12、selpg -s 10 -e 20 test.txt > out.txt 2>error.txt &
实行了重定向输出并在后台挂起了程序。

![12](https://github.com/imhejiamin/selpg/blob/master/my_images/12.png)

