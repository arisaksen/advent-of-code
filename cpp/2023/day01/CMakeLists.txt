# Header-only library for day01
add_library(day01_lib INTERFACE)
target_include_directories(day01_lib INTERFACE ${CMAKE_CURRENT_SOURCE_DIR})

# Executable
add_executable(day01 main.cc)
target_link_libraries(day01 PRIVATE day01_lib)
