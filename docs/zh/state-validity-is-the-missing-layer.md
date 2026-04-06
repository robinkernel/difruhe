# MCP 还不够：缺失的那一层是 State Validity

AI 基础设施正在快速走向一种新共识。

更多协议。  
更多工具。  
更多评测。  
更多互操作。  
更多 agent workflow。  

这些都重要。  
而且确实有价值。

但我认为，行业还缺一个更重要的层：

**state validity**

一个系统连接得很好，  
不代表它有资格行动。

它可以拥有正确的模型、  
正确的工具访问、  
正确的仓库指令、  
正确的 benchmark、  
正确的 workflow 集成——  

但它仍然可能处在**错误的状态**里。

它可能在读取过期的 handoff。  
可能站在已经失效的 authority frame 上行动。  
可能把早该失效的 permission 一路往后带。  
可能把历史 trace 当成当前 clearance。  
可能把旧的 entry surface 误当成 live one。  
也可能把一次局部同步误当成了有效的 present。

这不只是 capability 问题。

这是 **state validity** 问题。

而 state validity 会成为严肃 AI 系统里最重要的概念之一。

因为现在真正的问题，已经不只是：

**系统能不能完成任务？**

更难的问题是：

**这个系统在此时此刻，是否仍然拥有说话、决策、或改动正式世界的资格？**

这个门槛高得多。

而且本来就应该更高。

我们正在进入一个阶段：  
模型可以调用工具，  
工具可以触发 workflow，  
workflow 可以触碰生产系统，  
多个 agent 可以在循环里协作。

在这种世界里，最大的失败不只来自糟糕推理。

更常见的失败，会来自 **invalid continuation**。

一个已经过期的状态还在继续行动。  
一个 helper output 被过度提升。  
一个 handoff packet 看起来像当前，但其实不是。  
一个本该停留在 candidate-only 的任务，悄悄滑进了 execution。  
一条本该只属于历史的 trace，被再次拿来当 live permission。

这就是为什么我认为，  
connectivity 之后的下一层，不是“更多 autonomy”。

而是 **admissibility**。

在 AI 系统行动之前，  
必须先有一种方式判断：  
当前 state 是否仍然 admissible for action。

不是判断模型聪不聪明。  
不是判断工具能不能用。  
不是判断 benchmark 看起来够不够高。  
不是判断协议写得漂不漂亮。

而是判断：

**这个 state 还有效吗？**

这至少意味着五件事。

第一，**landing 不是 permission**。  
写进 formal host，不等于已经获得改动世界的资格。

第二，**adoption 不是 execution**。  
系统可以识别、吸收、登记一个对象，但这不等于它可以行动。

第三，**execution 不是持续 clearance**。  
即使一次动作是合法的，也不会自动生成永久延续的正当性。

第四，**history 不是 current authority**。  
trace、log、旧 packet 可以解释过去，但不能自动授权现在。

第五，**最新同步 state 必须高于更旧的 entry surface**。  
如果多个入口层互相冲突，系统必须先判定什么才是真正 current，然后 opener 才能定义下一步。

这正是很多 agent 基础设施现在还不完整的地方。

我们正在变得更会连接。  
更会编排。  
更会评测。  
更会调工具。  

但仍然缺少足够多的系统，在行动前先问一句：

**现在这个 present 还有效吗？**

这听起来很抽象。

其实一点都不抽象。

它决定：

一个 coding agent 能不能碰 repo。  
一个 workflow 能不能在 checkpoint 后继续。  
一个 handoff 能不能定义 next task。  
一个 helper agent 是否必须继续保持 candidate-only。  
一条 execution trace 是否必须先 demobilize，之后任何新动作才能重新开始。  
一个 automated loop 到底还是 lawful probe，还是已经漂进 unauthorized action。

换句话说：

**state validity 是 governance 从“政策文档”变成“runtime substrate”的地方。**

这也是为什么我认为，未来十年的 AI 基础设施会分成两类：

一类继续优化 throughput：  
更多工具、更多调用、更多 loop、更多 autonomous chain。

另一类则在建设一种更难的能力：  
精确判断系统在什么条件下，仍然保有行动资格。

前一类会跑得更快。  
后一类会更容易获得信任。

而 trust，最终会比单纯速度更容易复利。

所以，是的——  
协议重要。  
评测重要。  
指令面重要。  
互操作重要。

但它们都不是最后一层。

因为一个连接上的 agent，  
不会自动变成一个合法 actor。

缺失的那一层，是 **state validity**。

而我认为，这一层会成为下一代 AI 系统最关键的基础抽象之一。

**MCP 给了 AI 一个 socket。  
State validity 才决定它是否有资格触碰世界。**
