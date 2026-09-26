# 中国劳动法知识库（china-labor-advisor）

单一技能 `china-labor-advisor`（中国劳动法顾问），内部集成 **17 部法律法规全文**的结构化知识模块（`topics/`，共 **95 个章节文件**）与 **1 个应用型审查模块**（policy-review，10 个审查主题）。用户只需直接提问（"我被裁了"、"工伤了怎么办"、"公积金该交多少"、"帮我看看员工手册合不合法"），技能内的**场景路由表**自动定位到对应法律模块——无需用户了解任何法律名称或技能结构。

- **技能入口**: [china-labor-advisor/SKILL.md](china-labor-advisor/SKILL.md)
- **生成日期**: 2026-09-27
- **法律文本基准**: 劳动法（2018修正）、劳动合同法（2012修正）、社会保险法（2018修正）、调解仲裁法（2007通过，2008-05-01施行）、民法典（2020）、工伤保险条例（2010修订）、年休假条例（2007）+实施办法（2008）、住房公积金管理条例（2019修订）、保障农民工工资支付条例（2020）、劳动保障监察条例（2004）、职业病防治法（2018修正）、安全生产法（2021修正）、就业促进法（2015修正）、女职工劳动保护特别规定（2012）、残疾人保障法（2008修订，2018修正）、工会法（2021修正）、最高法劳动争议解释（一）（2021）+（二）（2025）

---

## 技能内部结构

```
china-labor-advisor/
├── SKILL.md                    ← 技能入口：场景路由表 + 法条定位表 + 主题索引
└── topics/                     ← 15 个法律知识模块（每个 = INDEX.md + chapters/ + 支撑文件：glossary/patterns/cheatsheet；china-labor-law 另含 worker-playbook）
    ├── china-labor-law/                劳动法（13章）· 框架法 + 劳动者维权行动手册
    ├── labor-contract-law/             劳动合同法（8章）· 合同/解除/N与2N/派遣
    ├── social-insurance-law/           社会保险法（12章）· 五险/征缴/待遇
    ├── labor-dispute-arbitration-law/  调解仲裁法（5章）· 仲裁前置/时效/一裁终局
    ├── labor-dispute-judicial-interpretations/  最高法解释一+二（8章）· 审判口径（2025最新）
    ├── civil-code/                     民法典（5章）· 协议撤销/人格权/侵权赔偿
    ├── work-injury-regulations/        工伤保险条例（4章）· 认定/伤残梯度/工亡
    ├── paid-annual-leave-regulations/  年休假条例+办法（2章）· 折算/300%补偿
    ├── housing-fund-regulations/       住房公积金管理条例（2章）· 缴存比例/提取/投诉
    ├── migrant-worker-wage-regulations/  保障农民工工资支付条例（6章）· 工资支付/清偿/工程建设制度
    ├── labor-inspection-regulations/   劳动保障监察条例（3章）· 监察职责/调查/加付赔偿
    ├── occupational-disease-prevention-law/  职业病防治法（6章）· 预防/健康检查/诊断鉴定
    ├── work-safety-law/                安全生产法（6章）· 安全义务/从业人员权利/事故处罚
    ├── employment-promotion-law/       就业促进法（4章）· 公平就业/反歧视/就业援助
    ├── female-worker-protection/       女职工劳动保护特别规定（3章）· 三期/产假/生育津贴
    └── policy-review/                  规章制度合规审查（应用层）· 四步流程+10主题红牌库

每个模块目录内：INDEX.md（模块索引，原 SKILL.md）+ chapters/ + glossary.md + patterns.md + cheatsheet.md
（china-labor-law 另含 worker-playbook.md 劳动者维权行动手册）
```

## 15 个知识模块 + 1 个应用模块

