#include<iostream>
#include "./headers/clearscreen.hpp"
#include "./headers/appdescription.hpp"
#include "./headers/conversionoptions.hpp"
using namespace std;

int main() {
  clearScreen();
  cout << "\n";
  appDescription();
  cout << "\n";
  conversionOptions();
  return 0;
}

