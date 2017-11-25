#include <iostream>
#include <string>
#include <vector>
#include <typeinfo>
#define ARRAY_SIZE(array) (sizeof((array))/sizeof((array[0])))

const int sizeOfData = 4;

struct Point{
  int x;
  int y;
  std::string label;
};
/*
class KNN{
private:
  int k;
  Point data[sizeOfData];
public:
  KNN(int k){
    this->k = k;
  }
  KNN(){
    k=3;
  }
  void loadData(Point* data){
    this -> data = data;
  }
  std::vector<int> euclideanDistance(Point p){
    std::vector<std::vector<int>> diffCoord;
    diffCoord.resize(sizeOfData, std::vector<int>(2, 0));
    for(int i=0; i < sizeOfData; i++){
      diffCoord[0][i] = p.x - data[i].x;
      diffCoord[0][i] = p.y - data[i].y;
    }
    //int[] d = diffCord ** 2
  }
};*/

int main(int argv, char **argc) {
  Point sum[sizeOfData];

  sum[0].x 	=	1;
  sum[0].y	= 2;

  sum[1].x	=	1;
  sum[1].y	= 4;

  sum[2].x	=	4;
  sum[2].y	=	1;

	sum[3].x	=	4;
  sum[3].y	= 2;

  sum[0].label	=	"Action";
  sum[1].label	=	"Action";
  sum[2].label 	= "Romantic";
  sum[3].label 	=	"Romantic";
  std::cout << typeid(sum).name() << std::endl;
}
