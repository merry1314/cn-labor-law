// consistency-check.go —— cn-labor-law 知识库口径一致性校验脚本
//
// 用法：go run consistency-check.go
//
// 校验项：
//   1. 模块总数 = 9 法律模块 + 1 policy-review = 10
//   2. 章节总数 = 59
//   3. 每个法律模块含 INDEX.md + chapters/ + glossary.md + patterns.md + cheatsheet.md
//   4. china-labor-law 另含 worker-playbook.md
//   5. SKILL.md 与 README.md 头部统计口径一致（9 部 / 59 章）
//   6. 工亡补助金动态数据 3 处均含「数据基准日」标记
//   7. README 与 SKILL.md 场景路由表一致（欠薪→ch05、加班费→ch04）
//   8. SKILL.md 五级定级含 🟢合规
//   9. DYNAMIC-DATA.md 已公布数据不含「预估」字样
//  10. 法条引用挂名正确（如第83条三款不挂在劳动争议调解仲裁法名下）
//
// 退出码：0 = 全部通过；非 0 = 存在不一致
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	topicsDir   = "china-labor-advisor/topics"
	skillPath   = "china-labor-advisor/SKILL.md"
	readmePath  = "README.md"
)

// 9 个法律知识模块
var lawModules = []string{
	"china-labor-law",
	"labor-contract-law",
	"social-insurance-law",
	"labor-dispute-arbitration-law",
	"labor-dispute-judicial-interpretations",
	"civil-code",
	"work-injury-regulations",
	"paid-annual-leave-regulations",
	"housing-fund-regulations",
}

// 每个模块预期章节数
var expectedChapters = map[string]int{
	"china-labor-law":                       13,
	"labor-contract-law":                    8,
	"social-insurance-law":                  12,
	"labor-dispute-arbitration-law":         5,
	"labor-dispute-judicial-interpretations": 8,
	"civil-code":                            5,
	"work-injury-regulations":               4,
	"paid-annual-leave-regulations":         2,
	"housing-fund-regulations":              2,
}

const (
	expectedLawModules    = 9
	expectedTotalChapters = 59
)

type check struct {
	name string
	ok   bool
	msg  string
}

func main() {
	var results []check

	results = append(results, checkModuleCount())
	results = append(results, checkChapterCount())
	results = append(results, checkModuleFiles())
	results = append(results, checkSkillReadmeConsistency())
	results = append(results, checkDynamicDataMarkers())
	results = append(results, checkRoutingConsistency())
	results = append(results, checkFiveLevelRating())
	results = append(results, checkDynamicDataNoEstimate())
	results = append(results, checkArticleAttribution())

	pass := 0
	for _, r := range results {
		status := "PASS"
		if !r.ok {
			status = "FAIL"
		}
		fmt.Printf("[%s] %s\n", status, r.name)
		if r.msg != "" {
			fmt.Printf("       %s\n", r.msg)
		}
		if r.ok {
			pass++
		}
	}

	fmt.Printf("\n%d/%d 项校验通过\n", pass, len(results))
	if pass < len(results) {
		os.Exit(1)
	}
}

// 1. 模块总数校验
func checkModuleCount() check {
	c := check{name: "模块总数 = 9 法律模块 + 1 policy-review"}

	entries, err := os.ReadDir(topicsDir)
	if err != nil {
		c.ok = false
		c.msg = fmt.Sprintf("无法读取 %s: %v", topicsDir, err)
		return c
	}

	// 统计子目录数
	dirCount := 0
	for _, e := range entries {
		if e.IsDir() {
			dirCount++
		}
	}

	// 应有 9 法律模块 + policy-review = 10
	if dirCount == expectedLawModules+1 {
		c.ok = true
	} else {
		c.ok = false
		c.msg = fmt.Sprintf("实际 %d 个模块目录，预期 %d 个", dirCount, expectedLawModules+1)
	}
	return c
}

