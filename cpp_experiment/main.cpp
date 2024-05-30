#include <chrono>
#include <random>
#include <iostream>
#include <algorithm>
#include <vector>

struct Point2i {
    int x;
    int y;
};

template<typename T>
struct Point {
    T x;
    T y;
};


bool operator<(Point2i p1, Point2i p2){
    return p1.x + p1.y < p2.x + p2.y;
}

template<typename T>
bool operator<(Point<T> p1, Point<T> p2){
    return p1.x + p1.y < p2.x + p2.y;
}

int main(){
    std::random_device dev;
    std::mt19937 rng(dev());
    int min_v = 0;
    int max_v = 100000;
    std::uniform_int_distribution<std::mt19937::result_type> dist(min_v, max_v);
    int n = 10000000;
    std::vector<Point<int>> v_template;
    std::vector<Point2i> v;

    for (int i = 0; i < n; ++i){
        int a = dist(rng);
        int b = dist(rng);
        v.push_back(Point2i{a, b});
        v_template.push_back(Point<int>{a, b});
    }
    auto start1 = std::chrono::high_resolution_clock::now();
    std::sort(v.begin(), v.end());
    auto stop1 = std::chrono::high_resolution_clock::now();
    auto duration1 = std::chrono::duration_cast<std::chrono::microseconds>(stop1 - start1);
    std::cout << duration1.count() << std::endl;

    auto start2 = std::chrono::high_resolution_clock::now();
    std::sort(v.begin(), v.end());
    auto stop2 = std::chrono::high_resolution_clock::now();
    auto duration2 = std::chrono::duration_cast<std::chrono::microseconds>(stop1 - start1);
    std::cout << duration2.count() << std::endl;

}

