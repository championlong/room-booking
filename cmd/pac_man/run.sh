#!/bin/bash

# 获取脚本当前路径
base_dir=`dirname $0`
# 进入脚本所在目录
cd $base_dir

# 下面写程序启动逻辑
$base_dir/main $1 $2
