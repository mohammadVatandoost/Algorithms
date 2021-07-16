#include <iostream>
#include <algorithm>>


int main() {
    std::vector<int> temp = {11, 23, 6,  4, 5, 8};
    std::sort(temp.begin(), temp.end());

    for(auto it = temp.begin(); it != temp.end(); it++) {
        std::cout<<*it <<", ";
    }
    std::cout<<std::endl;
    std::cout << "Hello, World!" << std::endl;
    return 0;
}
