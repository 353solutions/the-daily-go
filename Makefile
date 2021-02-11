%.html: %.md
	kramdown -i GFM \
	    --template template.erb \
	    --syntax-highlighter rouge  $< > $@

mdignore=playbook.md TODO.md
mdfiles=$(filter-out $(mdignore),$(wildcard *.md))
mdout=$(subst .md,.html,$(mdfiles))

class=tdg21
bucket=gs://353solutions/c


all: $(mdout) html

.PHONY: html
html:
	SRC=nlp genhtml

fresh: clean all

sync: all
	@rsync \
	    --exclude $(class).zip \
	    --exclude '*.md' \
	    --exclude '*.swp' \
	    --exclude .git \
	    --exclude .gitignore \
	    --exclude .idea \
	    --exclude .vscode \
	    --exclude Makefile \
	    --exclude solutions \
	    --exclude nlp/vendor \
	    --exclude quotes \
	    -av . /tmp/$(class)
	@gsutil -m rsync -r /tmp/$(class) $(bucket)/$(class)
	@gsutil -q -m acl -r ch -u AllUsers:R $(bucket)/$(class)


zip: all
	zip -r9 353-tdg.zip nlp html data README.html \
	    -x '*.cache*' \
	    -x '*.git/*' \
	    -x '*.gitkeep' \
	    -x '*.idea*' \
	    -x '*.swp' \
	    -x 'vendor/*' \
	    -x Makefile
