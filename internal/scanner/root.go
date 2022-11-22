package scanner

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

const maxFileSize = 99999

type IScanProcessor interface {
	Process(repoLink string) error
}

type ScanProcessor struct {
	FileSystem billy.Filesystem
}

type filepath struct {
	foldername string
	file       fs.FileInfo
}

func (sp *ScanProcessor) Process(repoLink string) ([]Finding, error) {
	if _, err := git.Clone(memory.NewStorage(), sp.FileSystem, &git.CloneOptions{
		URL: repoLink,
	}); err != nil {
		msg := fmt.Sprintf("processing message, error cloning: %s (%s)", err, repoLink)
		return nil, errors.New(msg)
	}

	files, err := sp.FileSystem.ReadDir("/")
	if err != nil {
		msg := fmt.Sprintf("processing message, error list files: %s", err)
		return nil, errors.New(msg)
	}

	findings, err := sp.scanFiles(files)
	if err != nil {
		return nil, err
	}

	return findings, nil
}

func (sp *ScanProcessor) scanFiles(files []fs.FileInfo) ([]Finding, error) {
	paths := []filepath{}
	for _, file := range files {
		paths = append(paths, filepath{
			foldername: "",
			file:       file,
		})
	}

	findings := []Finding{}
	for len(paths) > 0 {
		path := paths[len(paths)-1]
		fullpath := path.foldername + path.file.Name()
		paths = paths[:len(paths)-1]

		if path.file.IsDir() {
			subfiles, err := sp.FileSystem.ReadDir(fullpath)
			if err != nil {
				msg := fmt.Sprintf("processing message, error list files: %s", err)
				return nil, errors.New(msg)
			}
			for _, file := range subfiles {
				paths = append(paths, filepath{
					foldername: fullpath + "/",
					file:       file,
				})
			}

			continue
		}

		file, err := sp.FileSystem.Open(fullpath)
		if err != nil {
			msg := fmt.Sprintf("processing message, error opening file: %s (%s)", err, path)
			return nil, errors.New(msg)
		}

		dst := make([]byte, maxFileSize)
		n, err := file.Read(dst)
		if err != nil {
			msg := fmt.Sprintf("processing message, error reading file: %s (%s)", err, path)
			return nil, errors.New(msg)
		}

		content := string(dst[:n])

		newFindings := checkVulnerability(fullpath, content)
		findings = append(findings, newFindings...)

		if err = file.Close(); err != nil {
			msg := fmt.Sprintf("processing message, error closing file: %s (%s)", err, path)
			return nil, errors.New(msg)
		}
	}

	return findings, nil
}

func checkVulnerability(path string, content string) []Finding {
	findings := []Finding{}
	scanner := bufio.NewScanner(strings.NewReader(content))

	line := 0
	for scanner.Scan() {
		line += 1
		text := scanner.Text()

		for _, rule := range GetRules() {
			if rule.Checker(text) {
				finding := NewFinding(path, line, rule)
				findings = append(findings, *finding)
			}
		}
	}

	return findings
}
