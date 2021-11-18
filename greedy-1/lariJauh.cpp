#include <iostream>
using namespace std;

int main(){
    int n, k;   
    int nArray[100];
    int kArray[100];
    int cummulativeNArray[100];
    cin >> n >> k;
    for(int i = 0; i < n; i++){
        int inputN;
        cin >> inputN;
        cummulativeNArray[i] = cummulativeNArray[i-1] + inputN;

    }
    for(int i = 0; i < k; i++){
        cin >> kArray[i];
    }

    for(int i = 0; i < k; i++){

        int passedPatok = 0;
        for(int j = 0; j < n; j++){
            if(kArray[i] >= cummulativeNArray[j]){
                passedPatok++;
            }
        }
        cout << passedPatok << endl;
    }
    
    

    return 0

}