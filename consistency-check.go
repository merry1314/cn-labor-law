// consistency-check.go —— cn-labor-law 知识库口径一致性校验脚本（规则级）
//
// 用法：go run consistency-check.go
//
// 校验项（规则级，非特定缺陷硬编码）：
//  1. 模块总数 = 15 法律模块 + 1 policy-review = 16
//  2. 章节总数 = 87，各模块章节数匹配
//  3. 每个法律模块含 INDEX.md + chapters/ + glossary.md + patterns.md + cheatsheet.md
//  4. SKILL.md 与 README.md 头部统计口径一致（15 部 / 87 章）
//  5. 工亡补助金动态数据 3 处均含「数据基准日」标记
//  6. DYNAMIC-DATA 表格算术一致性：适用年度=收入年度+1，补助金=收入×20
//  7. 通用路由校验：所有 模块/chNN 引用的章节文件必须存在
//  8. SKILL.md 五级定级含 🟢合规
//  9. 法条引用越界扫描：条号不超过各法条数上限
//  10. 数据基准日新鲜度：距今不超过 13 个月
//  11. 条号体例一致性：正文条号统一为阿拉伯数字
//  12. 动态数据三处数值一致性：工亡补助金绝对金额在 3 处文件中一致
//
// 退出码：0 = 全部通过；非 0 = 存在不一致
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	topicsDir  = "china-labor-advisor/topics"
	skillPath  = "china-labor-advisor/SKILL.md"
	readmePath = "README.md"
)

// 17 个法律知识模块
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
	"migrant-worker-wage-regulations",
	"labor-inspection-regulations",
	"occupational-disease-prevention-law",
	"work-safety-law",
	"employment-promotion-law",
	"female-worker-protection",
	"disabled-persons-protection-law",
	"trade-union-law",
}

// 每个模块预期章节数
var expectedChapters = map[string]int{
	"china-labor-law":                        13,
	"labor-contract-law":                     8,
	"social-insurance-law":                   12,
	"labor-dispute-arbitration-law":          5,
	"labor-dispute-judicial-interpretations": 8,
	"civil-code":                             5,
	"work-injury-regulations":                4,
	"paid-annual-leave-regulations":          2,
	"housing-fund-regulations":               2,
	"migrant-worker-wage-regulations":        6,
	"labor-inspection-regulations":           3,
	"occupational-disease-prevention-law":    6,
	"work-safety-law":                        6,
	"employment-promotion-law":               4,
	"female-worker-protection":               3,
	"disabled-persons-protection-law":        4,
	"trade-union-law":                        4,
}

// 各法律条文数上限（用于法条号越界扫描）
var articleLimits = map[string]int{
	"劳动法":         107,
	"劳动合同法":       98,
	"社会保险法":       98,
	"调解仲裁法":       54,
	"劳动争议调解仲裁法":   54,
	"民法典":         1260,
	"工伤保险条例":      67,
	"年休假条例":       10,
	"年休假实施办法":     19,
	"公积金条例":       50,
	"住房公积金管理条例":   50,
	"解释一":         54,
	"解释二":         21,
	"农民工工资条例":     64,
	"保障农民工工资支付条例": 64,
	"劳动监察条例":      36,
	"劳动保障监察条例":    36,
	"职业病防治法":      87,
	"安全生产法":       119,
	"就业促进法":       69,
	"女职工劳动保护特别规定": 16,
	"女职工保护规定":     16,
	"残疾人保障法":      68,
	"工会法":         58,
}

