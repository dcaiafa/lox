Prism.languages.lox = {
  'comment': /\/\/.*|\/\*[\s\S]*?(?:\*\/|$)/,
  'string': {
    pattern: /'(?:\\.|[^\\'\r\n])*'/,
    greedy: true
  },
  'keyword': /@(discard|empty|emit|error|external|frag|left|lexer|list|macro|mode|parser|pop_mode|push_mode|right|start)\b/,
  'character-class': {
		pattern: /\[(?:\\.|[^\\\]\r\n])*\]/,
		greedy: true,
		alias: 'regex',
		inside: {
			'range': {
				pattern: /([^[]|(?:^|[^\\])(?:\\\\)*\\\[)-(?!\])/,
				lookbehind: true,
				alias: 'punctuation'
			},
			'escape': /\\(?:u(?:[a-fA-F\d]{4}|\{[a-fA-F\d]+\})|[pP]\{[=\w-]+\}|[^\r\nupP])/,
			'punctuation': /[\[\]]/
		}
	},
  'punctuation': /[:()|=]/
}
