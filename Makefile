

docker_build:
	docker build -t context_labs .

docker_run:
	docker run  context_labs

# The dependency check isn't working here.
uml_diagram: uml-1.png
	npx -p @mermaid-js/mermaid-cli mmdc -i uml.md -o uml.png