**常用命令：**
+ 对 Commit Message 进行过滤查找： `git log --oneline --grep “^feat|^fix|^perf”`

### Angular 规范

在 Angular 规范中，Commit Message 包含三个部分，分别是 Header、Body 和 Footer，格式如下：

```
<type>[optional scope]: <description>
// 空行
[optional body]
// 空行
[optional footer(s)]
```

其中，Header 是必需的，Body 和 Footer 可以省略。在以上规范中，<scope> 必须用括号 () 括起来， <type>[<scope>] 后必须紧跟冒号 ，冒号后必须紧跟空格，2 个空行也是必需的。

**一个符合 Angular 规范的 Commit Message 示例：**

```
fix($compile): couple of unit tests for IE9
# Please enter the Commit Message for your changes. Lines starting
# with '#' will be ignored, and an empty message aborts the commit.
# On branch master
# Changes to be committed:
# ...

Older IEs serialize html uppercased, but IE9 does not...
Would be better to expect case insensitive, unfortunately jasmine does
not allow to user regexps for throw expectations.

Closes #392
Breaks foo.bar api, foo.baz should be used instead
```

+ Header：包括三个字段：type（必选）、scope（可选）和 subject（必选）。


