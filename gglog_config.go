// edit by haozibi
// 放置 gglog 新增方法

package gglog

import (
	"strconv"
)

// 日志输入变量，控制 console 输出是否精简，默认为否
var IsSimple = false

// 设置日志console 输出是否精简，默认为否
func SetOutSimple(v bool) {
	IsSimple = v
}

//设置stderrThreshold值，只有大于等于 stderrThreshold 的级别才能正确输出，706行附近
// setOL => set output level
// 默认级别 ERROR
func SetOutLevel(value string) error {
	var threshold severity
	// Is it a known name?
	if v, ok := severityByName(value); ok {
		threshold = v
	} else {
		v, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		threshold = severity(v)
	}
	logging.stderrThreshold.set(threshold)
	return nil
}

// 设置日志输出目录，
func SetLogDir(dir string) {
	logDirs = append(logDirs, dir)
}
