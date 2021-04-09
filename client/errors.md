报错：
 ERROR  Failed to compile with 2 errors                                                                                                               9:00:35 ├F10: PM┤This dependency was not found:

* !!vue-style-loader!css-loader?
解决办法：
所以，解决方案就要根据情况而定，看你使用的CSS语言是什么,是 常规的 或者 less 或者 sass。

如果是 常规 的，执行 npm install stylus-loader css-loader style-loader --save-dev 安装依赖就行。
如果是 less 的，执行 npm install less less-loader --save-dev 安装依赖就行。
如果是 sass 的，执行 npm install sass sass-loader --save-dev 安装依赖就行。或者（$npm intall sass-loader --save ; $npm install node-sass --save）
但可能出现版本不匹配问题，可以参考这个版本