// 2. 章节总数校验
func checkChapterCount() check {
	c := check{name: "章节总数 = 59"}

	total := 0
	for _, mod := range lawModules {
		chDir := filepath.Join(topicsDir, mod, "chapters")
		entries, err := os.ReadDir(chDir)
		if err != nil {
			c.ok = false
			c.msg = fmt.Sprintf("无法读取 %s: %v", chDir, err)
			return c
		}
		count := 0
		for _, e := range entries {
			if !e.IsDir() && strings.HasPrefix(e.Name(), "ch") && strings.HasSuffix(e.Name(), ".md") {
				count++
			}
		}
		exp, ok := expectedChapters[mod]
		if !ok {
			c.ok = false
			c.msg = fmt.Sprintf("模块 %s 无预期章节数配置", mod)
			return c
		}
		if count != exp {
			c.ok = false
			c.msg = fmt.Sprintf("模块 %s 实际 %d 章，预期 %d 章", mod, count, exp)
			return c
		}
		total += count
	}

	// policy-review 有 10 个 review-guides（非章节，不计入 59）
	prDir := filepath.Join(topicsDir, "policy-review", "review-guides")
	if entries, err := os.ReadDir(prDir); err == nil {
		prCount := 0
		for _, e := range entries {
			if !e.IsDir() && strings.HasSuffix(e.Name(), ".md") {
				prCount++
			}
		}
		if prCount != 10 {
			c.ok = false
			c.msg = fmt.Sprintf("policy-review 实际 %d 个审查主题，预期 10 个", prCount)
			return c
		}
	}

	if total == expectedTotalChapters {
		c.ok = true
	} else {
		c.ok = false
		c.msg = fmt.Sprintf("实际总章节 %d，预期 %d", total, expectedTotalChapters)
	}
	return c
}

// 3. 各模块必备文件校验
func checkModuleFiles() check {
	c := check{name: "各模块必备文件结构"}

	requiredFiles := []string{"INDEX.md", "glossary.md", "patterns.md", "cheatsheet.md"}
	var missing []string

	for _, mod := range lawModules {
		modDir := filepath.Join(topicsDir, mod)
		// chapters 目录
		if _, err := os.Stat(filepath.Join(modDir, "chapters")); os.IsNotExist(err) {
			missing = append(missing, fmt.Sprintf("%s/chapters/", mod))
		}
		for _, f := range requiredFiles {
			if _, err := os.Stat(filepath.Join(modDir, f)); os.IsNotExist(err) {
				missing = append(missing, fmt.Sprintf("%s/%s", mod, f))
			}
		}
	}

	// china-labor-law 额外的 worker-playbook.md
	if _, err := os.Stat(filepath.Join(topicsDir, "china-labor-law", "worker-playbook.md")); os.IsNotExist(err) {
		missing = append(missing, "china-labor-law/worker-playbook.md")
	}

	// policy-review 必备文件
	prIndex := filepath.Join(topicsDir, "policy-review", "INDEX.md")
	if _, err := os.Stat(prIndex); os.IsNotExist(err) {
		missing = append(missing, "policy-review/INDEX.md")
	}

	if len(missing) == 0 {
		c.ok = true
	} else {
		c.ok = false
		c.msg = "缺失文件：" + strings.Join(missing, ", ")
	}
	return c
}

