---
name: china-labor-advisor
description: "中国劳动法顾问：劳动者与用人单位的劳动用工、社会保障全领域知识库（15部法律法规全文结构化+1个应用型审查模块）。覆盖：劳动合同（签订/二倍工资/无固定期限/试用期/解除终止/经济补偿N与2N/赔偿金）、裁员辞退开除、欠薪克扣工资、加班费、调岗降薪、劳务派遣、竞业限制/保密/服务期违约金、离职协议、五险一金（养老/医疗/失业/生育/工伤）、社保缴费比例与基数（交多少）、住房公积金（缴存比例/提取/贷款/单位不缴）、工伤认定与伤残待遇（48小时/停工留薪期/工亡三费）、年休假（折算/300%补偿）、社保补缴、职场性骚扰、个人信息/隐私、劳动仲裁诉讼（时效/一裁终局/证据举证）、离职协议撤销、退休返聘、农民工工资支付、劳动保障监察、职业病防治、安全生产、就业歧视（性别/乙肝/残疾/户籍）、女职工三期保护（产假98天/生育津贴/哺乳期）、**审查公司规章制度/员工手册是否合法**（上传制度文件→逐条定级→指出违法与损害权益条款→给应对策略）。用于回答'该不该赔/赔多少/交多少/怎么告/告谁/法院怎么判/制度合不合法'类劳动法律问题；2021-2025最新司法解释口径。"
---

<!-- argument-hint: [场景、主题、法条或章号，如"被裁员"、"二倍工资"、"工伤认定"、"第38条"、"竞业限制"] -->

# 中国劳动法顾问（china-labor-advisor）

**知识库**: 15 部法律法规全文结构化（劳动法/劳动合同法/社会保险法/调解仲裁法/最高法解释一+二/民法典/工伤保险条例/年休假条例+办法/住房公积金管理条例/保障农民工工资支付条例/劳动保障监察条例/职业病防治法/安全生产法/就业促进法/女职工劳动保护特别规定），共 87 章节文件 | **基准日**: 2026-09-26

## How to Use This Skill

用户的问题**直接按下方场景路由表**定位到对应知识模块，读取其 INDEX.md（及关联章节文件）后回答。**不要让用户选模块**——路由是本技能的职责。

- **场景类问题**（"我被裁了"/"工伤了怎么办"）→ 查**场景路由表**
- **法条类问题**（"第38条"/"解释二第19条"）→ 查**法条定位表**
- **计算类问题**（"赔多少"/"能拿几万"）→ 查对应场景 + cheatsheet 中的公式/梯度表
- **程序类问题**（"怎么告"/"找谁"/"时效"）→ topics/labor-dispute-arbitration-law + 司法解释
- **跨模块问题**（"车祸工伤双轨"）→ 场景路由表已标联动模块，**一起读**

---

## 场景路由表