| 模块 | 法律依据 | 章节 | 职责 |
|---|---|---|---|
| [china-labor-law](china-labor-advisor/topics/china-labor-law/INDEX.md) | 《劳动法》107条（2018修正） | 13 | **框架法**：就业、工时、工资、社保、安全卫生、特殊保护、争议、法律责任 |
| [labor-contract-law](china-labor-advisor/topics/labor-contract-law/INDEX.md) | 《劳动合同法》98条（2012修正） | 8 | **劳动关系特别法**：订立、二倍工资、无固定期限、解除/终止、N与2N、劳务派遣 |
| [social-insurance-law](china-labor-advisor/topics/social-insurance-law/INDEX.md) | 《社会保险法》98条（2018修正） | 12 | **社保特别法**：五险框架、征缴、待遇、基金、监督处罚 |
| [labor-dispute-arbitration-law](china-labor-advisor/topics/labor-dispute-arbitration-law/INDEX.md) | 《调解仲裁法》（2007通过） | 5 | **程序法**：仲裁前置、1年时效、一裁终局、开庭举证 |
| [labor-dispute-judicial-interpretations](china-labor-advisor/topics/labor-dispute-judicial-interpretations/INDEX.md) | 最高法解释（一）54条+（二）21条 | 8 | **审判口径**：受案、当事人、竞业、放弃社保无效、时效失权（2025-09-01最新） |
| [civil-code](china-labor-advisor/topics/civil-code/INDEX.md) | 《民法典》1260条（2021施行） | 5 | **民事基本法**（劳动关联）：协议撤销、格式条款、人格权（性骚扰/个人信息）、用人单位责任、工伤侵权双轨 |
| [work-injury-regulations](china-labor-advisor/topics/work-injury-regulations/INDEX.md) | 《工伤保险条例》67条（2010修订） | 4 | **行政法规**：工伤认定、48小时、伤残待遇梯度、工亡三费 |
| [paid-annual-leave-regulations](china-labor-advisor/topics/paid-annual-leave-regulations/INDEX.md) | 年休假条例10条+实施办法19条 | 2 | **行政规章**：5/10/15天档位、两道折算、300%未休补偿 |
| [housing-fund-regulations](china-labor-advisor/topics/housing-fund-regulations/INDEX.md) | 《住房公积金管理条例》50条（2026第三次修订，国令第844号） | 2 | **行政法规**：缴存比例5%–国家最高比例、月缴存额公式、提取九情形、不缴强制执行、骗提骗贷罚则 |
| [migrant-worker-wage-regulations](china-labor-advisor/topics/migrant-worker-wage-regulations/INDEX.md) | 《保障农民工工资支付条例》64条（2020施行） | 6 | **行政法规**：工资支付形式、清偿责任链、工程建设五项制度、失信惩戒、拒不支付劳动报酬罪移送 |
| [labor-inspection-regulations](china-labor-advisor/topics/labor-inspection-regulations/INDEX.md) | 《劳动保障监察条例》36条（2004施行） | 3 | **行政法规**：监察职责、管辖分工、调查程序、拖欠工资加付50%–100%、违法工时罚款 |
| [occupational-disease-prevention-law](china-labor-advisor/topics/occupational-disease-prevention-law/INDEX.md) | 《职业病防治法》87条（2018修正） | 6 | **法律**：前期预防三同时、职业健康检查、职业病诊断鉴定、工伤+民事赔偿 |
| [work-safety-law](china-labor-advisor/topics/work-safety-law/INDEX.md) | 《安全生产法》119条（2021修正） | 6 | **法律**：生产经营单位安全义务、从业人员六项权利、事故调查、事故罚款30万–2000万 |
| [employment-promotion-law](china-labor-advisor/topics/employment-promotion-law/INDEX.md) | 《就业促进法》69条（2015修正） | 4 | **法律**：平等就业权、反就业歧视（性别/残疾/乙肝/户籍）、公共就业服务、就业援助、歧视可直接起诉 |
| [female-worker-protection](china-labor-advisor/topics/female-worker-protection/INDEX.md) | 《女职工劳动保护特别规定》16条（2012） | 3 | **行政法规**：三期不得降薪辞退、产假98天+地方奖励假、生育津贴、哺乳期1小时、性骚扰防治 |
| [policy-review](china-labor-advisor/topics/policy-review/INDEX.md) | 规章制度合规审查方法论（应用层） | 10主题 | **应用型**：上传制度文件→五级定级逐条审查→指出违法/损害权益条款→应对策略 |