const (
	expectedLawModules    = 17
	expectedTotalChapters = 95
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
	results = append(results, checkDynamicDataArithmetic())
	results = append(results, checkRoutingGeneral())
	results = append(results, checkFiveLevelRating())
	results = append(results, checkArticleBounds())
	results = append(results, checkDataFreshness())
	results = append(results, checkArticleStyle())
	results = append(results, checkDynamicDataConsistency())

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
	c := check{name: "模块总数 = 17 法律模块 + 1 policy-review"}

	entries, err := os.ReadDir(topicsDir)
	if err != nil {
		c.ok = false
		c.msg = fmt.Sprintf("无法读取 %s: %v", topicsDir, err)
		return c
	}

	dirCount := 0
	for _, e := range entries {
		if e.IsDir() {
			dirCount++
		}
	}

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
	c := check{name: "章节总数 = 95"}

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
		if _, err := os.Stat(filepath.Join(modDir, "chapters")); os.IsNotExist(err) {
			missing = append(missing, fmt.Sprintf("%s/chapters/", mod))
		}
		for _, f := range requiredFiles {
			if _, err := os.Stat(filepath.Join(modDir, f)); os.IsNotExist(err) {
				missing = append(missing, fmt.Sprintf("%s/%s", mod, f))
			}
		}
	}

	if _, err := os.Stat(filepath.Join(topicsDir, "china-labor-law", "worker-playbook.md")); os.IsNotExist(err) {
		missing = append(missing, "china-labor-law/worker-playbook.md")
	}

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
	c := check{name: "SKILL.md / README.md 头部统计口径一致（17 部 / 95 章）"}

	skill, err := os.ReadFile(skillPath)
	if err != nil {
		c.ok = false
		c.msg = fmt.Sprintf("无法读取 %s: %v", skillPath, err)
		return c
	}
	readme, err := os.ReadFile(readmePath)
	if err != nil {
		c.ok = false
		c.msg = fmt.Sprintf("无法读取 %s: %v", readmePath, err)
		return c
	}

	skillStr := string(skill)
	readmeStr := string(readme)

	reSkill9 := regexp.MustCompile(`17 部法律法规`)
	reSkill59 := regexp.MustCompile(`95 章节文件`)
	if !reSkill9.MatchString(skillStr) {
		c.ok = false
		c.msg = "SKILL.md 未找到 \"17 部法律法规\""
		return c
	}
	if !reSkill59.MatchString(skillStr) {
		c.ok = false
		c.msg = "SKILL.md 未找到 \"95 章节文件\""
		return c
	}

	reReadme9 := regexp.MustCompile(`17 部法律法规`)
	reReadme59 := regexp.MustCompile(`95 个章节文件`)
	if !reReadme9.MatchString(readmeStr) {
		c.ok = false
		c.msg = "README.md 未找到 \"17 部法律法规\""
		return c
	}
	if !reReadme59.MatchString(readmeStr) {
		c.ok = false
		c.msg = "README.md 未找到 \"95 个章节文件\""
		return c
	}

	if strings.Contains(skillStr, "9 部法律法规") || strings.Contains(skillStr, "13 部法律法规") || strings.Contains(skillStr, "15 部法律法规") || strings.Contains(skillStr, "59 章节") || strings.Contains(skillStr, "80 章节") || strings.Contains(skillStr, "87 章节") {
		c.ok = false
		c.msg = "SKILL.md 仍含旧口径 \"9/13/15 部\" 或 \"59/80/87 章节\""
		return c
	}

	if strings.Contains(readmeStr, "9 部法律法规") || strings.Contains(readmeStr, "13 部法律法规") || strings.Contains(readmeStr, "15 部法律法规") || strings.Contains(readmeStr, "59 章节") || strings.Contains(readmeStr, "80 章节") || strings.Contains(readmeStr, "87 章节") {
		c.ok = false
		c.msg = "README.md 仍含旧口径 \"9/13/15 部\" 或 \"59/80/87 章节\""
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

//  6. DYNAMIC-DATA 表格算术一致性校验（规则级）
//     断言：适用年度 = 收入年度 + 1；补助金 = 收入 × 20
func checkDynamicDataArithmetic() check {
	c := check{name: "DYNAMIC-DATA 表格算术一致（适用年度=收入年度+1，补助金=收入×20）"}

	data, err := os.ReadFile("DYNAMIC-DATA.md")
	if err != nil {
		c.ok = false
		c.msg = fmt.Sprintf("无法读取 DYNAMIC-DATA.md: %v", err)
		return c
	}

	// 匹配表格数据行：| 收入年度 | 收入 | 适用年度 | 补助金 | 状态 |
	// 如：| 2023 | 51,821 元 | 2024 | 1,036,420 元（≈103.64 万） | 已公布... |
	lineRe := regexp.MustCompile(`^\|\s*(\d{4})\s*\|\s*([\d,]+)\s*元\s*\|\s*(\d{4})\s*\|\s*([\d,]+)\s*元`)
	lines := strings.Split(string(data), "\n")

	var errors []string
	for _, line := range lines {
		m := lineRe.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		incomeYear, _ := strconv.Atoi(m[1])
		incomeStr := strings.ReplaceAll(m[2], ",", "")
		income, _ := strconv.Atoi(incomeStr)
		applicableYear, _ := strconv.Atoi(m[3])
		grantStr := strings.ReplaceAll(m[4], ",", "")
		grant, _ := strconv.Atoi(grantStr)

		// 断言1：适用年度 = 收入年度 + 1
		if applicableYear != incomeYear+1 {
			errors = append(errors, fmt.Sprintf("收入年度%d的适用年度应为%d，实际为%d", incomeYear, incomeYear+1, applicableYear))
		}
		// 断言2：补助金 = 收入 × 20
		expectedGrant := income * 20
		if grant != expectedGrant {
			errors = append(errors, fmt.Sprintf("收入年度%d补助金应为%d（%d×20），实际为%d", incomeYear, expectedGrant, income, grant))
		}
	}

	if len(errors) == 0 {
		c.ok = true
	} else {
		c.ok = false
		c.msg = strings.Join(errors, "; ")
	}
	return c
}

//  7. 通用路由校验（规则级）
//     扫描 README.md 和 SKILL.md 中所有 模块名/chNN 引用，断言对应章节文件存在
func checkRoutingGeneral() check {
	c := check{name: "通用路由校验：所有 模块/chNN 引用的章节文件必须存在"}

	files := []string{readmePath, skillPath}
	// 匹配形如 module-name ch01 或 module-name/ch01 的引用
	refRe := regexp.MustCompile(`([a-z-]+)\s*ch(\d{2})`)

	var missing []string
	seen := map[string]bool{}

	for _, f := range files {
		content, err := os.ReadFile(f)
		if err != nil {
			continue
		}
		matches := refRe.FindAllStringSubmatch(string(content), -1)
		for _, m := range matches {
			mod := m[1]
			chNum := m[2]
			key := mod + "/ch" + chNum
			if seen[key] {
				continue
			}
			seen[key] = true

			// 只检查已知模块
			isKnown := false
			for _, lm := range lawModules {
				if mod == lm || strings.Contains(lm, mod) {
					isKnown = true
					// 优先用完整模块名
					if lm == mod || strings.HasPrefix(lm, mod) {
						mod = lm
					}
					break
				}
			}
			if !isKnown {
				continue
			}

			chaptersDir := filepath.Join(topicsDir, mod, "chapters")
			entries, err := os.ReadDir(chaptersDir)
			if err != nil {
				continue
			}
			found := false
			for _, e := range entries {
				if strings.HasPrefix(e.Name(), "ch"+chNum) && strings.HasSuffix(e.Name(), ".md") {
					found = true
					break
				}
			}
			if !found {
				missing = append(missing, key)
			}
		}
	}

	if len(missing) == 0 {
		c.ok = true
	} else {
		c.ok = false
		c.msg = "引用了不存在的章节：" + strings.Join(missing, ", ")
	}
	return c
}

// 8. SKILL.md 五级定级完整性校验
func checkFiveLevelRating() check {
	c := check{name: "SKILL.md 五级定级含 🟢合规"}

	skill, err := os.ReadFile(skillPath)
	if err != nil {
		c.ok = false
		c.msg = fmt.Sprintf("无法读取 %s: %v", skillPath, err)
		return c
	}
	skillStr := string(skill)

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

//  9. 法条引用越界扫描（规则级）
//     扫描全库"第X条"引用，断言条号不超过各法条数上限
func checkArticleBounds() check {
	c := check{name: "法条引用越界扫描：条号不超过各法条数上限"}

	// 匹配"法名 第X条"或"法名第X条"，支持阿拉伯数字
	refRe := regexp.MustCompile(`(劳动法|劳动合同法|社会保险法|调解仲裁法|劳动争议调解仲裁法|民法典|工伤保险条例|年休假条例|年休假实施办法|公积金条例|住房公积金管理条例|解释一|解释二)\s*第\s*(\d+)\s*条`)

	var violations []string

	// 扫描所有 .md 文件
	err := filepath.Walk(topicsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		matches := refRe.FindAllStringSubmatch(string(content), -1)
		for _, m := range matches {
			lawName := m[1]
			num, err := strconv.Atoi(m[2])
			if err != nil {
				continue
			}
			limit, ok := articleLimits[lawName]
			if !ok {
				continue
			}
			if num > limit {
				rel, _ := filepath.Rel(".", path)
				violations = append(violations, fmt.Sprintf("%s: %s第%d条（上限%d条）", rel, lawName, num, limit))
			}
		}
		return nil
	})
	if err != nil {
		c.ok = false
		c.msg = fmt.Sprintf("扫描失败: %v", err)
		return c
	}

	if len(violations) == 0 {
		c.ok = true
	} else {
		c.ok = false
		// 只显示前 5 条，避免输出过长
		if len(violations) > 5 {
			c.msg = strings.Join(violations[:5], "; ") + fmt.Sprintf("... 共 %d 处", len(violations))
		} else {
			c.msg = strings.Join(violations, "; ")
		}
	}
	return c
}

//  10. 数据基准日新鲜度校验
//     数据基准日距今超过 13 个月则告警
func checkDataFreshness() check {
	c := check{name: "数据基准日新鲜度：距今不超过 13 个月"}

	data, err := os.ReadFile("DYNAMIC-DATA.md")
	if err != nil {
		c.ok = false
		c.msg = fmt.Sprintf("无法读取 DYNAMIC-DATA.md: %v", err)
		return c
	}

	dateRe := regexp.MustCompile(`数据基准日\s*(\d{4})-(\d{2})-(\d{2})`)
	m := dateRe.FindStringSubmatch(string(data))
	if m == nil {
		c.ok = false
		c.msg = "DYNAMIC-DATA.md 未找到数据基准日"
		return c
	}

	year, _ := strconv.Atoi(m[1])
	month, _ := strconv.Atoi(m[2])
	day, _ := strconv.Atoi(m[3])
	baseDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

	// 13 个月 ≈ 395 天
	threshold := baseDate.AddDate(0, 13, 0)
	now := time.Now()

	if now.Before(threshold) {
		c.ok = true
	} else {
		c.ok = false
		c.msg = fmt.Sprintf("数据基准日 %s 距今已超过 13 个月，建议更新动态数据", baseDate.Format("2006-01-02"))
	}
	return c
}

//  11. 条号体例一致性校验
//     检查是否存在汉字条号（如"第三十六条"），应统一为阿拉伯数字（"第36条"）
func checkArticleStyle() check {
	c := check{name: "条号体例一致性：不使用汉字条号（统一阿拉伯数字）"}

	// 匹配"第" + 汉字数字 + "条"
	cnNumRe := regexp.MustCompile(`第[一二三四五六七八九十百千零两]+条`)

	var violations []string

	err := filepath.Walk(topicsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		if cnNumRe.MatchString(string(content)) {
			rel, _ := filepath.Rel(".", path)
			violations = append(violations, rel)
		}
		return nil
	})
	if err != nil {
		c.ok = false
		c.msg = fmt.Sprintf("扫描失败: %v", err)
		return c
	}

	if len(violations) == 0 {
		c.ok = true
	} else {
		c.ok = false
		c.msg = fmt.Sprintf("以下文件含汉字条号，应统一为阿拉伯数字：%s（共 %d 个文件）", strings.Join(violations[:5], ", "), len(violations))
	}
	return c
}

