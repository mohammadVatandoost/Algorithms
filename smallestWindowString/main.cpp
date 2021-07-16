#include <iostream>
#include <cstring>


std::string findSmallestWindows(std::string cont, std::string chars) {
//   const char *cont_array = cont.c_str();
    int chars_map[255];
    for(int i=0; i< 255; i++) {
        chars_map[i] = 0;
    }

    for(auto c = chars.begin(); c != chars.end(); c++ ) {
        chars_map[c.base()[0]]++;
    }
    std::string result;
    int index=0;
    for(int i=0; i<cont.length(); i++) {
        int chars_map_temp[255];
        std::memcpy(chars_map_temp, chars_map, 255* sizeof(int)) ;
        int counter = 0 ;

        std::string result_temp;
        int index_temp = 0;
        for(int j = index; j < cont.length(); j++ ) {
            char c = cont[j];
            if( (chars_map_temp[c] != 0)   ) {
                chars_map_temp[c]--;
                counter++;
                if (result_temp.empty()) {
                    index_temp = j;
                }
                result_temp.append(cont.c_str()+j, 1);
            } else if(counter > 0) {
                result_temp.append(cont.c_str()+j, 1);
            }
            if( counter == chars.length()) {
                break;
            }
        }
        index = index_temp+1;
        if (index == cont.length()) {
            break;
        }
        if( counter != chars.length()) {
            continue;
        }
        if(( result == "" && result_temp.length() >0 ) || (result != "" && result.length() > result_temp.length()) ) {
            result = result_temp;
        }

    }
    return result;
}


int main() {
    int chars_map[255];

    std::string chars = "tist";
    std::string cont =  "this is a test string";
    std::string answer = "t stri";

//
//
//    std::string result = "";
//    int index=0;
//    for(int i=0; i<cont.length(); i++) {
//       int chars_map_temp[255];
//       std::memcpy(chars_map_temp, chars_map, 255* sizeof(int)) ;
//       int counter = 0 ;
//
//       std::string result_temp;
//       int index_temp = 0;
//       for(int j = index; j < cont.length(); j++ ) {
//           char c = cont[j];
//          if( (chars_map_temp[c] != 0)   ) {
//             chars_map_temp[c]--;
//             counter++;
//             if (result_temp.empty()) {
//                 index_temp = j;
//             }
//              result_temp.append(cont.c_str()+j, 1);
//          } else if(counter > 0) {
//              result_temp.append(cont.c_str()+j, 1);
//          }
//          if( counter == chars.length()) {
//              break;
//          }
//       }
//        index = index_temp+1;
//        if (index == cont.length()) {
//            break;
//        }
//        if( counter != chars.length()) {
//            continue;
//        }
//       if(( result == "" && result_temp.length() >0 ) || (result != "" && result.length() > result_temp.length()) ) {
//           result = result_temp;
//       }
//
//    }


    std::cout << "result:"<< findSmallestWindows(cont, chars)  << std::endl;
    return 0;
}
