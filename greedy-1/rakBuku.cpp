#include <iostream>
using namespace std;

int main(){
    int nBebek;
    int cummulativeBebekArray[20000000000];
    int height;

    cin >> nBebek;
    cin >> height;

    cout << nBebek << " " << height << endl;

    for(int i = 0; i < nBebek; i++){
        int bebekHeight;
        cin >> bebekHeight;
        
        cummulativeBebekArray[i] = cummulativeBebekArray[i-1] + bebekHeight;
        
    }

    

    int minBebekAmount = 0;
    
    for(int i = 0; i < nBebek; i++){
        
        if(cummulativeBebekArray[i] <= height){
            minBebekAmount++;
        }
    }

    cout << minBebekAmount;

    return 0;
}