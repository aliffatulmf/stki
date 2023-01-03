BUILD_FOLDER = build

pre-build:
	mkdir $(BUILD_FOLDER)

build:
	go build -o $(BUILD_FOLDER)/stki main.go

clean:
	rm -rf $(BUILD_FOLDER)
