`$ python genprog_tests.py 被测EXEC 参考EXEC 用例输入文件`

` 输出：` `Test passed. or Test passed.`

- 值得注意的是`genprog_tests.py`，具有缓存`cache`文件`.genprog_test_cache.json`
- `genprog_tests.py`用于执行可执行的程序，并注入用例，并获取返回结果，并比较是否通过测试
- 利用的是标准程序比对法

- 可以看出我们的`introClass`的单个程序内容简明，但优势在于具有丰富的bug多特征
- [ ] 我的第一步任务是：用遗传算法生成用例并获取覆盖率

----

#### No.1 生成测试用例

- [x] `grades`程序集
  - `Makefile`的使用，递归编译文件
  - `gcovr`的使用，可以生成`xml`、`html`文件记录覆盖率，但是值得注意的是，
  该文件一旦生成会一直叉桩记录覆盖率，因此是累计覆盖率
  - 使用遗传算法生成测试用例是一个闭环的处理过程，即上一次生成的测试用例的
    覆盖率指标对下一次的用例生成具有指导作用，即启发式算法，因此设计思路是
    一个用例对应一个覆盖率，但其实也可以将这个所谓的`一个用例`的规模扩大，
    比如两个用例一组，对应一个覆盖率，（就是不知道适不适用张师兄的算法？）

#### No.2 故障定位信息矩阵收集
- [ ] `grade`程序集
  