// 4. SKILL.md 与 README.md 口径一致性
func checkSkillReadmeConsistency() check {
	c := check{name: "SKILL.md / README.md 头部统计口径一致（9 部 / 59 章）"}

	skill, err := os.ReadFile(skillPath)
	if err != nil {
		c.ok = false
		c.msg = fmt.Sprintf("无法读取 %s: %v", skillPath, err)
		return c
	}
	readme, err := os.ReadFile(readmePath)
	if er := err; er != nil {
		c.ok = false
		c.msg = fmt.Sprintf("无法读取 %s: %v", readmePath, er)
		return c
	}

	skillStr := string(skill)
	readmeStr := string(readme)

	// SKILL.md 应含 "9 部" 和 "59 章节"
	reSkill9 := regexp.MustCompile(`9 部法律法规`)
	reSkill59 := regexp.MustCompile(`59 章节文件`)
	if !reSkill9.MatchString(skillStr) {
		c.ok = false
		c.msg = "SKILL.md 未找到 \"9 部法律法规\""
		return c
	}
	if !reSkill59.MatchString(skillStr) {
		c.ok = false
		c.msg = "SKILL.md 未找到 \"59 章节文件\""
		return c
	}

	// README.md 应含 "9 部" 和 "59 个章节文件"
	reReadme9 := regexp.MustCompile(`9 部法律法规`)
	reReadme59 := regexp.MustCompile(`59 个章节文件`)
	if !reReadme9.MatchString(readmeStr) {
		c.ok = false
		c.msg = "README.md 未找到 \"9 部法律法规\""
		return c
	}
	if !reReadme59.MatchString(readmeStr) {
		c.ok = false
		c.msg = "README.md 未找到 \"59 个章节文件\""
		return c
	}

	// 不应再出现旧的 "8 部" 或 "57 章节"
	if strings.Contains(skillStr, "8 部法律法规") || strings.Contains(skillStr, "57 章节") {
		c.ok = false
		c.msg = "SKILL.md 仍含旧口径 \"8 部\" 或 \"57 章节\""
		return c
	}

	c.ok = true
	return c
}

// 5. 工亡补助金动态数据标记校验
func checkDynamicDataMarkers() check {
	c := check{name: "工亡补助金动态数据 3 处均含「数据基准日」标记"}

	files := []string{
		filepath.Join(topicsDir, "work-injury-regulations", "cheatsheet.md"),
		filepath.Join(topicsDir, "work-injury-regulations", "chapters", "ch03-assessment-benefits.md"),
		filepath.Join(topicsDir, "work-injury-regulations", "patterns.md"),
	}

	re := regexp.MustCompile(`数据基准日`)
	var missing []string

	for _, f := range files {
		content, err := os.ReadFile(f)
		if err != nil {
			missing = append(missing, fmt.Sprintf("%s（读取失败）", f))
			continue
		}
		if !re.MatchString(string(content)) {
			missing = append(missing, f)
		}
	}

	if len(missing) == 0 {
		c.ok = true
	} else {
		c.ok = false
		c.msg = "缺少「数据基准日」标记的文件：" + strings.Join(missing, ", ")
	}
	return c
}

// 7. README 与 SKILL.md 场景路由表一致性校验
//    劳动法 ch04=工作时间（含加班费第44条），ch05=工资（含欠薪）
func checkRoutingConsistency() check {
	c := check{name: "README 与 SKILL.md 路由表一致（欠薪→ch05、加班费→ch04）"}

	readme, err := os.ReadFile(readmePath)
	if err != nil {
		c.ok = false
		c.msg = fmt.Sprintf("无法读取 %s: %v", readmePath, err)
		return c
	}
	readmeStr := string(readme)

	// README 欠薪行应指向 ch05（工资），而非 ch04
	reBackpayWrong := regexp.MustCompile(`欠薪.*china-labor-law ch04[^0-9]`)
	reBackpayRight := regexp.MustCompile(`欠薪.*china-labor-law ch05`)
	if reBackpayWrong.MatchString(readmeStr) {
		c.ok = false
		c.msg = "README 欠薪路由指向 ch04（应为 ch05 工资章）"
		return c
	}
	if !reBackpayRight.MatchString(readmeStr) {
		c.ok = false
		c.msg = "README 欠薪路由未指向 ch05"
		return c
	}

	// README 加班费行应指向 ch04（工时），而非 ch03
	reOvertimeWrong := regexp.MustCompile(`加班费.*china-labor-law ch03[^0-9]`)
	reOvertimeRight := regexp.MustCompile(`加班费.*china-labor-law ch04`)
	if reOvertimeWrong.MatchString(readmeStr) {
		c.ok = false
		c.msg = "README 加班费路由指向 ch03（应为 ch04 工时章，第44条）"
		return c
	}
	if !reOvertimeRight.MatchString(readmeStr) {
		c.ok = false
		c.msg = "README 加班费路由未指向 ch04"
		return c
	}

	c.ok = true
	return c
}