| 用户场景 | 主模块（先读） | 联动模块（后读） |
|---|---|---|
| 被裁员/辞退/开除，算 N 还是 2N | [labor-contract-law](topics/labor-contract-law/INDEX.md) ch04 | [司法解释](topics/labor-dispute-judicial-interpretations/INDEX.md) ch04（47条未通知工会）、ch08（16–18条继续履行/空窗工资） |
| 没签劳动合同 / 二倍工资 | [labor-contract-law](topics/labor-contract-law/INDEX.md) ch02 | [司法解释](topics/labor-dispute-judicial-interpretations/INDEX.md) ch06（二倍工资精度/免责/两连签）、[民法典](topics/civil-code/INDEX.md) ch03（490条事实合同） |
| 被降薪/调岗（口头或单方） | [labor-contract-law](topics/labor-contract-law/INDEX.md) ch03 | [司法解释](topics/labor-dispute-judicial-interpretations/INDEX.md) ch04（43条1个月锁死）、[民法典](topics/civil-code/INDEX.md) ch03（544条变更不明） |
| 欠薪/克扣/拖欠工资 | [劳动法](topics/china-labor-law/INDEX.md) ch05 | [司法解释](topics/labor-dispute-judicial-interpretations/INDEX.md) ch01（15条欠条直诉）、ch04（49条保全）、[仲裁法](topics/labor-dispute-arbitration-law/INDEX.md)（离职起1年） |
| 加班费 | [劳动法](topics/china-labor-law/INDEX.md) ch04 | [司法解释](topics/labor-dispute-judicial-interpretations/INDEX.md) ch04（42条两步举证） |
| 公司没缴社保 / 签了放弃社保声明 | [社保法](topics/social-insurance-law/INDEX.md) ch07 | [司法解释](topics/labor-dispute-judicial-interpretations/INDEX.md) ch08（19条声明无效+解除补偿） |
| 公积金该交多少 / 单位不缴公积金 | [公积金条例](topics/housing-fund-regulations/INDEX.md) ch01（16、18条缴存公式与比例） | ch02（37、38条投诉与强制执行）、[社保法 cheatsheet](topics/social-insurance-law/cheatsheet.md)（五险比例速查） |
| 审查公司规章制度 / 员工手册合不合法 | [规章制度审查](topics/policy-review/INDEX.md)（四步审查流程+10主题红牌库） | 五级定级报告（🔴违法/🟠损害权益/🟡程序风险/🔵欠佳/🟢合规），逐条给法律依据与应对 |
| 工伤（受伤了/认定/待遇/赔偿） | [工伤条例](topics/work-injury-regulations/INDEX.md) ch02–ch03 | [社保法](topics/social-insurance-law/INDEX.md) ch04、[民法典](topics/civil-code/INDEX.md) ch05（第三人侵权双轨） |
| 上下班途中车祸 | [工伤条例](topics/work-injury-regulations/INDEX.md) ch02（14条六+举证包） | [民法典](topics/civil-code/INDEX.md) ch05（1213条保险顺序）、[劳动法](topics/china-labor-law/INDEX.md) ch04 |
| 竞业限制/保密协议/服务期 | [labor-contract-law](topics/labor-contract-law/INDEX.md) ch02 | [司法解释](topics/labor-dispute-judicial-interpretations/INDEX.md) ch07（适配原则/违约双付）、[民法典](topics/civil-code/INDEX.md) ch03（585条违约金酌减） |
| 离职协议/N+1私了想翻盘 | [司法解释](topics/labor-dispute-judicial-interpretations/INDEX.md) ch04（35条） | [民法典](topics/civil-code/INDEX.md) ch01（147–152条撤销+除斥期间） |
| 年休假（没休/折算/离职结算） | [年休假](topics/paid-annual-leave-regulations/INDEX.md) | [民法典](topics/civil-code/INDEX.md) ch03（497条格式条款） |
| 职场性骚扰 | [民法典](topics/civil-code/INDEX.md) ch04（1010条） | [劳动法](topics/china-labor-law/INDEX.md) ch06、[劳动合同法](topics/labor-contract-law/INDEX.md) ch03（38条二） |
| 查手机/装摄像头/收集个人信息 | [民法典](topics/civil-code/INDEX.md) ch04 | [劳动合同法](topics/labor-contract-law/INDEX.md) ch02（8条边界） |
| 换公司/换壳（集团内调动） | [司法解释](topics/labor-dispute-judicial-interpretations/INDEX.md) ch04（46条工龄合并） | ch05（3条混同用工）、ch06（10条两连签计次） |
| 怎么仲裁/诉讼/时效过了没 | [仲裁法](topics/labor-dispute-arbitration-law/INDEX.md) | [司法解释](topics/labor-dispute-judicial-interpretations/INDEX.md) ch08（20条失权）、[民法典](topics/civil-code/INDEX.md) ch02 |
| 告谁（皮包公司/挂靠/无照/包工头） | [司法解释](topics/labor-dispute-judicial-interpretations/INDEX.md) ch03、ch05 | [工伤条例](topics/work-injury-regulations/INDEX.md) ch04（66条穿透） |
| 劳务派遣（同工不同酬/退回/致害） | [labor-contract-law](topics/labor-contract-law/INDEX.md) ch05 | [民法典](topics/civil-code/INDEX.md) ch05（1191条二）、[年休假](topics/paid-annual-leave-regulations/INDEX.md)（14条派遣年假） |
| 试用期被辞退 | [labor-contract-law](topics/labor-contract-law/INDEX.md) ch02 | ch04（试用期解除限制） |
| 退休/返聘 | [司法解释](topics/labor-dispute-judicial-interpretations/INDEX.md) ch03（32条已废止） | [社保法](topics/social-insurance-law/INDEX.md) ch02（养老金）、[民法典](topics/civil-code/INDEX.md) ch05（1192条劳务） |

## 法条定位表

