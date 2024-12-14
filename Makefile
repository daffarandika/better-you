# <==> AUTORELOAD <==>
.PHONY: dev
dev:
	air

# <==> BUILD <==>
.PHONY: build
build:
	go build -o bin/app

# <==> RUN <==>
.PHONY: run
run:
	go run main.go

# <==> TAILDWIND <==>
.PHONY: tailwindcss
tailwindcss:
	bun run tailwindcss --config tailwind.config.js -i base.css -o static/css/styles.css --watch

# <==> TEMPL <==>
.PHONY: templ
templ:
	templ generate
