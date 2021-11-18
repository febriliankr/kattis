#include <iostream>
using namespace std;

int main(){

    int NUMBER;
    cin>>NUMBER;

    for (int ITERATION=0; ITERATION<NUMBER; ITERATION++){
    int n;
    cin>>n;
    int A[1005];
    memset(A, 0, sizeof A);
    int M[1005];
    memset(M, 0, sizeof M);
    int gap[1005];
    memset(gap, 0, sizeof gap);

    for (int i=0;i<n;i++){
        cin>>M[i];
        A[i]=M[i];
    }

    sort (A,A+n);

    int tmp = A[0];

    for (int i=1;i<n;i++){
        gap[i]= abs(A[i]-A[i-1]);
    }

    bool beda=false;

    for (int i=2;i<n;i++){
        if (gap[i]!=gap[i-1]){beda=true;}
    }

    if (beda){
        cout<<"non-arithmetic"<<endl;
    }

    bool permuted=false;
    if (beda==false){
        for (int i=0;i<n;i++){
            if ((M[i]!=A[i])&&( M[i]!=A[n-i-1])){
                cout<<"permuted arithmetic"<<endl; permuted=true;
                break;
                }
        }
    }
    if ((!permuted)&&(!beda)){cout<<"arithmetic"<<endl;}
}
return 0;
}

