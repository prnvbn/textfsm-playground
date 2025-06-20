/**
 * Jinja Language Definition for Monaco Editor
 *
 * Based on implementations from:
 * - https://github.com/zebbra/monaco-jinja2-example
 * - Microsoft Monaco Editor issue #4309 (https://github.com/microsoft/monaco-editor/issues/4309)
 * - Community contributions for Jinja syntax highlighting
 *
 * Adapted and enhanced for this project.
 */

export function registerJinjaLanguage() {
  if (!window.monaco) {
    console.warn("Monaco editor not available for Jinja language registration");
    return;
  }

  // Register the Jinja language
  monaco.languages.register({
    id: "jinja",
    extensions: [".jinja", ".jinja2", ".j2"],
    aliases: ["Jinja", "jinja", "jinja2"],
    mimetypes: ["text/jinja", "text/jinja2"],
  });

  // Define Jinja language configuration
  monaco.languages.setLanguageConfiguration("jinja", {
    comments: {
      blockComment: ["{#", "#}"],
    },
    brackets: [
      ["{#", "#}"],
      ["{%", "%}"],
      ["{{", "}}"],
      ["(", ")"],
      ["[", "]"],
      // HTML brackets
      ["<!--", "-->"],
      ["<", ">"],
    ],
    autoClosingPairs: [
      { open: "{# ", close: " #}" },
      { open: "{% ", close: " %}" },
      { open: "{{ ", close: " }}" },
      { open: "[", close: "]" },
      { open: "(", close: ")" },
      { open: '"', close: '"' },
      { open: "'", close: "'" },
    ],
    surroundingPairs: [
      { open: '"', close: '"' },
      { open: "'", close: "'" },
      // HTML
      { open: "<", close: ">" },
    ],
  });

  // Define Jinja language tokenizer using Monarch
  monaco.languages.setMonarchTokensProvider("jinja", {
    wordPattern: /(-?\d*\.\d\w*)|([^`~!@$^&*()=+[{]}\\|;:'",.<>\/\s]+)/g,

    defaultToken: "",
    tokenPostfix: "",
    ignoreCase: true,

    brackets: [
      { token: "delimiter.jinja", open: "{%", close: "%}" },
      { token: "delimiter.jinja", open: "{{", close: "}}" },
      { token: "delimiter.jinja", open: "{#", close: "#}" },
      { token: "delimiter.square.jinja", open: "[", close: "]" },
      { token: "delimiter.paren.jinja", open: "(", close: ")" },
    ],

    keywords: [
      // Control structures
      "if",
      "else",
      "elif",
      "endif",
      "for",
      "endfor",
      "in",
      "while",
      "endwhile",
      "break",
      "continue",

      // Blocks and macros
      "block",
      "endblock",
      "macro",
      "endmacro",
      "call",
      "endcall",
      "filter",
      "endfilter",
      "set",
      "endset",
      "with",
      "endwith",
      "without",
      "context",
      "autoescape",
      "endautoescape",
      "raw",
      "endraw",

      // Template inheritance
      "extends",
      "include",
      "import",
      "from",

      // Internationalization
      "trans",
      "endtrans",
      "pluralize",
      "endpluralize",

      // Control
      "as",
      "do",
      "loop",

      // Constants
      "true",
      "false",
      "none",
      "True",
      "False",
      "None",
    ],

    operators: [
      "and",
      "or",
      "not",
      "b-and",
      "b-xor",
      "b-or",
      "starts with",
      "ends with",
      "matches",
      "in",
      "is",
      "+",
      "-",
      "*",
      "**",
      "/",
      "//",
      "%",
      "==",
      "!=",
      "<",
      ">",
      ">=",
      "<=",
      "|",
      "~",
      ":",
      ".",
      "..",
      "?",
      "??",
      "=",
    ],

    tokenizer: {
      root: [
        // Whitespace
        [/\s+/, "white"],

        // Jinja comments
        [/{#/, "comment.jinja", "@commentState"],

        // Jinja blocks
        [/{%[-~]?/, "delimiter.jinja", "@blockState"],

        // Jinja variables
        [/{{[-~]?/, "delimiter.jinja", "@variableState"],

        // HTML tags
        [/<\/?[a-zA-Z][\w-]*/, "tag"],
        [/<\/?[a-zA-Z][\w-]*\s*>/, "tag"],

        // HTML attributes
        [/[a-zA-Z-]+(?=\s*=)/, "attribute.name"],
        [/"[^"]*"/, "attribute.value"],
        [/'[^']*'/, "attribute.value"],

        // Everything else
        [/[^{<]+/, "text"],
      ],

      commentState: [
        [/#}/, "comment.jinja", "@pop"],
        [/./, "comment.jinja"],
      ],

      blockState: [
        [/-?%}/, "delimiter.jinja", "@pop"],
        [
          /\b(if|else|elif|endif|for|endfor|while|endwhile|do|loop|break|continue|as|with|without|context|include|import|from|block|endblock|extends|macro|endmacro|call|endcall|filter|endfilter|set|endset|autoescape|endautoescape|raw|endraw|trans|endtrans|pluralize|endpluralize|true|false|none|True|False|None)\b/,
          "keyword.jinja",
        ],
        { include: "expression" },
      ],

      variableState: [
        [/-?}}/, "delimiter.jinja", "@pop"],
        { include: "expression" },
      ],

      expression: [
        // Whitespace
        [/\s+/, "white"],

        // Brackets and delimiters
        [/\(|\)|\[|\]|\{|\}/, "delimiter.jinja"],

        // Numbers
        [/\d+(\.\d+)?/, "number.jinja"],

        // Strings
        [/"([^#"\\]*(?:\\.[^#"\\]*)*)"/, "string.jinja"],
        [/'([^'\\]*(?:\\.[^'\\]*)*)'/, "string.jinja"],

        // Operators
        [/\b(and|or|not|b-and|b-xor|b-or)\b/, "operators.jinja"],
        [/\b(starts with|ends with|matches)\b/, "operators.jinja"],
        [/\b(in|is)\b/, "operators.jinja"],
        [
          /\+|-|\*{1,2}|\/{1,2}|%|==|!=|<|>|>=|<=|\||~|:|\.{1,2}|\?{1,2}|=/,
          "operators.jinja",
        ],

        // Variables and identifiers
        [
          /[^\W\d][\w]*/,
          {
            cases: {
              "@keywords": "keyword.jinja",
              "@default": "variable.jinja",
            },
          },
        ],
      ],
    },
  });
}
