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
