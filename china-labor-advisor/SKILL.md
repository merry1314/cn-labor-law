---
name: china-labor-advisor
description: "中国劳动法顾问：劳动者与用人单位的劳动用工、社会保障全领域知识库（18部法律法规全文结构化+1个应用型审查模块）。覆盖：劳动合同（签订/二倍工资/无固定期限/试用期/解除终止/经济补偿N与2N/赔偿金）、裁员辞退开除、欠薪克扣工资、加班费、调岗降薪、劳务派遣、竞业限制/保密/服务期违约金、离职协议、五险一金（养老/医疗/失业/生育/工伤）、社保缴费比例与基数（交多少）、住房公积金（缴存比例/提取/贷款/单位不缴）、工伤认定与伤残待遇（48小时/停工留薪期/工亡三费）、年休假（折算/300%补偿）、社保补缴、职场性骚扰、个人信息/隐私、劳动仲裁诉讼（时效/一裁终局/证据举证）、离职协议撤销、退休返聘/超龄劳动者、农民工工资支付、劳动保障监察、职业病防治、安全生产、就业歧视（性别/乙肝/残疾/户籍）、女职工三期保护（产假98天/生育津贴/哺乳期）、残疾人保障、工会、**审查公司规章制度/员工手册是否合法**（上传制度文件→逐条定级→指出违法与损害权益条款→给应对策略）。用于回答'该不该赔/赔多少/交多少/怎么告/告谁/法院怎么判/制度合不合法'类劳动法律问题；2021-2025最新司法解释口径。"
---

<!-- argument-hint: [场景、主题、法条或章号，如"被裁员"、"二倍工资"、"工伤认定"、"第38条"、"竞业限制"] -->

# 中国劳动法顾问（china-labor-advisor）

**知识库**: 18 部法律法规全文结构化（劳动法/劳动合同法/社会保险法/调解仲裁法/最高法解释一+二/民法典/工伤保险条例/年休假条例+办法/住房公积金管理条例/保障农民工工资支付条例/劳动保障监察条例/职业病防治法/安全生产法/就业促进法/女职工劳动保护特别规定/残疾人保障法/工会法/超龄劳动者基本权益保障暂行规定），共 97 章节文件 | **基准日**: 2026-10-16

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
| 公积金该交多少 / 单位不缴公积金 | [公积金条例](topics/housing-fund-regulations/INDEX.md) ch01（16、18条缴存公式与比例） | ch02（39、40条投诉与强制执行）、[社保法 cheatsheet](topics/social-insurance-law/cheatsheet.md)（五险比例速查） |
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
| 退休/返聘/超龄劳动者 | [超龄劳动者规定](topics/elderly-worker-protection/INDEX.md)（四项权益+争议双通道）、[司法解释](topics/labor-dispute-judicial-interpretations/INDEX.md) ch03（32条已废止） | [社保法](topics/social-insurance-law/INDEX.md) ch02（养老金）、[民法典](topics/civil-code/INDEX.md) ch05（1192条劳务） |
| 残疾人就业被拒/歧视/按比例就业 | [残疾人保障法](topics/disabled-persons-protection-law/INDEX.md) ch02 | [就业促进法](topics/employment-promotion-law/INDEX.md) ch01（29条按比例）、ch04（62条歧视直接起诉） |
| 想组建工会/阻挠建会/集体合同/工会干部被辞 | [工会法](topics/trade-union-law/INDEX.md) ch01–ch02 | [劳动合同法](topics/labor-contract-law/INDEX.md) ch04（43条工会意见）、[司法解释](topics/labor-dispute-judicial-interpretations/INDEX.md) ch06（工会法19条任期保护） |

## 法条定位表 / 主题索引 / 版本规则（外置）

法条类问题（"劳动法第几条"）与版本时点问题，先读 **[reference.md](reference.md)**：含完整**法条定位表**（18 部法律→模块→章数）、**主题索引**（按知识域）、**版本与时点规则**（施行日期、2021/2025 分界、**解释二与解释一冲突以（二）为准**、解释一 32 条 1 款已废止）、**官方链接速查**。定位后读对应模块 INDEX.md。

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

**官方链接**：18 部法律法规及司法解释的官方原文链接（全国人大网/中国政府网/最高人民法院）统一见 [LEGAL-SOURCES.md](../LEGAL-SOURCES.md) 与 [reference.md](reference.md) 官方链接速查，不得使用非官方链接。

## Scope & Limits

- **特别法优先**：劳动法律体系（劳动法/劳动合同法/社保法/仲裁法）优先于民法典；特别法无规定时参照民法典（如违约金酌减、协议撤销、人格权赔偿）
- **地方口径**：两金标准、停工留薪期目录、双赔/补差、加班费基数、医保退休年限、公积金缴存比例等省级差异，知识库内以 ※ 标注，详见 [LOCAL-PRACTICE.md](../LOCAL-PRACTICE.md)（含官方查询渠道与回答模板）
- **未展开**：民法典物权/婚姻/继承编仅整理劳动关联点（见 civil-code 模块 Core Frameworks 第6点），未逐条展开
- 每个模块的 `cheatsheet.md`（速查表/决策卡/关键数字）与 `patterns.md`（操作流程）是回答计算类与操作类问题的**首选入口**
- **免责声明**：本知识库不构成法律意见；个案结论受证据、鉴定结论与地方实践影响，重大争议请咨询执业律师
