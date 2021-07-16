#include <iostream>
#include <memory>
#include "my_shared_ptr.h"

int main() {
    std::shared_ptr<int> test= std::make_shared<int>() ;
    std::shared_ptr<int> test2(new int());
    std::shared_ptr<int> test3;
    *test = 10;
    std::cout << "value " << *test<<", use_count:"<< test.use_count() << std::endl;
    std::shared_ptr<int> test4 = test;
    std::cout << "value " << *test<<", use_count:"<< test.use_count() << std::endl;
    std::cout<< " ==================== "<<std::endl;
    My_Shared_ptr<int> my_test(new int(10));
    std::cout << "value " << *my_test<<", use_count:"<< my_test.use_count() << std::endl;
    My_Shared_ptr<int> my_test_copy = my_test;
    std::cout << "value " << *my_test<<", use_count:"<< my_test.use_count() << std::endl;
    std::cout << "my_test_copy value " << *my_test_copy<<", use_count:"<< my_test_copy.use_count() << std::endl;

    return 0;
}
