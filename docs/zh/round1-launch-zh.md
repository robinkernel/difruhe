# AI 意图与现实动作之间，缺了一层治理

大多数关于 AI 的讨论，仍然把注意力放在 intelligence 上：

- 模型能生成什么
- 模型能推理到什么程度
- 模型能多有效地调用工具
- 模型能在多大程度上替用户行动

但当 AI 系统从生成语言走向触发现实动作时，一个更深的问题开始出现：

**AI 意图与现实动作之间的边界**

在这条边界上，candidate output、valid state、admissible execution、legitimate action，不再是同一回事。

而这条边界，到现在仍然没有被清楚定义。

很多系统仍然把几个关键区分压扁在一起：

- 把 candidate 当成 action
- 把 output 当成 state
- 把 execution 当成 legitimacy
- 把 success 当成 permission

这种塌缩看起来可能很高效。
但它并不稳定。

Difruhe 从另一个前提出发：

**缺的不是更多 intelligence。**
**缺的是意图与现实动作之间的一层治理。**

Difruhe 不是一个工具。  
Difruhe 不是一个平台。  
Difruhe 不是一个 agent framework。

Difruhe 是一种结构语言，用来区分什么还只是被提出的内容，什么已经成为有效、可引用、并被允许进入现实的对象。

它当前先钉 5 个判断：

- candidate is not action
- state before intent
- admission before execution
- container validity matters
- execution success does not equal legitimacy

Difruhe 现在要做的，是在这个问题还没有大到无法忽视之前，先把这个问题空间命名出来。

如果 AI 真的要在现实世界里行动，那么“能产出”与“被允许进入现实”之间的那条线，就不能再继续保持隐含状态。

那条线需要结构。

Difruhe 想定义的，就是这一层。
