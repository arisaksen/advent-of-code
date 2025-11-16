#include "part_one.hpp"
#include <gtest/gtest.h>
#include <string_view>

TEST(PartOneTest, SimpleInput) {
    std::string_view input = "a1b2\nx3y4\n7z8";
    EXPECT_EQ(part_one::calibration_values(input), 124);
}

TEST(PartOneTest, EmptyInput) {
    std::string_view input = "";
    EXPECT_EQ(part_one::calibration_values(input), 0);
}

TEST(PartOneTest, SingleLine) {
    std::string_view input = "5x9";
    EXPECT_EQ(part_one::calibration_values(input), 59);
}