| 问到的法 | 模块 |
|---|---|
| 劳动法（第1–107条） | [topics/china-labor-law](topics/china-labor-law/INDEX.md) — 13章逐章对照 |
| 劳动合同法（第1–98条） | [topics/labor-contract-law](topics/labor-contract-law/INDEX.md) — 8章 |
| 社会保险法（第1–98条） | [topics/social-insurance-law](topics/social-insurance-law/INDEX.md) — 12章 |
| 住房公积金管理条例（7章47条） | [topics/housing-fund-regulations](topics/housing-fund-regulations/INDEX.md) — 2章 |
| 规章制度合规审查（应用层，10主题红牌库） | [topics/policy-review](topics/policy-review/INDEX.md) — review-guides |
| 劳动争议调解仲裁法 | [topics/labor-dispute-arbitration-law](topics/labor-dispute-arbitration-law/INDEX.md) — 5章 |
| 最高法解释（一）54条/（二）21条 | [topics/labor-dispute-judicial-interpretations](topics/labor-dispute-judicial-interpretations/INDEX.md) — 8章 |
| 民法典（1260条，劳动关联） | [topics/civil-code](topics/civil-code/INDEX.md) — 5章 |
| 工伤保险条例（67条） | [topics/work-injury-regulations](topics/work-injury-regulations/INDEX.md) — 4章 |
| 年休假条例+实施办法 | [topics/paid-annual-leave-regulations](topics/paid-annual-leave-regulations/INDEX.md) — 2章 |
| 保障农民工工资支付条例（64条） | [topics/migrant-worker-wage-regulations](topics/migrant-worker-wage-regulations/INDEX.md) — 6章 |
| 劳动保障监察条例（36条） | [topics/labor-inspection-regulations](topics/labor-inspection-regulations/INDEX.md) — 3章 |
| 职业病防治法（87条） | [topics/occupational-disease-prevention-law](topics/occupational-disease-prevention-law/INDEX.md) — 6章 |
| 安全生产法（119条） | [topics/work-safety-law](topics/work-safety-law/INDEX.md) — 6章 |
| 就业促进法（69条） | [topics/employment-promotion-law](topics/employment-promotion-law/INDEX.md) — 4章 |
| 女职工劳动保护特别规定（16条） | [topics/female-worker-protection](topics/female-worker-protection/INDEX.md) — 3章 |

**综合入口**：[劳动者维权行动手册](topics/china-labor-law/worker-playbook.md)（从证据到执行的完整流程）。

**法规原文核对**：[LEGAL-SOURCES.md](../LEGAL-SOURCES.md) — 15 部法律法规及司法解释的官方原文链接（全国人大网/中国政府网/最高人民法院），回答中引用法条时必须附上对应官方链接，便于用户核对原文。

---

## 主题索引（按知识域）

| 知识域 | 模块 | 高频主题 |
|---|---|---|
| 劳动关系基础 | china-labor-law | 就业平等、童工、工时（8h/44h/加班上限）、最低工资、女职工与未成年工特殊保护 |
| 劳动合同 | labor-contract-law | 书面合同、二倍工资、无固定期限、试用期、服务期、解除/终止、N/2N、经济补偿 |
| 社会保险 | social-insurance-law | 五险缴费、待遇、征缴稽核、滞纳金、未缴责任 |
| 争议程序 | labor-dispute-arbitration-law | 仲裁前置、调解、1年时效、一裁终局、开庭举证 |
| 审判口径 | labor-dispute-judicial-interpretations | 受案范围、当事人、终局裁决、撤销/不予执行、解释二新规（2025-09-01） |
| 民事基础 | civil-code | 协议效力与撤销、格式条款、3年时效、人格权（性骚扰/隐私/个人信息）、侵权赔偿 |
| 工伤保险 | work-injury-regulations | 认定三明治、48小时、鉴定十级、停工留薪、待遇梯度、工亡三费、未参保自付 |
| 休息休假 | paid-annual-leave-regulations | 5/10/15天、折算公式、300%补偿、书面放弃、派遣年假 |
| 住房公积金 | housing-fund-regulations | 缴存比例5%–12%、月缴存额公式、提取六情形、不缴强制执行 |

## 版本与时点规则

