add_executable(test_part_one test_day01_part_one.cpp)

target_include_directories(test_part_one
        PRIVATE
        ${CMAKE_SOURCE_DIR}/day01
)

target_link_libraries(test_part_one
        PRIVATE
        GTest::gtest_main
        part_one
)

# Discover the tests
include(GoogleTest)
gtest_discover_tests(test_part_one)
