#include <iostream>

#include "input.hpp"
#include "part_one.hpp"
#include "part_two.hpp"

int main() {
  std::cout << "The sum of calibration values is: " << part_one::calibration_values(input::day01) << "\n";
  std::cout << "The sum of calibration values with spelled out digits is: " << part_two::spelled_out(input::day01) << "\n";
}