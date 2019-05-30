%.html: %.md
	kramdown --template template.erb $< > $@

html/%.html: nlp/%.go
	mkdir -p $(shell dirname $@)
	pygmentize -Ofull,linenos=1,style=vs,lineanchors=l -l go -f html -o $@ $<

html/%.html: nlp/%.mod
	mkdir -p $(shell dirname $@)
	pygmentize -Ofull,linenos=1,style=vs,lineanchors=l -l go -f html -o $@ $<



mdignore=playbook.md TODO.md index.md
mdfiles=$(filter-out $(mdignore),$(wildcard *.md))
mdout=$(subst .md,.html,$(mdfiles))
gofiles=$(shell find nlp -type f -name '*.go' -not -path '*vendor*')
gohtml=$(subst .go,.html,$(gofiles))
goout=$(subst nlp/,html/,$(gohtml))
modfiles=$(shell find nlp -type f -name '*.mod')
modhtml=$(subst .mod,.html,$(modfiles))
modout=$(subst nlp/,html/,$(modhtml))
other=tdg.log

other_out=\
      html/.gitignore.html \
      html/Makefile.html \
      html/README.html \
      html/cmd/nlpd/Dockerfile.html \
      html/Dockerfile.test.html \
      html/.circleci/config.html

class=tdg
bucket=gs://353solutions/c


all: $(mdout) $(goout) $(modout) $(other_out)

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
	zip -r9 $(class).zip nlp html data $(other) README.html \
	    -x '*.cache*' \
	    -x '*.git/*' \
	    -x '*.gitkeep' \
	    -x '*.idea*' \
	    -x '*.swp' \
	    -x 'vendor/*' \
	    -x Makefile

upload-zip:
	@gsutil cp $(class).zip $(bucket)/$(class)/$(class).zip
	@gsutil -m -q acl -r ch -u AllUsers:R $(bucket)/$(class)/$(class.zip)

html/Makefile.html: nlp/Makefile
	    mkdir -p $(shell dirname $@)
	    pygmentize -Ofull,linenos=1,style=vs,lineanchors=l -l make -f html -o $@ $<

html/cmd/nlpd/Dockerfile.html: nlp/cmd/nlpd/Dockerfile
	    mkdir -p $(shell dirname $@)
	    pygmentize -Ofull,linenos=1,style=vs,lineanchors=l -l docker -f html -o $@ $<

html/Dockerfile.test.html: nlp/Dockerfile.test
	    mkdir -p $(shell dirname $@)
	    pygmentize -Ofull,linenos=1,style=vs,lineanchors=l -l docker -f html -o $@ $<

html/README.html: nlp/README.md
	    mkdir -p $(shell dirname $@)
	    pygmentize -Ofull,linenos=1,style=vs,lineanchors=l -l md -f html -o $@ $<

html/.gitignore.html: nlp/.gitignore
	    pygmentize -Ofull,linenos=1,style=vs,lineanchors=l -l text -f html -o $@ $<

html/.circleci/config.html: nlp/.circleci/config.yml
	    mkdir -p $(shell dirname $@)
	    pygmentize -Ofull,linenos=1,style=vs,lineanchors=l -l yaml -f html -o $@ $<