---

## 使用方法

**直接问场景即可**，技能自动路由（完整路由表见 [SKILL.md](china-labor-advisor/SKILL.md)）：

| 提问方式 | 示例 | 路由结果 |
|---|---|---|
| **场景** | "被裁员了怎么算补偿"、"没签劳动合同能要什么" | 劳动合同法 ch04/ch02 + 司法解释联动 |
| **法条** | "劳动合同法第38条"、"解释二第19条" | 法条定位表 → 对应模块 |
| **计算** | "工伤十级能赔多少"、"年假没休给多少钱" | 工伤条例 ch03 / 年休假 cheatsheet 公式 |
| **程序** | "时效过了吗"、"一裁终局什么意思" | 仲裁法 + 解释二20条 |
| **跨模块** | "上下班车祸，工伤和民事怎么并行" | 路由表已标联动：工伤条例 ch02 + 民法典 ch05 |

### 高频现实场景（新闻/生活中常见）

| 现实场景 | 直接这样问 | 涉及模块 |
|---|---|---|
| 公司裁员 N+1、"毕业"礼包 | "公司让我签协商解除协议，给的 N+1 合理吗？能要 2N 吗" | labor-contract-law ch04（协议解除/违法解除）+ g07 |
| "主动离职"被劝签辞职信 | "领导让我写个人原因辞职，我写了还能要补偿吗" | labor-contract-law ch04（38条被迫解除）+ g07（组合陷阱） |
| 拖欠/克扣工资、年底讨薪 | "公司拖欠三个月工资，投诉还是仲裁？能主张什么" | china-labor-law ch05（工资）+ 仲裁法（支付令/加付赔偿金） |
| 996/大小周、离职追讨加班费 | "离职一年了还能要加班费吗？没有打卡记录怎么办" | china-labor-law ch04（工时，第44条）+ 司法解释 ch04（42条两步举证）+ 仲裁法（时效） |
| 试用期被"随便"辞退 | "试用期第 5 天被辞退说我不合适，有赔偿吗" | labor-contract-law ch02（录用条件举证）+ g02 |
| 没签劳动合同要二倍工资 | "入职 8 个月没签合同，二倍工资能要几个月" | labor-contract-law ch02（最多11个月+时效）+ 解释二 |
| 调岗降薪、变相逼人走 | "公司把我从经理调去仓库看门，不去算旷工吗" | labor-contract-law ch04（合理性审查）+ g06 |
| 绩效末位被淘汰 | "绩效排名垫底就被开除，合法吗" | g06（末位≠不胜任，违法解除2N） |
| 竞业限制天价违约金 | "离职被前公司索赔 50 万竞业违约金，怎么办" | labor-contract-law ch02 + 解释一36–40条（未付补偿抗辩）+ g08 |
| 社保按最低基数缴 / 现金替代 | "公司按最低基数缴社保，发现金补贴，我亏了多少" | social-insurance-law ch07 + 解释二19条（约定无效）+ g10 |
| 公积金没缴/少缴 | "公司一直没给我缴公积金，找谁投诉" | housing-fund-regulations（12329 投诉→强制执行） |
| 工伤：工地/车间/送外卖受伤 | "送外卖摔骨折，平台没给我缴工伤保险怎么办" | work-injury-regulations ch02（未参保单位自担）+ g09 |
| 工伤私了协议签低了 | "工伤私了签了 3 万，后来评上十级伤残，能反悔吗" | work-injury-regulations ch03（待遇差额可再主张※） |
| 上班路上车祸 | "下班路上被车撞了，对方全责，工伤和车祸赔偿都能要吗" | work-injury-regulations ch02 + 民法典 ch05（双轨并行） |
| 年假从来不让休 | "工作 5 年从没休过年假，离职能补多少钱" | paid-annual-leave-regulations（300% 未休补偿折算） |
| 离职证明/扣工资要挟交接 | "离职时公司说不开离职证明、扣半个月工资" | labor-contract-law ch04（50条义务）+ g07 |
| 女职工三期被辞退/降薪 | "怀孕被调岗降薪，产假期间被辞退怎么办" | labor-contract-law ch04（42/45条解除禁令）+ g05/g07 |
| 员工手册霸王条款 | "公司手册写迟到罚款 500、离职返还年终奖，合法吗" | policy-review（上传手册→五级定级报告） |
| 仲裁前要准备什么证据 | "准备告公司，需要收集哪些证据" | 仲裁法（举证责任）+ china-labor-law worker-playbook |
| 劳务派遣/外包同工不同酬 | "我签的是派遣合同，和正式员工干一样活工资差一半" | labor-contract-law ch05（同工同酬） |

