all:
	rm -rf build
	mkdir build
	mkdir build/test
	cmake -S . -B build -G Ninja
	cmake --build build --parallel --verbose
	ctest -j 100 --test-dir build/test --output-on-failure
builder:
	cmake -S . -B build
	cmake --build build -- -j14
current:
	cmake -S . -B build
	cmake --build build -- -j14
	ctest -R <day> --test-dir build/test --output-on-failure
tester:
	ctest -j 100 --test-dir build/test --output-on-failure

# CMake = “Recipe creator” → figures out what needs to be cooked and generates instructions.
# Ninja = “Chef” → actually follows the recipe and cooks the food efficiently.