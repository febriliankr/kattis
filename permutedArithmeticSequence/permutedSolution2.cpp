#include <iostream>
using namespace std;

int main(){

    int n;
    cin >> n;
    for(int i = 0; i < n; i++){
        int k;
        cin >> k;


        int inputArray[1024];
        for (int j = 0; j < k; j++){
            
            cin >> inputArray[j];


        }

        int gap = 0;
        

        // Check selisih inputArray
        for(int j = 0; j < k-1; j++){
            int newGap;
            newGap = inputArray[j+1] - inputArray[j];
            
            
            if(gap != newGap){
                cout << "NO" << endl;    
            } else {
                gap = newGap;
            }


        }
        
    }
    return 0;
}