package prompt

import _ "embed"

//go:embed system.txt
var SystemPrompt string

//go:embed log_guide.txt
var LogGuidePrompt string
