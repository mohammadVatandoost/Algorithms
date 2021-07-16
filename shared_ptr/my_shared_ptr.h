//
// Created by mohammad on 7/16/21.
//

#ifndef SHARED_PTR_MY_SHARED_PTR_H
#define SHARED_PTR_MY_SHARED_PTR_H
#include <iostream>

template<class T>
class My_Shared_ptr {
public:
    explicit My_Shared_ptr(T *ob_ptr) {
        counter = new int(1);
        ptr = ob_ptr;
    }

    My_Shared_ptr( My_Shared_ptr& my_shared_ptr ) {
        counter = my_shared_ptr.counter;
        ptr = my_shared_ptr.ptr;
        (*counter)++;
    }

//    My_Shared_ptr& operator=(const T& ob_ptr) {
//        counter++;
//        return this;
//    };

    T operator*() {
        return *ptr;
    }

    int use_count() {
        return *counter;
    }

    ~My_Shared_ptr() {
        (*counter)--;
        if (*counter == 0) {
            std::cout<< "My_Shared_ptr object deleted"<<std::endl;
            delete ptr;
            delete counter;
        }
    }
private:
    int *counter;
    T *ptr;
};


#endif //SHARED_PTR_MY_SHARED_PTR_H
