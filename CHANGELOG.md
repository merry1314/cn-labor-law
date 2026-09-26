# 变更日志（CHANGELOG）

本文件记录 `cn-labor-law` 知识库的版本变更。遵循「可追溯、可回滚」原则，每次结构调整、口径修正、数据更新均应在此登记。

格式参考 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.1.0/)。

---

## [Unreleased]

### 优化
- **口径一致性**：修正 README.md 中「3 支撑文件」描述，明确各模块支撑文件为 glossary/patterns/cheatsheet，china-labor-law 另含 worker-playbook
- **动态数据可维护性**：为工亡补助金年度数值（3 处）增加「数据基准日」标记与集中更新指引
- **新增 `DYNAMIC-DATA.md`**：集中登记所有需年度更新的动态数据点、分布位置与更新步骤
- **新增 `CHANGELOG.md`**：本文件，建立版本变更记录机制
- **新增 `consistency-check.go`**：口径一致性校验脚本，自动核验模块数/章节数/文件引用

---

## 历史提交摘要（从 git log 整理）

| 提交 | 说明 |
|---|---|
| `08343b0` | feat(knowledge): 建成中国劳动法顾问技能库（china-labor-advisor） |
| `da5bf8e` | fix(knowledge): 修正司法解释错引并完成场景审查 |
| `63357c9` | docs(license): 添加 MIT 开源许可协议 |
| `92d9d91` | feat(knowledge): 新增住房公积金模块并补全社保缴费比例速查 |
| `d4a27d2` | fix(knowledge): 修正工亡补助金数据与 README 口径错误 |
| `40d8502` | feat(knowledge): 新增 policy-review 规章制度合规审查子模块 |
| `674fe1f` | docs: 免责声明去除生成工具表述 |
| `5515438` | docs(readme): 补充高频现实场景使用示例 |
| `64ec77d` | 新增 git 仓库忽略文件 |

### 已知口径变更记录

- **住房公积金模块**（commit `92d9d91`）：在初始构建后新增，导致 SKILL.md 头部统计（8 部/57 章）与实际（9 部/59 章）不一致，后经 `d4a27d2` 部分修正
- **工亡补助金数据**（commit `d4a27d2`）：曾修正数据与 README 口径错误，但动态数值仍硬编码于正文，本次优化引入年度更新机制
- **policy-review 模块**（commit `40d8502`）：新增应用型审查模块，使模块总数达 10 个（9 法律模块 + 1 应用模块）