| 文件 | 施行 | 备注 |
|---|---|---|
| 劳动法 | 1995-01-01 | 2018修正 |
| 劳动合同法 | 2008-01-01 | 2012修正 |
| 调解仲裁法 | 2008-05-01 | 2007通过 |
| 社会保险法 | 2011-07-01 | 2018修正 |
| 民法典 | **2021-01-01** | 此前事实按行为时法（旧九法废止） |
| 解释（一） | 2021-01-01 | 整合旧"四解释" |
| 解释（二） | **2025-09-01** | 与解释一冲突**以（二）为准**；解释一第32条1款（退休返聘=劳务）**已废止** |
| 工伤保险条例 | 2004-01-01 | 2010修订（未完成认定的从新） |
| 年休假条例/办法 | 2008-01-01 / 2008-09-18 | 公历年度 |

## 回答规范（法条引用）

回答用户问题时，凡涉及具体法律法规条文，**必须附上法条原文的官方链接**，以便用户核对原文、增强可信度。

**引用格式**（内联，紧跟条文号）：

```
《中华人民共和国劳动合同法》第四十七条（[查看原文](http://www.npc.gov.cn/zgrdw/npc/xinwen/lfgz/zxfl/2007-06/29/content_368169.htm)）规定，经济补偿按劳动者在本单位工作的年限，每满一年支付一个月工资的标准向劳动者支付。
```

**要点**：

1. **法律全称 + 条文号**：首次引用写全称（如《中华人民共和国劳动合同法》），后续可简称（如《劳动合同法》）
2. **官方链接**：链接来自 [LEGAL-SOURCES.md](../LEGAL-SOURCES.md) 中的官方来源（全国人大网/中国政府网/最高人民法院/人社部），不得使用非官方链接
3. **多个条文**：同一法律内的多个条文可共用一个链接，在条文号后标注；不同法律的条文分别标注各自链接
4. **计算/标准类**：涉及金额、比例、期限等法定标准时，除标注条文链接外，还应注明数据是否为动态值（如工亡补助金、社保缴费基数等，参见 [DYNAMIC-DATA.md](../DYNAMIC-DATA.md)）
5. **例外**：仅作背景提及、未作为结论依据的法条可不附链接；但作为结论核心依据的法条必须附链接

**官方链接速查**（完整清单见 [LEGAL-SOURCES.md](../LEGAL-SOURCES.md)）：

- 劳动法：http://www.npc.gov.cn/npc/c2/c30834/201905/t20190521_296651.html
- 劳动合同法：http://www.npc.gov.cn/zgrdw/npc/xinwen/lfgz/zxfl/2007-06/29/content_368169.htm
- 社会保险法：http://www.npc.gov.cn/zgrdw/npc/xinwen/2019-01/07/content_2070267.htm
- 调解仲裁法：http://www.npc.gov.cn/zgrdw/npc/xinwen/lfgz/zxfl/2007-12/29/content_1387809.htm
- 民法典：http://www.npc.gov.cn/c2/c30834/202006/t20200602_306457.html
- 工伤保险条例：https://www.mohrss.gov.cn/xxgk2020/fdzdgknr/zcfg/fg/202011/t20201103_394950.html
- 年休假条例：https://www.gov.cn/zhengce/zhengceku/2008-03/28/content_6636.htm
- 公积金条例：https://www.gov.cn/gongbao/content/2019/content_5468861.htm（2019版；2026年修订见 LEGAL-SOURCES.md）
- 解释（一）：https://www.court.gov.cn/fabu/xiangqing/282121.html
- 解释（二）：https://www.court.gov.cn/zixun/xiangqing/472691.html

## Scope & Limits

- **特别法优先**：劳动法律体系（劳动法/劳动合同法/社保法/仲裁法）优先于民法典；特别法无规定时参照民法典（如违约金酌减、协议撤销、人格权赔偿）
- **地方口径**：两金标准、停工留薪期目录、双赔/补差、加班费基数、医保退休年限、公积金缴存比例等省级差异，知识库内以 ※ 标注，详见 [LOCAL-PRACTICE.md](../../LOCAL-PRACTICE.md)（含官方查询渠道与回答模板）
- **未覆盖**：民法典物权/婚姻/继承编仅交叉提及；《就业促进法》《女职工劳动保护特别规定》等尚未建模块（相关条文已交叉引用）
- 每个模块的 `cheatsheet.md`（速查表/决策卡/关键数字）与 `patterns.md`（操作流程）是回答计算类与操作类问题的**首选入口**
- **免责声明**：本知识库不构成法律意见；个案结论受证据、鉴定结论与地方实践影响，重大争议请咨询执业律师
