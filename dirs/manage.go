package dirs

import (
	dircopy "github.com/otiai10/copy"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)


type Manager struct {
	dirToBackupPath string
	backupDirPath string
	backupCount int
}

func SetupBackupManager(dirToBackup, dirBackup string, backupCount int) (m *Manager) {
	m = &Manager{
		dirToBackupPath: dirToBackup,
		backupDirPath:   dirBackup,
		backupCount:     backupCount,
	}
	return m
}

func maxInIntSlice(s []int) int {
	max := 0
	for _, n := range s {
		if n > max {
			max = n
		}
	}
	return max
}

func minInIntSlice(s []int) int {
	min := math.MaxInt16
	for _, n := range s {
		if n < min {
			min = n
		}
	}
	return min
}

func getIntSuffix(s string) (suffix int) {
	re := regexp.MustCompile(".*[^\\d](\\d+)$")
	ss := re.FindStringSubmatch(s)
	if len(ss) == 2 {
		suffix, _ = strconv.Atoi(ss[1])
	}
	return suffix
}

func (m *Manager) getBackups() (newNum int) {
	var backupNums []int
	matches, err := filepath.Glob(m.backupDirPath + "*")
	for _, mt := range matches {
		backupNums = append(backupNums, getIntSuffix(mt))
	}
	newNum = maxInIntSlice(backupNums) + 1
	if len(matches) >= m.backupCount {
		var suffixDelete string
		if minInIntSlice(backupNums) == 0 {
			suffixDelete = ""
		} else {
			suffixDelete = strconv.Itoa(minInIntSlice(backupNums))
		}
		err = os.RemoveAll(m.backupDirPath + suffixDelete)
		if err != nil {
			log.Fatal(err)
		}
	}
	return newNum
}

func (m *Manager) Backup() {
	suffix := m.getBackups()
	err := dircopy.Copy(m.dirToBackupPath, m.backupDirPath + strconv.Itoa(suffix))
	if err != nil {
		log.Fatal(err)
	}
}