**模块内文件分工**：
- `INDEX.md` — 模块索引：核心框架 + 章节索引 + 主题索引 + 适用边界
- `chapters/chXX-*.md` — 逐章精读：Core Idea / Frameworks / Key Concepts / Mental Models / Anti-patterns / Worked Example / Key Takeaways / Connects To
- `glossary.md` — 术语表（按条文）
- `patterns.md` — 审查方法与决策流程（"遇到X怎么操作"）
- `cheatsheet.md` — 速查表、决策卡、关键数字卡、高频误判卡

---

## 跨模块联动速查（高频场景）

| 场景 | 主模块 | 联动模块 |
|---|---|---|
| 被裁员/辞退，算 N 还是 2N | labor-contract-law ch04 | 解释一47条（未通知工会=违法解除）、解释二16–18条（继续履行/空窗工资） |
| 没签劳动合同 | labor-contract-law ch02 | 解释二6–11条（二倍工资精度/免责/两连签）、民法典490条（事实合同） |
| 竞业限制 | labor-contract-law ch02 | 解释一36–40条（30%补偿/3个月解约）、解释二13–15条、民法典585条（违约金酌减） |
| 公司没缴社保 / 放弃社保声明 | social-insurance-law ch07 | 解释二19条（声明无效+38条解除+补缴后索回补贴） |
| 公积金该交多少 / 不缴公积金 | housing-fund-regulations（16/18条缴存、40条强制执行） | 社保法 cheatsheet（五险比例速查：12333 vs 12329 分流） |
| 欠薪 | china-labor-law ch05 | 解释一15条（欠条直诉）、49条（保全）、仲裁法（离职起1年） |
| 工伤 | work-injury-regulations | 社保法 ch04、民法典 ch05（第三人侵权双轨：1179/1213条） |
| 离职协议翻盘 | 解释一35条 | 民法典 ch01（147–152条撤销+除斥期间） |
| 年休假没休/离职结算 | paid-annual-leave-regulations | 仲裁法（1年时效）、民法典497条（格式条款无效） |
| 性骚扰/查手机/个人信息 | civil-code ch04 | 劳动法 ch06、劳动合同法8条 |
| 换壳/集团内调动 | 解释一46条（工龄合并） | 解释二3、10条（混同用工/两连签计次） |
| 告谁（皮包/挂靠/无照） | 解释 ch03/ch05 | 工伤条例66条（穿透）、解释一29条（出资人） |

---

## 分层适用关系

```
特别法优先：
  劳动法 / 劳动合同法 / 社会保险法 / 调解仲裁法 / 职业病防治法 / 安全生产法 / 就业促进法 / 女职工劳动保护特别规定（法律·特别法）
      ↓ 有规定从其规定；无规定参照 ↓
  民法典（民事基本法，补充层）
      ↓ 细化标准 ↓
  工伤保险条例 / 年休假条例+办法 / 住房公积金管理条例 / 保障农民工工资支付条例 / 劳动保障监察条例（行政法规）
      ↓ 审判展开 ↓
  最高法解释（一）（二）——回答"法院实际怎么判"
```

**关键时点**：
- **2021-01-01**：民法典施行（旧九法废止）——此前事实按**行为时法**
- **2025-09-01**：解释（二）施行——与解释（一）冲突**以（二）为准**；解释一第32条1款（退休返聘=劳务）**已废止**

