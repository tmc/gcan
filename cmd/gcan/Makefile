all: doc

doc:
	@go run -tags mkmanpage *.go mkmanpage | sed -e "/mkmanpage/{N;N;N;N;N;N;d;}" | sed '1 s/".*"/""/' > doc.troff
	@groff -t -e -mandoc -Tutf8 doc.troff | col -bx | sed '1{N;N;N;d;};$$d' | sed 's/^/\/\/ /' > doc.go
	@groff -mandoc -Thtml doc.troff | pandoc -f html -t markdown_github > README.md
	@echo "package main" >> doc.go
	@gofmt -w .
	@mv doc.troff gcan.1