// 8. SKILL.md 五级定级完整性校验（应含 🟢合规）
func checkFiveLevelRating() check {
	c := check{name: "SKILL.md 五级定级含 🟢合规"}

	skill, err := os.ReadFile(skillPath)
	if err != nil {
		c.ok = false
		c.msg = fmt.Sprintf("无法读取 %s: %v", skillPath, err)
		return c
	}
	skillStr := string(skill)

	// 五级定级行应同时包含 🔴🟠🟡🔵🟢 五级
	ratingLine := regexp.MustCompile(`五级定级.*`)
	match := ratingLine.FindString(skillStr)
	if match == "" {
		c.ok = false
		c.msg = "SKILL.md 未找到五级定级描述"
		return c
	}

	required := []string{"🔴", "🟠", "🟡", "🔵", "🟢"}
	var missing []string
	for _, level := range required {
		if !strings.Contains(match, level) {
			missing = append(missing, level)
		}
	}

	if len(missing) == 0 {
		c.ok = true
	} else {
		c.ok = false
		c.msg = "五级定级缺少：" + strings.Join(missing, " ")
	}
	return c
}

// 9. DYNAMIC-DATA.md 动态数据状态校验：已公布数据不得含「预估」字样
func checkDynamicDataNoEstimate() check {
	c := check{name: "DYNAMIC-DATA.md 已公布数据不含「预估」字样"}

	data, err := os.ReadFile("DYNAMIC-DATA.md")
	if err != nil {
		c.ok = false
		c.msg = fmt.Sprintf("无法读取 DYNAMIC-DATA.md: %v", err)
		return c
	}
	dataStr := string(data)

	// 2026 年度工亡补助金已由国家统计局 2026-01-19 公布，不得标注为预估
	if strings.Contains(dataStr, "预估") {
		c.ok = false
		c.msg = "DYNAMIC-DATA.md 含「预估」字样——已公布数据不应标注为预估"
		return c
	}

	c.ok = true
	return c
}

// 10. 法条引用挂名校验：第83条三款应归属社会保险法，非劳动争议调解仲裁法
func checkArticleAttribution() check {
	c := check{name: "法条引用挂名正确（第83条三款不挂在劳动争议调解仲裁法名下）"}

	ch10 := filepath.Join(topicsDir, "social-insurance-law", "chapters", "ch10-supervision.md")
	content, err := os.ReadFile(ch10)
	if err != nil {
		c.ok = false
		c.msg = fmt.Sprintf("无法读取 %s: %v", ch10, err)
		return c
	}
	contentStr := string(content)

	// 不得出现「劳动争议调解仲裁法: 第83条」之类的错误挂名
	wrongPattern := regexp.MustCompile(`劳动争议调解仲裁法.*第\s*83\s*条`)
	if wrongPattern.MatchString(contentStr) {
		c.ok = false
		c.msg = "ch10-supervision.md 把第83条挂在劳动争议调解仲裁法名下（实为社会保险法第83条）"
		return c
	}

	// 应明确归属社会保险法第83条
	rightPattern := regexp.MustCompile(`社会保险法.*第\s*83\s*条`)
	if !rightPattern.MatchString(contentStr) {
		c.ok = false
		c.msg = "ch10-supervision.md 未明确第83条归属社会保险法"
		return c
	}

	c.ok = true
	return c
}
