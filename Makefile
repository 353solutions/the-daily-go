%.html: %.md
	kramdown --template template.erb $< > $@

html/%.html: code/%.go
	mkdir -p $(shell dirname $@)
	pygmentize -Ofull,linenos=1,style=vs,lineanchors=l -l go -f html -o $@ $<

html/%.html: code/%.mod
	mkdir -p $(shell dirname $@)
	pygmentize -Ofull,linenos=1,style=vs,lineanchors=l -l go -f html -o $@ $<


mdignore=playbook.md TODO.md index.md
mdfiles=$(filter-out $(mdignore),$(wildcard *.md))
mdout=$(subst .md,.html,$(mdfiles))
gofiles=$(shell find code -type f -name '*.go' -not -path '*vendor*')
gohtml=$(subst .go,.html,$(gofiles))
goout=$(subst code/,html/,$(gohtml))
modfiles=$(shell find code -type f -name '*.mod')
modhtml=$(subst .mod,.html,$(modfiles))
modout=$(subst code/,html/,$(modhtml))
other=tdg.log
#other_out=\
#      html/nlp/cmd/nlpd/Dockerfile.html \
#      html/nlp/Makefile.html
class=tdg-b1
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
	    --exclude code/nlp/vendor \
	    -av . /tmp/$(class)
	@gsutil -m rsync -r /tmp/$(class) $(bucket)/$(class)
	@gsutil -q -m acl -r ch -u AllUsers:R $(bucket)/$(class)


zip: all
	zip -r9 $(class).zip code html data $(other) README.html \
	    -x '*.cache*' \
	    -x '*.git/*' \
	    -x '*.gitkeep' \
	    -x '*.idea*' \
	    -x '*.swp' \
	    -x Makefile

upload-zip:
	@gsutil cp $(class).zip $(bucket)/$(class)/$(class).zip
	@gsutil -m -q acl -r ch -u AllUsers:R $(bucket)/$(class)/$(class.zip)

#html/nlp/Makefile.html: code/nlp/Makefile
#        mkdir -p $(shell dirname $@)
#        pygmentize -Ofull,linenos=1,style=vs,lineanchors=l -l make -f html -o $@ $<

#html/nlp/cmd/nlpd/Dockerfile.html: code/nlp/cmd/nlpd/Dockerfile
#        mkdir -p $(shell dirname $@)
#        pygmentize -Ofull,linenos=1,style=vs,lineanchors=l -l docker -f html -o $@ $<