//  12. 动态数据三处数值一致性校验
//     工亡补助金绝对金额在 cheatsheet、ch03、patterns 三处必须一致
func checkDynamicDataConsistency() check {
	c := check{name: "动态数据三处数值一致：工亡补助金绝对金额在 3 处文件中一致"}

	files := []string{
		filepath.Join(topicsDir, "work-injury-regulations", "cheatsheet.md"),
		filepath.Join(topicsDir, "work-injury-regulations", "chapters", "ch03-assessment-benefits.md"),
		filepath.Join(topicsDir, "work-injury-regulations", "patterns.md"),
	}

	// 提取所有形如 1,036,420 的金额
	amountRe := regexp.MustCompile(`(\d{1,3}(?:,\d{3})+)元`)

	var fileAmounts [][]string
	for _, f := range files {
		content, err := os.ReadFile(f)
		if err != nil {
			c.ok = false
			c.msg = fmt.Sprintf("无法读取 %s: %v", f, err)
			return c
		}
		matches := amountRe.FindAllStringSubmatch(string(content), -1)
		var amounts []string
		for _, m := range matches {
			amounts = append(amounts, m[1])
		}
		fileAmounts = append(fileAmounts, amounts)
	}

	// 检查三处是否都包含相同的金额集合
	if len(fileAmounts) != 3 {
		c.ok = false
		c.msg = "未能提取到 3 处文件的金额"
		return c
	}

	// 取交集：检查每处是否都包含 1,036,420 / 1,083,760 / 1,130,040
	required := map[string]bool{
		"1,036,420": true,
		"1,083,760": true,
		"1,130,040": true,
	}

	var missing []string
	for i, amounts := range fileAmounts {
		amountSet := map[string]bool{}
		for _, a := range amounts {
			amountSet[a] = true
		}
		for req := range required {
			if !amountSet[req] {
				names := []string{"cheatsheet.md", "ch03-assessment-benefits.md", "patterns.md"}
				missing = append(missing, fmt.Sprintf("%s 缺 %s", names[i], req))
			}
		}
	}

	if len(missing) == 0 {
		c.ok = true
	} else {
		c.ok = false
		c.msg = strings.Join(missing, "; ")
	}
	return c
}
