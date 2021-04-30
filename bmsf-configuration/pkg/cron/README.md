定时任务框架
============

## 功能

- 支持简易语法描述如"@every 1s"表示每秒执行一次;
- 支持Unix/Linux Crontab语法如"0 * * * *";
- 单个Cron结构可管理多个定时任务;
- 支持代码中进行任务的注册、反注册，支持Cron整体的暂停、恢复、停止操作;
- 支持基于动态配置文件的描述，在cron.yaml中修改任务定义描述可实时生效, 若想停止指定某个任务则注释其对应描述即可;

## 说明

    任务严格按照描述的间隔进行执行，即使上轮任务还没有执行完成，新的任务执行也会开始，该原则与Unix/Linux Crontab一致,
    业务代码中需要自行根据场景需求做好重入，尤其是间隔较小的任务, 这种情况下需要在NeedRun函数中进行是否可以重入的判定!

## 示例

    参见example