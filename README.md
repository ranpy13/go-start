# create-go-app

This is a small CLI tool I wrote for myself to scaffold new Go projects. It’s not a framework or anything fancy — just a simple way to avoid writing the same boilerplate every time I start a new project. Just starting to learn and was trying to fish around for something like `cargo new insert-your-project-name-here` or maybe `create-react-app` (that is a good one) for go. Some were overkill, some were not working well, so I did my own. Mostly for myself. But here you go!

It creates a basic project structure with things like `.env`, `.gitignore`, `src/`, `utils/`, a simple logger, and initializes a Go module and a git repo.

Honestly, it's minimal on purpose. I didn’t want something like [`go-zero`](https://github.com/zeromicro/go-zero) or [`buffalo`](https://github.com/gobuffalo/buffalo) — just something that helps me get started quickly and keeps things clean.

---

## What It Generates

When you run it, it gives you a structure like this:

```

my-app/
├── .env
├── .env.example
├── .gitignore
├── go.mod
├── README.md
├── src/
│   └── main.go
├── utils/
│   └── helper.go
├── logger/
│   └── logger.go

````

The logger is a basic wrapper with `.Info()`, `.Debug()`, `.Error()` methods. It's just enough to get by and doesn't try to replace anything production-grade.

---

## How to Use It

You don’t need to install anything. As long as you have Go installed (1.16+), you can run it directly with:

```bash
go run github.com/ranpy13/go-start@latest my-project
````

Replace `my-project` with whatever you want your folder name to be.

Once it runs, just:

```bash
cd my-project
go run src/main.go
```

You’ll see some log messages to confirm everything's wired up.

---

## What's Inside

### `.env` and `.env.example`

You can store environment variables here. Not wired to anything automatically — you can plug in something like `github.com/joho/godotenv` if you need it. Now that I think about it, maybe that should have been included. Maybe, someday. todo: ??

### `logger/`

A tiny logger that prints to stdout/stderr with simple prefixes like `[INFO]`, `[DEBUG]`, `[ERROR]`. It’s enough for basic CLI or web tools.

### `utils/`

Just a placeholder for helper functions. Nothing magical.

### `src/main.go`

Entry point. Right now it just prints a few logger lines. Add your logic here.

Hmmm... you know coming from spring initializer and stuff, and from good programmign practices and TDD and DDD and what not, maybe I should add tests as well. But test what, in this demo?
todo: add test - take inspo from spring initializer, sample code, sample test

---

## Why I Made This

Every time I started a Go project, I’d end up setting up the same files manually — making folders, writing `.gitignore`, typing out `main.go`, etc. It got repetitive. This script just automates that part.

There are probably better tools out there. This isn’t trying to compete with them. It’s just simple and works for me.

If it helps someone else, that’s great. If not, that’s okay too.

---

## Future Plans

No real roadmap. If I find myself doing more things manually again (adding a router, logger upgrade, Dockerfile, whatever), I might add options like:

* `--http`
* `--docker`
* `--makefile`
* tests/ - sure thing, on my radar

But only if I actually need them. _(read: only if I am not feeling lazy)_

---

## Contributing

This is barely a “project”, so no pressure — but feel free to fork, modify, or open a PR if you think of something simple to add. Just keep it minimal.

---

## License

MIT — do whatever.
