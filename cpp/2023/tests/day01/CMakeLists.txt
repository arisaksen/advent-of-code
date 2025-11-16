add_executable(test_part_one test_day01_part_one.cpp)

target_link_libraries(test_part_one
        PRIVATE
        GTest::gtest_main
        day01_lib   # <--- link INTERFACE library, not part_one
)

include(GoogleTest)
gtest_discover_tests(test_part_one)
