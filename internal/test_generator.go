package internal

func GenerateTestFile(sourceFile string) string {
    filePath, err := filepath.Abs(filename)
    if err != nil {
        panic(fmt.Errorf("Generate tests: Error while getting filapath: %w", err))
	}
	ext := filepath.Ext(filePath)
	if ext != ".go" {
		panic(fmt.Errorf("Generate tests: Error file needs to be '.go': %w", err))
	}

	source, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Errorf("Generate tests: Error while reading file: %w", err))
	}

    // Формируем имя тестового файла
	testFile := strings.TrimSuffix(filePath, ext) + "_test.go"

	existing, err := os.ReadFile(target)
	if err == nil && !bytes.HasPrefix(existing, []byte(generatedTestMarker+"\n")) {
		panic("generate tests: refusing to overwrite manual tests: " + target)
	}
	if err != nil && !os.IsNotExist(err) {
		panic(fmt.Errorf("generate tests: %w", err))
	}

    return absPath, testFile, nil
}