---

## 原始法律文本

`_sources/` 目录保存 **17 部**法律法规的提取文本（共 **19 个** `.txt` 文件，年休假含条例+办法两个文件）。17 部已建模块**全部**有对应 `_sources` 原文：

**已建模块（17 部）**：

| 文件 | 内容 |
|---|---|
| labor-law-2018.pdf / labor-law-2018.txt | 劳动法（2018修正） |
| laodonghetongfa.txt | 劳动合同法 |
| shehuibaoxianfa.txt | 社会保险法 |
| tiaojiezhongcaifa.txt | 劳动争议调解仲裁法 |
| jieshiyi.txt / jieshier.txt | 最高法解释（一）/（二） |
| minfadian_full.txt | 民法典（1260条） |
| gongshangtiaoli.txt | 工伤保险条例 |
| nianxiujia-tiaoli.txt / nianxiujia-banfa.txt | 年休假条例 / 实施办法 |
| gongjijin-tiaoli.txt | 住房公积金管理条例（2026第三次修订，国令第844号，50条） |
| nongminggong-tiaoli.txt | 保障农民工工资支付条例 |
| jiancaitiaoli.txt | 劳动保障监察条例 |
| zhiyebingfangzhifa.txt | 职业病防治法（2018修正） |
| anquanshengchanfa.txt | 安全生产法（2021修正） |
| jiuyecujinfa.txt | 就业促进法（2015修正） |
| nvzhigong-tiaoli.txt | 女职工劳动保护特别规定（2012） |
| canjirenbaozhangfa.txt | 残疾人保障法（2008修订，2018修正） |
| gonghuifa.txt | 工会法（1992通过，2021修正） |

---

## 范围与免责声明

**覆盖**：17 部法律/法规全文，按劳动者维权视角深度展开（每条文配 When to use / Anti-patterns / Worked Example）。

**已知边界**：
- **地方口径**：两金标准（工伤医疗/就业补助金）、停工留薪期目录、双赔/补差、加班费基数细则、社保公积金具体费率、产假奖励假天数等以**省级规定与受诉法院口径**为准（知识库内以 ※ 标注），详见 [LOCAL-PRACTICE.md](LOCAL-PRACTICE.md)（含官方查询渠道与回答模板）
- **动态数据**：一次性工亡补助金（20倍×上年度城镇居民人均可支配收入）等逐年更新的标准，文中已列近年度参考值，**数据基准日 2026-09-27**；所有动态数据点的更新位置与方法见 [DYNAMIC-DATA.md](DYNAMIC-DATA.md)，使用时以国家统计局/人社部当年公布为准
- **法规原文核对**：回答中引用法条时附官方原文链接，17 部法律法规及司法解释的官方来源（全国人大网/中国政府网/最高人民法院）集中收录于 [LEGAL-SOURCES.md](LEGAL-SOURCES.md)，便于用户核对原文
- **未展开**：民法典物权/婚姻/继承编仅整理劳动关联点（见 civil-code 模块 Core Frameworks 第6点），未逐条展开
- **时效性**：生成于 2026-09-27，此后新司法解释/修订以官方文本为准

**免责声明与生成方式披露**：本知识库经人工校对（含法条编号扫描与引用核对），**未经执业律师独立复核**，为法律知识整理，**不构成法律意见**。个案结论受证据、伤残等级鉴定与地方实践影响，重大争议请咨询执业律师。

---

## 许可证

本项目采用 [MIT License](LICENSE) 开源许可协议。

- **允许**：自由使用、复制、修改、合并、发布、分发、再许可、商用
- **要求**：保留原始版权声明与许可声明
- **免责**：软件按"现状"提供，作者不承担任何担保与责任

法律文本的版权说明：`_sources/` 目录中的法律法规原文来自中国政府网等官方公开渠道，其文本本身不受本项目 MIT 协议约束（官方发布文本可自由传播）；MIT 协议适用于本项目的**结构化知识库内容**（SKILL.md、chapters/、支撑文件）与 README。
