<div align="center">

# ⛈ toofan

**a typing test for your terminal**

practice typing with real code · learn while you type · beat your pb

[![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=flat-square&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-bb9af7?style=flat-square)](LICENSE)

</div>

---

## features

- **two modes** — english words or real code snippets
- **code lessons** — practical, hand-written snippets you actually learn from
- **5 languages** — go, javascript, lua, shell, dart
- **infinite mode** — no timer, type the full snippet at your own pace
- **lesson picker** — choose what you want to practice, not just random code
- **live wpm** — see your speed as you type
- **personal bests** — tracked per duration and mode
- **5 themes** — tokyonight, gruvbox, rosépine, kanagawa, sakura
- **zero bloat** — just bubbletea + lipgloss, no other deps

## install

### go install
```sh
go install github.com/YOUR_USERNAME/toofan@latest
```

### build from source
```sh
git clone https://github.com/YOUR_USERNAME/toofan
cd toofan
make build
./toofan
```

### AUR (coming soon)
```sh
yay -S toofan
```

## usage

just run `toofan` and start typing. the test begins when you type the first character.

### keybindings

| key | action |
|---|---|
| `start typing` | begin test |
| `backspace` | fix mistakes |
| `esc` | restart current test |
| `tab` | change duration (∞/15/30/60/120s) |
| `ctrl+w` | toggle words ↔ code mode |
| `ctrl+l` | change programming language |
| `ctrl+o` | pick a specific lesson |
| `ctrl+t` | change theme |
| `ctrl+p` | open profile |
| `ctrl+c` | quit |

### code mode

code mode gives you real, practical code snippets to type — not random keywords. each snippet is a small, self-contained example that teaches a concept while building muscle memory for symbols like `{}`, `=>`, `()`, `<-`, etc.

you can pick a specific lesson with `ctrl+o` or let it pick randomly. set the timer to `∞` (infinite) and the test ends when you finish the snippet, no time pressure.

## how snippets work

snippets live as **native source files** inside `internal/lang/data/<language>/lessons/`. this means they're real `.go`, `.js`, `.sh` files — not strings buried in JSON.

each file has a single comment at the top for the topic heading:

```js
// Topic: Promises
function loadData(url) {
    return new Promise((resolve, reject) => {
        // ... actual code the user types
    });
}
```

the `// Topic:` line is extracted as a heading and shown in the UI. the user **only types the code** — comments are stripped.

## adding a language

1. create a directory: `internal/lang/data/yourname/lessons/`
2. drop in snippet files with native extensions:

```python
# Topic: List Comprehensions
squares = [x ** 2 for x in range(10)]
evens = [x for x in squares if x % 2 == 0]
print(evens)
```

3. rebuild — it's automatically picked up. no code changes needed.

supported comment styles: `//` (go, js, dart), `#` (shell, python), `--` (lua)

## adding snippets to an existing language

just add a new file to that language's `lessons/` folder. name it something descriptive:

```
internal/lang/data/javascript/lessons/08_closures.js
```

keep snippets short and practical. someone should be able to type the whole thing in 30-90 seconds and walk away understanding the concept.

## config

preferences are saved to `~/.toofan/config.txt` and persist between sessions:
- selected duration
- mode (words/code)
- language
- theme

test results are appended to `~/.toofan/results.txt`:
```
2026-04-01 22:18 |  85 wpm | 97.5% |  30s | words
2026-04-01 22:20 |  72 wpm | 94.2% |  60s | code:go
```

## themes

| theme | style |
|---|---|
| tokyonight | calm blues and purples |
| gruvbox | warm retro amber |
| rosépine | elegant dark pink |
| kanagawa | japanese ink tones |
| sakura | soft cherry blossom |

## project structure

```
toofan/
├── main.go                    # entry point
├── internal/
│   ├── model.go               # bubbletea model (state + controls)
│   ├── game.go                # typing engine + stats + infinite mode
│   ├── typing.go              # typing screen + topic display
│   ├── picker.go              # language, lesson, theme, duration pickers
│   ├── storage.go             # file persistence
│   ├── theme/                 # color palettes
│   └── lang/
│       ├── lang.go            # snippet parser + word loader
│       └── data/
│           ├── english/
│           │   └── words.txt  # word pool for word mode
│           ├── go/lessons/    # go code snippets
│           ├── javascript/lessons/
│           ├── shell/lessons/
│           ├── lua/lessons/
│           └── dart/lessons/
├── Makefile
├── LICENSE
└── README.md
```

## contributing

contributions welcome! especially:
- **new snippets** — short, practical, real-world code. see [adding snippets](#adding-snippets-to-an-existing-language)
- **new languages** — just a folder with lesson files
- **new themes** — one small Go file
- **bug fixes** and UX improvements

## dependencies

only two:
- [bubbletea](https://github.com/charmbracelet/bubbletea) — TUI framework
- [lipgloss](https://github.com/charmbracelet/lipgloss) — terminal styling

## license

[MIT](LICENSE)

