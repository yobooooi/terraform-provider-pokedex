default: install

build:
	go build -o terraform-provider-pokedex_v${version}

install: build
	mkdir -p ~/.terraform.d/plugins/terraform.local/local/pokedex/${version}/darwin_arm64
	mv terraform-provider-pokedex_v${version} ~/.terraform.d/plugins/terraform.local/local/pokedex/${version}/darwin_arm64/