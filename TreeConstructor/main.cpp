#include <iostream>
#include <string>
#include <set>
#include <sstream>

struct Edge{
    int parent;
    int child;
};

bool replace(std::string& str, const std::string& from, const std::string& to) {
    size_t start_pos = str.find(from);
    if(start_pos == std::string::npos)
        return false;
    str.replace(start_pos, from.length(), to);
    return true;
}

Edge parseEdge(std::string edge) {
    replace(edge, "(", "");
    replace(edge, ")", "");
    std::string str = "1,2,3,4,5,6";
    Edge result;

    std::stringstream ss(str);
    int j = 0 ;
    for (int i; ss >> i;) {
        if(j==0) {
            result.child = i;
        } else {
            result.parent = i;
        }
        j++;
        if (ss.peek() == ',')
            ss.ignore();
    }
    return result;
}


std::string TreeConstructor(std::string strArr[], int arrLength) {
    std::set<int> nodes_value;
    Edge *edges = new Edge[arrLength];
    for(int i=0; i<arrLength; i++) {
        edges[i] = parseEdge(strArr[i]);
        nodes_value.insert(edges[i].child);
        nodes_value.insert(edges[i].parent);
    }
    return strArr[0];
}

int main() {
    std::cout << "Hello, World!" << std::endl;
    std::string A[] = {"(1,2)", "(2,4)", "(5,7)", "(7,2)", "(9,5)"};
    int arrLength = 5 ;
    std::cout << TreeConstructor(A, arrLength);
    return 0;
}
